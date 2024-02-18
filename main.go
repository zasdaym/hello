package main

import (
	"context"
	"os/signal"
	"syscall"

	_ "embed"

	"github.com/oschwald/geoip2-golang"
	"github.com/rs/zerolog/log"
	"github.com/zasdaym/zmono/internal/config"
	"github.com/zasdaym/zmono/internal/http"
)

//go:embed GeoIP2-City.mmdb
var geoipCityDBBytes []byte

//go:embed GeoIP2-ISP.mmdb
var geoipISPDBBytes []byte

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := run(ctx); err != nil {
		log.Fatal().Err(err).Send()
	}
}

func run(ctx context.Context) error {
	cfg := config.Load()

	geoipCityDB, err := geoip2.FromBytes(geoipCityDBBytes)
	if err != nil {
		return err
	}
	defer geoipCityDB.Close()

	geoipISPDB, err := geoip2.FromBytes(geoipISPDBBytes)
	if err != nil {
		return err
	}
	defer geoipISPDB.Close()

	handler := http.Handler(geoipCityDB, geoipISPDB)
	return http.ListenAndServe(ctx, handler, cfg.Addr)
}
