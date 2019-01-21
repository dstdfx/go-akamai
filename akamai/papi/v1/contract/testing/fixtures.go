package testing

import "github.com/dstdfx/go-akamai/akamai/papi/v1/contract"

const (
	testListContractsRawResp = `
{
  "accountId": "act_1-1TJZFB",
  "contracts": {
    "items": [
      {
        "contractId": "ctr_1-1TJZH5",
        "contractTypeName": "Direct Customer"
      }
    ]
  }
}
`)


var (
	testListContractsExpected = []*contract.Contract{
		{
			ID: "ctr_1-1TJZH5",
			TypeName: "Direct Customer",
		},
	}
)
