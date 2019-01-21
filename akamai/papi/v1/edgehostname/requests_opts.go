package edgehostname

// GetOpts represent optional query params for get edge-hostname request.
type GetOpts struct {
	ContractID string `param:"contractId"`
	GroupID string `param:"groupId"`
	Options string `param:"options,omitempty"`
}

// ListOpts represent optional query params for list edge-hostname request.
type ListOpts struct {
	ContractID string `param:"contractId"`
	GroupID string `param:"groupId"`
	Options string `param:"options,omitempty"`
}

// CreateOpts represent optional query params for create edge-hostname request.
type CreateOpts struct {
	ContractID string `param:"contractId"`
	GroupID string `param:"groupId"`
	Options string `param:"options,omitempty"`
}

// CreateBody represents request body for edge-hostname creation.
type CreateBody struct {
	ProductID         string `json:"productId"`
	DomainPrefix      string `json:"domainPrefix"`
	DomainSuffix      string `json:"domainSuffix"`
	Secure            *bool   `json:"secure,omitempty"`
	IPVersionBehavior IPVersionBehaviorValue `json:"ipVersionBehavior"`
	SlotNumber        *int   `json:"slotNumber,omitempty"`
}