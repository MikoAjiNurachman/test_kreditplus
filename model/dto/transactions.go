package dto

type Transactions struct {
	CustomerID        int64   `json:"customer_id"`
	CreditAmount      float64 `json:"credit_amount"`
	ContractNumber    string  `json:"contract_number"`
	Year              int     `json:"year"`
	Tenor             int     `json:"tenor"`
	Otr               float64 `json:"otr"`
	AdminFee          float64 `json:"admin_fee"`
	InstallmentAmount float64 `json:"installment_amount"`
	Rate              int     `json:"annual_rate"` // bunga tahunan
	AssetName         string  `json:"asset_name"`
}

func (e *Transactions) ValidateRequest() (err StandardError) {

	if e.CustomerID == 0 {
		err = StandardError{}.GenerateEmptyField("customer_id")
		return
	}

	if e.CreditAmount == 0 {
		err = StandardError{}.GenerateEmptyField("credit_amount")
		return
	}

	if e.Year == 0 {
		err = StandardError{}.GenerateEmptyField("year")
		return
	}

	if e.Tenor == 0 {
		err = StandardError{}.GenerateEmptyField("tenor")
		return
	}

	if e.ContractNumber == "" {
		err = StandardError{}.GenerateEmptyField("contract_number")
		return
	}

	err = StandardError{}.GenerateNoError()
	return
}
