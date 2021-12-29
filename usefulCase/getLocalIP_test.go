package usefulCase

import (
	"fmt"
	"testing"
)

func TestGetLocalIP(t *testing.T) {
	ip, err := getLocalIp()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(ip)
}
