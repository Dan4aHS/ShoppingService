package main

import (
	"ShoppingExpensesService/internal/app"
	"ShoppingExpensesService/internal/config"
	"ShoppingExpensesService/internal/repo/postgres"
	"ShoppingExpensesService/internal/service"
	"ShoppingExpensesService/pkg/logging"
	"log"
)

func main() {
	cfg := config.GetConfig()
	logging.InitLogging()
	dbx, err := postgres.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}
	repo := postgres.NewRepository(dbx)
	serv := service.NewService(repo)
	application := app.NewApp(cfg, serv)
	application.Run()
}
