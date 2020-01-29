package pkg

import (
	"net"
	"testing"
)

func TestGetFreeAddress(t *testing.T) {
	target := net.ParseIP("192.168.61.2")
	ip, err := GetFreeAddress(net.IPNet{
		IP:   net.ParseIP(Network),
		Mask: net.CIDRMask(Mask, 32),
	}, []net.IP{net.ParseIP(Broadcast), net.ParseIP(Network), net.ParseIP(Gateway)})
	if err != nil {
		t.Error(err)
	}
	if !ip.Equal(target) {
		t.Error("ips don't match")
	}
}
