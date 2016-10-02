package handbookappdb

type AppLog struct {
	UserID      string `json:"UserId"`
	LogDateTime string
	LogName     string
	LogDataJson string
}
