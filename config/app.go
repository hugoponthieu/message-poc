package config

type AppConfig struct {
	Host     string
	Port     string
	MongoConfig MongoConfig
	AllowOrigin string
}
