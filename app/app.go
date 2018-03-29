package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shenglol/golang/config"
)

type App struct {
	config *config.Config
	router *mux.Router
}

func NewApp(config *config.Config) *App {
	app := &App{config, mux.NewRouter()}

	dbConn := config.DBConnection
	db := NewDB(dbConn.Host, dbConn.Port, dbConn.Database)
	app.configRouter(&Handler{db})

	return app
}

func (a *App) Run() {
	addr := a.config.Host + ":" + a.config.Port
	log.Fatal(http.ListenAndServe(addr, a.router))
}

func (a *App) configRouter(handler *Handler) {
	a.router.HandleFunc("/todos", handler.GetAllTodos).Methods("GET")
}
