package config

type DatabaseConfig struct {
	DriverName, DatabaseURI string
}

func NewPostgresConfig(dbDriver, dbURI string) *DatabaseConfig {
	return &DatabaseConfig{dbDriver, dbURI}
}
