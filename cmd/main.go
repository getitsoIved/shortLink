package main

import (
	"fmt"
	"net/http"

	"github.com/getitsoIved/shortLink/configs"
	"github.com/getitsoIved/shortLink/internal/auth"
	"github.com/getitsoIved/shortLink/internal/link"
	"github.com/getitsoIved/shortLink/internal/user"
	"github.com/getitsoIved/shortLink/pkg/db"
	"github.com/getitsoIved/shortLink/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositories
	linkRerository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)

	//Services
	authService := auth.NewAuthService(userRepository)

	//Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRerository,
	})

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Сервер слушает порт 8081")
	server.ListenAndServe()

}
