package main

import (
	db "kafka-go-getting-started/database"
	"kafka-go-getting-started/database/migration"
	"kafka-go-getting-started/internal/factory"
	"kafka-go-getting-started/internal/http"
	"kafka-go-getting-started/internal/middleware"
	// "kafka-go-getting-started/pkg/elasticsearch"
	"kafka-go-getting-started/pkg/util/env"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func init() {
	ENV := os.Getenv("ENV")
	env := env.NewEnv()
	env.Load(ENV)

	logrus.Info("Choosen environment " + ENV)
}

// @title kafka-go-getting-started
// @version 0.0.1
// @description This is a doc for kafka-go-getting-started.

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host localhost:3030
// @BasePath /
func main() {
	var PORT = os.Getenv("PORT")
	db.Init()
	migration.Init()
	// elasticsearch.Init()

	e := echo.New()
	middleware.Init(e)

	f := factory.NewFactory()
	http.Init(e, f)

	e.Logger.Fatal(e.Start(":" + PORT))
}
