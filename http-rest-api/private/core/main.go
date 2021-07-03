package main

import (
	"log"

	"github.com/Survereignty/spet-rest-api/private/core/api"
	"github.com/Survereignty/spet-rest-api/private/core/api/db"
	"github.com/Survereignty/spet-rest-api/private/core/lig"
)

func main() {

	// Инициализация логов
	if err := lig.Create("./logs/logs.txt"); err != nil {
		log.Fatal(err)
	}
	lig.Info("Logs is init")

	defer lig.Quit()

	// Настройки по умолчанию
	defaultConfig := &api.Config{
		Address: ":3001",
		DataBase: db.NewConfig(
			"host=localhost password=student dbname=spet sslmode=disable"),
	}

	// Инициализация настроек
	c, err := api.NewConfig(defaultConfig)
	if err != nil {
		lig.Сrash("Config file is not created", err)
	}
	lig.Info("Config file is init")

	// Инициализация сервера
	s := api.New(c)
	if err := s.Start(); err != nil {
		lig.Сrash("Server is not started", err)
	}
}
