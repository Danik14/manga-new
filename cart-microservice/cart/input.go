package cart

type CartInput struct {
	UserID   string `json:"user_id" binding:"required"`
	MangaID  string `json:"manga_id" binding:"required"`
	Quantity uint   `json:"quantity" binding:"required"`
}

type UpdateQuantityInput struct {
	CartID   string `json:"cart_id"`
	Quantity uint   `json:"quantity"`
}
