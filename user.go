package main

type task struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type user struct {
	Name    string  `json:"Name"`
	Address string  `json:"Address"`
	Phone   int     `json:"Phone"`
	Tasks   []*task `json:"Tasks"`
}

type users struct {
	Users []*user `json:"Users"`
}
