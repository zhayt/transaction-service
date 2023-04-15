package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/zhayt/transaction-service/model"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) CreateTransaction(e echo.Context) error {
	ctx, cancel := context.WithTimeout(e.Request().Context(), _timeoutContext)
	defer cancel()

	var transaction model.Transaction

	if err := e.Bind(&transaction); err != nil {
		h.log.Error("Bind error", zap.Error(err))
		return e.JSON(http.StatusBadRequest, err)
	}

	transactionID, err := h.transaction.CreateTransaction(ctx, transaction)
	if err != nil {
		h.log.Error("Create transaction error", zap.Error(err))
		return e.JSON(http.StatusInternalServerError, err)
	}

	h.log.Info("Transaction has been created", zap.Uint("id", transactionID))
	return e.JSON(http.StatusOK, transactionID)
}
