package main

import (
	srv "Interface_droch_3"
	"fmt"

	_ "Interface_droch_3/docs"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"

	"Interface_droch_3/internal/handler"
	"Interface_droch_3/internal/repository"
	"Interface_droch_3/internal/repository/postgres"
	"Interface_droch_3/internal/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	"log"
	"os"
)

// @title REST_API_ZAK
// @version 0.0.1
// @description Программа для обучения REST API

// @host localhost:8080
// @BasePath /
func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Ошибка инициализации конфига: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка загрузки переменной окружения: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: os.Getenv("DB_USER"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("Ошибка создания Postgres: %s", err.Error())
	}

	//rdb, err := redis_storage.NewRedisClient(redis_storage.Config{
	//	Addr: viper.GetString("rdb.address"),
	//})
	//
	//if err != nil {
	//	log.Fatalf("Ошибка создания Redis: %s", err.Error())
	//}

	//repo := repository.NewStorageUsersRedis(rdb)
	repo := repository.NewStorageUsersPostgres(db)
	services := service.NewServiceUsers(repo)
	handlers := handler.NewHandler(services)

	port := viper.GetString("srv_port")

	local := "Сервер запущен на: http://localhost:"
	fmt.Printf("%s%s\n", local, port)

	localSwag := fmt.Sprintf("Swagger: http://localhost:%s/api/docs/index.html", port)
	fmt.Println(localSwag)

	serv := new(srv.Server)
	if err = serv.Run(viper.GetString("srv_port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Не удалось запустить сервер: %s", err.Error())
	}

}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
