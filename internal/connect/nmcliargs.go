package connect

type WiFiBase struct {
	SSID    string
	ConName string
	Iface   string
}

type WiFiSecurity interface {
	nmcliArgs() []string
}

type WiFiConnection struct {
	Base     WiFiBase
	Security WiFiSecurity
}

type OpenSec struct {
	OWE bool
}

func (o OpenSec) nmcliArgs() []string {
	if o.OWE {
		return []string{
			"wifi-sec.key-mgmt", "owe",
		}
	}
	return []string{}
}

type PSKSec struct {
	Passphrase string
	SAE        bool
}

func (p PSKSec) nmcliArgs() []string {
	if p.SAE {
		return []string{
			"wifi-sec.key-mgmt", "sae",
			"wifi-sec.psk", p.Passphrase,
		}
	}
	return []string{
		"wifi-sec.key-mgmt", "wpa-psk",
		"wifi-sec.psk", p.Passphrase,
	}
}

type PEAPSec struct {
	Username string
	Password string
	CaCert   string
}

func (p PEAPSec) nmcliArgs() []string {
	args := []string{
		"wifi-sec.key-mgmt", "wpa-eap",
		"802-1x.eap", "peap",
		"802-1x.identity", p.Username,
		"802-1x.password", p.Password,
		"802-1x.phase2-auth", "mschapv2",
	}

	if p.CaCert != "" {
		args = append(args,
			"802-1x.ca-cert", p.CaCert,
		)
	}

	return args
}

type TLSSec struct {
	Identity    string
	ClientCert  string
	CaCert      string
	PrivateKey  string
	PrivKeyPass string
}

func (t TLSSec) nmcliArgs() []string {
	args := []string{
		"wifi-sec.key-mgmt", "wpa-eap",
		"802-1x.eap", "tls",
		"802-1x.identity", t.Identity,
		"802-1x.client-cert", t.ClientCert,
		"802-1x.private-key", t.PrivateKey,
		"802-1x.ca-cert", t.CaCert,
	}

	if t.PrivKeyPass != "" {
		args = append(args,
			"802-1x.private-key-password", t.PrivKeyPass,
		)
	}

	return args
}

func (w WiFiConnection) buildNmcliConnArgs() []string {
	args := []string{
		"connection", "add",
		"type", "wifi",
		"ifname", w.Base.Iface,
		"con-name", w.Base.ConName,
		"ssid", w.Base.SSID,
	}
	return append(args, w.Security.nmcliArgs()...)
}
