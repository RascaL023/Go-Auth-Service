package authority

import "context"

type Repository interface {

	FindByID(ctx context.Context, id int64) (*Authority, error)
	FindByIDs(ctx context.Context, ids []int64) ([]*Authority, error)

}
