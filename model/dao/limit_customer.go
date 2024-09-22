package dao

import "database/sql"

type LimitCustomer struct {
	ID         sql.NullInt64
	CustomerID sql.NullInt64
	Year       sql.NullInt64
	Tenor1     sql.NullInt64
	Tenor2     sql.NullInt64
	Tenor3     sql.NullInt64
	Tenor4     sql.NullInt64
	Tenor5     sql.NullInt64
	Tenor6     sql.NullInt64
	Tenor7     sql.NullInt64
	Tenor8     sql.NullInt64
	Tenor9     sql.NullInt64
	Tenor10    sql.NullInt64
	Tenor11    sql.NullInt64
	Tenor12    sql.NullInt64
	CreatedAt  sql.NullTime
	UpdatedAt  sql.NullTime
	Deleted    sql.NullBool
}

type ListLimitCustomer struct {
	CustomerNIK sql.NullString
	CustomerName sql.NullString
	ImageSelfie sql.NullString
	Year       sql.NullInt64
	Tenor1     sql.NullInt64
	Tenor2     sql.NullInt64
	Tenor3     sql.NullInt64
	Tenor4     sql.NullInt64
	Tenor5     sql.NullInt64
	Tenor6     sql.NullInt64
	Tenor7     sql.NullInt64
	Tenor8     sql.NullInt64
	Tenor9     sql.NullInt64
	Tenor10    sql.NullInt64
	Tenor11    sql.NullInt64
	Tenor12    sql.NullInt64
}