// Package edgegrid allows you to sign up your request to calling
// Akamai APIs. The package is based on:
// https://github.com/akamai/AkamaiOPEN-edgegrid-golang/tree/master/edgegrid
package edgegrid

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/dstdfx/go-akamai/akamai/internal/logger"
	"github.com/gofrs/uuid"
)

// AddRequestHeader sets the Authorization header to use Akamai Open API.
func AddRequestHeader(log logger.GenericLogger, config Config, req *http.Request) *http.Request {
	timestamp := makeEdgeTimeStamp()
	nonce := createNonce(log)

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Authorization", createAuthHeader(log, config, req, timestamp, nonce))
	return req
}

// Must be assigned the UTC time when the request is signed.
// Format of “yyyyMMddTHH:mm:ss+0000”
func makeEdgeTimeStamp() string {
	local := time.FixedZone("GMT", 0)
	t := time.Now().In(local)
	return fmt.Sprintf("%d%02d%02dT%02d:%02d:%02d+0000",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

// Must be assigned a nonce (number used once) for the request.
// It is a random string used to detect replayed request messages.
// A GUID is recommended.
func createNonce(log logger.GenericLogger) string {
	v, err := uuid.NewV4()
	if err != nil {
		log.Errorf("Generate UUID failed: %s", err)
		return ""
	}
	return v.String()
}

func stringMinifier(in string) (out string) {
	white := false
	for _, c := range in {
		if unicode.IsSpace(c) {
			if !white {
				out = out + " "
			}
			white = true
		} else {
			out = out + string(c)
			white = false
		}
	}
	return
}

func concatPathQuery(path, query string) string {
	if query == "" {
		return path
	}
	return fmt.Sprintf("%s?%s", path, query)
}

// createSignature is the base64-encoding of the SHA–256 HMAC of the
// data to sign with the signing key.
func createSignature(message, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	_, _ = h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func createHash(data string) string {
	h := sha256.Sum256([]byte(data))
	return base64.StdEncoding.EncodeToString(h[:])
}

func canonicalizeHeaders(config Config, req *http.Request) string {
	var unsortedHeader []string
	var sortedHeader []string
	for k := range req.Header {
		unsortedHeader = append(unsortedHeader, k)
	}
	sort.Strings(unsortedHeader)
	for _, k := range unsortedHeader {
		for _, sign := range config.HeaderToSign {
			if sign == k {
				v := strings.TrimSpace(req.Header.Get(k))
				sortedHeader = append(sortedHeader,
					fmt.Sprintf("%s:%s",
						strings.ToLower(k),
						strings.ToLower(stringMinifier(v))))
			}
		}
	}
	return strings.Join(sortedHeader, "\t")

}

// signingKey is derived from the client secret.
// The signing key is computed as the base64 encoding of the
// SHA–256 HMAC of the timestamp string (the field value included
// in the HTTP authorization header described above) with the client secret as the key.
func signingKey(config Config, timestamp string) string {
	key := createSignature(timestamp, config.ClientSecret)
	return key
}

// The content hash is the base64-encoded SHA–256 hash of the POST body.
// For any other request methods, this field is empty. But the tac separator (\t) must be included.
// The size of the POST body must be less than or equal to the value specified by the service.
// Any request that does not meet this criteria SHOULD be rejected during the signing process,
// as the request will be rejected by EdgeGrid.
func createContentHash(log logger.GenericLogger, config Config, req *http.Request) string {
	var (
		contentHash  string
		preparedBody string
		bodyBytes    []byte
	)
	if req.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(req.Body)
		req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		preparedBody = string(bodyBytes)
	}

	log.Debugf("Body is %s", preparedBody)
	if req.Method == http.MethodPost && len(preparedBody) > 0 {
		log.Debugf("Signing content: %s", preparedBody)
		if len(preparedBody) > config.MaxBody {
			log.Debugf("Data length %d is larger than maximum %d",
				len(preparedBody), config.MaxBody)

			preparedBody = preparedBody[0:config.MaxBody]
			log.Debugf("Data truncated to %d for computing the hash", len(preparedBody))
		}
		contentHash = createHash(preparedBody)
	}
	log.Debugf("Content hash is '%s'", contentHash)
	return contentHash
}

// The data to sign includes the information from the HTTP request that is
// relevant to ensuring that the request is authentic.
// This data set comprised of the request data combined with the authorization
// header value (excluding the signature field, but including the ; right before
// the signature field).
func signingData(log logger.GenericLogger, config Config, req *http.Request, authHeader string) string {

	dataSign := []string{
		req.Method,
		req.URL.Scheme,
		req.URL.Host,
		concatPathQuery(req.URL.Path, req.URL.RawQuery),
		canonicalizeHeaders(config, req),
		createContentHash(log, config, req),
		authHeader,
	}
	log.Debugf("Data to sign %s", strings.Join(dataSign, "\t"))
	return strings.Join(dataSign, "\t")
}

func signingRequest(log logger.GenericLogger,
	config Config,
	req *http.Request,
	authHeader,
	timestamp string) string {

	return createSignature(signingData(log, config, req, authHeader),
		signingKey(config, timestamp))
}

// The Authorization header starts with the signing algorithm moniker
// (name of the algorithm) used to sign the request.
// The moniker below identifies EdgeGrid V1, hash message authentication
// code, SHA–256 as the hash standard. This moniker is then followed by
// a space and an ordered list of name value pairs with each field separated
// by a semicolon.
func createAuthHeader(log logger.GenericLogger,
	config Config,
	req *http.Request,
	timestamp, nonce string) string {

	authHeader := fmt.Sprintf("EG1-HMAC-SHA256 client_token=%s;access_token=%s;timestamp=%s;nonce=%s;",
		config.ClientToken,
		config.AccessToken,
		timestamp,
		nonce,
	)
	log.Debugf("Unsigned authorization header: '%s'", authHeader)

	signedAuthHeader := fmt.Sprintf("%ssignature=%s",
		authHeader,
		signingRequest(log, config, req, authHeader, timestamp))

	log.Debugf("Signed authorization header: '%s'", signedAuthHeader)
	return signedAuthHeader
}
