package dao

import "database/sql"

type User struct {
	ID   sql.NullInt64
	Name sql.NullString
	Age  sql.NullInt32
}
