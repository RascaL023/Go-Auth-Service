package config

import (
	"errors"
	"fmt"
	"strings"
)

type Config struct {
	Server		Server
	Database	Database
}

type Server struct {
	Port		int
}

type Database struct {
	Driver		string
	Postgres 	Postgres
	SQLite 		SQLite
}

type Postgres struct {
	Host     	string
	Port     	int
	Database 	string
	User     	string
	Password 	string
	SSLMode  	string
}

type SQLite struct {
	Path 		string
}


func (c Config) Validate() error {
	switch strings.ToLower(c.Database.Driver) {
	case "memory":
		return nil

	case "sqlite":
		if c.Database.SQLite.Path == "" {
			return errors.New("sqlite path is required")
		}

	case "postgres":
		if c.Database.Postgres.Host == "" {
			return errors.New("postgres host is required")
		}
		if c.Database.Postgres.Port < 1 || c.Database.Postgres.Port > 65535 {
			return errors.New("invalid postgres port")
		}
		if c.Database.Postgres.User == "" {
			return errors.New("postgres user is required")
		}
		if c.Database.Postgres.Database == "" {
			return errors.New("postgres database name is required")
		}

	default:
		return fmt.Errorf(
			"unsupported database driver: %q",
			c.Database.Driver,
		)
	}

	return nil
}
