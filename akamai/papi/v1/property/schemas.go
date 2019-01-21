package property

// Property represents configuration that  provides the main way to control
// how edge servers respond to various kinds of requests for those assets.
type Property struct {
	// ID is property’s unique identifier.
	ID string `json:"propertyId"`

	// Name is descriptive name for the property.
	Name string `json:"propertyName"`

	// AccountID identifies the account under which the property was created.
	AccountID string `json:"accountId"`

	// ContractID identifies the contract under which the property was created.
	ContractID string `json:"contractId"`

	// GroupID identifies the group to which the property is assigned.
	GroupID string `json:"groupId"`

	// LatestVersion specifies the most recent version of the property.
	LatestVersion *int `json:"latestVersion"`

	// StagingVersion is the most recent version to be activated to the test network,
	// otherwise it's nil.
	StagingVersion *int `json:"stagingVersion"`

	// ProductionVersion is the most recent version to be activated to the production
	// network, otherwise it's nil.
	ProductionVersion *int `json:"productionVersion"`

	// AssetID is an alternative identifier for the property, for internal use.
	AssetID string `json:"assetId"`

	// Note is further descriptive commentary for the property.
	Note string `json:"note"`
}
