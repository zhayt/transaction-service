package storage

import (
	"context"
	"github.com/zhayt/transaction-service/config"
	"github.com/zhayt/transaction-service/model"
	"github.com/zhayt/transaction-service/storage/postgres"
	"go.uber.org/zap"
)

type IUserAccountStorage interface {
	CreateUserAccount(ctx context.Context, account model.UserAccount) (uint, error)
	GetUserAccountByID(ctx context.Context, accountID int) (model.UserAccount, error)
	GetUserAccountByCardNumber(ctx context.Context, cardNumber string) (model.UserAccount, error)
	UpdateUserAccountBalance(ctx context.Context, balance model.NewUserAccountBalance) (uint, error)
	DeleteUserAccount(ctx context.Context, accountID int) error
}

type ITransactionStorage interface {
	CreateTransaction(ctx context.Context, transaction model.Transaction) (uint, error)
}

type Storage struct {
	IUserAccountStorage
	ITransactionStorage
}

func NewStorage(logger *zap.Logger, cfg *config.Config) (*Storage, error) {
	db, err := postgres.Dial(cfg.DataBaseURL)
	if err != nil {
		return nil, err
	}

	return &Storage{
		IUserAccountStorage: postgres.NewUserAccountStorage(db, logger),
		ITransactionStorage: postgres.NewTransactionStorage(db, logger),
	}, nil
}
