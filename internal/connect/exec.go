// Package connect: SSID connection wizard
package connect

import (
	"fmt"
	"os/exec"
)

func RunNmcliConnAdd(connection WiFiConnection) error {
	c := exec.Command(
		"nmcli",
		connection.BuildNmcliConnArgs()...)
	err := c.Run()
	if err != nil {
		return fmt.Errorf("nmcli connection add failed: %s", err)
	}
	return nil
}

//	outStr := string(out)
//	uuid, err := extractUUID(outStr)
//	if err != nil {
//		return "", err
//	}
//	return uuid, nil

//func extractUUID(s string) (string, error) {
//	start := strings.LastIndex(s, "(")
//	end := strings.LastIndex(s, ")")
//
//	if start == -1 || end == -1 || end <= start {
//		return "", fmt.Errorf("could not extract UUID")
//	}
//
//	return s[start+1 : end], nil
//}

func RunNmcliConnUp(ssid string) error {
	//	if uuid == "" {
	//		return fmt.Errorf("UUID is empty, cannot connect")
	//	}
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

func RunNmcliConnShow() ([]byte, error) {
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

func RunNmcliConnDelete(ssid string) error {
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

func RunWpacliStatus(iface string) ([]byte, error) {
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
