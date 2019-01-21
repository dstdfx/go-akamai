package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/dstdfx/go-akamai/akamai/papi/v1/cpcode"
	th "github.com/dstdfx/go-akamai/internal/testhelper"
)

func TestCpCode_Create(t *testing.T){
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/cpcodes",
		RawRequest: testCreateCpCodeRawRequest,
		RawResponse: testCreateCpCodeRawResp,
		Method:      http.MethodPost,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testCreateCpCodeExpected

	actual, _, err := cpcode.Create(
		context.Background(),
		testEnv.Client,
		cpcode.CreateBody{
			ProductID:  th.TestProductID,
			CPCodeName: "SME WAA",
		},
		cpcode.CreateOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't create CP-Code")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.String {
		t.Errorf("expected String link to CP-Code, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestCpCode_Get(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/cpcodes/" + testGetCpCodeExpected.ID,
		RawResponse: testGetCpCodeRawResp,
		Method:      http.MethodGet,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testGetCpCodeExpected

	actual, _, err := cpcode.Get(
		context.Background(),
		testEnv.Client,
		testGetCpCodeExpected.ID,
		cpcode.GetOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get cp code")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to CP-Code, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestCpCode_GetByLink(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         testLinkToCpCode,
		RawResponse: testGetCpCodeByLinkRawResp,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testGetCpCodeByLinkExpected

	actual, _, err := cpcode.GetByLink(
		context.Background(),
		testEnv.Client,
		testLinkToCpCode)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get cp code")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to CP-Code, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestCpCode_List(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/cpcodes",
		RawResponse: testListCpCodesRawResp,
		Method:      http.MethodGet,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testListCpCodesExpected

	actual, _, err := cpcode.List(
		context.Background(),
		testEnv.Client,
		cpcode.ListOpts{
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
	if actualKind != reflect.Slice {
		t.Errorf("expected Slice of pointers to CP-Codes, but got %v", actualKind)
	}
	if len(actual) != 1 {
		t.Errorf("expected 1 CP-Code, but got %d", len(actual))
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}
