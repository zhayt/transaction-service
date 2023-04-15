package service

import (
	"context"
	"github.com/zhayt/transaction-service/model"
	"github.com/zhayt/transaction-service/storage"
	"go.uber.org/zap"
)

type IUserAccountService interface {
	CreateUserAccount(ctx context.Context, account model.UserAccount) (uint, error)
	GetUserAccountByID(ctx context.Context, accountID int) (model.UserAccount, error)
	GetUserAccountByCardNumber(ctx context.Context, cardNumber string) (model.UserAccount, error)
	UpdateUserAccountBalance(ctx context.Context, balance model.NewUserAccountBalance) (uint, error)
	DeleteUserAccount(ctx context.Context, accountID int) error
}

type ITransactionService interface {
	CreateTransaction(ctx context.Context, transaction model.Transaction) (uint, error)
}

type Manager struct {
	User        IUserAccountService
	Transaction ITransactionService
}

func NewManager(storage *storage.Storage, log *zap.Logger) *Manager {
	return &Manager{
		User:        NewUserAccountService(storage, log),
		Transaction: NewTransactionService(storage, log),
	}
}
