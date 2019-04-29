package testing

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/dstdfx/go-akamai/akamai/ccu/v3/cache"
	th "github.com/dstdfx/go-akamai/akamai/internal/testhelper"
)

func TestInvalidateByURL(t *testing.T){
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV3CCUClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/ccu/v3/invalidate/url/%s", cache.ProductionNetwork),
		RawResponse: testInvalidateCacheByURLRawResponse,
		RawRequest: testInvalidateCacheByURLRawRequest,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	})

	expected := testInvalidateCacheByURLExpected

	actual, _, err := cache.InvalidateByURL(
		context.Background(),
		testEnv.Client,
		cache.ProductionNetwork,
		cache.URLs{
			Objects: []string{
				"https://5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net/file.jpg",
				"https://5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net/video.mp4",
				"https://5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net/video2.mp4",
			},
		},)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't invalidate a cache")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to ccu response, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestInvalidateByURLWithHostname(t *testing.T){
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV3CCUClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/ccu/v3/invalidate/url/%s", cache.ProductionNetwork),
		RawResponse: testInvalidateCacheByURLWithHostnameRawResponse,
		RawRequest: testInvalidateCacheByURLWithHostnameRawRequest,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	})

	expected := testInvalidateCacheByURLWithHostnameExpected

	actual, _, err := cache.InvalidatebyURLWithHostname(
		context.Background(),
		testEnv.Client,
		cache.ProductionNetwork,
		cache.URLsWithHostname{
			Hostname: "5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net",
			Objects: []string{
				"/file.jpg",
				"/video.mp4",
				"/video2.mp4",
			},
		},)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't invalidate a cache")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to ccu response, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestInvalidateByCPCode(t *testing.T){
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV3CCUClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/ccu/v3/invalidate/cpcode/%s", cache.ProductionNetwork),
		RawResponse: testInvalidateCacheByCPCodeRawResponse,
		RawRequest: testInvalidateCacheByCPCodeRawRequest,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	})

	expected := testInvalidateCacheByCpCodeExpected

	actual, _, err := cache.InvalidateByCPCode(
		context.Background(),
		testEnv.Client,
		cache.ProductionNetwork,
		cache.CPCodes{
			Objects: []int{
				123,
				321,
			},
		},)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't invalidate a cache")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to ccu response, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestInvalidateByCacheTag(t *testing.T){
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV3CCUClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/ccu/v3/invalidate/tag/%s", cache.ProductionNetwork),
		RawResponse: testInvalidateCacheByCacheTagRawResponse,
		RawRequest: testInvalidateCacheByCacheTagRawRequest,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	})

	expected := testInvalidateCacheByCacheTagExpected

	actual, _, err := cache.InvalidateByCacheTag(
		context.Background(),
		testEnv.Client,
		cache.ProductionNetwork,
		cache.CacheTags{
			Objects: []string{
				"5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net",
			},
		},)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't invalidate a cache")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to ccu response, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestDeleteByURL(t *testing.T){
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV3CCUClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/ccu/v3/delete/url/%s", cache.ProductionNetwork),
		RawResponse: testDeleteCacheByURLRawResponse,
		RawRequest: testDeleteCacheByURLRawRequest,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	})

	expected := testDeleteCacheByURLExpected

	actual, _, err := cache.DeleteByURL(
		context.Background(),
		testEnv.Client,
		cache.ProductionNetwork,
		cache.URLs{
			Objects: []string{
				"https://5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net/file.jpg",
				"https://5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net/video.mp4",
				"https://5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net/video2.mp4",
			},
		},)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't delete a cache")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to ccu response, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestDeleteByURLWithHostname(t *testing.T){
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV3CCUClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/ccu/v3/delete/url/%s", cache.ProductionNetwork),
		RawResponse: testDeleteCacheByURLWithHostnameRawResponse,
		RawRequest: testDeleteCacheByURLWithHostnameRawRequest,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	})

	expected := testDeleteCacheByURLWithHostnameExpected

	actual, _, err := cache.DeleteByURLWithHostname(
		context.Background(),
		testEnv.Client,
		cache.ProductionNetwork,
		cache.URLsWithHostname{
			Hostname: "5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net",
			Objects: []string{
				"/file.jpg",
				"/video.mp4",
				"/video2.mp4",
			},
		},)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't delete a cache")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to ccu response, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestDeleteByCPCode(t *testing.T){
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV3CCUClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/ccu/v3/delete/cpcode/%s", cache.ProductionNetwork),
		RawResponse: testDeleteCacheByCPCodeRawResponse,
		RawRequest: testDeleteCacheByCPCodeRawRequest,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	})

	expected := testDeleteCacheByCpCodeExpected

	actual, _, err := cache.DeleteByCPCode(
		context.Background(),
		testEnv.Client,
		cache.ProductionNetwork,
		cache.CPCodes{
			Objects: []int{
				123,
				321,
			},
		},)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't delete a cache")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to ccu response, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestDeleteByCacheTag(t *testing.T){
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV3CCUClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/ccu/v3/delete/tag/%s", cache.ProductionNetwork),
		RawResponse: testDeleteCacheByCacheTagRawResponse,
		RawRequest: testDeleteCacheByCacheTagRawRequest,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	})

	expected := testDeleteCacheByCacheTagExpected

	actual, _, err := cache.DeleteByCacheTag(
		context.Background(),
		testEnv.Client,
		cache.ProductionNetwork,
		cache.CacheTags{
			Objects: []string{
				"5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net",
			},
		},)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't delete a cache")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to ccu response, but got %v", actualKind)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

