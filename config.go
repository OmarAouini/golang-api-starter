package main

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var CONFIG *Configuration

type Configuration struct {
	AppEnv      string `env:"APP_ENV"`
	RunningMode string `env:"RUNNING_MODE"`
	ApiSecret   string `env:"API_SECRET"`
}

func ConfigEnv() {
	err := godotenv.Load()
	if err == nil {
		logrus.Info(".env file found, going development mode")
	}
	if err != nil {
		logrus.Info(".env file not found, going production mode...")
	}
	configuration := Configuration{}
	if err := env.Parse(&configuration); err != nil {
		fmt.Printf("%+v\n", err)
	}
	logrus.Info("env vars loaded")
	CONFIG = &configuration
}

func ConfigLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
