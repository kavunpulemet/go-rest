package main

import (
	"RESTAPIService2/config"
	"RESTAPIService2/pkg/api"
	"RESTAPIService2/pkg/api/middlewares"
	"RESTAPIService2/pkg/repository"
	"RESTAPIService2/pkg/service/auth"
	"RESTAPIService2/pkg/service/item"
	"RESTAPIService2/pkg/service/list"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type App struct {
	ctx      context.Context
	server   *api.Server
	settings config.Settings
	postgres *sqlx.DB
}

func NewApp(ctx context.Context, settings config.Settings) *App {
	return &App{
		ctx:      ctx,
		settings: settings,
	}
}

func (a *App) InitDatabase() error {
	var err error
	a.postgres, err = sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		a.settings.Host, a.settings.Port, a.settings.Username, a.settings.DBName, a.settings.Password, a.settings.SSLMode))
	if err != nil {
		return err
	}

	err = a.postgres.Ping()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) InitService() {
	authService := auth.NewAuthorizationService(repository.NewAuthorizationPostgres(a.postgres))
	userIdentityMiddleware := middlewares.NewUserIdentityMiddleware(authService)
	todoLists := list.NewTodoListService(repository.NewTodoListPostgres(a.postgres))
	todoItems := item.NewTodoItemService(repository.NewTodoItemPostgres(a.postgres), todoLists)
	a.server = api.NewServer(userIdentityMiddleware)
	a.server.HandleAuth(authService)
	a.server.HandleLists(todoLists)
	a.server.HandleItems(todoItems)
}

func (a *App) Run() error {
	go func() {
		if err := a.server.Run(); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Println("run server")
	return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	err := a.server.Shutdown(ctx)
	if err != nil {
		logrus.Errorf("Не удалось отключить от сервера %w", err)
		return err
	}

	err = a.postgres.Close()
	if err != nil {
		logrus.Errorf("Не удалось отключиться от базы %w", err)
	}
	return nil
}
