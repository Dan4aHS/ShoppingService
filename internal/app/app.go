package app

import (
	"ShoppingExpensesService/internal/app/grpcapp"
	"ShoppingExpensesService/internal/config"
	"ShoppingExpensesService/internal/service"
	"log/slog"
)

type App struct {
	GRPCApp *grpcapp.App
}

func NewApp(cfg *config.Config, ps service.IProductService) *App {
	grpcApp := grpcapp.NewApp(ps, cfg)
	slog.Info("App initialized")
	return &App{
		GRPCApp: grpcApp,
	}
}

func (a *App) Run() {
	a.GRPCApp.Run()
}

func (a *App) Stop() {
	a.GRPCApp.Stop()
}
