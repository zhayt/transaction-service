package main

import (
	"fmt"
	"github.com/zhayt/transaction-service/config"
	"github.com/zhayt/transaction-service/logger"
	"github.com/zhayt/transaction-service/service"
	"github.com/zhayt/transaction-service/storage"
	"github.com/zhayt/transaction-service/transport/http"
	"github.com/zhayt/transaction-service/transport/http/handler"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// config
	var once sync.Once
	once.Do(config.PrepareEnv)

	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	// logger
	l, err := logger.Init(cfg)
	if err != nil {
		return fmt.Errorf("couldn't init logger: %w", err)
	}

	defer func(l *zap.Logger) {
		err = l.Sync()
		if err != nil {
			log.Fatal(err)
		}
	}(l)

	// storage
	repo, err := storage.NewStorage(l, cfg)
	if err != nil {
		return err
	}

	// service
	serv := service.NewManager(repo, l)

	// handler
	hand := handler.NewHandler(serv, l)

	// server

	httpSrv := http.NewServer(hand, cfg)

	l.Info("Start server")
	httpSrv.StartHTTPServer()

	// grace full shutdown
	osSignCh := make(chan os.Signal, 1)
	signal.Notify(osSignCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-osSignCh:
		l.Info("signal accepted: ", zap.String("signal", s.String()))
	case err = <-httpSrv.Notify:
		l.Info("server closing", zap.Error(err))
	}

	if err = httpSrv.GracefullyShutdown(); err != nil {
		return fmt.Errorf("error while shutting down server: %s", err)
	}

	return nil
}
