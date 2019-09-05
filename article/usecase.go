package article

import "github.com/arielizuardi/sph-backend-coding-challenge/model"

// Usecase defines article usecase contract
type Usecase interface {
	CreateArticle(a *model.Article) error
	GetArticleByID(id int) (*model.Article, error)
	GetAllArticle() ([]*model.Article, error)
}
