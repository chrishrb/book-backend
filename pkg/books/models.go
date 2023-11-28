package books

type Book struct {
	ID          string `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	NoOfPages   int    `db:"no_of_pages" json:"no_of_pages"`
	Year        int    `db:"year" json:"year"`
}
