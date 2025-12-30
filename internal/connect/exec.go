// Package connect: SSID connection wizard
package connect

import (
	"fmt"
	"io"
	"os/exec"
)

func runNmcliConnAdd(connection WiFiConnection) error {
	out, err := exec.Command(
		"nmcli",
		connection.buildNmcliConnArgs()...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("nmcli connection add failed: %v\n%s", err, string(out))
	}
	return nil
}

func runNmcliConnUp(ssid string) error {
	out, err := exec.Command(
		"nmcli",
		"connection",
		"up",
		ssid,
	).CombinedOutput()
	if err != nil {
		return fmt.Errorf("connection error: %v\n%s", err, string(out))
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
		return nil, fmt.Errorf("error listing connections: %v\n%s", err, string(out))
	}
	return out, nil
}

func runNmcliConnDelete(ssid string) error {
	out, err := exec.Command("nmcli",
		"connection",
		"delete",
		ssid).CombinedOutput()
	if err != nil {
		return fmt.Errorf("error deleting connection: %v\n%s", err, string(out))
	}
	return nil
}

func runWpacliScan(iface string) error {
	out, err := exec.Command("wpa_cli",
		"-i",
		iface, "scan").CombinedOutput()
	if err != nil {
		return fmt.Errorf("error running wpa_cli scan: %v\n%s", err, string(out))
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
		return nil, fmt.Errorf("error running wpa_cli scan scan_results: %v\n%s",
			err, string(out))
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

func runNmcliDevMonitor(iface string) (io.ReadCloser, func() error, error) {
	c := exec.Command("nmcli",
		"device",
		"monitor",
		iface)

	out, err := c.StdoutPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("error with nmcli output: %v", err)
	}

	errStart := c.Start()
	if errStart != nil {
		err := out.Close()
		if err != nil {
			return nil, nil, fmt.Errorf("error closing nmcli monitor: %v", err)
		}
		return nil, nil, fmt.Errorf("error starting nmcli monitor: %v", err)
	}

	stop := func() error {
		if c.Process == nil {
			return nil
		}
		return c.Process.Kill()
	}
	return out, stop, nil
}
