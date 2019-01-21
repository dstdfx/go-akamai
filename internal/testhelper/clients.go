package testhelper

import (
	"net/http"

	"github.com/dstdfx/go-akamai/akamai/papi"

	"github.com/dstdfx/go-akamai/akamai"
	"github.com/dstdfx/go-akamai/internal/logger"
	"github.com/dstdfx/go-akamai/internal/util"
)

// NewTestV1PAPIClient method initialize Akamai client for Property API v1.
func (tenv *TestEnv) NewTestV1PAPIClient() {
	apiVersion := "v1"

	client := &akamai.ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   util.BuildServiceEndpoint(
			tenv.Config.Host,
			papi.ServiceType,
			apiVersion),
		UserAgent:  akamai.DefaultUserAgent,
		Config: tenv.Config,
	}
	client.SetLogger(&logger.VoidLogger{})
	tenv.Client = client
}