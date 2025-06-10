package repository

import (
	"database/sql"
	"transaction-worker/internal/service-b/model"
)

type TransferRequestRepository interface {
	ProcessTransaction(model.TransferRequest) error
}

func NewTransferRequestRepository(db *sql.DB) TransferRequestRepository {
	return &Storage{
		db: db,
	}
}

func (s *Storage) ProcessTransaction(req model.TransferRequest) error {

	return nil
}
