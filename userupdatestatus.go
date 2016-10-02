package handbookappdb

import (
	"encoding/json"
	"time"
)

// UserUpdateStatus is the struct for a user update status
type UserUpdateStatus struct {
	ID                  string `json:"Id"`
	HandbookType        string
	UpdateNeeded        bool
	LastDateTimeChecked time.Time
	UpdateJson          string
}

// UdpateJsonMessage is the struct for
type UpdateJsonMessage struct {
	AddBookItemIds        []string
	DeleteBookItemIds     []string
	AddFullpageItemIds    []string
	DeleteFullpageItemIds []string
}

// SetUpdateJson is the function to set UpdateJson field
func (u *UserUpdateStatus) SetUpdateJson(j UpdateJsonMessage) (err error) {
	js, err := json.Marshal(j)
	if err != nil {
		return
	}
	u.UpdateJson = string(js)
	return
}

func (u *UserUpdateStatus) GetUpdateJson() (result UpdateJsonMessage, err error) {
	err = json.Unmarshal([]byte(u.UpdateJson), &result)
	return
}
