package main

import (
	"fmt"

	"github.com/jwil007/wifictl/internal/connect"
)

const (
	iface = "wlp0s20f3"
	ssid  = "Get off my LAN"
)

func main() {
	var conn connect.WiFiConnection
	conn.Base.SSID = ssid
	conn.Base.ConName = ssid
	conn.Base.Iface = iface
	conn.Security = connect.PSKSec{
		Passphrase: "Fargo1234",
		SAE:        false,
	}
	uuid, err := connect.RunNmcliConnAdd(conn)
	if err != nil {
		fmt.Printf("Error adding ssid: %s", err)
	}
	err1 := connect.RunNmcliConnUp(uuid)
	if err1 != nil {
		fmt.Printf("Error connecting: %s", err1)
	}
}
