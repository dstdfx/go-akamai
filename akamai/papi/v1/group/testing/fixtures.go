package testing

import "github.com/dstdfx/go-akamai/akamai/papi/v1/group"

var (
	testListGroupsRawResp = `
{
  "accountId": "act_1-1TJZFB",
  "accountName": "Example.com",
  "groups": {
    "items": [
      {
        "groupName": "Example.com-1-1TJZH5",
        "groupId": "grp_15225",
        "contractIds": [
          "ctr_1-1TJZH5"
        ]
      },
      {
        "groupName": "Test",
        "groupId": "grp_15231",
        "parentGroupId": "grp_15225",
        "contractIds": [
          "ctr_1-1TJZH5"
        ]
      }
    ]
  }
}
`)


var (
	testListGroupsExpected = []*group.Group{
		{
			ID: "grp_15225",
			Name: "Example.com-1-1TJZH5",
			ContractIDs: []string{
				"ctr_1-1TJZH5",
			},
		},
		{
			ID: "grp_15231",
			Name: "Test",
			ParentGroupID: "grp_15225",
			ContractIDs: []string{
				"ctr_1-1TJZH5",
			},
		},
	}
)
