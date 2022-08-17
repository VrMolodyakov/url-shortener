package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VrMolodyakov/url-shortener/internal/config"
	"github.com/VrMolodyakov/url-shortener/internal/handler"
	"github.com/VrMolodyakov/url-shortener/pkg/logging"
	"github.com/gorilla/mux"
)

type app struct {
	logger *logging.Logger
	cfg    *config.Config
	router *mux.Router
}

func NewApp(logger *logging.Logger, cfg *config.Config) *app {
	return &app{cfg: cfg, logger: logger}
}

func (a *app) Run() {
	a.startHttp()
}

func (a *app) startHttp() {
	a.logger.Info("start http server")
	a.initialize()
	a.initializeRouters()
	a.logger.Info("start listening...")
	port := fmt.Sprintf(":%v", a.cfg.Port)
	log.Fatal(http.ListenAndServe(port, a.router))
}

func (a *app) initialize() {
	//TODO: init
	a.router = mux.NewRouter()
	a.initializeRouters()
}

func (a *app) initializeRouters() {
	a.router.HandleFunc("/encode", handler.NewUrlHandler(a.logger).Encode)
}
