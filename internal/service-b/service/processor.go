package service

import (
	"transaction-worker/internal/service-b/model"
	"transaction-worker/internal/service-b/repository"
)

type TransferRequestService struct {
	repo repository.TransferRequestRep
}

func NewTransferRequestService(repo repository.TransferRequestRep) *TransferRequestService {
	return &TransferRequestService{repo: repo}
}

func (p *TransferRequestService) ProcessTransactionServi(req model.TransferRequest) error {
	// err := p.repo.

	return nil
}

// func (s *TransferRequestService) isDuplicateRequest(userID int, requestID string) (bool, error) {

// 	return true, nil
// }

// func (s *TransferRequestService) getUserBalance(userID int) (int, error) {

// 	return 0, nil
// }
