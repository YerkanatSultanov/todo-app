package main

import (
	"github.com/YerkanatSultanov/todo-app"
	"github.com/YerkanatSultanov/todo-app/pkg/handler"
	"github.com/YerkanatSultanov/todo-app/pkg/repository"
	"github.com/YerkanatSultanov/todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfigs(); err != nil {
		log.Fatalf("error in config files: %s", err)
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}

func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("configs")

	return viper.ReadInConfig()
}
