package cache


type Network string

const (
	StagingNetwork Network = "staging"
	ProductionNetwork Network = "production"
)

// CCUResponse represents common response that might be get by calling CCU API v3.
type CCUResponse struct {
	// DescribedBy represents the URL for the API’s machine readable
	// document. It describes the error code in more detail.
	DescribedBy string `json:"describedBy"`

	// Detail represents d etailed information about the HTTP status code
	// returned with the response.
	Detail string `json:"detail"`

	// EstimatedSeconds is the estimated number of seconds before the purge
	// is to complete.
	EstimatedSeconds *int `json:"estimatedSeconds"`

	// HTTPStatus is the HTTP code that indicates the status of the request
	// to invalidate or purge content. Successful requests yield a 201 code.
	HTTPStatus *int `json:"httpStatus"`

	// PurgeID is the unique identifier for the purge request.
	PurgeID string `json:"purgeId"`

	// SupportID is the identifier to provide Akamai Technical Support if issues arise.
	SupportID string `json:"supportId"`

	// Title describes the response type, for example, Rate Limit exceeded.
	Title string `json:"title"`
}