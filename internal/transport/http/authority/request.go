package authority

type AuthorityAdminRequest struct {
	Name		string	`json:"name,omitempty"`
	Description	*string	`json:"description,omitempty"`
}

type AuthorityAdminPutRequest struct {
	Name		string	`json:"name,omitempty"`
	Description	*string	`json:"description,omitempty"`
}

type AuthorityAdminPatchRequest struct {
	Name		*string	`json:"name,omitempty"`
	Description	*string	`json:"description,omitempty"`
}


func (p *AuthorityAdminPatchRequest) IsEmptyPatch() bool {
	return p.Name == nil && p.Description == nil
}
