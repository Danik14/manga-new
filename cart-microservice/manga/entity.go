package manga

type Manga struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseYear int    `json:"releaseYear"`
	Status      string `json:"status"`
	Price       int    `json:"price"`
	ChaptersNum int    `json:"chaptersNum"`
}
