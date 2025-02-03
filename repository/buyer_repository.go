package repository

import (
	"context"
	"database/sql"
	"errors"
	"oneiot-server/model/entity"
)

type IBuyerRepository interface {
	Create(ctx context.Context, buyerDetails entity.Buyer) (entity.Buyer, error)
	GetById(ctx context.Context, buyer entity.Buyer) (entity.Buyer, error)

	//Kita tidak dapat menghapus buyer karena terikat dengan orders sehingga tidak perlu untuk menghapusnya
	//Kita juga tidak dapat mengubah buyer ketika sudah dibuat
}

type BuyerRepository struct {
	db        *sql.DB
	tableName string
}

func (repository *BuyerRepository) GetById(ctx context.Context, buyerDetails entity.Buyer) (entity.Buyer, error) {
	query := "SELECT * FROM Buyers WHERE Id = ?"

	rows, err := repository.db.QueryContext(ctx, query, buyerDetails.Id)

	if err != nil {
		return entity.Buyer{}, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	if !rows.Next() {
		return entity.Buyer{}, errors.New("Tidak ditemukan buyer dengan id " + string(rune(buyerDetails.Id)))
	}

	err = rows.Scan(&buyerDetails.Id, &buyerDetails.FullName, &buyerDetails.Email, &buyerDetails.PhoneNumber, &buyerDetails.FullAddress, &buyerDetails.AdditionalNotes)

	if err != nil {
		return entity.Buyer{}, errors.New("Error saat scanning buyer details")
	}

	return buyerDetails, nil
}

func (repository *BuyerRepository) Create(ctx context.Context, buyerDetails entity.Buyer) (entity.Buyer, error) {
	query := "INSERT INTO Buyers(FullName, Email, PhoneNumber, FullAddress, AdditionalNotes) " +
		"VALUES(?, ?, ?, ?, ?)"

	execContext, err := repository.db.ExecContext(ctx, query, buyerDetails.FullName, buyerDetails.Email, buyerDetails.PhoneNumber, buyerDetails.FullAddress, buyerDetails.AdditionalNotes)

	if err != nil {
		return entity.Buyer{}, errors.New("Error saat memasukkan ke tabel buyer")
	}

	buyerDetails.Id, err = execContext.LastInsertId()

	if err != nil {
		return entity.Buyer{}, errors.New("Error saat mendapatkan id di tabel buyer")
	}

	return buyerDetails, nil
}

func NewBuyerRepository(db *sql.DB) IBuyerRepository {
	return &BuyerRepository{
		db:        db,
		tableName: "Buyers",
	}
}
