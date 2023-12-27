package main

import (
	srv "Interface_droch_3"
	_ "Interface_droch_3/docs"
	"Interface_droch_3/internal/handler"
	"Interface_droch_3/internal/repository"
	"Interface_droch_3/internal/repository/postgres"
	"Interface_droch_3/internal/service"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

// @title REST_API_ZAK
// @version 0.0.1
// @description REST API Training Program

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading environment variable: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("Postgres creation error: %s", err.Error())
	}

	//rdb, err := redis_storage.NewRedisClient(redis_storage.Config{
	//	Addr: os.Getenv("RDB_ADDRESS"),
	//})
	//
	//if err != nil {
	//	log.Fatalf("Redis creation error: %s", err.Error())
	//}

	//repo := repository.NewStorageUsersRedis(rdb)
	repo := repository.NewStorageUsersPostgres(db)
	services := service.NewServiceUsers(repo)
	handlers := handler.NewHandler(services)

	port := os.Getenv("SRV_PORT")

	localPrint := fmt.Sprintf("the server is running on: http://localhost:%s/", port)
	fmt.Println(localPrint)

	localSwag := fmt.Sprintf("swagger: http://localhost:%s/docs/index.html#/", port)
	fmt.Println(localSwag)

	serv := new(srv.Server)
	if err = serv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("the server could not be started: %s", err.Error())
	}

}

//func initConfig() error {
//	viper.AddConfigPath("configs")
//	viper.SetConfigName("config")
//	return viper.ReadInConfig()
//}
