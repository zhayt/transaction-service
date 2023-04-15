package postgres

import (
	"context"
	"github.com/zhayt/transaction-service/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserAccountStorage struct {
	db  *gorm.DB
	log *zap.Logger
}

func (r *UserAccountStorage) CreateUserAccount(ctx context.Context, account model.UserAccount) (uint, error) {
	result := r.db.WithContext(ctx).Omit("created_at").Create(&account)

	return account.ID, result.Error
}

func (r *UserAccountStorage) GetUserAccountByID(ctx context.Context, accountID int) (model.UserAccount, error) {
	var userAccount model.UserAccount

	if err := r.db.WithContext(ctx).Where("id = ?", accountID).First(&userAccount).Error; err != nil {
		return model.UserAccount{}, err
	}

	return userAccount, nil
}

func (r *UserAccountStorage) GetUserAccountByCardNumber(ctx context.Context, cardNumber string) (model.UserAccount, error) {
	var userAccount model.UserAccount

	if err := r.db.WithContext(ctx).Where("card_number = ?", cardNumber).First(&userAccount).Error; err != nil {
		return model.UserAccount{}, err
	}

	return userAccount, nil
}

func (r *UserAccountStorage) UpdateUserAccountBalance(ctx context.Context, balance model.NewUserAccountBalance) (uint, error) {
	result := r.db.WithContext(ctx).Table("user_accounts").Where("id = ?", balance.AccountID).Updates(map[string]interface{}{"balance": gorm.Expr("balance + ?", balance.ReplenishmentAmount)})

	return balance.AccountID, result.Error
}

func (r *UserAccountStorage) DeleteUserAccount(ctx context.Context, accountID int) error {
	return r.db.WithContext(ctx).Delete(&model.UserAccount{}, accountID).Error
}

func NewUserAccountStorage(db *gorm.DB, log *zap.Logger) *UserAccountStorage {
	return &UserAccountStorage{db: db, log: log}
}
