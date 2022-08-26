package main

import (
	"kafka-go-getting-started/config"
	"kafka-go-getting-started/consumer"
	"kafka-go-getting-started/controller"
	db "kafka-go-getting-started/database"
	"kafka-go-getting-started/database/migration"
	"kafka-go-getting-started/internal/factory"
	"kafka-go-getting-started/internal/http"
	"kafka-go-getting-started/internal/middleware"
	"kafka-go-getting-started/producer"
	"kafka-go-getting-started/routes"

	// "kafka-go-getting-started/pkg/elasticsearch"
	"kafka-go-getting-started/pkg/util/env"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)
var userController *controller.UserController

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

	config.CORSConfig(e)

	routes.GetUserApiRoutes(e, userController)

	e.Logger.Fatal(e.Start(":" + PORT))
}

func init() {
	p := config.InitKafkaProducer()
	producer := producer.NewProducer(p)
	userController = controller.NewUserController(producer)
	c := config.InitKafkaConsumer(config.UserConsumerGroup)
	consumer := consumer.NewConsumer(c)
	go consumer.Consume([]string{config.UserNotificationTopic})
}
