package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/dstdfx/go-akamai/akamai/papi/v1/group"
	th "github.com/dstdfx/go-akamai/internal/testhelper"

)

func TestListGroups(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/groups",
		RawResponse: testListGroupsRawResp,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testListGroupsExpected

	actual, _, err := group.List(context.Background(), testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get groups")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected Slice of pointers to groups, but got %v", actualKind)
	}
	if len(actual) != 2 {
		t.Errorf("expected 2 groups, but got %d", len(actual))
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}
