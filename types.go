package main

type User struct {
	Id string `json:"id"`
}

type Draft struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	Text   string `json:"text"`
}
