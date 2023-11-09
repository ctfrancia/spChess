package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/ctfrancia/spChess/cmd/routes"

	"github.com/gin-gonic/gin"
)

type config struct {
	port int
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

	/*
		app := &application{
			cfg:    cfg,
			logger: logger,
		}
	*/

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger()) // ??
	v1 := router.Group("/api/v1")
	{
		// v1.Use(AuthMiddleware())
		routes.UserRoutes(v1)
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("Starting %s server on %s", cfg.env, srv.Addr)

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)

	/*
		router := gin.Default()
		v1 := router.Group("/api/v1")
		{
			// v1.Use(AuthMiddleware())
			routes.UserRoutes(v1)
		}

		router.Run()
	*/
}
