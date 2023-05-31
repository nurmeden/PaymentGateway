package models

type Customer struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Adress string `json:"address"`
}
