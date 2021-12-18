package main

import (
	"context"
	"database/sql"
	"fmt"
	. "github.com/leyiqiang/home-library-server/config"
	"github.com/leyiqiang/home-library-server/controllers"
	"github.com/leyiqiang/home-library-server/repositories"
	"github.com/leyiqiang/home-library-server/services"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

var config Config

func init() {
	config = ReadConfig()
}

func main() {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(config)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	repo := repositories.NewPostgresRepo(db)
	service := services.NewService(repo)
	controller := controllers.Controller{Service: service}
	r := Routers(&controller)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port", config.Port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

func openDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		return nil, err
	}

	return db, nil
}
