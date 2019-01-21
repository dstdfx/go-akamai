package akamai

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dstdfx/go-akamai/internal/logger"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

const (
	// AppVersion is a version of the application.
	AppVersion = "0.0.1"

	// AppName is a global application name.
	AppName = "go-akamai"

	// DefaultUserAgent contains basic user agent that will be used in queries.
	DefaultUserAgent = AppName + "/" + AppVersion

	// defaultHTTPTimeout represents the default timeout (in seconds) for HTTP
	// requests.
	defaultHTTPTimeout = 120

	// defaultDialTimeout represents the default timeout (in seconds) for HTTP
	// connection establishments.
	defaultDialTimeout = 60

	// defaultKeepalive represents the default keep-alive period for an active
	// network connection.
	defaultKeepaliveTimeout = 60

	// defaultMaxIdleConns represents the maximum number of idle (keep-alive)
	// connections.
	defaultMaxIdleConns = 100

	// defaultIdleConnTimeout represents the maximum amount of time an idle
	// (keep-alive) connection will remain idle before closing itself.
	defaultIdleConnTimeout = 100

	// defaultTLSHandshakeTimeout represents the default timeout (in seconds)
	// for TLS handshake.
	defaultTLSHandshakeTimeout = 60

	// defaultExpectContinueTimeout represents the default amount of time to
	// wait for a server's first response headers.
	defaultExpectContinueTimeout = 1
)

// NewHTTPClient returns a reference to an initialized configured HTTP client.
func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout:   time.Second * defaultHTTPTimeout,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   defaultDialTimeout * time.Second,
				KeepAlive: defaultKeepaliveTimeout * time.Second,
			}).DialContext,
			MaxIdleConns:          defaultMaxIdleConns,
			IdleConnTimeout:       defaultIdleConnTimeout * time.Second,
			TLSHandshakeTimeout:   defaultTLSHandshakeTimeout * time.Second,
			ExpectContinueTimeout: defaultExpectContinueTimeout * time.Second,
		},
	}
}

// ServiceClient stores details that are needed to work with different Akamai APIs.
type ServiceClient struct {
	// HTTPClient represents an initialized HTTP client that will be used to do requests.
	HTTPClient *http.Client

	// Endpoint represents an endpoint that will be used in all requests.
	Endpoint string

	// UserAgent contains user agent that will be used in all requests.
	UserAgent string

	// Config represents Akamai Edge Grid config.
	Config *edgegrid.Config

	log logger.GenericLogger
}

// SetLogger sets logger to the given.
func (client *ServiceClient) SetLogger(log logger.GenericLogger) {
	client.log = log
}

// newRequest creates an HTTP request to Akamai APIs.
// A relative URL can be provided in path, which will be resolved to the
// Host specified in Config. If body is specified, it will be sent as the request body.
func (client *ServiceClient) newRequest(ctx context.Context,
	method, path string,
	body interface{}) (*http.Request, error){

	var (
		jsonBody io.Reader
		err     error
	)

	// Serialize body to io.Reader if presented
	if body != nil {
		v, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		jsonBody = bytes.NewReader(v)
	}

	u, err := url.Parse(client.Endpoint + path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), jsonBody)
	if err != nil {
		return nil, err
	}

	// Set context
	req = req.WithContext(ctx)

	// Set headers
	req.Header.Add("User-Agent", client.UserAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req = edgegrid.AddRequestHeader(*client.Config, req)

	return req, nil
}

// DoRequest creates an HTTP request and sends it to Akamai APIs.
// A relative URL can be provided in path, which will be resolved to the
// Host specified in Config. If body is specified, it will be sent as the request body.
func (client *ServiceClient) DoRequest(ctx context.Context,
	method, path string,
	body interface{}) (*ResponseResult, error) {

	client.log.Debugf("REQ %s %s", method, path)

	// Create new request
	req, err := client.newRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	// Execute request
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	client.log.Debugf("RESP %s %s %d", method, path, resp.StatusCode)

	respResult := &ResponseResult{
		Response: resp,
		Err:      nil,
	}

	// Handle error if presented
	if resp.StatusCode >= http.StatusBadRequest &&
		resp.StatusCode <= http.StatusNetworkAuthenticationRequired {
		respResult.Err = NewAPIError(respResult.Response)
	}

	return respResult, nil
}

// DoRequestWithCustomHeaders creates an HTTP request with re-written
// headers and sends it to Akamai APIs.
// A relative URL can be provided in path, which will be resolved to the
// Host specified in Config. If body is specified, it will be sent as the JSON request body.
func (client *ServiceClient) DoRequestWithCustomHeaders(ctx context.Context,
	method, path string,
	headers http.Header,
	body interface{}) (*ResponseResult, error) {

	client.log.Debugf("REQ %s %s", method, path)

	// Create new request
	req, err := client.newRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	// Add custom headers
	for header, values := range headers {
		req.Header.Set(header, strings.Join(values, ","))
	}

	// Execute request
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	client.log.Debugf("RESP %s %s %d", method, path, resp.StatusCode)

	respResult := &ResponseResult{
		Response: resp,
		Err:      nil,
	}

	if resp.StatusCode >= http.StatusBadRequest &&
		resp.StatusCode <= http.StatusNetworkAuthenticationRequired {
		respResult.Err = NewAPIError(respResult.Response)
	}

	return respResult, nil
}


// ResponseResult represents a result of a HTTP request.
// It embeds standard http.Response and adds a custom error description.
type ResponseResult struct {
	*http.Response

	// Err contains error that can be provided to a caller.
	Err *APIError
}

// ExtractResult allows to provide an object into which ResponseResult body will be extracted.
func (result *ResponseResult) ExtractResult(to interface{}) error {
	body, err := ioutil.ReadAll(result.Body)
	defer result.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, to)
	return err
}
