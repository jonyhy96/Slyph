package models

type Error struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}
