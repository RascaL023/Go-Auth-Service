package authority

import  "context"

type Repository interface {

	FindByID(context.Context, int64) (*Authority, error)
	FindByIDs(context.Context, []int64) ([]*Authority, error)
	ExistByName(string) bool

	Create(context.Context, Authority) (*Authority, error)
	Update(context.Context, Authority) (*Authority, error)
	DeleteByID(context.Context, int64) error

}
