package sqlite3

import "time"

func Datetime() string {
	return time.Now().Format(time.RFC3339)
}
