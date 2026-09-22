package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func sendTGMSG(token, chatID, message string) error {

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", token,
		chatID, url.QueryEscape(message))

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error sending message: ", err)
		return err
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	return nil
}
