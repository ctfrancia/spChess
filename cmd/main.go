package main

import (
	"flag"
	"os"

	"github.com/ctfrancia/spChess/cmd/routes"
)

type config struct {
	port string
	env  string
}

type application struct {
	cfg    config
	logger *slog.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment(development|staging|production)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		// v1.Use(AuthMiddleware())
		routes.UserRoutes(v1)
	}

	router.Run()
}
