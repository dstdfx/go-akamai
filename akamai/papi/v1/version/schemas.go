package version

import "time"

type NetworkStatus string

const (
	ActiveStatus NetworkStatus = "ACTIVE"
	InactiveStatus NetworkStatus = "INACTIVE"
	PendingStatus NetworkStatus = "PENDING"
)

// Version specifies the version of a property.
type Version struct {
	// ProductionStatus reflects whether the version has been activated to the production network.
	// If ACTIVE, the version is read-only. Otherwise it may be INACTIVE or PENDING.
	ProductionStatus NetworkStatus `json:"productionStatus"`

	// StagingStatus reflects whether the version has been activated to the test network.
	// If ACTIVE, the version is read-only. Otherwise it may be INACTIVE or PENDING.
	StagingStatus NetworkStatus `json:"stagingStatus"`

	// Etag is a digest with which to check the data’s integrity.
	Etag string `json:"etag"`

	// ProductID is the product assigned to the property when versioned.
	ProductID string `json:"productId"`

	// RuleFormat identifies the rule format currently assigned to the property version’s rule tree.
	RuleFormat string `json:"ruleFormat"`

	// Note is a descriptive log comment for the current version.
	Note string `json:"note"`

	// Read-only fields

	// EdgeHostname field. When searching for hostname or edgeHostname, this shows the
	// relevant edge hostname to which the active version currently applies.
	EdgeHostname string `json:"edgeHostname,omitempty"`

	// Hostname field. When searching for hostname or edgeHostname, this shows the
	// relevant property hostname to which the active version currently applies.
	Hostname string `json:"hostname,omitempty"`

	// ContractID identifies the contract under which the property version is active.
	ContractID string `json:"contractId,omitempty"`

	// AssetID provides an alternative property identifier, for internal use.
	AssetID string `json:"assetId,omitempty"`

	// AccountID identifies the account under which the property version is active.
	AccountID string `json:"accountId,omitempty"`

	// GroupID identifies the group under which the property version is active.
	GroupID string `json:"groupId,omitempty"`

	// PropertyID identifies the property to which the listed version belongs.
	PropertyID string `json:"propertyId,omitempty"`

	// PropertyName provides the name of the property to which the listed version belongs.
	PropertyName string `json:"propertyName,omitempty"`

	// PropertyVersion is a positive integer identifying the incremental version.
	PropertyVersion *int `json:"propertyVersion"`

	// UpdatedByUser represents the username associated with the new version.
	UpdatedByUser string `json:"updatedByUser"`

	// UpdateDate represents the date stamp of the version’s latest update.
	UpdatedDate *time.Time `json:"updatedDate"`
}
