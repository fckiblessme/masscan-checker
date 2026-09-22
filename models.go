package main

type Port struct {
	Port   int    `json:"port"`
	Proto  string `json:"proto"`
	Status string `json:"status"`
	Reason string `json:"reason"`
	Ttl    int    `json:"ttl"`
}

type Host struct {
	Ip        string `json:"ip"`
	Timestamp string `json:"timestamp"`
	Ports     []Port `json:"ports"`
}

type Config struct {
	Target   string `json:"target"`
	Ports    string `json:"ports"`
	Rate     string `json:"rate"`
	TGToken  string `json:"tg_token"`
	TGChatID string `json:"tg_chat_id"`
}
