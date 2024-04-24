package main

import (
	"RESTAPIService2/config"
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Error("error loading env variables: %s", err.Error())
	}

	settings, err := config.NewSettings()
	if err != nil {
		logrus.Fatalf("failed to read settings: %s", err.Error())
	}

	mainCtx := context.Background()
	ctx, cancel := context.WithCancel(mainCtx)
	defer cancel()

	app := NewApp(ctx, settings)
	if err = app.InitDatabase(); err != nil {
		logrus.Errorf(err.Error())
		return
	}

	app.InitService()

	if err = app.Run(); err != nil {
		logrus.Errorf(err.Error())
		return
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	<-sigChan

	if err = app.Shutdown(ctx); err != nil {
		logrus.Errorf(err.Error())
		return
	}

}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
