// Package connect: SSID connection wizard
package connect

import (
	"os/exec"
)

func RunWpacliScan(iface string) error {
	c := exec.Command("wpa_cli", "-i", iface, "scan")
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

func RunWpacliScanResults(iface string) ([]byte, error) {
	out, err := exec.Command(
		"wpa_cli",
		"-i",
		iface,
		"scan_results").CombinedOutput()
	if err != nil {
		return nil, err
	}
	return out, nil
}
