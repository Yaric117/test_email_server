package main

import (
	"fmt"
	"net/http"
	_ "testproject/db"
	"testproject/routes"
	"time"
)

// @contact.name  API Support
// @contact.url   https://example.com
// @contact.email email@example.com

// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Authorization
func main() {

	srv := http.Server{
		Handler:      routes.Router,
		Addr:         ":8080",
		WriteTimeout: 20 * time.Second,
		ReadTimeout:  50 * time.Second,
	}

	fmt.Printf("[server] %s\n", srv.ListenAndServe())
}
