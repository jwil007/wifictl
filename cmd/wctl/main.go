// Package main: Main entrypoint
package main

import (
	"fmt"

	"github.com/jwil007/wifictl/internal/connect"
)

const iface = "wlp0s20f3"

func main() {
	a := connect.RunWpacliScan(iface)
	if a != nil {
		fmt.Printf("error %v", a)
	}
	b := connect.RunWpacliScanResults(iface)
	if b != nil {
		fmt.Printf("error %v", b)
	}
}
