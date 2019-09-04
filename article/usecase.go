package article

// Usecase defines article usecase contract
type Usecase interface {
	CreateArticle(a *Article) error
	GetArticleByID(id int) (*Article, error)
	GetAllArticle() ([]*Article, error)
}
