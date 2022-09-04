package main

import (
	"math/rand"
	"time"

	"github.com/VrMolodyakov/url-shortener/internal"
	"github.com/VrMolodyakov/url-shortener/internal/config"
	"github.com/VrMolodyakov/url-shortener/pkg/logging"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cfg := config.GetConfig()
	logger := logging.GetLogger(cfg.Loglvl)
	app := internal.NewApp(logger, cfg)
	logger.Info("start App")
	app.Run()
}
