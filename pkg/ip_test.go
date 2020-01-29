package pkg

import (
	"net"
	"testing"
)

func TestGetFreeAddress(t *testing.T) {
	target := net.ParseIP("192.168.61.2")
	ip, err := GetFreeAddress(net.IPNet{
		IP:   net.ParseIP(NetworkFlag()),
		Mask: net.CIDRMask(MaskFlag(), 32),
	}, []net.IP{net.ParseIP(BroadcastFlag()), net.ParseIP(NetworkFlag()), net.ParseIP(GatewayFlag())})
	if err != nil {
		t.Error(err)
	}
	if !ip.Equal(target) {
		t.Error("ips don't match")
	}
}
