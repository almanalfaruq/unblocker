package model

type Response struct {
	Answer []Answer `json:"Answer,omitempty"`
}

type Answer struct {
	IP string `json:"data,omitempty"`
}

type ListIP []string
