package parser

import (
	"PetProject/models"
	"log"

	"github.com/gocolly/colly/v2"
)

func ParseNews(url string) ([]models.News, error) {
	collector := colly.NewCollector()
	var newsList []models.News

	collector.OnHTML(".titleline > a", func(e *colly.HTMLElement) {
		title := e.Text
		link := e.Attr("href")

		news := models.News{
			Title:title, 
			Link:link,
		}
		newsList = append(newsList, news)
	})

	collector.OnError(func(r *colly.Response, err error) {
		log.Println("Ошибка парсинга:", err)
	})

	err := collector.Visit(url)
	if err != nil {
		return nil, err
	}

	return newsList, nil
}
