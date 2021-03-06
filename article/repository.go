package article

import "github.com/arielizuardi/golang-backend-blog/model"

// Repository defines article repository contract
type Repository interface {
	CreateArticle(a *model.Article) error
	GetArticleByID(id int) (*model.Article, error)
	GetAllArticle() ([]*model.Article, error)
}
