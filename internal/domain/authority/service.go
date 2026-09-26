package authority

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) (*Service) {
	return &Service{
		repo: repo,
	}
}


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
