package cart

import (
	"cart-microservice/helper"
	"cart-microservice/manga"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type handler struct {
	cartService Service
}

func NewHandler(cartService Service) *handler {
	return &handler{cartService}
}

func (this *handler) GetCartByUUIDHandler(c *gin.Context) {
	uuid := c.Params.ByName("userUUID")

	cartGotten, err := this.cartService.GetCartByUUID(uuid)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get cart failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	cartResponse := []CartGottenFormatted{}

	for _, content := range cartGotten {
		client := &http.Client{}

		httpRequest, err := http.NewRequest("GET", "http://"+os.Getenv("MANGA_MICROSERVICE_URL")+fmt.Sprintf("/api/manga/%s", content.MangaID), nil)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			response := helper.APIResponse("Get cart failed", http.StatusBadRequest, "error", errorMessage)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		httpRequest.Header.Set("Content-Type", "application/json")

		response, err := client.Do(httpRequest)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			response := helper.APIResponse("Get cart failed", http.StatusBadRequest, "error", errorMessage)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		manga := manga.Manga{}
		json.NewDecoder(response.Body).Decode(&manga)
		cartFormatted := FormatMangaGotten(manga, content.Quantity)
		cartResponse = append(cartResponse, cartFormatted)

		response.Body.Close()
	}

	c.JSON(http.StatusOK, cartResponse)
	return
}

func (this *handler) AddItemByMangaUUIDHandler(c *gin.Context) {
	var cartInput CartInput

	err := c.ShouldBindJSON(&cartInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Add cart failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	cartAdded, err := this.cartService.AddItem(cartInput)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Add cart failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	c.JSON(http.StatusOK, cartAdded)
	return
}

func (this *handler) UpdateQuantityByCartUUIDHandler(c *gin.Context) {
	var updateQuantityInput UpdateQuantityInput

	err := c.ShouldBindJSON(&updateQuantityInput)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Add cart failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	uuid := updateQuantityInput.CartID
	quantity := updateQuantityInput.Quantity

	err = this.cartService.UpdateQuantityByCartUUID(uuid, quantity)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Update quantity failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	c.JSON(http.StatusOK, nil)
	return
}

func (this *handler) DeleteCartByUUIDHandler(c *gin.Context) {
	uuid := c.Params.ByName("cartUUID")

	err := this.cartService.DeleteCartByUUID(uuid)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Delete cart failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	c.JSON(http.StatusOK, nil)
	return
}

func (this *handler) DeleteUserCartByUUIDHandler(c *gin.Context) {
	uuid := c.Params.ByName("userUUID")

	err := this.cartService.DeleteUserCartByUUID(uuid)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Delete user cart failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	c.JSON(http.StatusOK, nil)
	return
}
