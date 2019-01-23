package testing

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	th "github.com/dstdfx/go-akamai/akamai/internal/testhelper"
	"github.com/dstdfx/go-akamai/akamai/papi/v1/rule"
	"github.com/stretchr/testify/assert"
)

func TestGetRuleTree(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithoutBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/properties/prp_511225/versions/1/rules",
		RawResponse: testGetRuleTreeRawResponse,
		Method:      http.MethodGet,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testGetRuleTreeExpected

	actual, _, err := rule.Get(
		context.Background(),
		testEnv.Client,
		"prp_511225",
		1,
		"v2018-09-12",
		rule.GetOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get a rule tree")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to rule tree, but got %v", actualKind)
	}

	assert.Equal(t, expected.Warnings, actual.Warnings)
	assert.Equal(t, expected.Errors, actual.Errors)
	assert.Equal(t, expected.RuleFormat, actual.RuleFormat)
	assert.Equal(t, expected.Rule.Name, actual.Rule.Name)
	assert.Equal(t, expected.Rule.Options, actual.Rule.Options)

	assert.Equal(t, len(expected.Rule.Behaviors), len(actual.Rule.Behaviors))
	assert.Equal(t, len(expected.Rule.Criteria), len(actual.Rule.Criteria))
	assert.Equal(t, len(expected.Rule.Children), len(actual.Rule.Children))

	// TODO: https://github.com/stretchr/testify/pull/168

	expectedRaw, _ := json.Marshal(expected)
	actualRaw, _ := json.Marshal(actual)

	assert.Equal(t, expectedRaw, actualRaw)
}

func TestUpdateRuleTree_Error(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/properties/prp_511226/versions/2/rules",
		RawResponse: testUpdateRuleTreeWithErrorsRawResponse,
		RawRequest: testUpdateRuleTreeWithErrorsRawRequest,
		Method:      http.MethodPut,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testUpdateRuleTreeWithErrorsExpected

	actual, _, err := rule.Update(
		context.Background(),
		testEnv.Client,
		"prp_511226",
		2,
		"v2018-09-12",
		rule.UpdateBody{
			Rule: &rule.Rule{
				Name: "default",
				Options: &rule.TopLevelRuleOptions{},
				Behaviors: []*rule.Behavior{
					{
						Name: "report",
						Options: rule.OptionValue{
							"logAcceptLanguage": false,
							"logCookies": "OFF",
							"logCustomLogField": false,
							"logHost": false,
							"logReferer": false,
							"logUserAgent": true,
						},
					},
					{
						Name: "caching",
						Options: rule.OptionValue{
							"behavior": "MAX_AGE",
							"mustRevalidate": false,
							"ttl": "7d",
						},
					},
					{
						Name: "tieredDistribution",
						Options: rule.OptionValue{
							"enabled": true,
							"tieredDistributionMap": "CH2",
						},
					},
					{
						Name: "cacheKeyQueryParams",
						Options: rule.OptionValue{
							"behavior": "IGNORE_ALL",
						},
					},
					{
						Name: "netSession",
						Options: rule.OptionValue{
							"enabled": false,
						},
					},
				},
				Children: []*rule.Rule{
					{
						Name: "Large File Optimization",
						CriteriaMustSatisfy: rule.CriteriaMustSatisfyAll,
						Behaviors: []*rule.Behavior{
							{
								Name: "largeFileOptimization",
								Options: rule.OptionValue{
									"enablePartialObjectCaching": "PARTIAL_OBJECT_CACHING",
									"enabled": true,
									"maximumSize": "16GB",
									"minimumSize": "100MB",
									"useVersioning": true,
								},
							},
						},
						Criteria: []*rule.Criteria{
							{
								Name: "fileExtension",
								Options: rule.OptionValue{
									"matchCaseSensitive": false,
									"matchOperator": "IS_ONE_OF",
									"values": []string{
										"exe",
										"bz2",
										"dmg",
										"gz",
										"iso",
										"mov",
										"pkg",
										"tar",
										"tgz",
										"wmv",
										"wma",
										"zip",
										"webp",
										"jxr",
										"hdp",
										"wdp",
									},
								},

							},
						},
						Children: []*rule.Rule{},
					},
					{
						Name: "Content Compression",
						CriteriaMustSatisfy: rule.CriteriaMustSatisfyAll,
						Behaviors: []*rule.Behavior{
							{
								Name: "gzipResponse",
								Options: rule.OptionValue{
									"behavior": "ALWAYS",
								},
							},
						},
						Criteria: []*rule.Criteria{
							{
								Name: "contentType",
								Options: rule.OptionValue{
									"matchCaseSensitive": false,
									"matchOperator": "IS_ONE_OF",
									"matchWildcard": true,
									"values": []string{
										"text/html*",
										"text/css*",
										"application/x-javascript*",
									},
								},
							},
						},
						Children: []*rule.Rule{},
					},
				},
			},
		},
		rule.UpdateOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't get a rule tree updated")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to the updated rule tree, but got %v", actualKind)
	}

	assert.Equal(t, expected.Warnings, actual.Warnings)
	assert.Equal(t, expected.Errors, actual.Errors)
	assert.Equal(t, expected.RuleFormat, actual.RuleFormat)
	assert.Equal(t, expected.Rule.Name, actual.Rule.Name)
	assert.Equal(t, expected.Rule.Options, actual.Rule.Options)

	assert.Equal(t, len(expected.Rule.Behaviors), len(actual.Rule.Behaviors))
	assert.Equal(t, len(expected.Rule.Criteria), len(actual.Rule.Criteria))
	assert.Equal(t, len(expected.Rule.Children), len(actual.Rule.Children))

	// TODO: https://github.com/stretchr/testify/pull/168
	// TODO: fix issue with comparing empty interfaces

	expectedRaw, _ := json.Marshal(expected)
	actualRaw, _ := json.Marshal(actual)

	assert.Equal(t, expectedRaw, actualRaw)
}

