package authority

import (
	"go-auth-service/internal/domain/authority"
	"strconv"
	"time"
)

type FormattedUnixTime int64

func ToAuthorityPublicResponse(a *authority.Authority) AuthorityPublicResponse {
	return AuthorityPublicResponse{
		ID: a.ID,
		Name: a.Name,
	}
}

func ToAuthorityAdminResponse(a *authority.Authority) AuthorityAdminResponse {
	return AuthorityAdminResponse{
		ID: a.ID,
		Name: a.Name,

		Description: a.Description,
		CreatedAt: FormattedUnixTime(a.CreatedAt),
		UpdatedAt: (*FormattedUnixTime)(a.UpdatedAt),
	}
}

func (t FormattedUnixTime) MarshalJSON() ([]byte, error) {
	formatted := time.Unix(int64(t), 0).UTC().Format("2006-01-02 15:04:05")
	return []byte(strconv.Quote(formatted)), nil
}
