package dao

import "database/sql"

type Transactions struct {
	ID                sql.NullInt64
	CustomerID        sql.NullInt64
	ContractNumber    sql.NullString
	OtrPrice          sql.NullFloat64
	AdminFee          sql.NullFloat64
	InstallmentAmount sql.NullFloat64
	RateAmount        sql.NullFloat64
	AssetName         sql.NullString
	TotalPayment      sql.NullFloat64
	CreatedAt         sql.NullTime
	UpdatedAt         sql.NullTime
	Deleted           sql.NullBool
}
