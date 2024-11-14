package main

import (
	"log"

	"github.com/Dormant512/url-compressor/backend/config"
	"github.com/Dormant512/url-compressor/backend/internal/handler"
	"github.com/Dormant512/url-compressor/backend/internal/repository"
	"github.com/Dormant512/url-compressor/backend/internal/service"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := repository.New(repository.Config{
		RedisDSN:           cfg.GetRedisDSN(),
		ClickHouseDSN:      cfg.GetClickHouseDSN(),
		ClickHouseDBName:   cfg.CHDBName,
		ClickHouseUsername: cfg.CHUser,
		ClickHousePassword: cfg.CHPassword,
	})
	if err != nil {
		log.Fatal(err)
	}

	svc := service.New(repo)
	router := handler.RegisterHandlers(svc)
	router.Run(cfg.GetAppAddress())
}
