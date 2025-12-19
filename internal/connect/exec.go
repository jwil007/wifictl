// Package connect: SSID connection wizard
package connect

import (
	"fmt"
	"os/exec"
)

func runNmcliConnAdd(connection WiFiConnection) error {
	c := exec.Command(
		"nmcli",
		connection.buildNmcliConnArgs()...)
	err := c.Run()
	if err != nil {
		return fmt.Errorf("nmcli connection add failed: %s", err)
	}
	return nil
}

func runNmcliConnUp(ssid string) error {
	c := exec.Command(
		"nmcli",
		"connection",
		"up",
		ssid,
	)
	err := c.Run()
	if err != nil {
		return fmt.Errorf("connection error: %s", err)
	}
	return nil
}

func runNmcliConnShow() ([]byte, error) {
	out, err := exec.Command("nmcli",
		"-t",
		"-f",
		"	NAME",
		"connection",
		"show").CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error listing connections")
	}
	return out, nil
}

func runNmcliConnDelete(ssid string) error {
	c := exec.Command("nmcli",
		"connection",
		"delete",
		ssid)
	err := c.Run()
	if err != nil {
		return fmt.Errorf("error deleting connection: %s", err)
	}
	return nil
}

func runWpacliScan(iface string) error {
	c := exec.Command("wpa_cli", "-i", iface, "scan")
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

func runWpacliScanResults(iface string) ([]byte, error) {
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

func runWpacliStatus(iface string) ([]byte, error) {
	out, err := exec.Command(
		"wpa_cli",
		"-i",
		iface,
		"status").CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error running wpa_cli status: %s", err)
	}
	return out, nil
}
