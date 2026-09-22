package main

import (
	"fmt"
	"net"
	"time"
)

func grabBanner(ip string, port int) string {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, time.Second*3)
	if err != nil {
		fmt.Println("Error dialing: ", err)
		return ""
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(time.Second * 2))
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading: ", err)
		return ""
	}

	return "Banner:" + string(buffer[:n])
}
