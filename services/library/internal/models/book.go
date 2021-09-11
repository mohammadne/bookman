package models

type Book struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	DateCreated string `json:"date_created,omitempty"`
}
