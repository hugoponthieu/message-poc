package config

type AppConfig struct {
	Host     string
	Port     string
	MongoConfig MongoConfig
	AllowOrigin string
	OidcBaseUrl string 
	Realm string
	ClientID string
}
