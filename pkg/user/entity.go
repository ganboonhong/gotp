package user

import "time"

const Table = "users"

type User struct {
	Id        int       `json:"id,omitempty"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
