package handbookappdb

type Page struct {
	ID         string
	EmbeddedID string
	ArticleID  string
	LinkTitle  string
	LinkIDs    []string
}