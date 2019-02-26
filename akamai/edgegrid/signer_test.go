package edgegrid

import (
	"bytes"
	"net/http"
	"regexp"
	"testing"

	"github.com/dstdfx/go-akamai/akamai/internal/logger"
	"github.com/stretchr/testify/assert"
)

var (
	testTimestamp = "20140321T19:34:21+0000"
	testNonce     = "nonce-xx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	testConfig    = Config{
		Host:         "https://akaa-baseurl-xxxxxxxxxxx-xxxxxxxxxxxxx.luna.akamaiapis.net/",
		AccessToken:  "akab-access-token-xxx-xxxxxxxxxxxxxxxx",
		ClientToken:  "akab-client-token-xxx-xxxxxxxxxxxxxxxx",
		ClientSecret: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx=",
		MaxBody:      2048,
		HeaderToSign: []string{
			"X-Test1",
			"X-Test2",
			"X-Test3",
		},
	}
	testURL                = "https://akamai-test.net"
	testVoidLogger = &logger.VoidLogger{}
)

func TestMakeEdgeTimeStamp(t *testing.T) {
	actual := makeEdgeTimeStamp()
	expected := regexp.MustCompile(`^\d{4}[0-1][0-9][0-3][0-9]T[0-2][0-9]:[0-5][0-9]:[0-5][0-9]\+0000$`)
	if !assert.Regexp(t, expected, actual) {
		t.Fatalf("regex do not match: expected %s - actual %s", expected, actual)
	}
}

func TestCreateNonce(t *testing.T) {
	actual := createNonce(testVoidLogger)
	for i := 0; i < 100; i++ {
		expected := createNonce(testVoidLogger)
		assert.NotEqual(t, actual, expected)
	}
}

func TestCreateAuthHeader(t *testing.T) {
	req, err  := http.NewRequest(http.MethodPost,
		testURL,
		bytes.NewBuffer([]byte("")))
	assert.NoError(t, err)
	authHeader := createAuthHeader(testVoidLogger, testConfig, req, testTimestamp, testNonce)
	assert.NotEqual(t, authHeader, "")
}
