package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	th "github.com/dstdfx/go-akamai/akamai/internal/testhelper"
	"github.com/dstdfx/go-akamai/akamai/papi/v1/hostname"
)

func TestGetHostnames(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/properties/prp_175780/versions/1/hostnames",
		RawResponse: testListHostnamesRawResponse,
		Method:      http.MethodGet,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testListHostnamesExpected

	actual, _, err := hostname.List(
		context.Background(),
		testEnv.Client,
		"prp_175780",
		1,
		hostname.ListOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get a list of property's hostnames")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected Slice of pointers to property's hostnames, but got %v", actualKind)
	}
	if len(actual) != 2 {
		t.Errorf("expected 2 property's hostnames, but got %d", len(actual))
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestUpdateHostnames(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/properties/prp_175780/versions/1/hostnames",
		RawResponse: testUpdateHostnamesRawResponse,
		RawRequest: testUpdateHostnamesRawRequest,
		Method:      http.MethodPut,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testUpdateHostnamesExpected

	actual, _, err := hostname.Update(
		context.Background(),
		testEnv.Client,
		"prp_175780",
		1,
		hostname.UpdateBody{
			Hostnames: []*hostname.Hostname{
				{
					CnameType: hostname.CnameTypeEdgeHostname,
					CnameFrom: "example.com",
					EdgeHostnameID: "ehn_895833",
					CnameTo: "example.com.edgesuite.net",

				},
				{
					CnameType: hostname.CnameTypeEdgeHostname,
					EdgeHostnameID: "ehn_895824",
					CnameFrom: "m.example.com",
				},
			},
		},
		hostname.UpdateOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get a list of updated property's hostnames")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected Slice of pointers to updated property's hostnames, but got %v",
			actualKind)
	}
	if len(actual) != 2 {
		t.Errorf("expected 2 updated property's hostnames, but got %d", len(actual))
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}
