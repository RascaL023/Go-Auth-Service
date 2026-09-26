package database

import (
	"context"
	"fmt"
	"go-auth-service/internal/config"
	"go-auth-service/internal/domain/authority"
	"go-auth-service/internal/infra/database/memory"
	"go-auth-service/internal/infra/database/postgre"
	"go-auth-service/internal/infra/database/sqlite"
	"net"
	"net/url"
	"strconv"
	"strings"
)

func ResolveAuthorityRepository(ctx context.Context, cfg config.Config) (authority.Repository, error) {
	switch strings.ToLower(cfg.Database.Driver){
	case "memory", "mockup":
		db := memory.NewConnection()
		return memory.New(db.DB), nil

	case "postgre", "postgres":
		url := url.URL{
			Scheme: cfg.Database.Driver,
			Host: 	net.JoinHostPort(cfg.Database.Postgres.Host, strconv.Itoa(cfg.Database.Postgres.Port)),
			Path: "/" + cfg.Database.Postgres.Database,
		}
		pool, err := postgre.NewConnection(ctx, url.String())
		if err != nil {
			return nil, err
		}

		return  postgre.New(pool), nil
	
	case "sqlite":
		db, err := sqlite.NewConnection(ctx, cfg.Database.SQLite.Path)
		if err != nil {
			return nil, err
		}
		
		return sqlite.New(db), nil

	default: 
		return  nil, fmt.Errorf("unknown database driver: %s", cfg.Database.Driver)
	}
}
