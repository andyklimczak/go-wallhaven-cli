package wallhaven

type SearchData struct {
	Data []Listing `json:"data"`
}

type Listing struct {
	ID       string `json:"id"`
	Path     string `json:"path"`
	FileType string `json:"file_type"`
}
