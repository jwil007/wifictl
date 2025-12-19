// Package connect: SSID connection wizard
package connect

func DoScan(iface string) ([]SSIDEntry, error) {
	err := runWpacliScan(iface)
	if err != nil {
		return nil, err
	}
	rawScan, err := runWpacliScanResults(iface)
	if err != nil {
		return nil, err
	}
	savedSSIDs, err := runNmcliConnShow()
	if err != nil {
		return nil, err
	}
	wpacliStatus, err := runWpacliStatus(iface)
	if err != nil {
		return nil, err
	}
	scanList, err := buildScanList(rawScan)
	if err != nil {
		return nil, err
	}
	groupedBySSID := groupBySSID(scanList)
	ssidList, err := buildSSIDList(groupedBySSID)
	if err != nil {
		return nil, err
	}
	ssidListSaved := checkIfSSIDSaved(savedSSIDs, ssidList)
	connectedSSID := getConnectedSSID(wpacliStatus)
	ssidListConnected := checkIfSSIDConn(connectedSSID, ssidListSaved)
	ssidListSorted := sortByRSSI(ssidListConnected)
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
	err := runNmcliConnAdd(conn)
	if err != nil {
		return err
	}
	err = runNmcliConnUp(conn.Base.SSID)
	if err != nil {
		return err
	}
	return nil
}

func DoConnectUp(ssid string) error {
	err := runNmcliConnUp(ssid)
	if err != nil {
		return err
	}
	return nil
}

func DoForgetSSID(ssid string) error {
	err := runNmcliConnDelete(ssid)
	if err != nil {
		return err
	}
	return nil
}