func TestUpdateRuleTree(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testEnv.NewTestV1PAPIClient()

	th.HandleReqWithBody(t, &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/papi/v1/properties/prp_511225/versions/1/rules",
		RawResponse: testUpdateRuleTreeRawResponse,
		RawRequest: testUpdateRuleTreeRawRequest,
		Method:      http.MethodPut,
		QueryParams: map[string]string{
			"contractId": th.TestContractID,
			"groupId": th.TestGroupID,
		},
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	expected := testUpdateRuleTreeExpected

	actual, _, err := rule.Update(
		context.Background(),
		testEnv.Client,
		"prp_511225",
		1,
		"v2018-09-12",
		rule.UpdateBody{
			Rule: &rule.Rule{
				Name: "default",
				Options: &rule.TopLevelRuleOptions{},
				Behaviors: []*rule.Behavior{
					{
						Name: "origin",
						Options: rule.OptionValue{
							"cacheKeyHostname": "REQUEST_HOST_HEADER",
							"compress": true,
							"enableTrueClientIp": true,
							"forwardHostHeader": "REQUEST_HOST_HEADER",
							"hostname": "test.dstdfx.space",
							"httpPort": 80,
							"httpsPort": 443,
							"originCertificate": "",
							"originSni": true,
							"originType": "CUSTOMER",
							"ports": "",
							"trueClientIpClientSetting": false,
							"trueClientIpHeader": "True-Client-IP",
							"verificationMode": "PLATFORM_SETTINGS",
						},
					},
					{
						Name: "cpCode",
						Options: rule.OptionValue{
							"value": rule.OptionValue{
								"id": 688480,
							},
						},
					},
					{
						Name: "report",
						Options: rule.OptionValue{
							"logAcceptLanguage": false,
							"logCookies": "OFF",
							"logCustomLogField": false,
							"logHost": false,
							"logReferer": false,
							"logUserAgent": true,
						},
					},
					{
						Name: "caching",
						Options: rule.OptionValue{
							"behavior": "MAX_AGE",
							"mustRevalidate": false,
							"ttl": "7d",
						},
					},
					{
						Name: "tieredDistribution",
						Options: rule.OptionValue{
							"enabled": true,
							"tieredDistributionMap": "CH2",
						},
					},
					{
						Name: "cacheKeyQueryParams",
						Options: rule.OptionValue{
							"behavior": "IGNORE_ALL",
						},
					},
					{
						Name: "netSession",
						Options: rule.OptionValue{
							"enabled": false,
						},
					},
				},
				Children: []*rule.Rule{
					{
						Name: "Large File Optimization",
						CriteriaMustSatisfy: rule.CriteriaMustSatisfyAll,
						Behaviors: []*rule.Behavior{
							{
								Name: "largeFileOptimization",
								Options: rule.OptionValue{
									"enablePartialObjectCaching": "PARTIAL_OBJECT_CACHING",
									"enabled": true,
									"maximumSize": "16GB",
									"minimumSize": "100MB",
									"useVersioning": true,
								},
							},
						},
						Criteria: []*rule.Criteria{
							{
								Name: "fileExtension",
								Options: rule.OptionValue{
									"matchCaseSensitive": false,
									"matchOperator": "IS_ONE_OF",
									"values": []string{
										"exe",
										"bz2",
										"dmg",
										"gz",
										"iso",
										"mov",
										"pkg",
										"tar",
										"tgz",
										"wmv",
										"wma",
										"zip",
										"webp",
										"jxr",
										"hdp",
										"wdp",
									},
								},

							},
						},
						Children: []*rule.Rule{},
					},
					{
						Name: "Content Compression",
						CriteriaMustSatisfy: rule.CriteriaMustSatisfyAll,
						Behaviors: []*rule.Behavior{
							{
								Name: "gzipResponse",
								Options: rule.OptionValue{
									"behavior": "ALWAYS",
								},
							},
						},
						Criteria: []*rule.Criteria{
							{
								Name: "contentType",
								Options: rule.OptionValue{
									"matchCaseSensitive": false,
									"matchOperator": "IS_ONE_OF",
									"matchWildcard": true,
									"values": []string{
										"text/html*",
										"text/css*",
										"application/x-javascript*",
									},
								},
							},
						},
						Children: []*rule.Rule{},
					},
				},
			},
		},
		rule.UpdateOpts{
			ContractID: th.TestContractID,
			GroupID: th.TestGroupID,
		})
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't get a rule tree updated")
	}

	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Ptr {
		t.Errorf("expected pointer to the updated rule tree, but got %v", actualKind)
	}

	assert.Equal(t, expected.Warnings, actual.Warnings)
	assert.Equal(t, expected.Errors, actual.Errors)
	assert.Equal(t, expected.RuleFormat, actual.RuleFormat)
	assert.Equal(t, expected.Rule.Name, actual.Rule.Name)
	assert.Equal(t, expected.Rule.Options, actual.Rule.Options)

	assert.Equal(t, len(expected.Rule.Behaviors), len(actual.Rule.Behaviors))
	assert.Equal(t, len(expected.Rule.Criteria), len(actual.Rule.Criteria))
	assert.Equal(t, len(expected.Rule.Children), len(actual.Rule.Children))

	// TODO: https://github.com/stretchr/testify/pull/168

	expectedRaw, _ := json.Marshal(expected)
	actualRaw, _ := json.Marshal(actual)

	assert.Equal(t, expectedRaw, actualRaw)
}
