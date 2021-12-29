package usefulCase

import (
	"errors"
	"net"
)

func getLocalIp() (string, error) {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addresses {
		if ipnNet, ok := addr.(*net.IPNet); ok && !ipnNet.IP.IsLoopback() {
			if ipnNet.IP.To4() != nil {
				return ipnNet.IP.String(), nil
			}
		}
	}
	return "", errors.New("can not find the client ip address")
}
