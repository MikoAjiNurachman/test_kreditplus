package CustomerRepository

import (
	"database/sql"
	"fmt"
	"kreditplus-api/model/dao"
	"kreditplus-api/model/dto"
)

type customerRepository struct {
	db        *sql.DB
	tableName string
}

// DeleteData implements CustomerRepository.
func (e *customerRepository) DeleteData(id int64) (err dto.StandardError) {
	param := []interface{}{
		id,
	}
	query := fmt.Sprintf(`
		update   
			%s 
		set deleted = true, updated_at = NOW()
		where id = $1 
	`, e.tableName)
	_, errS := e.db.Exec(query, param...)

	if errS != nil {
		err = err.GenerateInternalServerError(errS)
		return
	}

	err = err.GenerateNoError()
	return
}

// DetailData implements CustomerRepository.
func (e *customerRepository) DetailData(id int64) (dao.Customer, dto.StandardError) {
	query := fmt.Sprintf(`

        select 
            id, nik, full_name, legal_name, 
            birth_place, birth_date, sallary, 
            image_ktp, image_selfie, 
			created_at, updated_at, deleted  
		    from %s 
        where id = $1 and deleted = false for update 
    `, e.tableName)

	var temp dao.Customer
	errS := e.db.QueryRow(query, id).Scan(
		&temp.ID, &temp.NIK, &temp.FullName, &temp.LegalName,
		&temp.BirthPlace, &temp.BirthDate, &temp.Sallary,
		&temp.ImageKtp, &temp.ImageSelfie,
		&temp.CreatedAt, &temp.UpdatedAt, &temp.Deleted,
	)

	if errS == sql.ErrNoRows {
		return temp, dto.StandardError{}.GenerateNotFound(fmt.Sprintf("data with id %d not found", id))
	}

	if errS!= nil {
        return temp, dto.StandardError{}.GenerateInternalServerError(errS)
    }

	return temp, dto.StandardError{}.GenerateNoError()
}

// ListData implements CustomerRepository.
func (e *customerRepository) ListData() (result []dao.Customer, err dto.StandardError) {
	query := fmt.Sprintf(`
		select 
			nik, full_name, legal_name, 
			birth_place, birth_date, sallary, 
            image_ktp, image_selfie, deleted  
			from %s where deleted = false 
	`, e.tableName)

	rows, errS := e.db.Query(query)
	if errS != nil && errS != sql.ErrNoRows {
		return
	}

	for rows.Next() {
		var temp dao.Customer
		errS = rows.Scan(
			&temp.NIK, &temp.FullName, &temp.LegalName,
			&temp.BirthPlace, &temp.BirthDate, &temp.Sallary,
			&temp.ImageKtp, &temp.ImageSelfie, &temp.Deleted,
		)
		if errS != nil {
			return
		}
		result = append(result, temp)
	}

	err = err.GenerateNoError()
	return
}

// UpdateData implements CustomerRepository.
func (e *customerRepository) UpdateData(tx *sql.Tx, model dao.Customer) (err dto.StandardError) {
	param := []interface{}{
		model.ID.Int64,
		model.NIK.String,
		model.FullName.String,
		model.LegalName.String,
		model.BirthPlace.String,
		model.BirthDate.Time,
		model.Sallary.Int64,
		model.ImageKtp.String,
		model.ImageSelfie.String,
		model.UpdatedAt.Time,
		model.Deleted.Bool,
	}
	query := fmt.Sprintf(`
		update 
			%s
		set  
			nik = $2, full_name = $3, legal_name = $4, 
			birth_place = $5, birth_date = $6, sallary = $7, 
			image_ktp = $8, image_selfie = $9, updated_at = $10, deleted = $11 
		where id = $1 
	`, e.tableName)
	_, errS := tx.Exec(query, param...)

	if errS != nil {
		err = err.GenerateInternalServerError(errS)
		return
	}

	err = err.GenerateNoError()
	return
}

func (e customerRepository) UpsertData(tx *sql.Tx, model dao.Customer) (err dto.StandardError) {
	param := []interface{}{
		model.NIK.String,
		model.FullName.String,
		model.LegalName.String,
		model.BirthPlace.String,
		model.BirthDate.Time,
		model.Sallary.Int64,
		model.ImageKtp.String,
		model.ImageSelfie.String,
		model.Deleted.Bool,
	}
	query := fmt.Sprintf(`
		insert into %s 
			(nik, full_name, legal_name, birth_place, birth_date, sallary, image_ktp, image_selfie, deleted) 
		values
			($1, $2, $3, $4, $5, $6, $7, $8, $9) 
		on conflict (nik)
		do update 
		set 
			full_name = excluded.full_name,
			legal_name = excluded.legal_name,
            birth_place = excluded.birth_place,
            birth_date = excluded.birth_date,
            sallary = excluded.sallary,
            image_ktp = excluded.image_ktp,
			image_selfie = excluded.image_selfie,
			updated_at = now(),
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

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db, tableName: "customers"}

}
