package postgres

import (
	"context"
	"github.com/zhayt/transaction-service/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TransactionStorage struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewTransactionStorage(db *gorm.DB, log *zap.Logger) *TransactionStorage {
	return &TransactionStorage{db: db, log: log}
}

func (r *TransactionStorage) CreateTransaction(ctx context.Context, transaction model.Transaction) (uint, error) {
	tx := r.db.Begin()
	defer func() {
		if re := recover(); re != nil {
			r.log.Error("transaction error")
			tx.Rollback()
		}
	}()

	// отнимаем сумму покупки от счета пользователя
	if err := tx.WithContext(ctx).Table("user_accounts").Where("id = ?",
		transaction.AccountID).Updates(map[string]interface{}{"balance": gorm.Expr("balance - ?",
		transaction.Amount)}); err != nil {
		r.log.Error("Update user balance error", zap.Error(err.Error))
		tx.Rollback()
	}

	// создаем транзакцию
	if err := tx.Create(&transaction).Error; err != nil {
		r.log.Error("Create transaction error", zap.Error(err))
		tx.Rollback()
	}

	for _, item := range transaction.Items {
		item.TransactionID = transaction.ID
		if err := tx.Create(&item).Error; err != nil {
			r.log.Error("Create transaction item error", zap.Error(err))
			tx.Rollback()
		}
	}

	if err := tx.Commit().Error; err != nil {
		r.log.Error("Commit error", zap.Error(err))
		tx.Rollback()
	}

	r.log.Info("transaction has been created", zap.Uint("id", transaction.ID))
	return transaction.ID, nil
}
