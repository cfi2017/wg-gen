package pkg

import (
	"errors"
	"net"
	"strings"
)

type Peers []Peer

func (pp Peers) IPs() []net.IP {
	ips := make([]net.IP, len(pp))
	for i, peer := range pp {
		ips[i] = peer.IP
	}
	return ips
}

type Peer struct {
	PrivateKey Key
	PublicKey  Key
	IP         net.IP
}

func GeneratePeerWithPublicKey(pk string) (peer Peer, err error) {
	peer.PublicKey, err = ParseKey(pk)
	if err != nil {
		return
	}
	peers, err := ParseConfigFile(ConfigFile)
	if err != nil {
		return
	}
	subnet := net.IPNet{
		IP:   net.ParseIP(Network),
		Mask: net.CIDRMask(Mask, 32),
	}
	peer.IP, err = GetFreeAddress(subnet, append(peers.IPs(), net.ParseIP(Network), net.ParseIP(Gateway), net.ParseIP(Broadcast)))
	return
}

func GeneratePeer() (peer Peer, err error) {
	peer.PrivateKey, err = GeneratePrivateKey()
	if err != nil {
		return
	}
	peer.PublicKey = peer.PrivateKey.PublicKey()
	peers, err := ParseConfigFile(ConfigFile)
	if err != nil {
		return
	}
	subnet := net.IPNet{
		IP:   net.ParseIP(Network),
		Mask: net.CIDRMask(Mask, 32),
	}
	peer.IP, err = GetFreeAddress(subnet, append(peers.IPs(), net.ParseIP(Network), net.ParseIP(Gateway), net.ParseIP(Broadcast)))
	return
}

func Parse(in string) (p Peer, err error) {
	in = strings.TrimSpace(in)
	lines := strings.Split(in, "\n")
	if len(lines) != 4 {
		err = errors.New("malformed peer")
		return
	}
	fields, err := parseFields(lines[1:])
	if err != nil {
		return
	}
	p.PublicKey, err = ParseKey(fields["PublicKey"])
	if err != nil {
		return
	}
	p.IP = net.ParseIP(strings.Split(fields["AllowedIPs"], "/")[0])
	return
}

func parseFields(lines []string) (fields map[string]string, err error) {
	fields = make(map[string]string)
	for _, line := range lines {
		parts := strings.Split(line, " = ")
		if len(parts) != 2 {
			return nil, errors.New("invalid part length")
		}
		fields[parts[0]] = strings.TrimSpace(parts[1])
	}
	return
}
