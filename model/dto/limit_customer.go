package dto

import "time"

type LimitCustomer struct {
	ID         int64     `json:"id"`
	CustomerID int64     `json:"customer_id"`
	Year       int64     `json:"year"`
	Tenor1     int64     `json:"tenor_1"`
	Tenor2     int64     `json:"tenor_2"`
	Tenor3     int64     `json:"tenor_3"`
	Tenor4     int64     `json:"tenor_4"`
	Tenor5     int64     `json:"tenor_5"`
	Tenor6     int64     `json:"tenor_6"`
	Tenor7     int64     `json:"tenor_7"`
	Tenor8     int64     `json:"tenor_8"`
	Tenor9     int64     `json:"tenor_9"`
	Tenor10    int64     `json:"tenor_10"`
	Tenor11    int64     `json:"tenor_11"`
	Tenor12    int64     `json:"tenor_12"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Deleted    bool      `json:"deleted"`
}

type ListLimitCustomer struct {
	CustomerNIK  string `json:"customer_nik"`
	CustomerName string `json:"customer_name"`
	ImageSelfie  string `json:"image_selfie"`
	Year         int64  `json:"year"`
	Tenor1       int64  `json:"tenor_1"`
	Tenor2       int64  `json:"tenor_2"`
	Tenor3       int64  `json:"tenor_3"`
	Tenor4       int64  `json:"tenor_4"`
	Tenor5       int64  `json:"tenor_5"`
	Tenor6       int64  `json:"tenor_6"`
	Tenor7       int64  `json:"tenor_7"`
	Tenor8       int64  `json:"tenor_8"`
	Tenor9       int64  `json:"tenor_9"`
	Tenor10      int64  `json:"tenor_10"`
	Tenor11      int64  `json:"tenor_11"`
	Tenor12      int64  `json:"tenor_12"`
}

func (e *LimitCustomer) ValidateUpsert() (err StandardError) {

	if e.CustomerID < 1 {
		err = StandardError{}.GenerateEmptyField("customer_id")
		return
	}

	err = StandardError{}.GenerateNoError()
	return
}
