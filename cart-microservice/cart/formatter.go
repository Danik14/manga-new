package cart

import (
	"cart-microservice/manga"
)

type CartGottenFormatted struct {
	Manga    manga.Manga `json:"manga"`
	Quantity uint        `json:"quantity"`
}

func FormatMangaGotten(manga manga.Manga, quantity uint) CartGottenFormatted {
	cartGottenFormatted := CartGottenFormatted{}

	cartGottenFormatted.Manga = manga
	cartGottenFormatted.Quantity = quantity

	return cartGottenFormatted
}
