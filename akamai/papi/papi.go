package papi

import "github.com/dstdfx/go-akamai/akamai"

const (
	// ServiceType contains the name of the Akamai service which this
	// package is intended for.
	ServiceType = "papi"

	// UserAgent contains the user agent for all versions of the Akamai Property API client.
	UserAgent = akamai.DefaultUserAgent
)
