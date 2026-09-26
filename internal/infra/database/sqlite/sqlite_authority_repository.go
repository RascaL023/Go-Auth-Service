package sqlite

import (
	"context"
	"database/sql"
	entity "go-auth-service/internal/domain/authority"
)

type SqliteAuthorityRepository struct {
	DB		*sql.DB
}

func New(db *sql.DB) *SqliteAuthorityRepository {
	return &SqliteAuthorityRepository{
		DB: db,
	}
}


func (s *SqliteAuthorityRepository) FindByID(
	ctx context.Context, 
	id int64,
) (*entity.Authority, error) {
	// TODO: make implementation

	return nil, nil
}

func (s *SqliteAuthorityRepository) FindByIDs(
	ctx context.Context, 
	ids []int64,
) ([]*entity.Authority, error) {
	// TODO: make implementation

	return nil, nil
}
