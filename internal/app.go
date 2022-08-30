package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/VrMolodyakov/url-shortener/internal/adapters/db/urlDb"
	"github.com/VrMolodyakov/url-shortener/internal/config"
	"github.com/VrMolodyakov/url-shortener/internal/domain/service"
	"github.com/VrMolodyakov/url-shortener/internal/handler"
	"github.com/VrMolodyakov/url-shortener/pkg/client/redis"
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
	a.logger.Info("start listening...")
	port := fmt.Sprintf(":%v", a.cfg.Port)
	log.Fatal(http.ListenAndServe(port, a.router))
}

func (a *app) initialize() {
	a.logger.Debug("start init handler")
	rdCfg := redis.GetRdConfig(a.cfg.Redis.Password, a.cfg.Redis.Host, a.cfg.Redis.Port, a.cfg.Redis.DbNumber)
	rdClient, err := redis.NewClient(context.Background(), &rdCfg)
	a.checkErr(err)
	repo := urlDb.NewUrlRepository(a.logger, rdClient)
	shortener := service.NewShortener(a.logger)
	urlService := service.NewUrlService(repo, shortener)
	a.router = mux.NewRouter()
	a.initializeRouters(urlService)
}

func (a *app) initializeRouters(service handler.UrlService) {
	h := handler.NewUrlHandler(a.logger, service)
	a.router.HandleFunc("/encode", h.EncodeUrl).Methods(http.MethodPost)
	a.router.HandleFunc("/{shortUrl}", h.DecodeUrl).Methods(http.MethodGet)
	a.router.HandleFunc("/encode/custom", h.EncodeCustomUrl).Methods(http.MethodPost)
}

func (a *app) checkErr(err error) {
	if err != nil {
		a.logger.Fatal(err)
	}
}
