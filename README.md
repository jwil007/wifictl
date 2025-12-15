# wifictl (wctl)
The goal of this project is to build a suite of Wi-Fi command line utilities for Linux. In the Linux CLI, you can already do pretty much whatever you need as a Wi-Fi engineer, but it is difficult the precise syntax, or which CLI tools you need to use.

For the most part, these utilities will be wrappers around existing user space tools (nmcli, wpa_ci, iw). 

The binary name is wctl - so you will invoke these tools using a pattern such as `wctl connect`.

## wctl connect
A TUI that presents a table with available SSIDs, with the ability to configure security (if needed) and connect. Also supports your standard operations such as forgetting saved SSIDs.
it parses wpa_cli scan data and uses nmcli to connect - this prevents this utility from fighting with Network Manager. It also means it requires network manager, and may require sudo.
