// Package connect: SSID connection wizard
package connect

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

	ssidList, err := BuildSSIDList(groupedBySSID)
	if err != nil {
		return err
	}

	// sort by RSSI - will prob remove later
	ssidListSorted := SortByRSSI(ssidList)

	Tui(ssidListSorted)

	return nil
}
