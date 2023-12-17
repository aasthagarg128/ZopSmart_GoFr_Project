package models

// Book represents the structure of a book
type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Publisher string `json:"publisher"`
}

