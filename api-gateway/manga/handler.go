package manga

import (
	"api-gateway/helper"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Test(c *gin.Context) {
	client := &http.Client{}

	var testInput = struct {
		Message string `json:"message"`
	}{
		Message: "manga microservice ok",
	}
	request, err := json.Marshal(testInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "bad"})
		return
	}

	httpRequest, err := http.NewRequest("POST", "http://"+os.Getenv("MANGA_MICROSERVICE_URL")+"/api/manga/test", bytes.NewBuffer(request))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "bad"})
		return
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		fmt.Println("+++++++Negry++++++")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "bad"})
		return
	}

	var testOutput = struct {
		Message string `json:"message"`
	}{}

	json.NewDecoder(httpResponse.Body).Decode(&testOutput)
	c.JSON(http.StatusOK, testOutput)

	httpResponse.Body.Close()
	return
}

func (this *handler) GetAllMangasHandler(c *gin.Context) {
	client := &http.Client{}

	httpRequest, err := http.NewRequest("GET", "http://"+os.Getenv("MANGA_MICROSERVICE_URL")+"/api/manga", nil)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get mangas failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get mangas failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	mangaResponse := []MangaGottenFormatted{}
	json.NewDecoder(httpResponse.Body).Decode(&mangaResponse)

	response := helper.APIResponse("Get mangas success", http.StatusOK, "success", mangaResponse)
	c.JSON(http.StatusOK, response)

	httpResponse.Body.Close()
	return
}

func (this *handler) GetMangasByTitleHandler(c *gin.Context) {
	titleRequest := c.Params.ByName("title")

	client := &http.Client{}

	httpRequest, err := http.NewRequest("GET", "http://"+os.Getenv("MANGA_MICROSERVICE_URL")+fmt.Sprintf("/api/manga/title/%s", &titleRequest), nil)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get mangas by title failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get mangas by title failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	mangaResponse := []MangaGottenFormatted{}
	json.NewDecoder(httpResponse.Body).Decode(&mangaResponse)

	response := helper.APIResponse("Get mangas by title success", http.StatusOK, "success", mangaResponse)
	c.JSON(http.StatusOK, response)

	httpResponse.Body.Close()
	return
}

func (this *handler) GetMangasByAuthorHandler(c *gin.Context) {
	authorRequest := c.Params.ByName("author")

	client := &http.Client{}

	httpRequest, err := http.NewRequest("GET", "http://"+os.Getenv("MANGA_MICROSERVICE_URL")+fmt.Sprintf("/api/manga/author/%s", &authorRequest), nil)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get mangas by author failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get mangas by author failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	mangaResponse := []MangaGottenFormatted{}
	json.NewDecoder(httpResponse.Body).Decode(&mangaResponse)

	response := helper.APIResponse("Get mangas by author success", http.StatusOK, "success", mangaResponse)
	c.JSON(http.StatusOK, response)

	httpResponse.Body.Close()
	return
}
