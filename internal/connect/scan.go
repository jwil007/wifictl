// Package connect: SSID connection wizard
package connect

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

type ScanResult struct {
	BSSID   string
	Freq    int
	RSSI    int
	SecType []string
	SSID    string
}

type SSIDList struct {
	SSID    string
	RSSI    int
	SecType []string
	Bands   []string
}

func BuildScanList(iface string) ([]ScanResult, error) {
	err := RunWpacliScan(iface)
	if err != nil {
		return nil, err
	}

	r, err := RunWpacliScanResults(iface)
	if err != nil {
		return nil, err
	}

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
		}
		rssi, err := strconv.Atoi(parts[2])
		if err != nil {
			log.Printf("Invalid RSSI %v", err)
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

func GroupBySSID(scanList []ScanResult) map[string][]ScanResult {
	groupedBySSID := make(map[string][]ScanResult)

	for _, scan := range scanList {
		groupedBySSID[scan.SSID] = append(groupedBySSID[scan.SSID], scan)
	}
	return groupedBySSID
}
