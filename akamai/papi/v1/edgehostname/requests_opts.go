package edgehostname

// GetOpts represent optional query params for get edge-hostname request.
type GetOpts struct {
	// ContractID is identifier for the contract.
	ContractID string `param:"contractId"`

	// GroupID is identifier for the group.
	GroupID string `param:"groupId"`

	// Options represents comma-separated list of options to enable; mapDetails
	// enables extra mapping-related information.
	Options string `param:"options,omitempty"`
}

// ListOpts represent optional query params for list edge-hostname request.
type ListOpts struct {
	// ContractID is identifier for the contract.
	ContractID string `param:"contractId"`

	// GroupID is identifier for the group.
	GroupID string `param:"groupId"`

	// Options represents comma-separated list of options to enable; mapDetails
	// enables extra mapping-related information.
	Options string `param:"options,omitempty"`
}

// CreateOpts represent optional query params for create edge-hostname request.
type CreateOpts struct {
	// ContractID is identifier for the contract.
	ContractID string `param:"contractId"`

	// GroupID is identifier for the group.
	GroupID string `param:"groupId"`

	// Options represents comma-separated list of options to enable; mapDetails
	// enables extra mapping-related information.
	Options string `param:"options,omitempty"`
}

// CreateBody represents request body for edge-hostname creation.
type CreateBody struct {
	// ProductID is an identifier of the the product assigned to the property.
	ProductID         string `json:"productId"`

	// DomainPrefix is the origin domain portion of the edge hostname. An edge hostname consists
	// of a customer-specific namePrefix such as www.example.com and an Akamai-specific
	// domainSuffix such as edgesuite.net. The edge hostname’s final DNS CNAME combines the two,
	// for example, www.example.com.edgesuite.net.
	DomainPrefix      string `json:"domainPrefix"`

	// DomainSuffix is the Akamai-specific portion of the edge hostname, for example, edgesuite.net.
	DomainSuffix      string `json:"domainSuffix"`

	// Secure flag indicates if the edge hostname is to be used with SSL.
	// The default value for POST requests is false. Setting secure:true for new edge
	// hostnames is not supported.
	Secure            *bool   `json:"secure,omitempty"`

	// SecureNetwork specifies the type of security for the new edge hostname.
	// With STANDARD_TLS specified, specify a domainSuffix of edgesuite.net.
	// With SHARED_CERT specified, specify a domainSuffix of akamaized.net.
	// With ENHANCED_TLS specified, you need to specify a certEnrollmentId value along
	// with a domainSuffix of edgekey.net.
	SecureNetwork SecureNetworkValue `json:"secureNetwork"`

	// IPVersionBehavior indicates which version of the IP protocol to use: IPV4 for version 4 only,
	// or IPV6_COMPLIANCE for both 4 and 6. The default value for requests is IPV4.
	IPVersionBehavior IPVersionBehaviorValue `json:"ipVersionBehavior"`

	// SlotNumber represents the slot number for ESSL (secure) properties.
	SlotNumber        *int   `json:"slotNumber,omitempty"`

	// CertEnrollmentID represents an identifier of the enrollment.
	// When creating an Enhanced TLS edge hostname, this sets the certificate enrollment ID.
	// Specify this on POST, with secureNetwork set to ENHANCED_TLS. To obtain a value programmatically,
	// run the Certificate Provisioning System API’s List enrollments operation.
	// Choose the appropriate enrollment, strip the leading path expression from its location member,
	// and use that value as the certEnrollmentId.
	CertEnrollmentID *int `json:"certEnrollmentId,omitempty"`
}