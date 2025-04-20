package main

import (
	"fmt"
	"log"
	"net/http"

	userApi "github.com/M-kos/crm-user/internal/api/go"
	"github.com/M-kos/crm-user/internal/config"
)

func main() {
	log.Printf("Server started")
	c := config.New()

	fmt.Println("Config: ", c)

	AuthAPIService := userApi.NewAuthAPIService()
	AuthAPIController := userApi.NewAuthAPIController(AuthAPIService)

	UserAPIService := userApi.NewUserAPIService()
	UserAPIController := userApi.NewUserAPIController(UserAPIService)

	router := userApi.NewRouter(AuthAPIController, UserAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
