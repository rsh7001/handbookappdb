package handbookappdb

type Book struct {
	ID                 string `json:"Id"`
	Title              string
	StartingBookPageID string `json:"StartingId"`
	OrderIndex         int    `json:"Order"`
}