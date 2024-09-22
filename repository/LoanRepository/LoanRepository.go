package LoanRepository

import (
	"database/sql"
	"fmt"
	"kreditplus-api/model/dao"
	"kreditplus-api/model/dto"
)

type loanRepository struct {
	db *sql.DB
	tableName string
}

// GetLimitTenor implements LoanRepository.
func (e *loanRepository) GetLimitTenor(customerID, tenor, year int64) (limit int64, err dto.StandardError) {
	
	if tenor < 1 || tenor > 12 {
		err = dto.StandardError{}.GenerateInvalidRequest()
		return
	}

	param := []interface{}{
		customerID,
        year,
	}

	query := fmt.Sprintf(`
		select 
			tenor_1, tenor_2, tenor_3, tenor_4,
			tenor_5, tenor_6, tenor_7, tenor_8,
            tenor_9, tenor_10, tenor_11, tenor_12 
		from %s 
		where customer_id = $1 and year = $2 limit 1 for update 
	`, "limit_customers")

	var temp dao.LimitCustomer
	errS := e.db.QueryRow(query, param...).Scan(
		&temp.Tenor1, &temp.Tenor2, &temp.Tenor3, &temp.Tenor4,
		&temp.Tenor5, &temp.Tenor6, &temp.Tenor7, &temp.Tenor8,
		&temp.Tenor9, &temp.Tenor10, &temp.Tenor11, &temp.Tenor12,
	)

	if errS == sql.ErrNoRows {
		err = dto.StandardError{}.GenerateNotFound("no found limit in any tenor")
		return
	}

	switch tenor {
		case 1:
			limit = temp.Tenor1.Int64
		case 2:
			limit = temp.Tenor2.Int64
        case 3:
			limit = temp.Tenor3.Int64
		case 4:
			limit = temp.Tenor4.Int64
        case 5:
			limit = temp.Tenor5.Int64
		case 6:
			limit = temp.Tenor6.Int64
        case 7:
			limit = temp.Tenor7.Int64
		case 8:
			limit = temp.Tenor8.Int64
        case 9:
            limit = temp.Tenor9.Int64
        case 10:
			limit = temp.Tenor10.Int64
		case 11:
			limit = temp.Tenor11.Int64
        case 12:
            limit = temp.Tenor12.Int64
	}


	if errS!= nil {
		err = dto.StandardError{}.GenerateInternalServerError(errS)
        return
    }

	return
}

func (e *loanRepository) SaveTransactions(tx *sql.Tx, model dao.Transactions) (err dto.StandardError) {
	param := []interface{}{
		model.CustomerID.Int64,
		model.ContractNumber.String,
		model.OtrPrice.Float64, 
		model.AdminFee.Float64, 
		model.InstallmentAmount.Float64,
		model.RateAmount.Float64,
		model.AssetName.String,
        model.TotalPayment.Float64,
        model.UpdatedAt.Time,
        model.Deleted.Bool,
	}
	query := fmt.Sprintf(`
		insert into %s 
			(
				customer_id, contract_number, otr_price, admin_fee,
				installment_amount, rate_amount, asset_name, total_payment, 
                updated_at, deleted 
			) 
		values
			(
				$1, $2, $3, $4, 
				$5, $6, $7, $8, 
				$9, $10
			) 
		on conflict (contract_number)
		do nothing 
	`, e.tableName)
	_, errS := tx.Exec(query, param...)

	if errS != nil {
		err = err.GenerateInternalServerError(errS)
		return
	}

	err = err.GenerateNoError()
	return
}

func NewLoanRepository(db *sql.DB) LoanRepository {
	return &loanRepository{db: db, tableName: "transactions"}
}
