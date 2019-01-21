package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/dstdfx/go-akamai/akamai/papi/v1/contract"
	th "github.com/dstdfx/go-akamai/internal/testhelper"
)

func TestListContracts(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/contracts",
		RawResponse: testListContractsRawResp,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testListContractsExpected

	actual, _, err := contract.List(context.Background(), testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get contracts")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected Slice of pointers to contracts, but got %v", actualKind)
	}
	if len(actual) != 1 {
		t.Errorf("expected 1 contract, but got %d", len(actual))
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}
