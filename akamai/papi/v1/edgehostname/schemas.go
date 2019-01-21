package edgehostname

type IPVersionBehaviorValue string

const (
	IPVersionBehaviorV4 IPVersionBehaviorValue = "IPV4"
	IPVersionBehaviorV4V6 IPVersionBehaviorValue = "IPV6_COMPLIANCE"
)

// EdgeHostname represents an Edge Hostname resource
type EdgeHostname struct {
	ID         			   string      `json:"edgeHostnameId,omitempty"`
	EdgeHostnameDomain     string      `json:"edgeHostnameDomain,omitempty"`
	ProductID              string      `json:"productId"`
	DomainPrefix           string      `json:"domainPrefix"`
	DomainSuffix           string      `json:"domainSuffix"`
	Secure                 bool        `json:"secure,omitempty"`
	IPVersionBehavior      IPVersionBehaviorValue      `json:"ipVersionBehavior,omitempty"`
	MapDetailsSerialNumber int         `json:"mapDetails:serialNumber,omitempty"`
	MapDetailsSlotNumber   int         `json:"mapDetails:slotNumber,omitempty"`
	MapDetailsMapDomain    string      `json:"mapDetails:mapDomain,omitempty"`
}