package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/materkov/shortener/redirecter"
)

type Options struct {
	Addr               string `long:"addr" env:"ADDR" default:":8000" description:"HTTP server addr"`
	RedisURL           string `long:"redis-url" env:"REDIS_URL" default:"127.0.0.1:6379" description:"Redis URL"`
	PostgresConnString string `long:"postgres-conn" env:"POSTGRES_CONN" default:"user=postgres dbname=shortener sslmode=disable" description:"Postgres connection string"`
}

func main() {
	opts := Options{}
	flags.Parse(&opts)

	s := redirecter.NewStorePersistent(opts.PostgresConnString)
	pubsub := redirecter.NewPubsub(opts.RedisURL)
	app := redirecter.NewApp(s, pubsub)
	api := redirecter.NewAPI(app, opts.Addr)

	api.Serve()
}
