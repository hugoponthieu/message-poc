package cli

import (
	"context"
	"message/app"
	"message/config"
	"message/seeder"
	"os"

	"github.com/urfave/cli/v3"
)

func GetConfig() (*config.AppConfig, error) {
	config := &config.AppConfig{}

	cmd := &cli.Command{
		Name:  "message-service",
		Usage: "A message service application",
		Commands: []*cli.Command{
			{
				Name:  "seed",
				Usage: "Seed the message service",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "mseed",
						Value: 15,
						Usage: "the number of messages to seed",
					},
					&cli.IntFlag{
						Name:  "batch",
						Value: 15,
						Usage: "the number of messages to seed",
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					mseed := cmd.Int("mseed")
					batch := cmd.Int("batch")

					app_seeder, err := app.InitApp(*config)
					if err != nil {
						return err
					}
					seeder := seeder.NewSeeder(*app_seeder)

					seeder.Seed(int(mseed), int(batch))
					// Use mseed value here
					return nil
				},
			},
			{
				Name:  "start",
				Usage: "Start the message service",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					server, err := app.InitApp(*config)
					if err != nil {
						return err
					}
					err = server.Start()
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			// App flags
			&cli.StringFlag{
				Name:        "host",
				Sources:     cli.EnvVars("APP_HOST"),
				Value:       "0.0.0.0",
				Destination: &config.Host,
				Usage:       "host to listen on",
			},
			&cli.StringFlag{
				Name:        "port",
				Sources:     cli.EnvVars("APP_PORT"),
				Value:       "8080",
				Destination: &config.Port,
				Usage:       "port to listen on",
			},
			// MongoDB flags
			&cli.StringFlag{
				Name:        "mongo-host",
				Sources:     cli.EnvVars("MONGO_HOST"),
				Value:       "localhost:27017",
				Destination: &config.MongoConfig.Host,
				Usage:       "MongoDB host",
			},
			&cli.StringFlag{
				Name:        "mongo-username",
				Sources:     cli.EnvVars("MONGO_USERNAME"),
				Value:       "mongo",
				Destination: &config.MongoConfig.Username,
				Usage:       "MongoDB username",
			},
			&cli.StringFlag{
				Name:        "mongo-password",
				Sources:     cli.EnvVars("MONGO_PASSWORD"),
				Value:       "admin",
				Destination: &config.MongoConfig.Password,
				Usage:       "MongoDB password",
			},
			&cli.StringFlag{
				Name:        "mongo-database",
				Sources:     cli.EnvVars("MONGO_DATABASE"),
				Value:       "message_db",
				Destination: &config.MongoConfig.Database,
				Usage:       "MongoDB database name",
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		return nil, err
	}

	return config, nil
}
