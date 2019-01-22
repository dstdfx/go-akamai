package testing

import "github.com/dstdfx/go-akamai/akamai/papi/v1/hostname"

const (
	testListHostnamesRawResponse = `
{
    "accountId": "act_1-1TJZFB",
    "contractId": "ctr_1-1TJZH5",
    "groupId": "grp_15225",
    "propertyId": "prp_175780",
    "propertyVersion": 1,
    "etag": "6aed418629b4e5c0",
    "hostnames": {
        "items": [
            {
                "cnameType": "EDGE_HOSTNAME",
                "edgeHostnameId": "ehn_895822",
                "cnameFrom": "example.com",
                "cnameTo": "example.com.edgesuite.net"
            },
            {
                "cnameType": "EDGE_HOSTNAME",
                "edgeHostnameId": "ehn_895833",
                "cnameFrom": "m.example.com",
                "cnameTo": "m.example.com.edgesuite.net"
            }
        ]
    }
}
`

	testUpdateHostnamesRawResponse = `
{
    "accountId": "act_1-1TJZFB",
    "contractId": "ctr_1-1TJZH5",
    "groupId": "grp_15225",
    "propertyId": "prp_175780",
    "propertyVersion": 1,
    "etag": "6aed418629b4e5c0",
    "hostnames": {
        "items": [
            {
                "cnameType": "EDGE_HOSTNAME",
                "edgeHostnameId": "ehn_895822",
                "cnameFrom": "example.com",
                "cnameTo": "example.com.edgesuite.net"
            },
            {
                "cnameType": "EDGE_HOSTNAME",
                "edgeHostnameId": "ehn_895833",
                "cnameFrom": "m.example.com",
                "cnameTo": "m.example.com.edgesuite.net"
            }
        ]
    }
}
`
)


const (
	testUpdateHostnamesRawRequest = `
[
    {
        "cnameType": "EDGE_HOSTNAME",
        "cnameFrom": "example.com",
		"edgeHostnameId": "ehn_895833",
        "cnameTo": "example.com.edgesuite.net"
    },
    {
        "cnameType": "EDGE_HOSTNAME",
        "edgeHostnameId": "ehn_895824",
        "cnameFrom": "m.example.com"
    }
]
`
)

var (
	testListHostnamesExpected = []*hostname.Hostname{
		{
			CnameType: hostname.CnameTypeEdgeHostname,
			EdgeHostnameID: "ehn_895822",
			CnameFrom: "example.com",
			CnameTo: "example.com.edgesuite.net",

		},
		{
			CnameType: hostname.CnameTypeEdgeHostname,
			EdgeHostnameID: "ehn_895833",
			CnameFrom: "m.example.com",
			CnameTo: "m.example.com.edgesuite.net",

		},
	}
	testUpdateHostnamesExpected = []*hostname.Hostname{
		{
			CnameType: hostname.CnameTypeEdgeHostname,
			EdgeHostnameID: "ehn_895822",
			CnameFrom: "example.com",
			CnameTo: "example.com.edgesuite.net",

		},
		{
			CnameType: hostname.CnameTypeEdgeHostname,
			EdgeHostnameID: "ehn_895833",
			CnameFrom: "m.example.com",
			CnameTo: "m.example.com.edgesuite.net",

		},
	}
)
