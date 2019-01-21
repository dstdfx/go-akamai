package contract

// Contract represents a fixed term of service for Akamai user.
type Contract struct {
	ID string `json:"contractId"`
	TypeName string `json:"contractTypeName"`
}
