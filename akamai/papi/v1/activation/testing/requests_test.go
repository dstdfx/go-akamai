package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	th "github.com/dstdfx/go-akamai/akamai/internal/testhelper"
	"github.com/dstdfx/go-akamai/akamai/internal/util"
	"github.com/dstdfx/go-akamai/akamai/papi/v1/activation"
)

func TestListActivations(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/properties/prp_511225/activations",
		RawResponse: testListActivationsRawResponse,
		Method:      http.MethodGet,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testListActivationsExpected

	actual, _, err := activation.List(
		context.Background(),
		testEnv.Client,
		"prp_511225",
		activation.ListOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get a list of property's activations")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to a list of property's activations, but got %v",
			actualKind)
	}
	if len(actual) != 1 {
		t.Errorf("expected 1 property's activation, but got %d", len(actual))
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestGetActivation(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/properties/prp_511225/activations/atv_6804950",
		RawResponse: testGetActivationRawResponse,
		Method:      http.MethodGet,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testGetActivationExpected

	actual, _, err := activation.Get(
		context.Background(),
		testEnv.Client,
		"prp_511225",
		"atv_6804950",
		activation.GetOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get a property's activation")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to property's activation, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestGetByLinkActivation(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         testActivationLink,
		RawResponse: testGetActivationByLinkRawResponse,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testGetActivationByLinkExpected

	actual, _, err := activation.GetByLink(
		context.Background(),
		testEnv.Client,
		testActivationLink)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get a property's activation")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to property's activation, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestCreateActivation(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/properties/prp_511225/activations",
		RawResponse: testCreateActivationRawResponse,
		RawRequest: testCreateActivationRawRequest,
		Method:      http.MethodPost,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testCreateActivationExpected

	actual, _, err := activation.Create(
		context.Background(),
		testEnv.Client,
		"prp_511225",
		activation.CreateBody{
			PropertyVersion: 1,
			Network: activation.NetworkProduction,
			Note: "test from go",
			NotifyEmails: []string{"dstdfx@example.ru"},
			AcknowledgeAllWarnings: util.BoolPtr(true),
		},
		activation.CreateOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't create a property's activation")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.String {
		t.Errorf("expected string link to property's activation, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestCancelActivation(t *testing.T){
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/properties/prp_511225/activations/atv_6804950",
		RawResponse: testCancelActivationRawResponse,
		Method:      http.MethodDelete,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testCancelActivationExpected

	actual, _, err := activation.Cancel(
		context.Background(),
		testEnv.Client,
		"prp_511225",
		"atv_6804950",
		activation.CancelOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't cancel a property's activation")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to property's activation, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestCancelActivation_AlreadyCanceled(t *testing.T){
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/properties/prp_511225/activations/atv_6804950",
		Method:      http.MethodDelete,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusNoContent,
		CallFlag:    &endpointCalled,
	})

	_, _, err := activation.Cancel(
		context.Background(),
		testEnv.Client,
		"prp_511225",
		"atv_6804950",
		activation.CancelOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't cancel a property's activation")
	}
}
