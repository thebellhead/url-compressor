package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	PostURL(longURL string) (string, error)
	GetURL(shortURL string) (string, error)
}

type Service struct {
	repository Repository
}

func New(repo Repository) *Service {
	return &Service{
		repository: repo,
	}
}

type GetURLResponse struct {
	LongURL string `json:"long_url"`
}

type PostURLParams struct {
	LongURL string `json:"long_url"`
}

type PostURLResponse struct {
	ShortURL string `json:"short_url,omitempty"`
}

// @title          URL Compressor
// @version        1.0
// @description    A URL shortening service implemented in Go with Gin
// @contact.name   Grigorev Mikhail
// @contact.url    https://telegram.me/dormant512
// @contact.email  mikegrig@inbox.ru
// @host           localhost:1228
// @BasePath       /

// CompressURL   godoc
// @Summary      Shortens given URL
// @Description  Responds with the shortened URL in JSON
// @Produce      json
// @Param        long_url_json  body      service.PostURLParams  true  "Long URL JSON"
// @Success      200  {object} service.PostURLResponse
// @Router       /compressor [post]
func (s *Service) CompressURL(c *gin.Context) {
	params := new(PostURLParams)
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot parse json body"})
		return
	}
	if params.LongURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty url provided"})
		return
	}
	shortURL, err := s.repository.PostURL(params.LongURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "repository failed with PostURL"})
		return
	}
	c.JSON(http.StatusOK, PostURLResponse{ShortURL: shortURL})
}

// RedirectURL   godoc
// @Summary      Redirects short URL
// @Description  Responds with nothing, redirects to long URL
// @Param        short_url  path      string  true  "short URL"
// @Success      302
// @Router       /compressor/{short_url} [get]
func (s *Service) RedirectURL(c *gin.Context) {
	shortURL := c.Param("compressed_url")
	if shortURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty url provided"})
		return
	}
	redirectURL, err := s.repository.GetURL(shortURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "repository failed with GetURL"})
		return
	}
	c.Redirect(http.StatusFound, redirectURL)
}
