package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	th "github.com/dstdfx/go-akamai/akamai/internal/testhelper"
	"github.com/dstdfx/go-akamai/akamai/papi/v1/product"
)

func TestListProducts(t *testing.T){
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/products",
		RawResponse: testListProductsRawResponse,
		Method:      http.MethodGet,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testListProductsExpected

	actual, _, err := product.List(
		context.Background(),
		testEnv.Client,
		product.ListOpts{ContractID: th.TestContractID})
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get list of products")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to products, but got %v", actualKind)
	}
	if len(actual) != 1 {
		t.Errorf("expected 1 product, but got %d", len(actual))
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}
