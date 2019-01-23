package rule

type ValidationMode string

const (
	FastValidateMode ValidationMode = "fast"
	FullValidateMode ValidationMode = "full"
)

// GetOpts represents optional query params for get rule tree request.
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

	// ValidateMode field, when setting this to fast performs a quick validation
	// check based on the provided JSON. This is faster than the default full validation,
	// which performs more extensive checks on the converted XML metadata configuration.
	ValidateMode ValidationMode `json:"validateMode"`

	// ValidateRules field. Set to true by default. When false, skips validation tests that
	// would identify potential problems within the response object’s errors and warnings arrays.
	ValidateRules bool `json:"validateRules"`
}

// UpdateOpts represents optional query params for update rule tree request.
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

	// ValidateMode field, when setting this to fast performs a quick validation
	// check based on the provided JSON. This is faster than the default full validation,
	// which performs more extensive checks on the converted XML metadata configuration.
	ValidateMode ValidationMode `json:"validateMode"`

	// ValidateRules field. Set to true by default. When false, skips validation tests that
	// would identify potential problems within the response object’s errors and warnings arrays.
	ValidateRules bool `json:"validateRules"`

	// DryRun flag allows for a dry run in order to gather any possible errors without
	// saving the rule tree.
	DryRun bool `json:"dryRun"`
}

// UpdateBody represents update request body.
type UpdateBody struct {
	// Rule represents updated rule object to be applied.
	Rule *Rule `json:"rules"`
}
