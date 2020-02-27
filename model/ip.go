package model

type Response struct {
	Answer []Answer `json:"Answer"`
}

type Answer struct {
	Data string `json:"data"`
}

type ListIP []string
