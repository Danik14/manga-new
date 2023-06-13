package manga

type RequestModel struct {
	MangaID  string `json:"mangaId"`
	Quantity uint   `json:"quantity"`
}
