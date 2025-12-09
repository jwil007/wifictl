// Package connect: SSID connection wizard
package connect

import (
	"fmt"
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

func BuildScanList(iface string) ([]ScanResult, error) {
	scanErr := RunWpacliScan(iface)
	if scanErr != nil {
		return nil, scanErr
	}

	r, scanRErr := RunWpacliScanResults(iface)
	if scanRErr != nil {
		return nil, scanRErr
	}

	linesRaw := strings.Split(string(r), "\n")

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
			fmt.Println(part)
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
	for _, scan := range scanList {
		fmt.Println(scan.BSSID, "\n", scan.Freq, "\n", scan.RSSI, "\n", scan.SecType, "\n", scan.SSID)
	}
	return scanList, nil
}
