// Package connect: SSID connection wizard
package connect

import (
	"fmt"
	"strconv"
)

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
		fmt.Println("wpacli scan error")
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
	fmt.Println("checked for connected SSID")
	ssidListConnected := CheckIfSSIDConn(connectedSSID, ssidListSaved)
	fmt.Println(strconv.Itoa(len(ssidListConnected)))
	fmt.Println("build ssid list with connected ssid")
	ssidListSorted := SortByRSSI(ssidListConnected)
	fmt.Println("sorted SSID list by RSSI")
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
	err := RunNmcliConnAdd(conn)
	if err != nil {
		return err
	}
	err = RunNmcliConnUp(conn.Base.SSID)
	if err != nil {
		return err
	}
	return nil
}

func DoConnectUp(ssid string) error {
	err := RunNmcliConnUp(ssid)
	if err != nil {
		return err
	}
	return nil
}

func DoForgetSSID(ssid string) error {
	err := RunNmcliConnDelete(ssid)
	if err != nil {
		return err
	}
	return nil
}
