package pkg

import (
	"net"

	. "github.com/spf13/viper"
)

func NetworkFlag() string {
	return GetString("network")
}

func GatewayFlag() string {
	return GetString("gateway")
}

func BroadcastFlag() string {
	return GetString("broadcast")
}

func MaskFlag() int {
	return GetInt("mask")
}

func ConfigFileFlag() string {
	return GetString("wg-config")
}

func DNSFlag() []net.IP {
	raw := GetStringSlice("dns")
	ips := make([]net.IP, len(raw))
	for i, ip := range raw {
		ips[i] = net.ParseIP(ip)
	}
	return ips
}

func PublicKeyFlag() string {
	return GetString("pubkey")
}

func NetworksFlag() []net.IPNet {
	raw := GetStringSlice("networks")
	networks := make([]net.IPNet, len(raw))
	for i, n := range raw {
		_, network, err := net.ParseCIDR(n)
		if err != nil {
			panic(err)
		}
		networks[i] = *network
	}
	return networks
}

func EndpointFlag() string {
	return GetString("endpoint")
}
