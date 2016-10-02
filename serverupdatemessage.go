package handbookappdb

import (
	"time"
)

// ServerUpdateMessage is the structure for delivering udpates to mobile app
type ServerUpdateMessage struct {
	ID                int `json:"Id"`
	Time              time.Time
	Action            string
	ArticleID         string
	ArticleTitle      string
	ArticleContent    string
	BookpageID        string
	BookpageArticleID string
	BookpageLinkTitle string
	BookpageLinkIDs   []string
	BookID            string
	BookTitle         string
	BookStartingID    string
	BookOrder         int
	FullPageID        string
	FullPageContent   string
	FullPageTitle     string
}
