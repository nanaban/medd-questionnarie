package main

import (
	"context"
	"flag"

	"questionnarie/pkg/handler"
	"questionnarie/pkg/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

var (
	flagAddr string
	flagConn string
)

func main() {
	flag.StringVar(&flagAddr, "addr", ":2222", "server listen address")
	flag.StringVar(&flagConn, "conn", "postgresql://postgres:postgrespass@localhost:5433/test4?sslmode=disable", "DB connection URL")
	flag.Parse()

	logger := logrus.New()

	repo, err := repository.New(flagConn)
	if err != nil {
		logger.Fatalf("create repository err: %s", err)
	}

	e := echo.New()
	e.Use(middleware.Recover()) // recover middleware
	v1 := e.Group("/v1")

	ctx := context.TODO()
	h := handler.NewHandler(ctx, repo)
	h.Register(v1)

	logger.Infof("server running at %s", flagAddr)
	if err := e.Start(flagAddr); err != nil {
		logger.Fatal(err)
	}
}
