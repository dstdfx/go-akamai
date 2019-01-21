package property

// CreateOpts represent optional query params for create property request.
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

// ListOpts represent optional query params for list properties request.
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

// GetOpts represent optional query params for get property request.
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
}

// DeleteOpts represent optional query params for delete property request.
type DeleteOpts struct {
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

// CreateBody represents request body for property creation.
type CreateBody struct {
	ProductID string `json:"productId"`
	PropertyName string `json:"propertyName"`

	// CloneFrom identifies another property instance to clone when making a POST request
	// to create a new property. The cloned property must share the same contract and group.
	CloneFrom *CloneFrom `json:"cloneFrom,omitempty"`
}

// CloneFrom identifies another property instance to clone when making a POST request
// to create a new property. The cloned property must share the same contract and group.
type CloneFrom struct {
	// PropertyID specifies the property to clone.
	PropertyID string `json:"propertyId"`

	// Version specifies the version of the property to clone.
	Version int `json:"version"`

	// VersionEtag performs an etag data integrity check on the original property.
	VersionEtag *string `json:"cloneFromVersionEtag,omitempty"`

	// CopyHostnames indicates if the same set of hostnames is assigned to the new property,
	// false by default.
	CopyHostnames *bool `json:"copyHostnames,omitempty"`
}
