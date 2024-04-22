package main

import (
	"ShoppingExpensesService/internal/app"
	"ShoppingExpensesService/internal/config"
	"ShoppingExpensesService/internal/repo/postgres"
	"ShoppingExpensesService/internal/service"
	"log"
)

func main() {
	cfg := config.GetConfig()
	dbx, err := postgres.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}
	repo := postgres.NewRepository(dbx)
	serv := service.NewService(repo)
	application := app.NewApp(cfg, serv)
	application.Run()
}
