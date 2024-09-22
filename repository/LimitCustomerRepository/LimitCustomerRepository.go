package LimitCustomerRepository

import (
	"database/sql"
	"fmt"
	"kreditplus-api/model/dao"
	"kreditplus-api/model/dto"
)

type limitCustomerRepository struct {
	db        *sql.DB
	tableName string
}

// ListLimitCustomer implements LimitCustomerRepository.
func (e *limitCustomerRepository) ListLimitCustomer() (result []dao.ListLimitCustomer, err dto.StandardError) {
	query := fmt.Sprintf(`
		select 
			c.nik, c.full_name, c.image_selfie, lc.year, 
			lc.tenor_1, lc.tenor_2, lc.tenor_3, lc.tenor_4, 
			lc.tenor_5, lc.tenor_6, lc.tenor_7, lc.tenor_8,
			lc.tenor_9, lc.tenor_10, lc.tenor_11, lc.tenor_12 
			from %s lc 
			left join %s c 
				on c.id = lc.customer_id 
			where lc.deleted = false and c.deleted = false 
	`, e.tableName, "customers")

	rows, errS := e.db.Query(query)
	if errS != nil && errS != sql.ErrNoRows {
		return
	}

	for rows.Next() {
		var temp dao.ListLimitCustomer
		errS = rows.Scan(
			&temp.CustomerNIK, &temp.CustomerName, &temp.ImageSelfie, 
			&temp.Year,
            &temp.Tenor1, &temp.Tenor2, &temp.Tenor3, &temp.Tenor4,
            &temp.Tenor5, &temp.Tenor6, &temp.Tenor7, &temp.Tenor8,
            &temp.Tenor9, &temp.Tenor10, &temp.Tenor11, &temp.Tenor12,
		)
		if errS != nil {
			return
		}
		result = append(result, temp)
	}

	err = err.GenerateNoError()
	return
}

// UpsertData implements LimitCustomerRepository.
func (e *limitCustomerRepository) UpsertData(tx *sql.Tx, model dao.LimitCustomer) (err dto.StandardError) {
	param := []interface{}{
		model.CustomerID.Int64,
		model.Year.Int64,
		model.Tenor1.Int64,
		model.Tenor2.Int64,
		model.Tenor3.Int64,
		model.Tenor4.Int64,
		model.Tenor5.Int64,
		model.Tenor6.Int64,
		model.Tenor7.Int64,
		model.Tenor8.Int64,
		model.Tenor9.Int64,
		model.Tenor10.Int64,
		model.Tenor11.Int64,
		model.Tenor12.Int64,
		model.UpdatedAt.Time,
		model.Deleted.Bool,
	}
	query := fmt.Sprintf(`
		insert into %s 
			(
				customer_id, year, tenor_1, tenor_2, 
				tenor_3, tenor_4, tenor_5, tenor_6, 
				tenor_7, tenor_8, tenor_9, tenor_10, 
				tenor_11, tenor_12, updated_at, deleted
			) 
		values
			(
				$1, $2, $3, $4, 
				$5, $6, $7, $8, 
				$9, $10, $11, $12, 
				$13, $14, $15, $16 
			) 
		on conflict (customer_id, year)
		do update 
		set 
			tenor_1 = excluded.tenor_1,
			tenor_2 = excluded.tenor_2,
            tenor_3 = excluded.tenor_3,
            tenor_4 = excluded.tenor_4,
            tenor_5 = excluded.tenor_5,
            tenor_6 = excluded.tenor_6,
			tenor_7 = excluded.tenor_7,
            tenor_8 = excluded.tenor_8,
            tenor_9 = excluded.tenor_9,
            tenor_10 = excluded.tenor_10,
            tenor_11 = excluded.tenor_11,
			tenor_12 = excluded.tenor_12,
            updated_at = NOW(),  
			deleted = excluded.deleted
	`, e.tableName)
	_, errS := tx.Exec(query, param...)

	if errS != nil {
		err = err.GenerateInternalServerError(errS)
		return
	}

	err = err.GenerateNoError()
	return
}

func NewLimitCustomerRepository(db *sql.DB) LimitCustomerRepository {
	return &limitCustomerRepository{db: db, tableName: "limit_customers"}
}
