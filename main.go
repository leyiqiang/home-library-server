package main

import (
	. "github.com/leyiqiang/home-library-server/config"
	"github.com/leyiqiang/home-library-server/controllers"
	"github.com/leyiqiang/home-library-server/repositories"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

var config Config

func init() {
	config.Read()
}

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	repo := repositories.NewMongoRepo()
	controller := controllers.Controller{Repo: repo}
	r := Routers(&controller)

	srv := &http.Server{
		Addr:         config.Server.Port,
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port", config.Server.Port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
