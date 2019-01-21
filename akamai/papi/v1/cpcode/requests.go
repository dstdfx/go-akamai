package cpcode

import (
	"context"
	"net/http"
	"strings"

	"github.com/dstdfx/go-akamai/akamai"
	"github.com/dstdfx/go-akamai/internal/util"
)

const cpCodeURL = "/cpcodes"

// Create requests a creation of the CP-Code.
func Create(ctx context.Context,
	client *akamai.ServiceClient,
	body CreateBody,
	opts CreateOpts) (string, *akamai.ResponseResult, error) {

	url := cpCodeURL
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
		CPCodeLink string `json:"cpcodeLink"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return "", respResult, respResult.Err
	}

	return v.CPCodeLink, respResult, nil
}

// Get requests a getting of the specific CP-Code.
func Get(ctx context.Context,
	client *akamai.ServiceClient,
	cpCodeID string,
	opts GetOpts) (*CPCode, *akamai.ResponseResult, error){

	url := strings.Join([]string{cpCodeURL, cpCodeID}, "/")
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
		CPCodes struct{
			Items []*CPCode `json:"items"`
		} `json:"cpcodes"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, respResult, err
	}

	return v.CPCodes.Items[0], respResult, nil
}

// GetByLink requests a getting of the specific CP-Code by the link.
func GetByLink(ctx context.Context,
	client *akamai.ServiceClient,
	link string) (*CPCode, *akamai.ResponseResult, error){

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
		CPCodes struct{
			Items []*CPCode `json:"items"`
		} `json:"cpcodes"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, respResult, respResult.Err
	}

	return v.CPCodes.Items[0], respResult, nil
}

// List returns a list of available CP-Codes.
func List(ctx context.Context,
	client *akamai.ServiceClient,
	opts ListOpts) ([]*CPCode, *akamai.ResponseResult, error) {

	url := cpCodeURL
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
		CPCodes struct{
			Items []*CPCode `json:"items"`
		} `json:"cpcodes"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, respResult, err
	}

	return v.CPCodes.Items, respResult, nil
}
