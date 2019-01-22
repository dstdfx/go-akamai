package version

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dstdfx/go-akamai/akamai"
	"github.com/dstdfx/go-akamai/akamai/internal/util"
)

const versionURL = "/properties/%s/versions"

// Create requests a creation of a new property version based on any previous version.
// All data from the createFromVersion populates the new version, including its
// rules and hostnames. Specify createFromVersionEtag if you want to ensure
// you are creating from a version that has not changed since you fetched it.
func Create(ctx context.Context,
	client *akamai.ServiceClient,
	propertyID string,
	body CreateBody,
	opts CreateOpts) (string, *akamai.ResponseResult, error) {

	url := fmt.Sprintf(versionURL, propertyID)
	queryParams, err := util.BuildQueryParameters(opts)
	if err != nil {
		return "", nil, err
	}

	if queryParams != "" {
		url = strings.Join([]string{url, queryParams}, "?")
	}

	respResult, err := client.DoRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return "", nil, err
	}
	if respResult.Err != nil {
		return "", respResult, respResult.Err
	}

	v := struct {
		PropertyVersionLink string `json:"versionLink"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return "", nil, err
	}

	return v.PropertyVersionLink, respResult, nil
}

// Get requests a getting of the specific version of the property.
func Get(ctx context.Context,
	client *akamai.ServiceClient,
	propertyVersion int,
	propertyID string,
	opts GetOpts) (*Version, *akamai.ResponseResult, error){

	url := fmt.Sprintf(versionURL + "/%d", propertyID, propertyVersion)
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
		Versions struct{
			Items []*Version `json:"items"`
		} `json:"versions"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, nil, err
	}

	return v.Versions.Items[0], respResult, nil
}

// GetByLink requests a getting of the specific version of the property by link.
func GetByLink(ctx context.Context,
	client *akamai.ServiceClient,
	link string) (*Version, *akamai.ResponseResult, error){

	if strings.HasPrefix(link, "/papi/v0") {
		link = strings.TrimPrefix(link, "/papi/v0")
	}

	if strings.HasPrefix(link, "/papi/v1") {
		link = strings.TrimPrefix(link, "/papi/v1")
	}

	respResult, err := client.DoRequest(ctx, http.MethodGet, link, nil)
	if err != nil {
		return nil, nil, err
	}
	if respResult.Err != nil {
		return nil, respResult, respResult.Err
	}

	v := struct {
		Versions struct{
			Items []*Version `json:"items"`
		} `json:"versions"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, nil, err
	}

	return v.Versions.Items[0], respResult, nil
}

// GetLatest requests a getting of the latest version of the property.
//
// Returns either the latest property version overall, or the latest one active on
// the production or staging networks. By default, the response yields the property
// version with the highest number. Specifying an activatedOn of STAGING or PRODUCTION
// yields the version that’s currently active on either network.
func GetLatest(ctx context.Context,
	client *akamai.ServiceClient,
	propertyID string,
	opts GetOpts) (*Version, *akamai.ResponseResult, error){

	url := fmt.Sprintf(versionURL + "/latest", propertyID)
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
		Versions struct{
			Items []*Version `json:"items"`
		} `json:"versions"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, nil, err
	}

	return v.Versions.Items[0], respResult, nil
}

// List returns a list of the property's versions.
func List(ctx context.Context,
	client *akamai.ServiceClient,
	propertyID string,
	opts ListOpts) ([]*Version, *akamai.ResponseResult, error){

	url := fmt.Sprintf(versionURL, propertyID)
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
		Versions struct{
			Items []*Version `json:"items"`
		} `json:"versions"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, nil, err
	}

	return v.Versions.Items, respResult, nil
}
