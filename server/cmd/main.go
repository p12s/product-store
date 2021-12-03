package main

import (
	"context"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	"github.com/p12s/product-store/server/internal/config"
	"github.com/p12s/product-store/server/internal/repository"
	"github.com/p12s/product-store/server/internal/server"
	"github.com/p12s/product-store/server/internal/service"
	"github.com/p12s/product-store/server/internal/transport/grpc"
	"github.com/sirupsen/logrus"
)

const MONGODB_TIMEOUT = 10 * time.Second

// @title Store app gRPC-client
// @version 0.0.1
// @description Simple application for loading/getting products
func main() {
	runtime.GOMAXPROCS(1)
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error reading env variables from file: %s\n", err.Error())
	}
	cfg, err := config.New()
	if err != nil {
		logrus.Fatalf("error loading env variables: %s\n", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), MONGODB_TIMEOUT)
	defer cancel()
	db, err := repository.NewMongoDB(ctx, repository.Config{
		Username: cfg.Db.User,
		Password: cfg.Db.Password,
		Url:      cfg.Db.Uri,
	})
	if err != nil {
		logrus.Fatalf("failed to initialize mongodb: %s\n", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := grpc.NewHandler(services, cfg.File.SaveDir)
	srv := server.New(handlers)

	logrus.Printf("Server started on port %d. %s", cfg.Server.Port, time.Now().Format(time.RFC3339))

	srv.ListenAndServe(cfg.Server.Port)
}
