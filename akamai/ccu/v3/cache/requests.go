package cache

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dstdfx/go-akamai/akamai"
)

const (
	invalidateBaseURL = "/invalidate"
	invalidateByURL = invalidateBaseURL + "/url/%s"
	invalidateByCPCode = invalidateBaseURL + "/cpcode/%s"
	invalidateByCacheTag = invalidateBaseURL + "/tag/%s"

	deleteBaseURL = "/delete"
	deleteByURL = deleteBaseURL + "/url/%s"
	deleteByCPCode = deleteBaseURL + "/cpcode/%s"
	deleteByCacheTag = deleteBaseURL + "/tag/%s"
)

// InvalidateByURL invalidates content on the selected URL for the selected network.
// You should consider invalidating content by default. This keeps each object in cache
// until the version on your origin server is newer. Deletion retrieves the object
// regardless, which can dramatically increase the load on your origin server and would
// prevent Akamai from serving the old content if your origin is unreachable.
func InvalidateByURL(ctx context.Context,
	client *akamai.ServiceClient,
	network Network, urls URLs) (*CCUResponse, *akamai.ResponseResult, error){

	url := fmt.Sprintf(invalidateByURL, network)

	respResult, err := client.DoRequest(ctx, http.MethodPost, url, urls)
	if err != nil {
		return nil, nil, err
	}
	if respResult.Err != nil {
		return nil, respResult, respResult.Err
	}

	ccuResp := CCUResponse{}

	if err = respResult.ExtractResult(&ccuResp); err != nil {
		return nil, nil, err
	}

	return &ccuResp, respResult, nil
}

// InvalidatebyURLWithHostname invalidates content on the selected URL related to the given hostname
// for the selected network.
// You should consider invalidating content by default. This keeps each object in cache
// until the version on your origin server is newer. Deletion retrieves the object
// regardless, which can dramatically increase the load on your origin server and would
// prevent Akamai from serving the old content if your origin is unreachable.
func InvalidatebyURLWithHostname(ctx context.Context,
	client *akamai.ServiceClient,
	network Network, urls URLsWithHostname) (*CCUResponse, *akamai.ResponseResult, error) {

	url := fmt.Sprintf(invalidateByURL, network)

	respResult, err := client.DoRequest(ctx, http.MethodPost, url, urls)
	if err != nil {
		return nil, nil, err
	}
	if respResult.Err != nil {
		return nil, respResult, respResult.Err
	}

	ccuResp := CCUResponse{}

	if err = respResult.ExtractResult(&ccuResp); err != nil {
		return nil, nil, err
	}

	return &ccuResp, respResult, nil
}


// InvalidateByCPCode invalidates content on the selected CP code for the selected network.
// You should consider invalidating content by default. This keeps each object in cache until
// the version on your origin server is newer. Deletion retrieves the object regardless,
// which can dramatically increase the load on your origin server and would prevent
// Akamai from serving the old content if your origin is unreachable.
func InvalidateByCPCode(ctx context.Context,
	client *akamai.ServiceClient,
	network Network, cpCodes CPCodes) (*CCUResponse, *akamai.ResponseResult, error){

	url := fmt.Sprintf(invalidateByCPCode, network)

	respResult, err := client.DoRequest(ctx, http.MethodPost, url, cpCodes)
	if err != nil {
		return nil, nil, err
	}
	if respResult.Err != nil {
		return nil, respResult, respResult.Err
	}

	ccuResp := CCUResponse{}

	if err = respResult.ExtractResult(&ccuResp); err != nil {
		return nil, nil, err
	}

	return &ccuResp, respResult, nil
}


// InvalidateByCacheTag invalidates content on the selected set of cache tags for the selected network.
// You should consider invalidating content by default. This keeps each object in cache until
// the version on your origin server is newer. Deletion retrieves the object regardless, which
// can dramatically increase the load on your origin server and would prevent Akamai from serving
// the old content if your origin is unreachable.
func InvalidateByCacheTag(ctx context.Context,
	client *akamai.ServiceClient,
	network Network, cacheTags CacheTags) (*CCUResponse, *akamai.ResponseResult, error) {

	url := fmt.Sprintf(invalidateByCacheTag, network)

	respResult, err := client.DoRequest(ctx, http.MethodPost, url, cacheTags)
	if err != nil {
		return nil, nil, err
	}
	if respResult.Err != nil {
		return nil, respResult, respResult.Err
	}

	ccuResp := CCUResponse{}

	if err = respResult.ExtractResult(&ccuResp); err != nil {
		return nil, nil, err
	}

	return &ccuResp, respResult, nil
}


// DeleteByURL deletes content on the selected URL for the selected network.
func DeleteByURL(ctx context.Context,
	client *akamai.ServiceClient,
	network Network, urls URLs) (*CCUResponse, *akamai.ResponseResult, error){

	url := fmt.Sprintf(deleteByURL, network)

	respResult, err := client.DoRequest(ctx, http.MethodPost, url, urls)
	if err != nil {
		return nil, nil, err
	}
	if respResult.Err != nil {
		return nil, respResult, respResult.Err
	}

	ccuResp := CCUResponse{}

	if err = respResult.ExtractResult(&ccuResp); err != nil {
		return nil, nil, err
	}

	return &ccuResp, respResult, nil
}

// DeleteByURLWithHostname deletes content on the selected URL related to the given hostname
// for the selected network.
func DeleteByURLWithHostname(ctx context.Context,
	client *akamai.ServiceClient,
	network Network, urls URLsWithHostname) (*CCUResponse, *akamai.ResponseResult, error){

	url := fmt.Sprintf(deleteByURL, network)

	respResult, err := client.DoRequest(ctx, http.MethodPost, url, urls)
	if err != nil {
		return nil, nil, err
	}
	if respResult.Err != nil {
		return nil, respResult, respResult.Err
	}

	ccuResp := CCUResponse{}

	if err = respResult.ExtractResult(&ccuResp); err != nil {
		return nil, nil, err
	}

	return &ccuResp, respResult, nil
}

// DeleteByCPCode deletes content on the selected CP code for the selected network.
func DeleteByCPCode(ctx context.Context,
	client *akamai.ServiceClient,
	network Network, cpCodes CPCodes) (*CCUResponse, *akamai.ResponseResult, error){

	url := fmt.Sprintf(deleteByCPCode, network)

	respResult, err := client.DoRequest(ctx, http.MethodPost, url, cpCodes)
	if err != nil {
		return nil, nil, err
	}
	if respResult.Err != nil {
		return nil, respResult, respResult.Err
	}

	ccuResp := CCUResponse{}

	if err = respResult.ExtractResult(&ccuResp); err != nil {
		return nil, nil, err
	}

	return &ccuResp, respResult, nil
}

// DeleteByCacheTag deletes content on the selected set of cache tags for the selected network.
func DeleteByCacheTag(ctx context.Context,
	client *akamai.ServiceClient,
	network Network, cacheTags CacheTags) (*CCUResponse, *akamai.ResponseResult, error){

	url := fmt.Sprintf(deleteByCacheTag, network)

	respResult, err := client.DoRequest(ctx, http.MethodPost, url, cacheTags)
	if err != nil {
		return nil, nil, err
	}
	if respResult.Err != nil {
		return nil, respResult, respResult.Err
	}

	ccuResp := CCUResponse{}

	if err = respResult.ExtractResult(&ccuResp); err != nil {
		return nil, nil, err
	}

	return &ccuResp, respResult, nil
}