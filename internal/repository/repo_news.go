package repository

import (
	"PetProject/models"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	
)

type NewsRepository struct {
	db *pgxpool.Pool
}

func NewNewsRepos(db *pgxpool.Pool) *NewsRepository {
	return &NewsRepository{db: db}
}


func (r *NewsRepository) GetNews(ctx context.Context, limit, offset int) ([]models.News, error) {
	var newsList []models.News

	query := `SELECT id,title, link FROM news ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		log.Println(" Ошибка выполнения запроса:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var news models.News
		if err := rows.Scan(&news.ID,&news.Title, &news.Link); err!=nil{
			log.Println("Ошибка при Сканировании", err)
		}

		newsList = append(newsList, news)
		}
			return newsList, nil
}


func (r *NewsRepository) SaveNews(ctx context.Context, news models.News) error {
	query := `INSERT INTO news(title, link, created_at)VALUES ( $1,$2, NOW())`
	_,err := r.db.Exec(ctx, query, news.Title, news.Link)

	if err !=nil{
		log.Println(" Ошибка при сохранении новости:", err)
	}
	return err
}
