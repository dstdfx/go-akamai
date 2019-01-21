package group

import (
	"context"
	"net/http"

	"github.com/dstdfx/go-akamai/akamai"
)

// List returns list of available groups.
func List(ctx context.Context, client *akamai.ServiceClient) ([]*Group, *akamai.ResponseResult, error) {
	respResult, err := client.DoRequest(ctx, http.MethodGet, "/groups", nil)
	if err != nil {
		return nil, nil, err
	}
	if respResult.Err != nil {
		return nil, respResult, respResult.Err
	}

	v := struct {
		Groups struct{
			Items []*Group `json:"items"`
		} `json:"groups"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, nil, err
	}

	return v.Groups.Items, respResult, nil
}
