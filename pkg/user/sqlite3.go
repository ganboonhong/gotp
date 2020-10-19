package user

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) Find(id int) (u *User, err error) {
	return &User{}, nil
}

func (r *repo) Store(u *User) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	stmt, err := tx.Prepare(`
		INSERT INTO ` + Table + `
		(name, created_at)
		VALUES (?, ?, datetime('now'))
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.Name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	fmt.Println(id)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
