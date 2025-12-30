package main

import (
	"fmt"

	"github.com/jwil007/wifictl/internal/connect"
)

func main() {
	status, err := connect.MonitorConnection("wlp0s20f3")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", status)
}
