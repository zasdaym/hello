package config

import (
	"flag"
)

type Config struct {
	Addr  string
	Debug bool
}

func Load() Config {
	var result Config
	flag.StringVar(&result.Addr, "addr", ":8080", "Listen address")
	flag.BoolVar(&result.Debug, "debug", false, "Enable debug mode")
	flag.Parse()
	return result
}
