// Package connect: SSID connection wizard
package connect

import "fmt"

// Connect orchestrates connection process
func Connect(iface string) error {
	err := RunWpacliScan(iface)
	if err != nil {
		return err
	}

	rawScan, err := RunWpacliScanResults(iface)
	if err != nil {
		return err
	}

	scanList, err := BuildScanList(rawScan)
	if err != nil {
		return err
	}

	groupedBySSID := GroupBySSID(scanList)

	// debug print, will remove
	for ssid, items := range groupedBySSID {
		fmt.Println("SSID:", ssid)
		for _, it := range items {
			fmt.Printf("  %+v\n", it)
		}
	}

	return nil
}
