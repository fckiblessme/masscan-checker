package main

import (
	"fmt"
	"os/exec"
)

func runScan(config Config) ([]byte, error) {
	cmd := exec.Command("sudo", "masscan", "-p"+config.Ports, config.Target, "--rate="+config.Rate, "-oJ", "result.json")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing sudo")
		return nil, err
	}
	return output, nil
}
