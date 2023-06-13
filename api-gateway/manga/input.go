package manga

type RequestModel struct {
	MangaID  string `json:"manga_id"`
	Quantity uint   `json:"quantity"`
}
