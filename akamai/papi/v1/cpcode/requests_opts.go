package cpcode

// CreateOpts represent optional query params for create CP-Code request.
type CreateOpts struct {
	ContractID string `param:"contractId"`
	GroupID string `param:"groupId"`
}

// GetOpts represent optional query params for get CP-Code request.
type GetOpts struct {
	ContractID string `param:"contractId"`
	GroupID string `param:"groupId"`
}

// ListOpts represent optional query params for list CP-Code request.
type ListOpts struct {
	ContractID string `param:"contractId"`
	GroupID string `param:"groupId"`
}

// CreateBody represents create CP-Code request body.
type CreateBody struct {
	ProductID string `json:"productId"`
	CPCodeName string `json:"cpcodeName"`
}

