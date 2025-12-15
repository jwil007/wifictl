# wifictl (wctl)
This project will be a suite of Wi-Fi command line utilities for Linux. In the Linux CLI, you can already do pretty much whatever you need as a Wi-Fi engineer, but it's a challenge to recall the precise syntax for the various tools you'd utilize (nmcli, wpa_cli, iw, etc). The goal if this project is not to reinvent the wheel, but to build wrappers around these tools which improve ease of use.

The binary name is wctl - so you will invoke these tools using a pattern such as `wctl connect`.

This project is written in Golang - mostly as a learning experience for myself and ease of binary distribution. Also because I'm not smart enough to learn Rust or C :)

## wctl connect
A TUI that presents a table with available SSIDs, with the ability to configure security (if needed) and connect. it also supports your standard operations such as forgetting saved SSIDs.
it parses wpa_cli scan data and uses nmcli to connect - this prevents this utility from fighting with Network Manager. It also means it depends on Network Manager, and may require sudo.
