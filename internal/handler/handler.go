package handler

import (
	parser "PetProject/internal/parsers"
	"PetProject/internal/service"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewsHandler struct {
	service *service.NewsService
}

func NewNewsHandler(service *service.NewsService) *NewsHandler {
	return &NewsHandler{service: service}
}

func (h *NewsHandler) GetNews(c *gin.Context) {
	ctx := context.Background()
	limit := 10
	offset := 0

	news, err := h.service.GetNews(ctx, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения новостей"})
		return
	}

	c.JSON(http.StatusOK, news)
}

func (h *NewsHandler) ParseNews(c *gin.Context) {
	ctx := context.Background()
	url := "https://news.ycombinator.com/"

	newsList, err := parser.ParseNews(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка парсинга"})
		return
	}

	go func() {
		for _, news := range newsList{
			err := h.service.SaveNews(ctx,news)
			if err!=nil{
				log.Println("Ошибка с сохранением новости:", err)
			}
		}
	}()
	c.JSON(http.StatusOK, gin.H{"message": "Парсинг завершён", "news_count": len(newsList)})
}