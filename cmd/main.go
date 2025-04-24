package main

import (
	"fmt"
	"net/http"

	userApi "github.com/M-kos/crm-user/internal/api/go"
	"github.com/M-kos/crm-user/internal/config"
	"github.com/M-kos/crm-user/internal/logger"
)

func main() {
	c := config.New()
	l := logger.NewLogger()

	AuthAPIService := userApi.NewAuthAPIService()
	AuthAPIController := userApi.NewAuthAPIController(AuthAPIService)

	UserAPIService := userApi.NewUserAPIService()
	UserAPIController := userApi.NewUserAPIController(UserAPIService)

	router := userApi.NewRouter(AuthAPIController, UserAPIController)

	l.Info("Server started", "port", c.Port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", c.Port), router)
	if err != nil {
		l.Error(err.Error())
	}
}
