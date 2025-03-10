package cli

import (
	"context"
	"message/config"
	"os"

	"github.com/urfave/cli/v3"
)

func GetConfig() (*config.AppConfig, error) {
	config := &config.AppConfig{}

	cmd := &cli.Command{
		Name:  "message-service",
		Usage: "A message service application",
		Flags: []cli.Flag{
			// App flags
			&cli.StringFlag{
				Name: "host",
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
