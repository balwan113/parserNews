-- +goose Up
CREATE TABLE news (
                      id SERIAL PRIMARY KEY,       -- Уникальный ID новости
                      title TEXT NOT NULL,         -- Заголовок новости
                      link TEXT NOT NULL,          -- Ссылка на источник
                      created_at TIMESTAMP DEFAULT NOW() -- Дата добавления
);

DROP TABLE news;
-- +goose Down

