package http

import (
	"context"
	"net/http"

	"github.com/oschwald/geoip2-golang"
	"golang.org/x/sync/errgroup"
)

func Handler(geoipCityDB *geoip2.Reader, geoipISPDB *geoip2.Reader) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /bmi", bmiHandler())
	mux.HandleFunc("/echo", echoHandler())
	mux.HandleFunc("GET /geoip/{ip}", geoipHandler(geoipCityDB, geoipISPDB))
	mux.HandleFunc("GET /health", healthHandler())
	return mux
}

func ListenAndServe(ctx context.Context, handler http.Handler, addr string) error {
	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	var g errgroup.Group
	g.Go(func() error {
		return srv.ListenAndServe()
	})
	g.Go(func() error {
		<-ctx.Done()
		return srv.Shutdown(context.Background())
	})
	err := g.Wait()
	if err == http.ErrServerClosed {
		return nil
	}
	return err
}
