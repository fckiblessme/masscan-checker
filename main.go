package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {

	for {
		data, err := os.ReadFile("config.json")
		if err != nil {
			fmt.Println("Error reading config.json")
			continue
		}

		var config Config
		err = json.Unmarshal(data, &config)
		if err != nil {
			fmt.Println("Error parsing config.json")
			continue
		}

		rotateResultFile()

		output, err := runScan(config)
		if err != nil {
			fmt.Println("Error running scan:", err)
			continue
		}
		fmt.Println(string(output))

		hosts, err := readHosts("result.json")
		if err != nil {
			fmt.Println("Error reading hosts:", err)
			continue
		}

		prevHosts, err := readHosts("previousRes.json")
		if err != nil {
			fmt.Println("Error reading previous hosts:", err)
			continue
		}

		prevMap := buildPortMap(prevHosts)
		for _, h := range hosts {
			for _, p := range h.Ports {
				key := fmt.Sprintf("%s:%d", h.Ip, p.Port)
				if _, exists := prevMap[key]; !exists {
					fmt.Println("Founded: ", h.Ip, p.Port)
					banner := grabBanner(h.Ip, p.Port)
					sendTGMSG(config.TGToken, config.TGChatID, (fmt.Sprintf("%s:%d:%s", h.Ip, p.Port, banner)))

				}
			}
		}
		time.Sleep(10 * time.Minute)
	}

}
