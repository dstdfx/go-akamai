package testhelper

import (
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

const (
	TestClientSecret = "cHYFZMDzULZb6b7TSi+JfalaBExw85+bmRHC705jP6A="
	TestClientToken = "akac-uf3nvl8qhfa26fns-scnrjt2fyqh3jnvn"
	TestAccessToken = "akac-h68622kxzf7ep3f3-uvegmgoluaosnmqc"
	TestMaxBody = 131072
)

// NewTestEdgeGridConfig creates fake Akamai Edge Grid config.
func NewTestEdgeGridConfig(host string) *edgegrid.Config{
	return &edgegrid.Config{
		Host: host,
		ClientSecret: TestClientSecret,
		ClientToken: TestClientToken,
		AccessToken: TestAccessToken,
		MaxBody: TestMaxBody,
	}
}
