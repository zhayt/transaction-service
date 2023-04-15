package handler

import (
	"github.com/zhayt/transaction-service/service"
	"go.uber.org/zap"
	"time"
)

const _timeoutContext = 5 * time.Second

type Handler struct {
	user service.IUserAccountService
	log  *zap.Logger
}

func NewHandler(service *service.Manager, log *zap.Logger) *Handler {
	return &Handler{user: service.User, log: log}
}
