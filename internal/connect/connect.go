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

	savedSSIDs, err := RunNmcliConnShow()
	if err != nil {
		return err
	}

	wpacliStatus, err := RunWpacliStatus(iface)
	if err != nil {
		return err
	}

	scanList, err := BuildScanList(rawScan)
	if err != nil {
		return err
	}

	groupedBySSID := GroupBySSID(scanList)

	ssidList, err := BuildSSIDList(groupedBySSID)
	if err != nil {
		return err
	}

	ssidListSaved := CheckIfSSIDSaved(savedSSIDs, ssidList)

	connectedSSID := GetConnectedSSID(wpacliStatus)

	ssidListConnected := CheckIfSSIDConn(connectedSSID, ssidListSaved)

	// sort by RSSI - will prob remove later
	ssidListSorted := SortByRSSI(ssidListConnected)

	for _, entry := range ssidListSorted {
		fmt.Printf("%+v\n", entry)
	}

	return nil
}
