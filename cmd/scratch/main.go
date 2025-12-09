package main

import (
	"fmt"

	"github.com/jwil007/wifictl/internal/connect"
)

const iface = "wlp0s20f3"

func main() {
	scanList, err := connect.BuildScanList(iface)
	if err != nil {
		fmt.Println(err)
	}
	groupedBySSID := connect.GroupBySSID(scanList)

	for ssid, items := range groupedBySSID {
		fmt.Println("SSID:", ssid)
		for _, it := range items {
			fmt.Printf("  %+v\n", it)
		}
	}
}
