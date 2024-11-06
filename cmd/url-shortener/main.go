package main

import (
	"log/slog"
	"urlshortener/internal/config"
	"urlshortener/pkg/logger"
)

func main() {
	cfg := config.MustLoadConf()

	log := logger.SetupLogger(cfg.Env)
	log.Info("info logs", slog.String("env", cfg.Env))
	log.Debug("debug logs")
}
