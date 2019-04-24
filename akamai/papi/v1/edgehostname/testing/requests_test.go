package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	th "github.com/dstdfx/go-akamai/akamai/internal/testhelper"
	"github.com/dstdfx/go-akamai/akamai/internal/util"
	"github.com/dstdfx/go-akamai/akamai/papi/v1/edgehostname"
)

func TestGetEdgeHostnames(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/edgehostnames/" + testGetEdgeHostnameExpected.ID,
		RawResponse: testGetEdgeHostnameRawResponse,
		Method:      http.MethodGet,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testGetEdgeHostnameExpected

	actual, _, err := edgehostname.Get(
		context.Background(),
		testEnv.Client,
		expected.ID,
		edgehostname.GetOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get cp codes")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to edge-hostname, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestGetByLinkEdgeHostnames(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         testEdgeHostnameLink,
		RawResponse: testGetByLinkEdgeHostnameRawResponse,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testGetByLinkEdgeHostnameExpected

	actual, _, err := edgehostname.GetByLink(
		context.Background(),
		testEnv.Client,
		testEdgeHostnameLink)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get cp codes")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to edge-hostname, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestCreateEdgeHostnames(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/edgehostnames",
		RawResponse: testCreateEdgeHostnameRawResponse,
		RawRequest: testCreateEdgeHostnameRawRequest,
		Method:      http.MethodPost,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testCreateEdgeHostnameExpected

	actual, _, err := edgehostname.Create(
		context.Background(),
		testEnv.Client,
		edgehostname.CreateBody{
			ProductID: "prd_HTTP_Downloads",
			DomainPrefix: "www.example.com",
			DomainSuffix: "edgesuite.net",
			Secure: util.BoolPtr(true),
			SecureNetwork: edgehostname.SecureNetworkStandartTLS,
			IPVersionBehavior: edgehostname.IPVersionBehaviorV4,
			SlotNumber: util.IntPtr(12345),
		},
		edgehostname.CreateOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get cp codes")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.String {
		t.Errorf("expected string link toedge-hostname, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestListEdgeHostnames(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/edgehostnames",
		RawResponse: testListEdgeHostnamesRawResponse,
		Method:      http.MethodGet,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testListEdgeHostnamesExpected

	actual, _, err := edgehostname.List(
		context.Background(),
		testEnv.Client,
		edgehostname.ListOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get list edge-hostnames")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to edge-hostnames, but got %v", actualKind)
	}
	if len(actual) != 2 {
		t.Errorf("expected 2 edge-hostnames, but got %d", len(actual))
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}
