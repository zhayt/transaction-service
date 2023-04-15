package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/zhayt/transaction-service/model"
	"github.com/zhayt/transaction-service/storage"
	"go.uber.org/zap"
)

type TransactionService struct {
	transaction *storage.Storage
	log         *zap.Logger
}

func NewTransactionService(transaction *storage.Storage, log *zap.Logger) *TransactionService {
	return &TransactionService{transaction: transaction, log: log}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, transaction model.Transaction) (uint, error) {
	s.log.Info("Start create transaction", zap.String("card-number", transaction.CardNumber))
	userAccount, err := s.transaction.GetUserAccountByCardNumber(ctx, transaction.CardNumber)
	if err != nil {
		return 0, fmt.Errorf("couldn't get user account by card number: %w", err)
	}

	if userAccount.Balance < transaction.Amount {
		return 0, fmt.Errorf("not enough money in balance: %w", errors.New("client error"))
	}

	transaction.AccountID = userAccount.ID

	return s.transaction.CreateTransaction(ctx, transaction)
}
