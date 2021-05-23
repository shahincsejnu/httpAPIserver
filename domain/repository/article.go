package repository

import "httpAPIserver/domain/entity"

type ArticleRepository interface {
	// FetchAll retrieves all Article in the db
	FetchAll() ([]*entity.Article, error)

	// FetchByID retrieves article with the given id from the db
	FetchByID(id string) (*entity.Article, error)

	// Create creates an article in the db
	Create(article *entity.Article) (*entity.Article, error)

	// Update updates and article in the db
	Update(id string, article *entity.Article) (*entity.Article, error)

	// Delete deletes an article from the db
	Delete(id string) error
}
