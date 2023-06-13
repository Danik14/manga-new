package cart

type Cart struct {
	ID       string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID   string `gorm:"type:uuid" json:"user_id"`
	MangaID  string `gorm:"type:uuid" json:"manga_id"`
	Quantity uint   `gorm:"type:uint" json:"quantity"`
}

type CartGottenFormatted struct {
	Manga    Manga `json:"manga"`
	Quantity uint  `json:"quantity"`
}

type Manga struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseYear int    `json:"releaseYear"`
	Status      string `json:"status"`
	ChaptersNum int    `json:"chaptersNum"`
	Author      string `json:"author"`
}

type UpdateQuantityInput struct {
	CartID   string `json:"cart_id"`
	Quantity uint   `json:"quantity"`
}
