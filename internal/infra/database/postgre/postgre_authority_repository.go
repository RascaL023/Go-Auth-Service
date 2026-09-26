package postgre

import (
	"context"
	entity "go-auth-service/internal/domain/authority"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgreAuthorityRepository struct {
	pool	*pgxpool.Pool
}

func New(pool *pgxpool.Pool) *PostgreAuthorityRepository {
	return &PostgreAuthorityRepository{
		pool: pool,
	}
}


func (p *PostgreAuthorityRepository) FindByID(
	ctx context.Context, 
	id int64,
) (*entity.Authority, error) {
	// TODO: make implementation

	return nil, nil
}

func (p *PostgreAuthorityRepository) FindByIDs(
	ctx context.Context, 
	ids []int64,
) ([]*entity.Authority, error) {
	// TODO: make implementation

	return nil, nil
}

