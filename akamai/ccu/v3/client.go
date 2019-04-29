package v3

import (
	"github.com/dstdfx/go-akamai/akamai"
	"github.com/dstdfx/go-akamai/akamai/ccu"
	"github.com/dstdfx/go-akamai/akamai/edgegrid"
	"github.com/dstdfx/go-akamai/akamai/internal/logger"
	"github.com/dstdfx/go-akamai/akamai/internal/util"
)

// APIVersion represents current API version.
const APIVersion = "v3"

// NewV3CCUClient initializes a new Akamai client for  Content Control Utility API v3.
func NewV3CCUClient(config *edgegrid.Config, log ...logger.GenericLogger) *akamai.ServiceClient {
	client := &akamai.ServiceClient{
		HTTPClient: akamai.NewHTTPClient(),
		UserAgent: ccu.UserAgent,
		Endpoint: util.BuildServiceEndpoint(
			config.Host,
			ccu.ServiceType,
			APIVersion),
		Config: config,
	}

	v := logger.SelectLogger(log...)
	client.SetLogger(v)
	return client
}
