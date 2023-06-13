package manga

type MangaGottenFormatted struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseYear int    `json:"releaseYear"`
	Status      string `json:"status"`
	ChaptersNum int    `json:"chaptersNum"`
	Price       int    `json:"price"`
	Author      string `json:"author"`
}

func FormatMangasGotten(allMangasGotten []Manga) []MangaGottenFormatted {
	allMangasFormatted := []MangaGottenFormatted{}

	for _, mangaGotten := range allMangasGotten {
		mangaFormatted := MangaGottenFormatted{}

		mangaFormatted.ID = mangaGotten.ID
		mangaFormatted.Title = mangaGotten.Title
		mangaFormatted.Description = mangaGotten.Description
		mangaFormatted.ReleaseYear = mangaGotten.ReleaseYear
		mangaFormatted.Status = mangaGotten.Status
		mangaFormatted.ChaptersNum = mangaGotten.ChaptersNum
		mangaFormatted.Author = mangaGotten.Author
		mangaFormatted.Price = mangaGotten.Price
		allMangasFormatted = append(allMangasFormatted, mangaFormatted)
	}

	return allMangasFormatted
}
