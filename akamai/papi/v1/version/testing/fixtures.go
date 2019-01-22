package testing

import (
	"time"

	"github.com/dstdfx/go-akamai/akamai/internal/util"
	"github.com/dstdfx/go-akamai/akamai/papi/v1/version"
)

const (
	testCreateVersionRawResponse = `
{
    "versionLink": "/papi/v1/properties/prp_173136/versions/2?contractId=ctr_1-1TJZH5&groupId=grp_15225"
}
`

	testGetVersionRawResponse = `
{
    "propertyId": "prp_173136",
    "propertyName": "981.catalog.amenai.net",
    "accountId": "act_1-1TJZFB",
    "contractId": "ctr_1-1TJZH5",
    "groupId": "grp_15225",
    "assetId": "aid_101",
    "versions": {
        "items": [
            {
                "propertyVersion": 1,
                "updatedByUser": "afaden",
                "updatedDate": "2015-03-02T15:06:13Z",
                "productionStatus": "ACTIVE",
                "stagingStatus": "INACTIVE",
                "etag": "71573b922a87abc3",
                "productId": "prd_Alta",
                "ruleFormat": "latest",
                "note": "initial version"
            }
        ]
    }
}
`
	testGetVersionByLinkRawResponse = `
{
    "propertyId": "prp_173136",
    "propertyName": "981.catalog.amenai.net",
    "accountId": "act_1-1TJZFB",
    "contractId": "ctr_1-1TJZH5",
    "groupId": "grp_15225",
    "assetId": "aid_101",
    "versions": {
        "items": [
            {
                "propertyVersion": 1,
                "updatedByUser": "afaden",
                "updatedDate": "2015-03-02T15:06:13Z",
                "productionStatus": "ACTIVE",
                "stagingStatus": "INACTIVE",
                "etag": "71573b922a87abc3",
                "productId": "prd_Alta",
                "ruleFormat": "latest",
                "note": "initial version"
            }
        ]
    }
}
`

	testGetLatestVersionRawResponse = `
{
    "propertyId": "prp_173136",
    "propertyName": "981.catalog.amenai.net",
    "accountId": "act_1-1TJZFB",
    "contractId": "ctr_1-1TJZH5",
    "groupId": "grp_15225",
    "assetId": "aid_101",
    "versions": {
        "items": [
            {
                "propertyVersion": 1,
                "updatedByUser": "afaden",
                "updatedDate": "2015-03-02T15:06:13Z",
                "productionStatus": "ACTIVE",
                "stagingStatus": "INACTIVE",
                "etag": "71573b922a87abc3",
                "productId": "prd_Alta",
                "ruleFormat": "latest",
                "note": "initial version"
            }
        ]
    }
}
`
	testListVersionsRawResponse = `
{
    "propertyId": "prp_173136",
    "propertyName": "981.catalog.amenai.net",
    "accountId": "act_1-1TJZFB",
    "contractId": "ctr_1-1TJZH5",
    "groupId": "grp_15225",
    "assetId": "aid_101",
    "versions": {
        "items": [
            {
                "propertyVersion": 2,
                "updatedByUser": "amenai",
                "updatedDate": "2015-03-02T15:06:13Z",
                "productionStatus": "INACTIVE",
                "stagingStatus": "ACTIVE",
                "etag": "5891b5b522d5df08",
                "productId": "prd_Alta",
                "note": "updated caching"
            },
            {
                "propertyVersion": 1,
                "updatedByUser": "afaden",
                "updatedDate": "2015-03-02T15:06:13Z",
                "productionStatus": "ACTIVE",
                "stagingStatus": "INACTIVE",
                "etag": "71573b922a87abc3",
                "productId": "prd_Alta",
                "note": "initial version"
            }
        ]
    }
}`
)

const (
	testCreateVersionRawRequest = `
{
    "createFromVersion": 1,
    "createFromVersionEtag": "2641910c585cf67b"
}
`
	testCreateVersionExpected = "/papi/v1/properties/prp_173136/versions/2?contractId=ctr_1-1TJZH5&groupId=grp_15225"
	testVersionLink = "/papi/v1/properties/prp_173136/versions/2"
)

var (
	testTime = time.Date(2015, 03,02, 15, 06,13, 0, time.UTC)
	testGetVersionExpected = &version.Version{
		PropertyVersion: util.IntPtr(1),
		UpdatedByUser: "afaden",
		UpdatedDate: &testTime,
		ProductionStatus: version.ActiveStatus,
		StagingStatus: version.InactiveStatus,
		Etag: "71573b922a87abc3",
		ProductID: "prd_Alta",
		RuleFormat: "latest",
		Note: "initial version",
	}
	testGetVersionByLinkExpected = &version.Version{
		PropertyVersion: util.IntPtr(1),
		UpdatedByUser: "afaden",
		UpdatedDate: &testTime,
		ProductionStatus: version.ActiveStatus,
		StagingStatus: version.InactiveStatus,
		Etag: "71573b922a87abc3",
		ProductID: "prd_Alta",
		RuleFormat: "latest",
		Note: "initial version",
	}
	testGetLatestVersionExpected = &version.Version{
		PropertyVersion: util.IntPtr(1),
		UpdatedByUser: "afaden",
		UpdatedDate: &testTime,
		ProductionStatus: version.ActiveStatus,
		StagingStatus: version.InactiveStatus,
		Etag: "71573b922a87abc3",
		ProductID: "prd_Alta",
		RuleFormat: "latest",
		Note: "initial version",
	}
	testListVersionsExpected = []*version.Version{
		{
			PropertyVersion: util.IntPtr(2),
			UpdatedByUser: "amenai",
			UpdatedDate: &testTime,
			ProductionStatus: version.InactiveStatus,
			StagingStatus: version.ActiveStatus,
			Etag: "5891b5b522d5df08",
			ProductID: "prd_Alta",
			Note: "updated caching",
		},
		{
			PropertyVersion: util.IntPtr(1),
			UpdatedByUser: "afaden",
			UpdatedDate: &testTime,
			ProductionStatus: version.ActiveStatus,
			StagingStatus: version.InactiveStatus,
			Etag: "71573b922a87abc3",
			ProductID: "prd_Alta",
			Note: "initial version",
		},
	}
)
