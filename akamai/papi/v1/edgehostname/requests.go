package edgehostname

import (
	"context"
	"net/http"
	"strings"

	"github.com/dstdfx/go-akamai/akamai"
	"github.com/dstdfx/go-akamai/akamai/internal/util"
)

const edgeHostnameURL = "/edgehostnames"

// Get requests a getting of the specific edge-hostname.
func Get(ctx context.Context,
	client *akamai.ServiceClient,
	edgeHostnameID string,
	opts GetOpts) (*EdgeHostname, *akamai.ResponseResult, error) {

	url := strings.Join([]string{edgeHostnameURL, edgeHostnameID}, "/")
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
		EdgeHostnames struct {
			Items []*EdgeHostname `json:"items"`
		} `json:"edgeHostnames"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, nil, err
	}

	return v.EdgeHostnames.Items[0], respResult, nil
}

// GetByLink requests a getting of the specific edge-hostname by link.
func GetByLink(ctx context.Context,
	client *akamai.ServiceClient,
	link string) (*EdgeHostname, *akamai.ResponseResult, error) {

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
		EdgeHostnames struct {
			Items []*EdgeHostname `json:"items"`
		} `json:"edgeHostnames"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, nil, err
	}

	return v.EdgeHostnames.Items[0], respResult, nil
}

// Create requests a creation of the edge-hostname.
func Create(ctx context.Context,
	client *akamai.ServiceClient,
	body CreateBody,
	opts CreateOpts) (string, *akamai.ResponseResult, error) {

	url := edgeHostnameURL
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
		EdgeHostnameLink string `json:"edgeHostnameLink"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return "", nil, err
	}

	return v.EdgeHostnameLink, respResult, nil
}

// List returns a list of edge-hostnames available under a contract.
func List(ctx context.Context,
	client *akamai.ServiceClient,
	opts ListOpts) ([]*EdgeHostname, *akamai.ResponseResult, error) {

	url := edgeHostnameURL
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
		EdgeHostnames struct {
			Items []*EdgeHostname `json:"items"`
		} `json:"edgeHostnames"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, nil, err
	}

	return v.EdgeHostnames.Items, respResult, nil
}
