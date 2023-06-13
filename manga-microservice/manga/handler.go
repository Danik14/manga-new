package manga

import (
	"manga-microservice/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	mangaService Service
}

func NewHandler(mangaService Service) *handler {
	return &handler{mangaService}
}

func (h *handler) GetAllMangasHandler(c *gin.Context) {
	// Call Process
	mangas, err := h.mangaService.GetAllMangas()
	if err != nil {
		errorMessage := gin.H{"message": "Invalid query"}
		response := helper.APIResponse("Get products failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, mangas)
	// return
}

func (h *handler) GetMangasByTitleHandler(c *gin.Context) {
	// Get path params
	titleReq := c.Params.ByName("title")

	// Call Process
	mangas, err := h.mangaService.GetMangasByTitle(titleReq)
	if err != nil {
		errorMessage := gin.H{"message": "Invalid query"}
		response := helper.APIResponse("Get mangas failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, mangas)
}

func (h *handler) GetMangasByAuthorHandler(c *gin.Context) {
	authorReq := c.Params.ByName("author")

	mangas, err := h.mangaService.GetMangasByAuthor(authorReq)
	if err != nil {
		errorMessage := gin.H{"message": "Invalid query"}
		response := helper.APIResponse("Get mangas failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, mangas)
}

func (h *handler) GetMangaByUUID(c *gin.Context) {
	uuid := c.Params.ByName("mangaUUID")

	manga, err := h.mangaService.GetMangaByUUID(uuid)
	if err != nil {
		errorMessage := gin.H{"message": "Invalid query"}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	c.JSON(http.StatusOK, manga)
}

func (h *handler) GetTotalByUUIDHandler(c *gin.Context) {
	var requestObjects []RequestModel
	err := c.ShouldBindJSON(&requestObjects)
	if err != nil {
		errorMessage := gin.H{"message": "Invalid query"}
		response := helper.APIResponse("Get total failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	total, err := h.mangaService.GetTotal(requestObjects)
	if err != nil {
		errorMessage := gin.H{"message": "Invalid query"}
		response := helper.APIResponse("Get total failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, gin.H{"total": total})
}
