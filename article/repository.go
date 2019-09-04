package article

// Repository defines article repository contract
type Repository interface {
	CreateArticle(a *Article) error
	GetArticleByID(id int) (*Article, error)
	GetAllArticle() ([]*Article, error)
}
