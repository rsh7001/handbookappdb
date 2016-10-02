package handbookappdb

type Article struct {
	ID          string `json:"Id"`
	EmbeddedID  string
	ArticleType string
	Title       string
	Summary     string
	HTMLContent string
}