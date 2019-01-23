package activation

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dstdfx/go-akamai/akamai"
	"github.com/dstdfx/go-akamai/akamai/internal/util"
)

const activationURL = "/properties/%s/activations"


// List returns a list of activations for the property.
func List(ctx context.Context,
	client *akamai.ServiceClient,
	propertyID string,
	opts ListOpts) ([]*Activation, *akamai.ResponseResult, error) {

	url := fmt.Sprintf(activationURL, propertyID)
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
		Activations struct {
			Items []*Activation `json:"items"`
		} `json:"activations"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, nil, err
	}

	return v.Activations.Items, respResult, nil
}


// Get requests a getting of the specific property's activation.
func Get(ctx context.Context,
	client *akamai.ServiceClient,
	propertyID string,
	activationID string,
	opts GetOpts) (*Activation, *akamai.ResponseResult, error) {

	url := fmt.Sprintf(activationURL + "/%s",
		propertyID,
		activationID)
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
		Activations struct {
			Items []*Activation `json:"items"`
		} `json:"activations"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, nil, err
	}

	return v.Activations.Items[0], respResult, nil
}

// GetByLink requests a getting of the specific property's activation by link.
func GetByLink(ctx context.Context,
	client *akamai.ServiceClient,
	link string) (*Activation, *akamai.ResponseResult, error) {

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
		Activations struct {
			Items []*Activation `json:"items"`
		} `json:"activations"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, nil, err
	}

	return v.Activations.Items[0], respResult, nil
}

// Create requests a creation of the property's activation.
func Create(ctx context.Context,
	client *akamai.ServiceClient,
	propertyID string,
	body CreateBody,
	opts CreateOpts) (string, *akamai.ResponseResult, error) {

	url := fmt.Sprintf(activationURL, propertyID)
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
		ActivationLink string `json:"activationLink"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return "", nil, err
	}

	return v.ActivationLink, respResult, nil
}
