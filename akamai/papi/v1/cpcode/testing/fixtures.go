package testing

import (
	"time"

	"github.com/dstdfx/go-akamai/akamai/papi/v1/cpcode"
)

const (
	testListCpCodesRawResp = `
{
  "accountId": "act_1-1TJZFB",
  "contractId": "ctr_1-1TJZFW",
  "groupId": "grp_15225",
  "cpcodes": {
    "items": [
      {
        "cpcodeId": "cpc_33190",
        "cpcodeName": "SME WAA",
        "productIds": [
          "prd_HTTP_Downloads"
        ],
        "createdDate": "2015-03-02T15:06:13Z"
      }
    ]
  }
}
`
	testCreateCpCodeRawResp = `
{
  "cpcodeLink": "/papi/v0/cpcodes/cpc_33190?contractId=ctr_1-1TJZFW&groupId=grp_15166"
}

`
	testGetCpCodeRawResp = `
{
  "accountId": "act_1-1TJZFB",
  "contractId": "ctr_1-1TJZFW",
  "groupId": "grp_15166",
  "cpcodes": {
    "items": [
      {
        "cpcodeId": "cpc_33190",
        "cpcodeName": "SME WAA",
        "productIds": [
          "prd_HTTP_Downloads"
        ],
        "createdDate": "2015-03-02T15:06:13Z"
      }
    ]
  }
}
`
	testGetCpCodeByLinkRawResp = `
{
  "accountId": "act_1-1TJZFB",
  "contractId": "ctr_1-1TJZFW",
  "groupId": "grp_15166",
  "cpcodes": {
    "items": [
      {
        "cpcodeId": "cpc_33190",
        "cpcodeName": "SME WAA",
        "productIds": [
          "prd_HTTP_Downloads"
        ],
        "createdDate": "2015-03-02T15:06:13Z"
      }
    ]
  }
}
`

)

const testCreateCpCodeRawRequest = `
{
  "productId": "prd_HTTP_Downloads",
  "cpcodeName": "SME WAA"
}
`

const testLinkToCpCode = "/papi/v1/cpcodes/cpc_33190"


var (
	testTime = time.Date(2015, 03,02, 15, 06,13, 0, time.UTC)
	testListCpCodesExpected = []*cpcode.CPCode{
		{
			ID: "cpc_33190",
			Name: "SME WAA",
			ProductIDs: []string{
				"prd_HTTP_Downloads",
			},
			CreatedDate: &testTime,
		},
	}
	testGetCpCodeExpected = &cpcode.CPCode{
			ID: "cpc_33190",
			Name: "SME WAA",
			ProductIDs: []string{
				"prd_HTTP_Downloads",
			},
			CreatedDate: &testTime,
	}
	testGetCpCodeByLinkExpected = &cpcode.CPCode{
		ID: "cpc_33190",
		Name: "SME WAA",
		ProductIDs: []string{
			"prd_HTTP_Downloads",
		},
		CreatedDate: &testTime,
	}
	testCreateCpCodeExpected = "/papi/v0/cpcodes/cpc_33190?contractId=ctr_1-1TJZFW&groupId=grp_15166"
)
