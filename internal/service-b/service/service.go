package service

import (
	"transaction-worker/internal/service-b/model"
	"transaction-worker/internal/service-b/repository"
)

type TransferRequest interface {
	ProcessTransactionServi(req model.TransferRequest) error
}

type Service struct {
	TransferRequest
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TransferRequest: NewTransferRequestService(repos.TransferRequestRep),
	}
}
