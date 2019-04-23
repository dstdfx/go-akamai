package edgehostname

type IPVersionBehaviorValue string

const (
	IPVersionBehaviorV4 IPVersionBehaviorValue = "IPV4"
	IPVersionBehaviorV4V6 IPVersionBehaviorValue = "IPV6_COMPLIANCE"
)

type SecureNetworkValue string

const (
	SecureNetworkStandartTLS SecureNetworkValue = "STANDARD_TLS"
	SecureNetworkSharedCert SecureNetworkValue = "SHARED_CERT"
	SecureNetworkEnhancedTLS SecureNetworkValue = "ENHANCED_TLS"
)

// EdgeHostname represents an Edge Hostname resource
type EdgeHostname struct {
	// ID is edge hostname's unique identifier.
	ID         			   string      `json:"edgeHostnameId,omitempty"`

	// EdgeHostnameDomain represents the full edge domain name
	// formed from the domainPrefix and domainSuffix.
	EdgeHostnameDomain     string      `json:"edgeHostnameDomain,omitempty"`

	// ProductID is an identifier of the the product assigned to the property.
	ProductID              string      `json:"productId"`

	// DomainPrefix is the origin domain portion of the edge hostname. An edge hostname consists
	// of a customer-specific namePrefix such as www.example.com and an Akamai-specific
	// domainSuffix such as edgesuite.net. The edge hostname’s final DNS CNAME combines the two,
	// for example, www.example.com.edgesuite.net.
	DomainPrefix           string      `json:"domainPrefix"`

	// DomainSuffix is the Akamai-specific portion of the edge hostname, for example, edgesuite.net.
	DomainSuffix           string      `json:"domainSuffix"`

	// Secure flag indicates if the edge hostname is to be used with SSL.
	Secure                 bool        `json:"secure,omitempty"`

	// IPVersionBehavior indicates which version of the IP protocol to use: IPV4 for version 4 only,
	// or IPV6_COMPLIANCE for both 4 and 6. The default value for requests is IPV4.
	IPVersionBehavior      IPVersionBehaviorValue      `json:"ipVersionBehavior,omitempty"`

	// Additional meta information.
	MapDetailsSerialNumber int         `json:"mapDetails:serialNumber,omitempty"`
	MapDetailsSlotNumber   int         `json:"mapDetails:slotNumber,omitempty"`
	MapDetailsMapDomain    string      `json:"mapDetails:mapDomain,omitempty"`
}