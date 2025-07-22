package main

import (
	"fmt"
	"net/http"

	"github.com/getitsoIved/shortLink/configs"
	"github.com/getitsoIved/shortLink/internal/auth"
	"github.com/getitsoIved/shortLink/internal/link"
	"github.com/getitsoIved/shortLink/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositories
	linkRerository := link.NewLinkRepository(db)

	//Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRerository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Сервер слушает порт 8081")
	server.ListenAndServe()

}
