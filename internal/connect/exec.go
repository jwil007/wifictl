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
	c := exec.Command("wpa_cli", "dev", iface, "scan")
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}
