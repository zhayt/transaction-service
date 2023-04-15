package service

import (
	"context"
	"github.com/zhayt/transaction-service/model"
	"github.com/zhayt/transaction-service/storage"
	"go.uber.org/zap"
)

type UserAccountService struct {
	user storage.IUserAccountStorage
	log  *zap.Logger
}

func (s *UserAccountService) CreateUserAccount(ctx context.Context, account model.UserAccount) (uint, error) {
	// валидация данных

	return s.user.CreateUserAccount(ctx, account)
}

func (s *UserAccountService) GetUserAccountByID(ctx context.Context, accountID int) (model.UserAccount, error) {
	return s.user.GetUserAccountByID(ctx, accountID)
}

func (s *UserAccountService) GetUserAccountByCardNumber(ctx context.Context, cardNumber string) (model.UserAccount, error) {
	return s.user.GetUserAccountByCardNumber(ctx, cardNumber)
}

func (s *UserAccountService) UpdateUserAccountBalance(ctx context.Context, balance model.NewUserAccountBalance) (uint, error) {
	return s.user.UpdateUserAccountBalance(ctx, balance)
}

func (s *UserAccountService) DeleteUserAccount(ctx context.Context, accountID int) error {
	return s.user.DeleteUserAccount(ctx, accountID)
}

func NewUserAccountService(user storage.IUserAccountStorage, log *zap.Logger) *UserAccountService {
	return &UserAccountService{user: user, log: log}
}
