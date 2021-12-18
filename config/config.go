package config

import (
	"flag"
	"fmt"
)

type Config struct {
	Port int
	Env  string
	DB   struct {
		DSN string
	}
}

func ReadConfig() Config {
	var cfg Config
	flag.IntVar(&cfg.Port, "Port", 3000, "Server Port to listen on")
	flag.StringVar(&cfg.Env, "Env", "development", "application environment(development | production)")
	flag.StringVar(&cfg.DB.DSN, "DSN", "postgresql://root:root@my_db/go_movies?sslmode=disable", "Postgres connection string")
	flag.Parse()
	fmt.Println("Running")
	return cfg
}
