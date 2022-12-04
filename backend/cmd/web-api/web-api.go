package main

import (
	"os"
	"time"

	"github.com/aaronclaydon/whensmybin/pkg/api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

type ProvidersConfigFile struct {
	Test      string
	Providers []ProviderConfig
}
type ProviderConfig struct {
	ID       string
	Provider string
	Config   map[string]interface{}
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).Level(zerolog.InfoLevel)

	if os.Getenv("WHENSMYBIN_DEBUG") == "YES" {
		log.Logger = log.Logger.Level(zerolog.DebugLevel)
	}

	app := &cli.App{
		Name: "web-api",
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "run web api server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "listen",
						Value: ":8080",
						Usage: "listen target for the web server",
					},
				},
				Action: func(c *cli.Context) error {
					api.SetupServer(c.String("listen"))

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}
