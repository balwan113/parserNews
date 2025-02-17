package main

import (
	"PetProject/internal/handler"
	"PetProject/internal/repository"
	"PetProject/internal/service"
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDB() (*pgxpool.Pool, error) {
	// Загружаем .env файл
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("ошибка загрузки .env файла: %w", err)
	}

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга конфигурации БД: %w", err)
	}

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к БД: %w", err)
	}

	return db, nil
}

func main() {
	db, err := ConnectDB()
	if err != nil {
		fmt.Println("Ошибка подключения к базе данных:", err)
		return
	}

	repo := repository.NewNewsRepos(db)
	svc := service.NewNewsService(repo)
	handler := handler.NewNewsHandler(svc)

	route := gin.Default()
	route.GET("/news", handler.GetNews)
	route.POST("/parse", handler.ParseNews)

	route.Run(":8080")
}
