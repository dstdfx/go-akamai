package hostname

// CnameTypeValue is used to create an "enum" of possible Hostname.CnameType values
type CnameTypeValue string

// CnameTypeEdgeHostname Hostname.CnameType value EDGE_HOSTNAME
const CnameTypeEdgeHostname CnameTypeValue = "EDGE_HOSTNAME"


// Hostname specifies an edge hostname that is applied to a property version.
type Hostname struct {
	// CnameType is the hostname for edge content, corresponding to the edge hostname
	// object’s edgeHostnameDomain member.
	CnameType CnameTypeValue `json:"cnameType"`

	// EdgeHostnameID is a unique identifier for the edge hostname.
	EdgeHostnameID string `json:"edgeHostnameId,omitempty"`

	// CnameFrom is the original origin’s hostname.
	CnameFrom string `json:"cnameFrom"`

	// CnameTo is the hostname for edge content, corresponding to
	// the edge hostname object’s edgeHostnameDomain member.
	CnameTo string `json:"cnameTo,omitempty"`
}
