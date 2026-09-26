package config

func Load() (Config, error) {
	appPort, err := envInt("APP_PORT", 8080)
	if err != nil {
		return Config{}, err
	}

	dbPort, err := envInt("DB_POSTGRES_PORT", 5432)
	if err != nil {
		return Config{}, err
	}
	
	cfg := Config{
		Server: Server{
			Port: appPort,
		},

		Database: Database{
			Driver: env("DB_DRIVER", ""),
			Postgres: Postgres{
				Host: env("DB_POSTGRES_HOST", "localhost"),
				Port: dbPort,
				Database: env("DB_POSTGRES_DATABASE", ""),
				User: env("DB_POSTGRES_USER", ""),
				Password: env("DB_POSTGRES_PASSWORD", ""),
				SSLMode: env("DB_POSTGRES_SSLMODE", "disable"),
			},
			SQLite: SQLite{
				Path: env("DB_SQLITE_PATH", ""),
			},
		},
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
