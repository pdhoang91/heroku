package entities

type Account struct {
	ID      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}
