package testing

import "github.com/dstdfx/go-akamai/akamai/papi/v1/edgehostname"

const (
	testGetEdgeHostnameRawResponse = `
{
  "accountId": "act_1-1TJZFB",
  "contractId": "ctr_1-1TJZH5",
  "groupId": "grp_15225",
  "edgeHostnames": {
    "items": [
      {
        "edgeHostnameId": "ehn_887436",
        "edgeHostnameDomain": "example.com.edgesuite.net",
        "productId": "prd_HTTP_Downloads",
        "domainPrefix": "example.com",
        "domainSuffix": "edgesuite.net",
        "secure": false,
        "ipVersionBehavior": "IPV4",
        "mapDetails:serialNumber": 1951,
        "mapDetails:mapDomain": "a1951.g.akamai.net"
      }
    ]
  }
}
`
	testGetByLinkEdgeHostnameRawResponse = `
{
  "accountId": "act_1-1TJZFB",
  "contractId": "ctr_1-1TJZH5",
  "groupId": "grp_15225",
  "edgeHostnames": {
    "items": [
      {
        "edgeHostnameId": "ehn_887436",
        "edgeHostnameDomain": "example.com.edgesuite.net",
        "productId": "prd_HTTP_Downloads",
        "domainPrefix": "example.com",
        "domainSuffix": "edgesuite.net",
        "secure": false,
        "ipVersionBehavior": "IPV4",
        "mapDetails:serialNumber": 1951,
        "mapDetails:mapDomain": "a1951.g.akamai.net"
      }
    ]
  }
}
`
	testListEdgeHostnamesRawResponse = `
{
  "accountId": "act_1-1TJZFB",
  "contractId": "ctr_1-1TJZH5",
  "groupId": "grp_15225",
  "edgeHostnames": {
    "items": [
      {
        "edgeHostnameId": "ehn_887437",
        "edgeHostnameDomain": "example2.com.edgesuite.net",
        "productId": "prd_HTTP_Downloads",
        "domainPrefix": "example2.com",
        "domainSuffix": "edgesuite.net",
        "secure": false,
        "ipVersionBehavior": "IPV4",
        "mapDetails:serialNumber": 1952,
        "mapDetails:mapDomain": "a1952.g.akamai.net"
      },
      {
        "edgeHostnameId": "ehn_887436",
        "edgeHostnameDomain": "example.com.edgesuite.net",
        "productId": "prd_HTTP_Downloads",
        "domainPrefix": "example.com",
        "domainSuffix": "edgesuite.net",
        "secure": false,
        "ipVersionBehavior": "IPV4",
        "mapDetails:serialNumber": 1951,
        "mapDetails:mapDomain": "a1951.g.akamai.net"
      }
    ]
  }
}
`
	testCreateEdgeHostnameRawResponse = `
{
  "edgeHostnameLink": "/papi/v1/edgehostnames/ehn_1332?contractId=ctr_1-1TJZH5&grp_15225"
}

`
)

const (
	testCreateEdgeHostnameRawRequest = `
{
  "productId": "prd_HTTP_Downloads",
  "domainPrefix": "www.example.com",
  "domainSuffix": "edgesuite.net",
  "secure": true,
  "ipVersionBehavior": "IPV4",
  "slotNumber": 12345
}
`
	testEdgeHostnameLink = "/papi/v1/edgehostnames/ehn_1332"
)

var (
	testCreateEdgeHostnameExpected = "/papi/v1/edgehostnames/ehn_1332?contractId=ctr_1-1TJZH5&grp_15225"

	testGetEdgeHostnameExpected = &edgehostname.EdgeHostname{
		ID: "ehn_887436",
		EdgeHostnameDomain: "example.com.edgesuite.net",
		ProductID: "prd_HTTP_Downloads",
		DomainPrefix: "example.com",
		DomainSuffix: "edgesuite.net",
		Secure: false,
		IPVersionBehavior: edgehostname.IPVersionBehaviorV4,
		MapDetailsSerialNumber: 1951,
		MapDetailsMapDomain: "a1951.g.akamai.net",
	}

	testGetByLinkEdgeHostnameExpected = &edgehostname.EdgeHostname{
		ID: "ehn_887436",
		EdgeHostnameDomain: "example.com.edgesuite.net",
		ProductID: "prd_HTTP_Downloads",
		DomainPrefix: "example.com",
		DomainSuffix: "edgesuite.net",
		Secure: false,
		IPVersionBehavior: edgehostname.IPVersionBehaviorV4,
		MapDetailsSerialNumber: 1951,
		MapDetailsMapDomain: "a1951.g.akamai.net",
	}

	testListEdgeHostnamesExpected = []*edgehostname.EdgeHostname{
		{
			ID: "ehn_887437",
			EdgeHostnameDomain: "example2.com.edgesuite.net",
			ProductID: "prd_HTTP_Downloads",
			DomainPrefix: "example2.com",
			DomainSuffix: "edgesuite.net",
			Secure: false,
			IPVersionBehavior: edgehostname.IPVersionBehaviorV4,
			MapDetailsSerialNumber: 1952,
			MapDetailsMapDomain: "a1952.g.akamai.net",
		},
		{
			ID: "ehn_887436",
			EdgeHostnameDomain: "example.com.edgesuite.net",
			ProductID: "prd_HTTP_Downloads",
			DomainPrefix: "example.com",
			DomainSuffix: "edgesuite.net",
			Secure: false,
			IPVersionBehavior: edgehostname.IPVersionBehaviorV4,
			MapDetailsSerialNumber: 1951,
			MapDetailsMapDomain: "a1951.g.akamai.net",
		},
	}
)