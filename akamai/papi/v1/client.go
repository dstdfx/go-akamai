package v1

import (
	"github.com/dstdfx/go-akamai/akamai"
	"github.com/dstdfx/go-akamai/akamai/edgegrid"
	"github.com/dstdfx/go-akamai/akamai/internal/logger"
	"github.com/dstdfx/go-akamai/akamai/internal/util"
	"github.com/dstdfx/go-akamai/akamai/papi"
)

// APIVersion represents current API version.
const APIVersion = "v1"

// NewV1PAPIClient initializes a new Akamai client for Property API v1.
func NewV1PAPIClient(config *edgegrid.Config, log ...logger.GenericLogger) *akamai.ServiceClient {
	client := &akamai.ServiceClient{
		HTTPClient: akamai.NewHTTPClient(),
		UserAgent: papi.UserAgent,
		Endpoint: util.BuildServiceEndpoint(
			config.Host,
			papi.ServiceType,
			APIVersion),
		Config: config,
	}

	v := logger.SelectLogger(log...)
	client.SetLogger(v)
	return client
}