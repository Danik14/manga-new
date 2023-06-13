package main

import (
	"manga-microservice/database"
	"manga-microservice/manga"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/api/manga/test", func(c *gin.Context) {
		var testInput = struct {
			Message string `json:"message"`
		}{}

		err := c.ShouldBindJSON(&testInput)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "bad"})
			return
		}

		c.JSON(http.StatusOK, testInput)
		return
	})

	db := database.StartConnection()

	var (
		mangaRepository = manga.NewRepository(db)
		mangaService    = manga.NewService(mangaRepository)
		mangaHandler    = manga.NewHandler(mangaService)
	)
	router.GET("/api/manga", mangaHandler.GetAllMangasHandler)
	router.GET("/api/manga/title/:title", mangaHandler.GetMangasByTitleHandler)
	router.GET("/api/manga/author/:author", mangaHandler.GetMangasByAuthorHandler)

	router.GET("/api/manga/:mangaUUID", mangaHandler.GetMangaByUUID)
	router.POST("/api/manga/total", mangaHandler.GetTotalByUUIDHandler)

	router.Run(":8080")
}
