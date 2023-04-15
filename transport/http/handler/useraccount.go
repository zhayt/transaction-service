package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/zhayt/transaction-service/model"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func (h *Handler) CreateUserAccount(e echo.Context) error {
	ctx, cancel := context.WithTimeout(e.Request().Context(), _timeoutContext)
	defer cancel()

	var userAccount model.UserAccount

	if err := e.Bind(&userAccount); err != nil {
		h.log.Error("Bind error", zap.Error(err))
		return e.JSON(http.StatusBadRequest, err)
	}

	accountID, err := h.user.CreateUserAccount(ctx, userAccount)
	if err != nil {
		h.log.Error("Create user account error", zap.Error(err))
		return e.JSON(http.StatusInternalServerError, err)
	}

	h.log.Info("User account has been created", zap.Uint("id", accountID))
	return e.JSON(http.StatusOK, accountID)
}

func (h *Handler) UpdateUserAccount(e echo.Context) error {
	ctx, cancel := context.WithTimeout(e.Request().Context(), _timeoutContext)
	defer cancel()

	var replenishmentAmount model.NewUserAccountBalance

	if err := e.Bind(&replenishmentAmount); err != nil {
		h.log.Error("Bind error", zap.Error(err))
		return e.JSON(http.StatusBadRequest, err)
	}

	userAccount, err := h.user.GetUserAccountByCardNumber(ctx, replenishmentAmount.CardNumber)
	if err != nil {
		h.log.Error("Get user account by card number error", zap.Error(err))
		return e.JSON(http.StatusBadRequest, err)
	}

	h.log.Info("User found with card number", zap.String("name", userAccount.Name), zap.Float64("balance", userAccount.Balance))

	replenishmentAmount.AccountID = userAccount.ID

	userAccount.ID, err = h.user.UpdateUserAccountBalance(ctx, replenishmentAmount)
	if err != nil {
		h.log.Error("Update user account balance error", zap.Error(err))
		return e.JSON(http.StatusInternalServerError, err)
	}

	userAccount.Balance += replenishmentAmount.ReplenishmentAmount
	h.log.Info("User balance has been updated", zap.Uint("account_id", userAccount.ID), zap.Float64("balance", userAccount.Balance))

	return e.JSON(http.StatusOK, userAccount)
}

func (h *Handler) ShowUserAccountInfo(e echo.Context) error {
	ctx, cancel := context.WithTimeout(e.Request().Context(), _timeoutContext)
	defer cancel()

	accountId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		h.log.Error("Param error", zap.Error(err))
		return e.JSON(http.StatusNotFound, err)
	}

	userAccount, err := h.user.GetUserAccountByID(ctx, accountId)
	if err != nil {
		h.log.Error("Get user bu account id error", zap.Error(err))
		return e.JSON(http.StatusNotFound, err)
	}

	h.log.Info("User account has been found", zap.Uint("accountId", userAccount.ID))
	return e.JSON(http.StatusOK, userAccount)
}

func (h *Handler) DeleteUserAccount(e echo.Context) error {
	ctx, cancel := context.WithTimeout(e.Request().Context(), _timeoutContext)
	defer cancel()

	userAccountID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		h.log.Error("Param error", zap.Error(err))
		return e.JSON(http.StatusNotFound, err)
	}

	if err = h.user.DeleteUserAccount(ctx, userAccountID); err != nil {
		h.log.Error("Delete user account error", zap.Error(err))
		return e.JSON(http.StatusInternalServerError, err)
	}

	h.log.Info("User account has been deleted", zap.Int("id", userAccountID))
	return e.JSON(http.StatusOK, userAccountID)
}
