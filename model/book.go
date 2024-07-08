package model

type Book struct {
	UUID      string `json:"uuid"`
	Title     string `json:"title"`
	ISBN      string `json:"isbn"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Year      int    `json:"year"`
	Category  string `json:"category"`
	Location  string `json:"location"`
	Eksemplar int    `json:"eksemplar"`
}
