package authority

import (
	"context"
	"fmt"
	"go-auth-service/internal/domain"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) (*Service) {
	return &Service{
		repo: repo,
	}
}


// ========== READ ==========

func (s *Service) GetByID(
	ctx context.Context, 
	id int64,
) (*Authority, error) {
	authority, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return  authority, nil
}

func (s *Service) GetByIDs(
	ctx context.Context,
	ids []int64,
) ([]*Authority, error) {
	authorities, err := s.repo.FindByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	return authorities, nil
}

// ========== WRITE ==========

func (s *Service) Create(
	ctx context.Context,
	entity Authority,
) (*Authority, error) {
	if err := s.throwIfNameExist(entity.Name); err != nil {
		return nil, err
	}

	authority, err := s.repo.Create(ctx, entity)
	if err != nil {
		return nil, err
	}

	return authority, nil
}

func (s *Service) Update(
	ctx context.Context, 
	entity Authority,
) (*Authority, error) {
	existing, err := s.repo.FindByID(ctx, entity.ID)
	if  err != nil {
		return nil, err
	}

	if existing.Name != entity.Name {
		if err := s.throwIfNameExist(entity.Name); err != nil {
			return nil, err
		}
	}

	updated, err := s.repo.Update(ctx, entity)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *Service) DeleteByID(
	ctx context.Context,
	id int64,
) error {
	err := s.repo.DeleteByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// ========== UTIL ==========

func (s *Service) throwIfNameExist(name string) error {
	if exist := s.repo.ExistByName(name); exist {
		return domain.NewAppError(
			domain.ErrDuplicate,
			fmt.Sprintf(
				"Authority with name %s already exist, name must be unique",
				name,
				),
			)
	}

	return nil
}
