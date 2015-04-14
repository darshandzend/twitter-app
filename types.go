package main

type user struct {
	Id string `json:"id"`
}

type draft struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	Text   string `json:"text"`
}
