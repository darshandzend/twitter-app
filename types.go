package main

type User struct {
	Id string `json:"id"`
}

type Draft struct {
	Id     string `json:"id"`
	UserId string `json:"-"`
	Text   string `json:"text"`
}
