package user

const Table = "users"

type User struct {
	ID        uint   `json:"id,omitempty"`
	Name      string `json:"name"`
	CreatedAt int    `json:"created_at,omitempty"`
}
