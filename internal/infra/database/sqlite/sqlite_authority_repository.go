package sqlite

import (
	"context"
	"database/sql"
	"go-auth-service/internal/domain"
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


// =========== READ ===========

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

func (s *SqliteAuthorityRepository) ExistByName(authorityName string) bool {

	return false
}

// =========== WRITE ===========

func (s *SqliteAuthorityRepository) Create(
	ctx context.Context,
	authority entity.Authority,
) (*entity.Authority, error) {
	// TODO: make implementation sqlite create

	return nil, domain.NewAppError(
		domain.ErrInternal,  
		"Not implemented yet",
	)
}

func (s *SqliteAuthorityRepository) Update(
	ctx context.Context,
	authority entity.Authority,
) (*entity.Authority, error) {
	// TODO: make implementation sqlite update

	return nil, domain.NewAppError(
		domain.ErrInternal,  
		"Not implemented yet",
	)
}
func (s *SqliteAuthorityRepository) DeleteByID(
	ctx context.Context,
	id int64,
) error {
	// TODO: make implementation sqlite delete

	return domain.NewAppError(
		domain.ErrInternal,  
		"Not implemented yet",
	)
}

// =========== UTIL ===========

