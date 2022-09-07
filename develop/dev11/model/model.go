package model

import "time"

// Event структура события
type Event struct {
	UserID      int       `json:"user_id"`
	Date        time.Time `json:"date"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

// Request структура для считывания запроса
type Request struct {
	UserID      int    `json:"user_id"`
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
