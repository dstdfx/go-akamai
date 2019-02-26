// Package edgegrid allows you to sign up your request to calling
// Akamai APIs. The package is based on:
// https://github.com/akamai/AkamaiOPEN-edgegrid-golang/tree/master/edgegrid
package edgegrid

// Config provides all the necessary fields to
// create authorization header for calling Akamai APIs.
type Config struct {
	Host         string
	ClientToken  string
	ClientSecret string
	AccessToken  string
	HeaderToSign []string
	MaxBody      int
}
