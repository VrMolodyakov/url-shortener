package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/VrMolodyakov/url-shortener/internal/adapters/db/urlDb"
	"github.com/VrMolodyakov/url-shortener/internal/config"
	"github.com/VrMolodyakov/url-shortener/internal/domain/service"
	"github.com/VrMolodyakov/url-shortener/internal/handler"
	"github.com/VrMolodyakov/url-shortener/pkg/client/redis"
	"github.com/VrMolodyakov/url-shortener/pkg/logging"
	"github.com/VrMolodyakov/url-shortener/pkg/shutdown"
	"github.com/gorilla/mux"
)

const (
	writeTimeout = 15 * time.Second
	readTimeout  = 15 * time.Second
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
	a.logger.Debug("start init handler")
	rdCfg := redis.GetRdConfig(a.cfg.Redis.Password, a.cfg.Redis.Host, a.cfg.Redis.Port, a.cfg.Redis.DbNumber)
	rdClient, err := redis.NewClient(context.Background(), &rdCfg)
	a.checkErr(err)
	repo := urlDb.NewUrlRepository(a.logger, rdClient)
	shortener := service.NewShortener(a.logger)
	urlService := service.NewUrlService(repo, shortener)
	a.router = mux.NewRouter()
	handler := handler.NewUrlHandler(a.logger, urlService)
	handler.InitRoutes(a.router)
	a.logger.Info("start listening...")
	port := fmt.Sprintf(":%s", a.cfg.Port)
	server := &http.Server{
		Addr:         port,
		Handler:      a.router,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}
	a.checkErr(err)
	go shutdown.Graceful([]os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM}, rdClient, server)
	if err := server.ListenAndServe(); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			a.logger.Warn("server shutdown")
		default:
			a.logger.Fatal(err)
		}
	}
	a.logger.Info("app shutdown")
}

func (a *app) checkErr(err error) {
	if err != nil {
		a.logger.Fatal(err)
	}
}
