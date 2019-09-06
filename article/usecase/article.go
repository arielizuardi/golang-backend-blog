package usecase

import (
	"github.com/arielizuardi/golang-backend-blog/article"
	"github.com/arielizuardi/golang-backend-blog/model"
	"gopkg.in/go-playground/validator.v9"
)

// Usecase implements article usecase
type Usecase struct {
	repository article.Repository
	validate   *validator.Validate
}

// NewArticleUsecase return new instances of usecase
func NewArticleUsecase(r article.Repository) *Usecase {
	validate := validator.New()
	return &Usecase{repository: r, validate: validate}
}

// CreateArticle create an article
func (u *Usecase) CreateArticle(a *model.Article) error {
	if err := u.validate.Struct(a); err != nil {
		return err
	}

	return u.repository.CreateArticle(a)
}

// GetArticleByID get article by given id
func (u *Usecase) GetArticleByID(id int) (*model.Article, error) {
	return u.repository.GetArticleByID(id)
}

// GetAllArticle return all articles stored in repository
func (u *Usecase) GetAllArticle() ([]*model.Article, error) {
	return u.repository.GetAllArticle()
}
