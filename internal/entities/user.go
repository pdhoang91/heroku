package entities

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	AccountIDs []int  `json:"account_ids"`
}
