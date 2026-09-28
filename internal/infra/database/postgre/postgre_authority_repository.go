package postgre

import (
	"context"
	"go-auth-service/internal/domain"
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


// =========== READ ===========

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

func (p *PostgreAuthorityRepository) ExistByName(authorityName string) bool {

	return false
}

// =========== WRITE ===========

func (p *PostgreAuthorityRepository) Create(
	ctx context.Context,
	authority entity.Authority,
) (*entity.Authority, error) {
	// TODO: make implementation postgre create

	return nil, domain.NewAppError(
		domain.ErrInternal,  
		"Not implemented yet",
	)
}

func (p *PostgreAuthorityRepository) Update(
	ctx context.Context,
	authority entity.Authority,
) (*entity.Authority, error) {
	// TODO: make implementation postgre update

	return nil, domain.NewAppError(
		domain.ErrInternal,  
		"Not implemented yet",
	)
}
func (p *PostgreAuthorityRepository) DeleteByID(
	ctx context.Context,
	id int64,
) error {
	// TODO: make implementation postgre delete

	return domain.NewAppError(
		domain.ErrInternal,  
		"Not implemented yet",
	)
}

// =========== UTIL ===========

