package hostname

// ListOpts represents optional query params for list hostnames request.
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

	// ValidateHostnames is flag when false (the default), skips validation tests that would identify potential
	// hostname-related problems within the response object’s errors and warnings arrays.
	ValidateHostnames bool `param:"validateHostnames"`
}

// UpdateOpts represents optional query params for update hostnames request.
type UpdateOpts struct {
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

	// ValidateHostnames is flag when false (the default), skips validation tests that would identify potential
	// hostname-related problems within the response object’s errors and warnings arrays.
	ValidateHostnames bool `param:"validateHostnames"`
}

// UpdateBody represents update request body for update hostnames request.
type UpdateBody struct {
	// Hostnames represent an updated list of property's hostnames.
	Hostnames []*Hostname `json:"hostnames"`
}
