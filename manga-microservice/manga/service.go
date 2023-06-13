package manga

type Service interface {
	GetAllMangas() ([]Manga, error)
	GetMangasByTitle(title string) ([]Manga, error)
	GetMangasByAuthor(title string) ([]Manga, error)
	GetMangaByUUID(id string) (Manga, error)
	GetTotal(requestObjects []RequestModel) (int, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllMangas() ([]Manga, error) {
	mangas, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return mangas, nil
}

func (s *service) GetMangasByTitle(title string) ([]Manga, error) {
	mangas, err := s.repository.GetByTitle(title)
	if err != nil {
		return nil, err
	}

	return mangas, nil
}

func (s *service) GetMangasByAuthor(author string) ([]Manga, error) {
	mangas, err := s.repository.GetByAuthor(author)
	if err != nil {
		return nil, err
	}

	return mangas, nil
}

func (s *service) GetMangaByUUID(id string) (Manga, error) {
	manga, err := s.repository.GetByUUID(id)
	if err != nil {
		return manga, err
	}

	return manga, nil
}

func (s *service) GetTotal(requestObjects []RequestModel) (int, error) {
	type total struct {
		Price    int
		Quantity uint
	}

	var totalObjects []total

	for _, c := range requestObjects {
		manga, err := s.repository.GetByUUID(c.MangaID)
		if err != nil {
			return -1, err
		}

		totalObject := total{}
		totalObject.Price = manga.Price
		// totalObject.Quantity = c.Quantity

		totalObjects = append(totalObjects, totalObject)
	}

	var totalPayment int
	for _, c := range totalObjects {
		totalPayment += int(c.Quantity) * c.Price
	}

	return totalPayment, nil
}
