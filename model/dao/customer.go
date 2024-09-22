package dao

import "database/sql"

type Customer struct {
	ID          sql.NullInt64
	NIK         sql.NullString
	FullName    sql.NullString
	LegalName   sql.NullString
	BirthPlace  sql.NullString
	BirthDate   sql.NullTime
	Sallary     sql.NullInt64
	ImageKtp    sql.NullString
	ImageSelfie sql.NullString
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	Deleted     sql.NullBool
}
