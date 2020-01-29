package pkg

import (
	"bufio"
	"net"
	"os"
	"strings"
	"text/template"
)

type Server struct {
	DNS       []net.IP
	PublicKey string
	Networks  []net.IPNet
	Endpoint  string
}

type Config struct {
	Peer   Peer
	Server Server
}

const (
	ServerTemplateString = `
[Peer]
PublicKey = {{ .Peer.PublicKey.String }}
AllowedIPs = {{ .Peer.IP.String }}/32
PersistentKeepAlive = 60
`
	ClientTemplateString = `
[Interface]
PrivateKey = {{ .Peer.PrivateKey.String }}
Address = {{ .Peer.IP.String }}/32
{{if .Server.DNS}}DNS = {{ range $index, $element := .Server.DNS }}{{if $index}},{{end}}{{$element.String}}{{end}}{{end}}

[Peer]
PublicKey = {{ .Server.PublicKey }}
AllowedIPs = {{ range $index, $element := .Server.Networks }}{{if $index}},{{end}}{{$element}}{{end}}
Endpoint = {{ .Server.Endpoint }}
PersistentKeepalive = 60
`
)

func (c Config) ServerConfig() (out string, err error) {
	t, err := template.New("server").Parse(ServerTemplateString)
	if err != nil {
		return
	}
	w := &strings.Builder{}
	err = t.Execute(w, c)
	return w.String(), nil
}

func (c Config) ClientConfig() (out string, err error) {
	t, err := template.New("client").Parse(ClientTemplateString)
	if err != nil {
		return
	}
	w := &strings.Builder{}
	err = t.Execute(w, c)
	return w.String(), nil

}

func GetDefaultServer() (s Server) {
	s.Endpoint = EndpointFlag()
	s.PublicKey = PublicKeyFlag()
	s.Networks = NetworksFlag()
	s.DNS = DNSFlag()
	return
}

func ParseConfigFile(name string) (Peers, error) {
	peers := make(Peers, 0)
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	// advance to first peer
	for scanner.Scan() && scanner.Text() != "[Peer]" {
	}

	buffer := "[Peer]\n"
	for scanner.Scan() {
		line := scanner.Text()
		if line == "[Peer]" {
			peer, err := Parse(buffer)
			if err != nil {
				return nil, err
			}
			peers = append(peers, peer)
			buffer = ""
		}
		buffer += line + "\n"
	}
	// last peer
	if strings.HasPrefix(strings.TrimSpace(buffer), "[Peer]") {
		peer, err := Parse(strings.TrimSpace(buffer))
		if err != nil {
			return nil, err
		}
		peers = append(peers, peer)
	}
	return peers, nil
}

func AppendToConfigFile(name string, peer Peer) error {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	out, err := Config{
		Peer: peer,
	}.ServerConfig()
	if err != nil {
		return err
	}
	_, err = file.WriteString(out)
	return err
}
