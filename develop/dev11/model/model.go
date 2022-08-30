package model

import "time"

type Event struct {
	UserId      int       `json:"user_id"`
	Date        time.Time `json:"date"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

type Request struct {
	UserId      int    `json:"user_id"`
	Date        string `json:"date"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

/*
{
"user_id": 1,
"date": "2022-01-10",
"title": "asdddcxzk",
"description": "dsa1212asdsa"
}

*/
