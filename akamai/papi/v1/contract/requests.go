package contract

import (
	"context"
	"net/http"

	"github.com/dstdfx/go-akamai/akamai"
)

const contractURL = "/contracts"

// List returns list of available contracts.
func List(ctx context.Context, client *akamai.ServiceClient) ([]*Contract, *akamai.ResponseResult, error){

	respResult, err := client.DoRequest(ctx, http.MethodGet, contractURL, nil)
	if err != nil {
		return nil, nil, err
	}
	if respResult.Err != nil {
		return nil, respResult, respResult.Err
	}

	v := struct {
		Contracts struct{
			Items []*Contract `json:"items"`
		} `json:"contracts"`
	}{}

	if err = respResult.ExtractResult(&v); err != nil {
		return nil, respResult, err
	}

	return v.Contracts.Items, respResult, nil
}

