package usecase

import (
	"github.com/arielizuardi/sph-backend-coding-challenge/article"
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
func (u *Usecase) CreateArticle(a *article.Article) error {
	if err := u.validate.Struct(a); err != nil {
		return err
	}

	return u.repository.CreateArticle(a)
}

// GetArticleByID get article by given id
func (u *Usecase) GetArticleByID(id int) (*article.Article, error) {
	return u.repository.GetArticleByID(id)
}

// GetAllArticle return all articles stored in repository
func (u *Usecase) GetAllArticle() ([]*article.Article, error) {
	return u.repository.GetAllArticle()
}
