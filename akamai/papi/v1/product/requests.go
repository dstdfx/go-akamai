package product

import (
	"context"
	"net/http"
	"strings"

	"github.com/dstdfx/go-akamai/akamai"
	"github.com/dstdfx/go-akamai/akamai/internal/util"
)

const productURL = "/products"

// List lists the set of products that are available under a given contract.
func List(ctx context.Context,
	client *akamai.ServiceClient,
	opts ListOpts) ([]*Product, *akamai.ResponseResult, error) {

	url := productURL
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
		Products struct {
			Items []*Product `json:"items"`
		} `json:"products"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, nil, err
	}

	return v.Products.Items, respResult, nil
}
