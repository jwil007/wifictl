// Package connect: SSID connection wizard
package connect

import (
	"os/exec"
)

type ScanList struct {
	BSSID   string
	SSID    string
	RSSI    int
	SecType string
	Freq    int
}

func RunWpacliScan(iface string) error {
	c := exec.Command("wpa_cli", "-i", iface, "scan")
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

func RunWpacliScanResults(iface string) error {
	c := exec.Command("wpa_cli", "-i", iface, "scan_results")
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}
