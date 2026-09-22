package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func rotateResultFile() error {
	_, err := os.Stat("result.json")
	if os.IsNotExist(err) {
		fmt.Println("No result.json")
	} else {
		err := os.Rename("result.json", "previousRes.json")
		if err != nil {
			fmt.Println("Error renaming result.json")
			return err
		}
	}
	return nil
}

func readHosts(fileName string) ([]Host, error) {
	_, err := os.Stat(fileName)
	var prevData []byte
	if os.IsNotExist(err) {
		fmt.Println("No ", fileName)
	} else {
		prevData, err = os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error reading ", fileName)
		}
	}

	if len(prevData) == 0 {
		return []Host{}, nil
	}

	var prevHosts []Host
	err = json.Unmarshal(prevData, &prevHosts)
	if err != nil {
		fmt.Println("Error parsing ", fileName, err)
		return nil, err
	}
	return prevHosts, nil
}
