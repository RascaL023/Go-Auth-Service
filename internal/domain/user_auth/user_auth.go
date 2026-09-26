package user_auth

import "go-auth-service/internal/domain/role"

type UserAuth struct {
	ID				int64
	Email 			string
	HashedPassword 	string
	Roles 			[]role.Role

	CreatedAt 		int64
	UpdatedAt 		*int64
	DeletedAt 		*int64
}
