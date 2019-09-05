package article

import "github.com/arielizuardi/sph-backend-coding-challenge/model"

// Repository defines article repository contract
type Repository interface {
	CreateArticle(a *model.Article) error
	GetArticleByID(id int) (*model.Article, error)
	GetAllArticle() ([]*model.Article, error)
}
