package testhelper

import (
	"net/http"
	"net/http/httptest"

	"github.com/dstdfx/go-akamai/akamai"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

const (
	TestContractID = "ctr_C-10UBXF3"
	TestGroupID    = "grp_52301"
	TestPropertyID = "prp_509694"
	TestProductID  = "prd_HTTP_Downloads"
	TestRuleFormat = "v2018-09-12"
)

// TestEnv represents a testing environment for all resources.
type TestEnv struct {
	Mux    *http.ServeMux
	Server *httptest.Server
	Client *akamai.ServiceClient
	Config *edgegrid.Config
}

// SetupTestEnv prepares the new testing environment.
func SetupTestEnv() *TestEnv {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	testEnv := &TestEnv{
		Mux:    mux,
		Server: server,
		Config: NewTestEdgeGridConfig(server.URL),
	}
	return testEnv
}

// TearDownTestEnv releases the testing environment.
func (tenv *TestEnv) TearDownTestEnv() {
	tenv.Server.Close()
	tenv.Server = nil
	tenv.Mux = nil
	tenv.Client = nil
	tenv.Config = nil
}
