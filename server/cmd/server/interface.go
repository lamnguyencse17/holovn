package main

type chatData struct {
	Name      string `json:"name"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	IsTl      bool   `json:"is_tl"`
}