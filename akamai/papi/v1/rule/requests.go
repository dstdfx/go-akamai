package rule

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dstdfx/go-akamai/akamai"
	"github.com/dstdfx/go-akamai/akamai/internal/util"
)

const ruleURL = "/properties/%s/versions/%d/rules"

// Get requests a getting of the entire rule tree for a property version.
func Get(ctx context.Context,
	client *akamai.ServiceClient,
	propertyID string,
	propertyVersion int,
	ruleFormat string,
	opts GetOpts) (*Tree, *akamai.ResponseResult, error) {

	url := fmt.Sprintf(ruleURL, propertyID, propertyVersion)

	queryParams, err := util.BuildQueryParameters(opts)
	if err != nil {
		return nil, nil, err
	}

	if queryParams != "" {
		url = strings.Join([]string{url, queryParams}, "?")
	}

	// Re-write Accept header to be able to use specific rule format.
	// Details: https://developer.akamai.com/api/core_features/property_manager/v1.html#updaterf
	headers := http.Header{
		"Accept": []string{
			fmt.Sprintf("application/vnd.akamai.papirules.%s+json", ruleFormat),
		},
	}

	respResult, err := client.DoRequestWithCustomHeaders(ctx,
		http.MethodGet,
		url,
		headers, nil)
	if err != nil {
		return nil, nil, err
	}
	if respResult.Err != nil {
		return nil, respResult, respResult.Err
	}

	v := Tree{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, nil, err
	}

	return &v, respResult, nil
}

// Update requests a updating of the property's rule tree.
func Update(ctx context.Context,
	client *akamai.ServiceClient,
	propertyID string,
	propertyVersion int,
	ruleFormat string,
	body UpdateBody,
	opts UpdateOpts) (*Tree, *akamai.ResponseResult, error){

	url := fmt.Sprintf(ruleURL, propertyID, propertyVersion)

	queryParams, err := util.BuildQueryParameters(opts)
	if err != nil {
		return nil, nil, err
	}

	if queryParams != "" {
		url = strings.Join([]string{url, queryParams}, "?")
	}

	// Re-write Content-type header to be able to use specific rule format.
	// Details: https://developer.akamai.com/api/core_features/property_manager/v1.html#updaterf
	headers := http.Header{
		"Content-Type": []string{
			fmt.Sprintf("application/vnd.akamai.papirules.%s+json", ruleFormat),
		},
	}

	respResult, err := client.DoRequestWithCustomHeaders(ctx,
		http.MethodPut,
		url,
		headers,
		body)
	if err != nil {
		return nil, nil, err
	}
	if respResult.Err != nil {
		return nil, respResult, respResult.Err
	}

	v := Tree{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, nil, err
	}

	return &v, respResult, nil
}

// GetDigest retrieves the rule tree’s Etag without the rule tree object.
// Etag will be presented in response headers.
func GetDigest(ctx context.Context,
	client *akamai.ServiceClient,
	propertyID string,
	propertyVersion int,
	opts GetOpts) (*akamai.ResponseResult, error) {

	url := fmt.Sprintf(ruleURL, propertyID, propertyVersion)

	queryParams, err := util.BuildQueryParameters(opts)
	if err != nil {
		return nil, err
	}

	if queryParams != "" {
		url = strings.Join([]string{url, queryParams}, "?")
	}

	respResult, err := client.DoRequest(ctx, http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}
	if respResult.Err != nil {
		return respResult, respResult.Err
	}

	return respResult, nil
}