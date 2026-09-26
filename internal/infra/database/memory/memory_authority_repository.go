package memory

import (
	"context"
	"fmt"
	"go-auth-service/internal/domain"
	entity "go-auth-service/internal/domain/authority"
)

type MemoryAuthorityRepository struct {
	DB map[int64]*entity.Authority
}

func New(db map[int64]*entity.Authority) *MemoryAuthorityRepository {
	return &MemoryAuthorityRepository{
		DB: db,
	}
}

func (m *MemoryAuthorityRepository) FindByID(
	ctx context.Context,
	id int64,
) (*entity.Authority, error) {
	data, ok := m.DB[id]
	if !ok {
		return nil, domain.NewAppError(
			domain.ErrNotFound,
			fmt.Sprintf("Cannot found authority with ID: %d", id),
		)
	}

	return data, nil
}

func (m *MemoryAuthorityRepository) FindByIDs(
	ctx context.Context,
	ids []int64,
) ([]*entity.Authority, error) {
	var datas []*entity.Authority
	for _, id := range ids {
		data, ok := m.DB[id]
		if !ok {
			return nil, domain.NewAppError(
				domain.ErrNotFound,
				fmt.Sprintf("Cannot found authority with ID: %d", id),
			)
		}

		datas = append(datas, data)
	}

	return datas, nil
}
