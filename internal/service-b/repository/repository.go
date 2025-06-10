package repository

import (
	"database/sql"
	"transaction-worker/internal/service-b/model"
)

type TransferRequestRep interface {
	ProcessTransaction(req model.TransferRequest) error
}

type Repository struct {
	TransferRequestRep
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		TransferRequestRep: NewTransferRequestRepository(db),
	}
}
