package order

type CartInput struct {
	MangaID  string `json:"manga_id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Quantity uint   `json:"quantity"`
}
