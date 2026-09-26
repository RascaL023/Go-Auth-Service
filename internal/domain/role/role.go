package role

import "go-auth-service/internal/domain/authority"

type Role struct {
	ID			int64
	Name 		string
	Description *string
	Authorities []authority.Authority

	CreatedAt 	int64
	UpdatedAt 	*int64
	DeletedAt 	*int64
}
