package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	th "github.com/dstdfx/go-akamai/akamai/internal/testhelper"
	"github.com/dstdfx/go-akamai/akamai/internal/util"
	"github.com/dstdfx/go-akamai/akamai/papi/v1/property"
)

func TestCreateProperty(t *testing.T){
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/properties",
		RawRequest: testCreatePropertyRawRequest,
		RawResponse: testCreatePropertyRawResponse,
		Method:      http.MethodPost,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testCreatePropertyExpected

	actual, _, err := property.Create(
		context.Background(),
		testEnv.Client,
		property.CreateBody{
			ProductID:  th.TestProductID,
			PropertyName: "my.new.property.com",
		},
		property.CreateOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't create property")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.String {
		t.Errorf("expected string link to property, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestCreateProperty_Cloned(t *testing.T){
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/properties",
		RawRequest: testClonePropertyRawRequest,
		RawResponse: testClonePropertyRawResponse,
		Method:      http.MethodPost,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testClonePropertyExpected

	actual, _, err := property.Create(
		context.Background(),
		testEnv.Client,
		property.CreateBody{
			ProductID:  th.TestProductID,
			PropertyName: "my.new.property.com",
			CloneFrom: &property.CloneFrom{
				PropertyID: "prp_175780",
				Version: 2,
				VersionEtag: util.StringPtr("a9dfe78cf93090516bde891d009eaf57"),
				CopyHostnames: util.BoolPtr(true),
			},
		},
		property.CreateOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't create property")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.String {
		t.Errorf("expected string link to property, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestGetProperty(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/properties/" + testGetPropertyExpected.ID,
		RawResponse: testGetPropertyRawResponse,
		Method:      http.MethodGet,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testGetPropertyExpected

	actual, _, err := property.Get(
		context.Background(),
		testEnv.Client,
		expected.ID,
		property.GetOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get a property")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to property, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestGetByLinkProperty(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         testPropertyLink,
		RawResponse: testGetPropertyByLinkRawResponse,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testGetPropertyByLinkExpected

	actual, _, err := property.GetByLink(
		context.Background(),
		testEnv.Client,
		testPropertyLink)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get a property")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to property, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestListProperties(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/properties",
		RawResponse: testListPropertiesRawRequest,
		Method:      http.MethodGet,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testListPropertiesExpected

	actual, _, err := property.List(
		context.Background(),
		testEnv.Client,
		property.ListOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get a property")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected Slice of pointers to properties, but got %v", actualKind)
	}
	if len(actual) != 2 {
		t.Errorf("expected 2 properties, but got %d", len(actual))
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestDeleteProperty(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/properties/prp_112233",
		RawResponse: testDeletePropertyResponse,
		Method:      http.MethodDelete,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	// TODO: check response
	_, err := property.Delete(
		context.Background(),
		testEnv.Client,
		"prp_112233",
		property.DeleteOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get a property")
	}
}