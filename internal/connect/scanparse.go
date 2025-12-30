// Package connect: SSID connection wizard
package connect

import (
	"bufio"
	"fmt"
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ScanResult struct {
	BSSID   string
	Freq    int
	RSSI    int
	SecType []string
	SSID    string
}

type SSIDEntry struct {
	SSID       string
	RSSI       int
	BSSIDCount int
	SecType    []string
	Bands      []string
	Saved      bool
	Connected  bool
}

type ConnectionStatus struct {
	Duration time.Duration
	Status   string
	Failure  bool
}

func buildScanList(rawScan []byte) ([]ScanResult, error) {
	r := string(rawScan)

	linesRaw := strings.Split(string(r), "\n")
	// Remove first line from array, since it's a header
	lines := slices.Delete(linesRaw, 0, 1)
	var scanList []ScanResult

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) < 5 {
			// this skips lines wtih null SSID
			continue
		}
		// SSIDs can contain spaces, so find where SSID begins
		// and join as its own element
		flags := parts[3]
		idx := strings.Index(line, flags)
		ssid := strings.TrimSpace(line[idx+len(flags):])

		// trim flags element to extract SecType
		flagParts := strings.Split(flags, "][")
		var secType []string
		for _, part := range flagParts {
			if strings.Contains(part, "WPA") {
				s := strings.Replace(part, "[", "", 1)
				secType = append(secType, s)
			}
		}

		// Convert ASCII to ints for Freq and RSSI
		freq, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Printf("Invalid freq %v", err)
			return nil, err
		}
		rssi, err := strconv.Atoi(parts[2])
		if err != nil {
			log.Printf("Invalid RSSI %v", err)
			return nil, err
		}

		scanList = append(scanList, ScanResult{
			BSSID:   parts[0],
			Freq:    freq,
			RSSI:    rssi,
			SecType: secType,
			SSID:    ssid,
		})
	}
	return scanList, nil
}

func groupBySSID(scanList []ScanResult) map[string][]ScanResult {
	groupedBySSID := make(map[string][]ScanResult)

	for _, scan := range scanList {
		groupedBySSID[scan.SSID] = append(groupedBySSID[scan.SSID], scan)
	}
	return groupedBySSID
}

func buildSSIDList(groupedBySSID map[string][]ScanResult) ([]SSIDEntry, error) {
	var ssidList []SSIDEntry
	for ssid, items := range groupedBySSID {
		// create slices for each attribute
		var rssiList []int
		var bssidList []string
		var secList [][]string
		var freqList []int

		for _, item := range items {
			rssiList = append(rssiList, item.RSSI)
			bssidList = append(bssidList, item.BSSID)
			secList = append(secList, item.SecType)
			freqList = append(freqList, item.Freq)
		}
		maxRssi, err := findMax(rssiList)
		if err != nil {
			return nil, err
		}

		bssidCount := len(bssidList)
		secType := secList[0]

		bands, err := processBands(freqList)
		if err != nil {
			return nil, err
		}

		ssidList = append(ssidList, SSIDEntry{
			SSID:       ssid,
			RSSI:       maxRssi,
			BSSIDCount: bssidCount,
			SecType:    secType,
			Bands:      bands,
		})

	}
	return ssidList, nil
}

func findMax(s []int) (int, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("cannot find max of empty slice")
	}
	maxVal := s[0]
	for _, item := range s {
		if item > maxVal {
			maxVal = item
		}
	}
	return maxVal, nil
}

func processBands(freqList []int) ([]string, error) {
	var bandsRaw []string
	var bands []string

	for _, freq := range freqList {
		switch {
		case freq >= 2412 && freq <= 2484:
			bandsRaw = append(bandsRaw, "2.4GHz")

		case freq >= 5160 && freq <= 5885:
			bandsRaw = append(bandsRaw, "5GHz")

		case freq >= 5935 && freq <= 7115:
			bandsRaw = append(bandsRaw, "6GHz")
		default:
			return nil, fmt.Errorf("invalid freq: %v", freq)
		}
	}

	if slices.Contains(bandsRaw, "2.4GHz") {
		bands = append(bands, "2.4GHz")
	}
	if slices.Contains(bandsRaw, "5GHz") {
		bands = append(bands, "5GHz")
	}
	if slices.Contains(bandsRaw, "6GHz") {
		bands = append(bands, "6GHz")
	}

	return bands, nil
}

func checkIfSSIDSaved(savedSSIDs []byte, ssidList []SSIDEntry) []SSIDEntry {
	s := string(savedSSIDs)
	savedSSIDList := strings.Split(s, "\n")

	for i := range ssidList {
		if slices.Contains(savedSSIDList, ssidList[i].SSID) {
			ssidList[i].Saved = true
		}
	}
	return ssidList
}

func getConnectedSSID(wpacliStatus []byte) string {
	s := string(wpacliStatus)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		parts := strings.Split(line, "=")
		if parts[0] == "ssid" {
			connectedSSID := parts[1]
			return connectedSSID
		}
	}
	return ""
}

func checkIfSSIDConn(connectedSSID string, ssidList []SSIDEntry) []SSIDEntry {
	for i := range ssidList {
		if ssidList[i].SSID == connectedSSID {
			ssidList[i].Connected = true
			return ssidList
		}
	}
	return ssidList
}

func sortByRSSI(ssidList []SSIDEntry) []SSIDEntry {
	sort.Slice(ssidList, func(i, j int) bool {
		return ssidList[i].RSSI > ssidList[j].RSSI
	})
	return ssidList
}

func MonitorConnection(iface string) (ConnectionStatus, error) {
	start := time.Now()
	timeout := time.NewTimer(30 * time.Second)
	defer timeout.Stop()

	out, stop, err := runNmcliDevMonitor(iface)
	if err != nil {
		return ConnectionStatus{}, err
	}

	lines := make(chan string)
	scanErr := make(chan error, 1)

	go func() {
		defer close(lines)
		scanner := bufio.NewScanner(out)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		scanErr <- scanner.Err()
	}()

	for {
		select {
		case <-timeout.C:
			t := time.Now()
			elapsed := t.Sub(start)
			_ = stop()
			return ConnectionStatus{
				Duration: elapsed,
				Status:   "connection timed out",
				Failure:  true,
			}, nil

		case line, ok := <-lines:
			if !ok {
				if err := <-scanErr; err != nil {
					return ConnectionStatus{}, err
				}
				t := time.Now()
				elapsed := t.Sub(start)
				_ = stop()
				return ConnectionStatus{
					Duration: elapsed,
					Status:   "monitor ended",
					Failure:  true,
				}, nil

			}
			switch {
			case strings.Contains(line, ": connected"):
				fmt.Println("connection success")
				t := time.Now()
				elapsed := t.Sub(start)
				_ = stop()

				return ConnectionStatus{
					Duration: elapsed,
					Status:   line,
					Failure:  false,
				}, nil

			case strings.Contains(line, "connection failed"):
				fmt.Println("connection failure")
				t := time.Now()
				elapsed := t.Sub(start)
				_ = stop()
				return ConnectionStatus{
					Duration: elapsed,
					Status:   line,
					Failure:  true,
				}, nil
			}
		}
	}
}
