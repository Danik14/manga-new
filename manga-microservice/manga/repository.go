package manga

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]Manga, error)
	GetByTitle(title string) ([]Manga, error)
	GetByAuthor(author string) ([]Manga, error)
	GetByUUID(id string) (Manga, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Manga, error) {
	mangas := []Manga{}
	err := r.db.Find(&mangas).Error
	if err != nil {
		return nil, err
	}

	return mangas, nil
}

func (r *repository) GetByTitle(title string) ([]Manga, error) {
	mangas := []Manga{}
	err := r.db.Find(&mangas).Where("title = ?", title).Error
	if err != nil {
		return nil, err
	}

	return mangas, nil
}

func (r *repository) GetByAuthor(title string) ([]Manga, error) {
	mangas := []Manga{}
	err := r.db.Find(&mangas).Where("author = ?", title).Error
	if err != nil {
		return nil, err
	}

	return mangas, nil
}

func (r *repository) GetByUUID(id string) (Manga, error) {
	manga := Manga{}
	err := r.db.Find(&manga).Where("id = ?", id).Error
	if err != nil {
		return manga, err
	}

	return manga, nil
}
