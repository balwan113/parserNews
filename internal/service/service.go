package service

import (
	"PetProject/internal/repository"
	"PetProject/models"
	"context"
	"log"
)

type NewsService struct {
	repo *repository.NewsRepository
}

func NewNewsService(repo *repository.NewsRepository) *NewsService {
	return &NewsService{repo: repo}
}

func (s *NewsService) GetNews(ctx context.Context, limit, offset int) ([]models.News, error) {
	return s.repo.GetNews(ctx, limit, offset)
}

func (s *NewsService) SaveNews(ctx context.Context, news models.News) error {
	log.Println("Новость сохранена:", news.Title, news.Link)
	return s.repo.SaveNews(ctx, news)
}
