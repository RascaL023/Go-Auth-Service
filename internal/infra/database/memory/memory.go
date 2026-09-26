package memory

import entity "go-auth-service/internal/domain/authority"

func NewConnection() *MemoryAuthorityRepository {
	return &MemoryAuthorityRepository{
		DB: make(map[int64]*entity.Authority),
	}
}
