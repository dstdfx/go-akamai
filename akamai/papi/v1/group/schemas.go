package group

// Group represents an object which control access to properties and
// help to consolidate reporting functions.
type Group struct {
	ID string `json:"groupId"`
	Name string `json:"groupName"`
	ParentGroupID string `json:"parentGroupId,omitempty"`
	ContractIDs []string `json:"contractIds"`
}
