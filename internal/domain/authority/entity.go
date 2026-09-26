package authority

type Authority struct {
	ID			int64
	Name 		string
	Description *string

	CreatedAt 	int64
	UpdatedAt 	*int64
	DeletedAt 	*int64
}
