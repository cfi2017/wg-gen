package pkg

import (
	"errors"
	"net"
)

func GetFreeAddress(network net.IPNet, used []net.IP) (net.IP, error) {
	for ip := dupIP(network.IP); network.Contains(ip); Increment(ip) {
		found := false
		for _, u := range used {
			if ip.Equal(u) {
				found = true
				break
			}
		}
		if !found {
			return ip, nil
		}
	}
	return nil, errors.New("no free ip")
}

func dupIP(ip net.IP) net.IP {
	dup := make(net.IP, len(ip))
	copy(dup, ip)
	return dup
}

func Increment(ip net.IP) {
	// for each segment (reverse)
	for j := len(ip) - 1; j >= 0; j-- {
		// increment segment
		ip[j]++
		// if segment didn't wrap around break, else increment next segment
		if ip[j] > 0 {
			break
		}
	}
}
