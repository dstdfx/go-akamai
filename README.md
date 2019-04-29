# go-akamai
[![Build Status](https://travis-ci.org/dstdfx/go-akamai.svg?branch=master)](https://travis-ci.org/dstdfx/go-akamai)
[![Coverage Status](https://coveralls.io/repos/github/dstdfx/go-akamai/badge.svg?branch=master)](https://coveralls.io/github/dstdfx/go-akamai?branch=master)

Golang SDK for accessing different Akamai service APIs.
https://developer.akamai.com/api

Supported APIs:
- [Property Manager API v1](https://developer.akamai.com/api/core_features/property_manager/v1.html)
- [Fast Purge API v3](https://developer.akamai.com/api/core_features/fast_purge/v3.html)

## Getting started ##

### Installation ###

You can install ```go-akamai``` as a Go package:
````bash
go get github.com/dstdfx/go-akamai
````

### Authentication ####

To work with Akamai APIs you need to create API client with specific
access rules. More information [here](https://developer.akamai.com/legacy/introduction/Prov_Creds.html).

### Usage example ###

```go
// Example: cache invalidation for specific URLs.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dstdfx/go-akamai/akamai/ccu/v3"
	"github.com/dstdfx/go-akamai/akamai/ccu/v3/cache"
	"github.com/dstdfx/go-akamai/akamai/edgegrid"
)

func main(){
	// Fill it up with your credentials
	cfg := &edgegrid.Config{
		Host:         "<your-host>",
		ClientSecret: "<your-client-secret>",
		ClientToken:  "<your-client-token>",
		AccessToken:  "<your-access-token>",
		MaxBody:      131072,
	}

	client := v3.NewV3CCUClient(cfg)

	// Invalidate cache for specific URLs for production network.
	// Default context just for the sake of simplicity.
	resp, _, err := cache.InvalidateByURL(context.Background(),
		client,
		cache.ProductionNetwork,
		cache.URLs{
			Objects:[]string{
				"https://1136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net/example_video.mp4",
				"https://4136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net/example_video_4k.mp4",
			},
	})
	if err != nil {
            panic(err)
	}
	fmt.Printf("%+v\n", resp)
}

```

## Contributing ##
Please read [CONTRIBUTING.md](.github/CONTRIBUTING.md) for details.


## License ##
This library is distributed under the MIT license found in the [LICENSE](./LICENSE) file.
