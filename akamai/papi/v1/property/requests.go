package property

import (
	"context"
	"net/http"
	"strings"

	"github.com/dstdfx/go-akamai/akamai"
	"github.com/dstdfx/go-akamai/akamai/internal/util"
)

const propertyURL = "/properties"

// List returns list of the properties.
func List(ctx context.Context,
	client *akamai.ServiceClient,
	opts ListOpts) ([]*Property, *akamai.ResponseResult, error) {

	url := propertyURL
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
		Properties struct{
			Items []*Property `json:"items"`
		} `json:"properties"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, respResult, err
	}

	return v.Properties.Items, respResult, nil
}

// Get requests a getting of the specific property.
func Get(ctx context.Context,
	client *akamai.ServiceClient,
	propertyID string,
	opts GetOpts) (*Property, *akamai.ResponseResult, error) {

	url := strings.Join([]string{propertyURL, propertyID}, "/")
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
		Properties struct{
			Items []*Property `json:"items"`
		} `json:"properties"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, respResult, err
	}

	return v.Properties.Items[0], respResult, nil
}

// GetByLink requests a getting of the specific property by link.
func GetByLink(ctx context.Context,
	client *akamai.ServiceClient,
	link string) (*Property, *akamai.ResponseResult, error) {

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
		Properties struct{
			Items []*Property `json:"items"`
		} `json:"properties"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, respResult, err
	}

	return v.Properties.Items[0], respResult, nil
}

// Create requests a creation of the property.
func Create(ctx context.Context,
	client *akamai.ServiceClient,
	body CreateBody,
	opts CreateOpts) (string, *akamai.ResponseResult, error) {

	url := propertyURL
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
		PropertyLink string `json:"propertyLink"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return "", respResult, err
	}

	return v.PropertyLink, respResult, nil
}

// Delete requests a deletion of the specific property.
func Delete(ctx context.Context,
	client *akamai.ServiceClient,
	propertyID string,
	opts DeleteOpts) (*akamai.ResponseResult, error) {

	url := strings.Join([]string{propertyURL, propertyID}, "/")
	queryParams, err := util.BuildQueryParameters(opts)
	if err != nil {
		return nil, err
	}

	if queryParams != "" {
		url = strings.Join([]string{url, queryParams}, "?")
	}

	respResult, err := client.DoRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	if respResult.Err != nil {
		return respResult, respResult.Err
	}

	return respResult, nil
}
