package api

import (
	"net/http"

	"github.com/Survereignty/spet-rest-api/private/core/api/db"
	"github.com/Survereignty/spet-rest-api/private/core/lig"
	"github.com/gorilla/mux"
)

// Объект с конфигурацией API сервера
type Api struct {
	config *Config
	router *mux.Router
	store  *db.Store
}

// Создание экземпляра API
func New(config *Config) *Api {
	return &Api{
		config: config,
		router: mux.NewRouter(),
	}
}

// Запуск API
func (s *Api) Start() error {
	if err := s.Store(); err != nil {
		return err
	}

	s.Routers()

	lig.Listen("Server listen port", s.config.Address)

	return http.ListenAndServe(s.config.Address, s.router)
}

func (s *Api) Store() error {
	store := db.New(s.config.DataBase)

	if err := store.Open(); err != nil {
		return err
	}

	s.store = store

	lig.Info("Database is connected")

	return nil
}
