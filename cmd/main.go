package main

import (
	"fmt"
	"net/http"

	"github.com/getitsoIved/shortLink/configs"
	"github.com/getitsoIved/shortLink/internal/auth"
	"github.com/getitsoIved/shortLink/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
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
