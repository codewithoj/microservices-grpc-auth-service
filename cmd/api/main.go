package main

import (
	"fmt"
	"log"
	"net/http"
)

// 1. routes.go file - check
// 2. handler.go file - check
// 3. main.go file - implement server

const webPort = "8081"

type Config struct{}

func main() {
	log.Println("starting authentication service")
	app := &Config{}

	//define our server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	log.Println("starting authentication service on port: ", webPort)
	// start and listen on specificport
	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}
