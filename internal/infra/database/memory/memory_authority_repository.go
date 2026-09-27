package memory

import (
	"context"
	"fmt"
	"go-auth-service/internal/domain"
	entity "go-auth-service/internal/domain/authority"
	"sync"
)

type MemoryAuthorityRepository struct {
	db 			map[int64]*entity.Authority
	mtx			sync.RWMutex
	nameIndex	map[string]*entity.Authority
	lastID		int64
}

func New(db map[int64]*entity.Authority) (*MemoryAuthorityRepository, error) {
	repo := &MemoryAuthorityRepository{
		db: db,
		nameIndex: make(map[string]*entity.Authority),
		lastID: 1,
	}

	err := repo.build()
	return repo, err
}

// =========== READ ===========

func (m *MemoryAuthorityRepository) FindByID(
	ctx context.Context,
	id int64,
) (*entity.Authority, error) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	data, ok := m.db[id]
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
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	for _, id := range ids {
		data, ok := m.db[id]
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
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	_, exist := m.nameIndex[authorityName]
	return exist
}

// =========== WRITE ===========

func (m *MemoryAuthorityRepository) Create(
	ctx context.Context,
	authority entity.Authority,
) (*entity.Authority, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if _, exist := m.nameIndex[authority.Name]; exist {
		return nil, domain.NewAppError(
			domain.ErrDuplicate,
			fmt.Sprintf("Authority with name %s already exist", authority.Name),
		)
	}

	authority.ID = m.lastID
	m.lastID++

	saved := m.save(authority)
	return &saved, nil
}

func (m *MemoryAuthorityRepository) Update(
	ctx context.Context,
	authority entity.Authority,
) (*entity.Authority, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

    old, exists := m.db[authority.ID]
    if !exists {
        return nil, domain.NewAppError(
            domain.ErrNotFound,
            fmt.Sprintf(
                "Authority with ID %d doesn't exist",
                authority.ID,
            ),
        )
    }

    if old.Name != authority.Name {
        if _, exists := m.nameIndex[authority.Name]; exists {
            return nil, domain.NewAppError(
                domain.ErrDuplicate,
                fmt.Sprintf(
                    "Authority with name %s already exist",
                    authority.Name,
                ),
            )
        }

        delete(m.nameIndex, old.Name)
    }

	saved := m.save(authority)
	return &saved, nil
}

// =========== UTIL ===========

func (m *MemoryAuthorityRepository) save(data entity.Authority) entity.Authority {
	m.db[data.ID] = &data
	m.nameIndex[data.Name] = &data
	return data
}

func (m *MemoryAuthorityRepository) build() (error) {
	if len(m.db) == 0 {
		return nil
	}

	nameIdx := make(map[string]*entity.Authority, len(m.db))
	var maxID int64
	for _, data := range m.db {
		if maxID < data.ID {
			maxID = data.ID
		}

		if _, exist := nameIdx[data.Name]; exist {
			return domain.NewAppError(
				domain.ErrDuplicate,
				"Name is unique, failed to build authorityRepo",
			)
		}

		nameIdx[data.Name] = data
	}

	m.lastID = maxID + 1
	m.nameIndex = nameIdx
	return nil
}
