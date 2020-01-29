package pkg

import (
	"bufio"
	"os"
)

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
	return peers, nil
}

func AppendToConfigFile(name string, peer Peer) error {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	out, err := peer.String()
	if err != nil {
		return err
	}
	_, err = file.WriteString(out)
	return err
}
