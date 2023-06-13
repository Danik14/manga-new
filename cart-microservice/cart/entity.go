package cart

type Cart struct {
	ID       string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID   string `gorm:"type:uuid" json:"user_id"`
	MangaID  string `gorm:"type:uuid" json:"manga_id"`
	Quantity uint   `gorm:"type:uint" json:"quantity"`
}
