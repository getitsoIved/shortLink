package main

import (
	"fmt"
	"net/http"

	"github.com/getitsoIved/shortLink/configs"
	"github.com/getitsoIved/shortLink/internal/auth"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Сервер слушает порт 8081")
	server.ListenAndServe()

}
