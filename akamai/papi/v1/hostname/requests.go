package hostname

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dstdfx/go-akamai/akamai"
	"github.com/dstdfx/go-akamai/akamai/internal/util"
)

const hostnameURL  = "/properties/%s/versions/%d/hostnames"

// List returns list of all the hostnames assigned to a property version.
func List(ctx context.Context,
	client *akamai.ServiceClient,
	propertyID string,
	propertyVersion int,
	opts ListOpts) ([]*Hostname, *akamai.ResponseResult, error) {

	url := fmt.Sprintf(hostnameURL, propertyID, propertyVersion)
	queryParams, err := util.BuildQueryParameters(opts)
	if err != nil {
		return nil, nil, err
	}

	if queryParams != "" {
		url = strings.Join([]string{url, queryParams}, "?")
	}

	respResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if respResult.Err != nil {
		return nil, respResult, respResult.Err
	}

	v := struct {
		Hostnames struct {
			Items []*Hostname `json:"items"`
		} `json:"hostnames"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, nil, err
	}

	return v.Hostnames.Items, respResult, nil
}

// Update modifies the set of hostname entries for a property version.
// For each hostname entry, the cnameFrom is required along with either the cnameTo or edgeHostnameID.
func Update(ctx context.Context,
	client *akamai.ServiceClient,
	propertyID string,
	propertyVersion int,
	body UpdateBody,
	opts UpdateOpts) ([]*Hostname, *akamai.ResponseResult, error) {

	url := fmt.Sprintf(hostnameURL, propertyID, propertyVersion)
	queryParams, err := util.BuildQueryParameters(opts)
	if err != nil {
		return nil, nil, err
	}

	if queryParams != "" {
		url = strings.Join([]string{url, queryParams}, "?")
	}

	respResult, err := client.DoRequest(ctx, http.MethodPut, url, body.Hostnames)
	if err != nil {
		return nil, nil, err
	}

	if respResult.Err != nil {
		return nil, respResult, respResult.Err
	}

	v := struct {
		Hostnames struct {
			Items []*Hostname `json:"items"`
		} `json:"hostnames"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, nil, err
	}

	return v.Hostnames.Items, respResult, nil
}
