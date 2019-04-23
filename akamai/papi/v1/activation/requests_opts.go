package activation

// GetOpts represents optional query params for get activation request.
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

// CancelOpts represents optional query params for cancel activation request.
type CancelOpts struct {
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

// ListOpts represents optional query params for list activation request.
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

// CreateOpts represents optional query params for create activation request.
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

// CreateBody represents a request body for activation create request.
type CreateBody struct {
	// Type field. Either ACTIVATE or DEACTIVATE. The default is ACTIVATE.
	// Any new activation automatically deactivates the current activation. Note that
	// if you were to POST a DEACTIVATE type on an active property, it would no longer
	// serve any traffic. You would need to modify (de-Akamaize) your DNS configuration
	// and specify a different way to field the traffic.
	ActivationType      Value        `json:"activationType,omitempty"`

	// Network represents the network to activate on, either STAGING or PRODUCTION.
	Network             NetworkValue `json:"network"`

	// NotifyEmails is a list of email address strings to notify when the activation status changes.
	NotifyEmails        []string     `json:"notifyEmails"`

	// AcknowledgeWarnings is a list of msg_-prefixed string IDs acknowledging any warnings
	// noted in responses to previous activation requests
	AcknowledgeWarnings []string     `json:"acknowledgeWarnings,omitempty"`

	// AcknowledgeAllWarnings flag. When POSTing an activation, allows you to skip acknowledging each warning.
	// With this enabled, you can omit the acknowledgeWarnings array.
	AcknowledgeAllWarnings *bool `json:"acknowledgeAllWarnings,omitempty"`

	// UseFastFallback flag. After activating a property and finding it causes problems, POSTing another activation within
	// one hour with useFastFallback enabled quickly rolls back to the previous version.
	// This option is only available for activations where canFastFallback is true.
	UseFastFallback *bool `json:"useFastFallback,omitempty"`

	// PropertyVersion is the property version targeted with activation. Once activated,
	// you can no longer modify that version of the property.
	PropertyVersion     int          `json:"propertyVersion"`

	// Note assigns a log message to the activation request.
	Note                string       `json:"note,omitempty"`
}