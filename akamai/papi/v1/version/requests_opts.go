package version

// CreateOpts represents optional query params for create property version request.
type CreateOpts struct {
	// ContractID is identifier for the contract. The parameter is optional if the
	// property has been provisioned under only one contract. Otherwise you need
	// to specify it along with the GroupID. (In other operations that don’t specify
	// a propertyId URL parameter, this parameter is always required.)
	ContractID string `param:"contractId"`

	// GroupID is identifier for the group. The parameter is optional if the property
	// has been provisioned under only one group. Otherwise you need to specify it along
	// with the ContractID. (In other operations that don’t specify a propertyId URL parameter,
	// this parameter is always required.)
	GroupID string `param:"groupId"`
}

// GetOpts represent optional query params for get property version request.
type GetOpts struct {
	// ContractID is identifier for the contract. The parameter is optional if the
	// property has been provisioned under only one contract. Otherwise you need
	// to specify it along with the GroupID. (In other operations that don’t specify
	// a propertyId URL parameter, this parameter is always required.)
	ContractID string `param:"contractId"`

	// GroupID is identifier for the group. The parameter is optional if the property
	// has been provisioned under only one group. Otherwise you need to specify it along
	// with the ContractID. (In other operations that don’t specify a propertyId URL parameter,
	// this parameter is always required.)
	GroupID string `param:"groupId"`


	ActivatedOn string `param:"activatedOn"`
}

// ListOpts represents optional query params for list property version request.
type ListOpts struct {
	// ContractID is identifier for the contract. The parameter is optional if the
	// property has been provisioned under only one contract. Otherwise you need
	// to specify it along with the GroupID. (In other operations that don’t specify
	// a propertyId URL parameter, this parameter is always required.)
	ContractID string `param:"contractId"`

	// GroupID is identifier for the group. The parameter is optional if the property
	// has been provisioned under only one group. Otherwise you need to specify it along
	// with the ContractID. (In other operations that don’t specify a propertyId URL parameter,
	// this parameter is always required.)
	GroupID string `param:"groupId"`
}

// CreateBody represents request body for a property creation.
type CreateBody struct {
	// CreateFromVersion is the property version on which to base the new version.
	CreateFromVersion int `json:"createFromVersion"`

	// CreateFromVersionEtag is the data digest of the version on which to
	// base the new version.
	CreateFromVersionEtag string `json:"createFromVersionEtag,omitempty"`
}
