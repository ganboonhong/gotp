package user

const Table = "users"

type User struct {
	ID        uint   `json:"id,omitempty"`
	Account   string `json:"account"`
	Password  string `json:"password"`
	CreatedAt int    `json:"created_at,omitempty"`
}
