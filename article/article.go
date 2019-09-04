package article

// Article represents article
type Article struct {
	ID      int    `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
	Author  string `json:"author" db:"author"`
}
