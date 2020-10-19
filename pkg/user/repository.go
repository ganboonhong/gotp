package user

type Reader interface {
	Find(id int) (*User, error)
}

type Writer interface {
	Store(user *User) (int, error)
}

type Repository interface {
	Reader
	Writer
}
