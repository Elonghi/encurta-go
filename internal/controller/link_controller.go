package controller

import (
	"net/http"

	"github.com/elonghi/encurtago/internal/requests"
	"github.com/elonghi/encurtago/internal/responses"
	"github.com/elonghi/encurtago/internal/service"
	"github.com/gin-gonic/gin"
)

type LinkController struct {
	service *service.LinkService
}

func NewLinkController(linkService *service.LinkService) *LinkController {
	return &LinkController{
		service: linkService,
	}
}

func (LinkController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func (l LinkController) Shorten(c *gin.Context) {

	var linkRequest requests.LinkRequest
	err := c.ShouldBindJSON(&linkRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	shortenUrl, err := l.service.ShortenURL(linkRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	response := responses.SetLinkResponse(shortenUrl)

	c.JSON(200, gin.H{
		"success":     true,
		"shorten_url": response,
	})
}

func (l LinkController) Unshorten(c *gin.Context) {
	unshortenUrl, err := l.service.UnshortenURL(c.Query("url"))

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"unshorten_url": unshortenUrl,
	})
}

func (l LinkController) Redirect(c *gin.Context) {
	shortURL := c.Param("shortURL")
	unshortenUrl, err := l.service.UnshortenURL(shortURL)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	http.Redirect(c.Writer, c.Request, unshortenUrl, http.StatusMovedPermanently)
}
