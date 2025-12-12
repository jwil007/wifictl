package main

import (
	"fmt"

	"github.com/jwil007/wifictl/internal/connect"
)

const (
	iface = "wlp0s20f3"
)

func main() {
	err := connect.Connect(iface)
	if err != nil {
		fmt.Printf("error: %+s", err)
	}
}
