package main

import (
	"fmt"
	"log/slog"
	"net/http"

	userApi "github.com/M-kos/crm-user/internal/api/go"
	"github.com/M-kos/crm-user/internal/config"
	"github.com/M-kos/crm-user/internal/logger"
	"github.com/M-kos/crm-user/internal/storage"
)

func main() {
	config := config.New()
	log := logger.NewLogger()
	_ = storage.NewStorage(config)

	AuthAPIService := userApi.NewAuthAPIService()
	AuthAPIController := userApi.NewAuthAPIController(AuthAPIService)

	UserAPIService := userApi.NewUserAPIService()
	UserAPIController := userApi.NewUserAPIController(UserAPIService)

	router := userApi.NewRouter(AuthAPIController, UserAPIController)

	log.Info("Server started", slog.Int("port", config.Port))

	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router)
	if err != nil {
		log.Error(err.Error())
	}
}
