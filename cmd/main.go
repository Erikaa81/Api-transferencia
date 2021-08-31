package main

import (
	"github.com/Erikaa81/Banco-api/app"
	"github.com/Erikaa81/Banco-api/controllers/exit"
	"github.com/Erikaa81/Banco-api/controllers/logger"
	"github.com/Erikaa81/Banco-api/controllers/server"
	"github.com/Erikaa81/Banco-api/models"
	"github.com/Erikaa81/Banco-api/routes"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var api *app.App

func initenv() error {
	// capturando variáveis de ambiente
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal("Falha ao carregar: ", viper.ConfigFileUsed())
	}
	return err
}

func initapp() error {
	logrus.Info("Usando arquivo config: ", viper.ConfigFileUsed())
	// armazenando configurações em um struct app
	var err error
	api, err = app.GetApp()
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return err
}

func initdb() error {
	// migrando os schemas do DB
	err := api.DB.Client.AutoMigrate(&models.Accounts{}, &models.Transfer{})
	if err != nil {
		logrus.Fatal("Erro na migração dos dados")
	}
	return err
}

func init() {
	if initenv() == nil {
		if initapp() == nil {
			if initdb() == nil {
				if api.Cfg.GetDebugMode() == "true" {
					logrus.Warn("Banking rodando em modo Debug")
				} else {
					logrus.Warn("Banking rodando em modo Silent")
				}
			}
		}
	}
}

func main() {

	defer api.DB.CloseDB()

	srv := server.
		GetServer().
		WithAddr(api.Cfg.GetAPIPort()).
		WithRouter(routes.GetRouter(api)).
		WithLogger(logger.Error)

	go func() {
		api.Log.Info("Iniciando servidor na porta ", api.Cfg.GetAPIPort())
		if err := srv.StartServer(); err != nil {
			api.Log.Fatal(err.Error())
		}
	}()

	exit.Init(func() {
		if err := srv.CloseServer(); err != nil {
			api.Log.Error(err.Error())
		}

		if err := api.DB.CloseDB(); err != nil {
			api.Log.Error(err.Error())
		}
	})
}
