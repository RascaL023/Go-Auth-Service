package memory

import (
	"context"
	"fmt"
	"go-auth-service/internal/domain"
	entity "go-auth-service/internal/domain/authority"
)

type MemoryAuthorityRepository struct {
	DB 		map[int64]*entity.Authority
	LastID	int64
}

func New(db map[int64]*entity.Authority) *MemoryAuthorityRepository {
	return &MemoryAuthorityRepository{
		DB: db,
		LastID: 1,
	}
}

// =========== READ ===========

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

	copy := *data
	return &copy, nil
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

		copy := *data
		datas = append(datas, &copy)
	}

	return datas, nil
}

func (m *MemoryAuthorityRepository) ExistByName(authorityName string) bool {
	for _, value  := range m.DB {
		if authorityName == value.Name {
			return true
		}
	}
	
	return false
}

// =========== WRITE ===========

func (m *MemoryAuthorityRepository) Create(
	ctx context.Context,
	authority entity.Authority,
) (*entity.Authority, error) {
	result := m.upsert(authority)
	m.LastID++
	return &result, nil
}

func (m *MemoryAuthorityRepository) Update(
	ctx context.Context,
	authority entity.Authority,
) (*entity.Authority, error) {
	result := m.upsert(authority)
	return &result, nil
}

// =========== UTIL ===========

func (m *MemoryAuthorityRepository) upsert(data entity.Authority) entity.Authority {
	var id int64

	if data.ID != 0 {
		id = data.ID
	} else {
		id = m.LastID
	}

	m.DB[id] = &data
	return *m.DB[id]
}
