package models

import "time"

type ChatRoom struct {
	Id       int
	Name 	string
}

type Message struct {
	Id		int
	Type	int
	Contain string
	Time 	time.Time
}