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

	testSearchVersionsRawResponse = `
{
    "versions": {
        "items": [
            {
                "accountId": "act_F-AC-1519477",
                "assetId": "aid_10627973",
                "contractId": "ctr_C-10UBXF3",
                "groupId": "grp_52301",
                "productionStatus": "ACTIVE",
                "propertyId": "prp_509694",
                "propertyName": "my-property",
                "propertyVersion": 2,
                "stagingStatus": "INACTIVE",
                "updatedByUser": "dstdfx@example.ru",
                "updatedDate": "2015-03-02T15:06:13Z"
            },
            {
                "accountId": "act_F-AC-1519477",
                "assetId": "aid_10627973",
                "contractId": "ctr_C-10UBXF3",
                "groupId": "grp_52301",
                "productionStatus": "INACTIVE",
                "propertyId": "prp_509694",
                "propertyName": "my-property",
                "propertyVersion": 1,
                "stagingStatus": "ACTIVE",
                "updatedByUser": "dstdfx@example.ru",
                "updatedDate": "2015-03-02T15:06:13Z"
            }
        ]
    }
}
`
)

const (
	testCreateVersionRawRequest = `
{
    "createFromVersion": 1,
    "createFromVersionEtag": "2641910c585cf67b"
}
`
	testSearchVersionRawRequest = `
{
    "propertyName": "my-property"
}`
	testCreateVersionExpected = "/papi/v1/properties/prp_173136/versions/2?contractId=ctr_1-1TJZH5&groupId=grp_15225"
	testVersionLink = "/papi/v1/properties/prp_173136/versions/2"
)

var (
	testTime = time.Date(2015, 03,02, 15, 06,13, 0, time.UTC)
	testGetVersionExpected = &version.Version{
		PropertyVersion: util.IntPtr(1),
		UpdatedByUser: "afaden",
		UpdatedDate: &testTime,
		ProductionStatus: version.StatusActive,
		StagingStatus: version.StatusInactive,
		Etag: "71573b922a87abc3",
		ProductID: "prd_Alta",
		RuleFormat: "latest",
		Note: "initial version",
	}
	testGetVersionByLinkExpected = &version.Version{
		PropertyVersion: util.IntPtr(1),
		UpdatedByUser: "afaden",
		UpdatedDate: &testTime,
		ProductionStatus: version.StatusActive,
		StagingStatus: version.StatusInactive,
		Etag: "71573b922a87abc3",
		ProductID: "prd_Alta",
		RuleFormat: "latest",
		Note: "initial version",
	}
	testGetLatestVersionExpected = &version.Version{
		PropertyVersion: util.IntPtr(1),
		UpdatedByUser: "afaden",
		UpdatedDate: &testTime,
		ProductionStatus: version.StatusActive,
		StagingStatus: version.StatusInactive,
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
			ProductionStatus: version.StatusInactive,
			StagingStatus: version.StatusActive,
			Etag: "5891b5b522d5df08",
			ProductID: "prd_Alta",
			Note: "updated caching",
		},
		{
			PropertyVersion: util.IntPtr(1),
			UpdatedByUser: "afaden",
			UpdatedDate: &testTime,
			ProductionStatus: version.StatusActive,
			StagingStatus: version.StatusInactive,
			Etag: "71573b922a87abc3",
			ProductID: "prd_Alta",
			Note: "initial version",
		},
	}
	testSearchVersionsExpected = []*version.Version{
		{
			AccountID: "act_F-AC-1519477",
			AssetID: "aid_10627973",
			ContractID: "ctr_C-10UBXF3",
			GroupID: "grp_52301",
			ProductionStatus: version.StatusActive,
			StagingStatus: version.StatusInactive,
			PropertyID: "prp_509694",
			PropertyName: "my-property",
			PropertyVersion: util.IntPtr(2),
			UpdatedByUser: "dstdfx@example.ru",
			UpdatedDate: &testTime,
		},
		{
			AccountID: "act_F-AC-1519477",
			AssetID: "aid_10627973",
			ContractID: "ctr_C-10UBXF3",
			GroupID: "grp_52301",
			ProductionStatus: version.StatusInactive,
			StagingStatus: version.StatusActive,
			PropertyID: "prp_509694",
			PropertyName: "my-property",
			PropertyVersion: util.IntPtr(1),
			UpdatedByUser: "dstdfx@example.ru",
			UpdatedDate: &testTime,
		},
	}
)
