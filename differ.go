package main

import "fmt"

func buildPortMap(hosts []Host) map[string]bool {

	res := make(map[string]bool)
	for _, h := range hosts {
		for _, p := range h.Ports {
			res[fmt.Sprintf("%s:%d", h.Ip, p.Port)] = true
		}
	}
	return res
}
