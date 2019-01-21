package testing

import (
	"github.com/dstdfx/go-akamai/akamai/internal/util"
	"github.com/dstdfx/go-akamai/akamai/papi/v1/property"
)

const (
	testCreatePropertyRawResponse = `
{
  "propertyLink": "/papi/v0/properties/prp_173137?contractId=ctr_1-1TJZH5&groupId=grp_15225"
}
`
	testClonePropertyRawResponse = `
{
  "propertyLink": "/papi/v0/properties/prp_173137?contractId=ctr_1-1TJZH5&groupId=grp_15225"
}
`
	testGetPropertyRawResponse = `
{
  "properties": {
    "items": [
      {
        "accountId": "act_1-1TJZFB",
        "contractId": "ctr_1-1TJZH5",
        "groupId": "grp_15225",
        "propertyId": "prp_175780",
        "propertyName": "example.com",
        "latestVersion": 2,
        "stagingVersion": 1,
        "productionVersion": null,
        "assetId": "aid_101",
        "note": "Notes about example.com"
      }
    ]
  }
}

`
	testGetPropertyByLinkRawResponse = `
{
  "properties": {
    "items": [
      {
        "accountId": "act_1-1TJZFB",
        "contractId": "ctr_1-1TJZH5",
        "groupId": "grp_15225",
        "propertyId": "prp_175780",
        "propertyName": "example.com",
        "latestVersion": 2,
        "stagingVersion": 1,
        "productionVersion": null,
        "assetId": "aid_101",
        "note": "Notes about example.com"
      }
    ]
  }
}

`

	testListPropertiesRawRequest = `
{
  "properties": {
    "items": [
     {
        "accountId": "act_1-1TJZFB",
        "contractId": "ctr_1-1TJZH5",
        "groupId": "grp_15225",
        "propertyId": "prp_175780",
        "propertyName": "example.com",
        "latestVersion": 2,
        "stagingVersion": 1,
        "productionVersion": null,
        "assetId": "aid_101",
        "note": "Notes about example.com"
      },
      {
        "accountId": "act_1-1TJZFB",
        "contractId": "ctr_1-1TJZH5",
        "groupId": "grp_15225",
        "propertyId": "prp_275780",
        "propertyName": "example2.com",
        "latestVersion": 2,
        "stagingVersion": 1,
        "productionVersion": null,
        "assetId": "aid_101",
        "note": "Notes about example2.com"
      }
    ]
  }
}
`
	testDeletePropertyResponse = `
{
    "message": "Deletion Successful."
}`
)


const (
	testCreatePropertyRawRequest = `
{
  "productId": "prd_HTTP_Downloads",
  "propertyName": "my.new.property.com"
}
`
	testClonePropertyRawRequest = `
{
  "productId": "prd_HTTP_Downloads",
  "propertyName": "my.new.property.com",
  "cloneFrom": {
    "propertyId": "prp_175780",
    "version": 2,
    "cloneFromVersionEtag": "a9dfe78cf93090516bde891d009eaf57",
    "copyHostnames": true
  }
}
`
	testPropertyLink = "/papi/v1/properties/prp_173137"

)

var (
	testCreatePropertyExpected = "/papi/v0/properties/prp_173137?contractId=ctr_1-1TJZH5&groupId=grp_15225"
	testClonePropertyExpected = "/papi/v0/properties/prp_173137?contractId=ctr_1-1TJZH5&groupId=grp_15225"
	testGetPropertyExpected = &property.Property{
		ID: "prp_175780",
		Name: "example.com",
		AccountID: "act_1-1TJZFB",
		ContractID: "ctr_1-1TJZH5",
		GroupID: "grp_15225",
		LatestVersion: util.IntPtr(2),
		StagingVersion: util.IntPtr(1),
		AssetID: "aid_101",
		Note: "Notes about example.com",
	}
	testGetPropertyByLinkExpected = &property.Property{
		ID: "prp_175780",
		Name: "example.com",
		AccountID: "act_1-1TJZFB",
		ContractID: "ctr_1-1TJZH5",
		GroupID: "grp_15225",
		LatestVersion: util.IntPtr(2),
		StagingVersion: util.IntPtr(1),
		AssetID: "aid_101",
		Note: "Notes about example.com",
	}
	testListPropertiesExpected = []*property.Property{
		{
			ID: "prp_175780",
			Name: "example.com",
			AccountID: "act_1-1TJZFB",
			ContractID: "ctr_1-1TJZH5",
			GroupID: "grp_15225",
			LatestVersion: util.IntPtr(2),
			StagingVersion: util.IntPtr(1),
			AssetID: "aid_101",
			Note: "Notes about example.com",
		},
		{
			ID: "prp_275780",
			Name: "example2.com",
			AccountID: "act_1-1TJZFB",
			ContractID: "ctr_1-1TJZH5",
			GroupID: "grp_15225",
			LatestVersion: util.IntPtr(2),
			StagingVersion: util.IntPtr(1),
			AssetID: "aid_101",
			Note: "Notes about example2.com",
		},
	}
)