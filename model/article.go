package model

// Article represents article
type Article struct {
	ID      int    `json:"id" db:"id"`
	Title   string `json:"title" db:"title" validate:"required"`
	Content string `json:"content" db:"content" validate:"required"`
	Author  string `json:"author" db:"author" validate:"required"`
}

// TableName set table name in DB
func (a *Article) TableName() string {
	return "article"
}
