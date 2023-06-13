package manga

import "gorm.io/gorm"

type Manga struct {
	gorm.Model
	ID          string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Title       string `gorm:"type:varchar(100)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	ReleaseYear int    `gorm:"type:int" json:"releaseYear"`
	Status      string `gorm:"type:varchar(50)" json:"status"`
	ChaptersNum int    `gorm:"type:int" json:"chaptersNum"`
	Price       int    `gorm:"type:int" json:"price"`
	Author      string `gorm:"type:varchar(150)" json:"author"`
}
