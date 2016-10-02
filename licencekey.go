package handbookappdb

type LicenceKey struct {
	ID           string `json:"Id"`
	HandbookType string
	UserID       string `json:"UserId"`
}
