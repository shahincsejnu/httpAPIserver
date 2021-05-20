package in_memory

import (
	"httpAPIserver/domain/entity"
	"httpAPIserver/domain/exception"
	"httpAPIserver/domain/repository"
)

type inMemoryArticleRepository struct {
	data map[string]*entity.Article
}

func (i *inMemoryArticleRepository) FetchAll() ([]*entity.Article, error) {
	articles := make([]*entity.Article, 0)

	for _, article := range i.data {
		articles = append(articles, article)
	}

	return articles, nil
}

func (i *inMemoryArticleRepository) FetchByID(id string) (*entity.Article, error) {
	if article, ok := i.data[id]; ok {
		return article, nil
	}

	return nil, exception.ErrNotFound
}

func (i *inMemoryArticleRepository) Create(article *entity.Article) (*entity.Article, error) {
	// check if article already exist with the id
	if _, ok := i.data[article.ID]; ok {
		return nil, exception.ErrConflict
	}
	i.data[article.ID] = article
	return article, nil
}

func (i *inMemoryArticleRepository) Update(id string, article *entity.Article) (*entity.Article, error) {
	// check for article existence.
	if _, ok := i.data[id]; !ok {
		return nil, exception.ErrNotFound
	}

	// delete the index
	delete(i.data, id)

	// insert new
	return i.Create(article)
}

func (i *inMemoryArticleRepository) Delete(id string) error {
	// check for article existence, delete the index if exist
	if _, ok := i.data[id]; ok {
		delete(i.data, id)
		return nil
	}
	return exception.ErrNotFound
}

// NewInMemoryArticleRepository returns an in memory implementation of ArticleRepository
func NewInMemoryArticleRepository() repository.ArticleRepository {
	data := make(map[string]*entity.Article)
	return &inMemoryArticleRepository{data: data}
}
