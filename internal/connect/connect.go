// Package connect: SSID connection wizard
package connect

func DoScan(iface string) ([]SSIDEntry, error) {
	err := RunWpacliScan(iface)
	if err != nil {
		return nil, err
	}
	rawScan, err := RunWpacliScanResults(iface)
	if err != nil {
		return nil, err
	}
	savedSSIDs, err := RunNmcliConnShow()
	if err != nil {
		return nil, err
	}
	wpacliStatus, err := RunWpacliStatus(iface)
	if err != nil {
		return nil, err
	}
	scanList, err := BuildScanList(rawScan)
	if err != nil {
		return nil, err
	}
	groupedBySSID := GroupBySSID(scanList)
	ssidList, err := BuildSSIDList(groupedBySSID)
	if err != nil {
		return nil, err
	}
	ssidListSaved := CheckIfSSIDSaved(savedSSIDs, ssidList)
	connectedSSID := GetConnectedSSID(wpacliStatus)
	ssidListConnected := CheckIfSSIDConn(connectedSSID, ssidListSaved)
	ssidListSorted := SortByRSSI(ssidListConnected)
	return ssidListSorted, nil
}

func DoConnect(iface string, ssidEntry SSIDEntry, sec WiFiSecurity) error {
	base := WiFiBase{
		SSID:    ssidEntry.SSID,
		ConName: ssidEntry.SSID,
		Iface:   iface,
	}
	conn := WiFiConnection{
		Base:     base,
		Security: sec,
	}
	uuid, err := RunNmcliConnAdd(conn)
	if err != nil {
		return err
	}
	err = RunNmcliConnUp(uuid)
	if err != nil {
		return err
	}
	return nil
}
