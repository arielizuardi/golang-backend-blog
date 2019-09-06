package repository

import (
	"github.com/arielizuardi/golang-backend-blog/model"
	"github.com/jinzhu/gorm"
)

// PostgresArticleRepository implements article repository
type PostgresArticleRepository struct {
	conn *gorm.DB
}

// NewPostgresArticleRepository return new instances of postgres article repository
func NewPostgresArticleRepository(conn *gorm.DB) *PostgresArticleRepository {
	return &PostgresArticleRepository{conn: conn}
}

// CreateArticle ...
func (r *PostgresArticleRepository) CreateArticle(a *model.Article) error {
	return r.conn.Create(&a).Error
}

// GetArticleByID ...
func (r *PostgresArticleRepository) GetArticleByID(id int) (*model.Article, error) {
	var result model.Article
	if err := r.conn.Where(model.Article{ID: id}).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}

// GetAllArticle ...
func (r *PostgresArticleRepository) GetAllArticle() ([]*model.Article, error) {
	var result []*model.Article
	if err := r.conn.Find(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}
