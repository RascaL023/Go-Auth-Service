package authority

type AuthorityPublicResponse struct {
	ID			int64				`json:"id"`
	Name		string				`json:"name"`
}

type AuthorityAdminResponse struct {
	ID			int64				`json:"id"`
	Name		string				`json:"name"`
	Description	*string				`json:"description"`

	CreatedAt 	FormattedUnixTime  	`json:"createdAt"`
	UpdatedAt 	*FormattedUnixTime	`json:"updatedAt"`
	DeletedAt 	*FormattedUnixTime	`json:"deletedAt"`
} 
