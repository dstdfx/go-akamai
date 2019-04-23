package testing

import (
	"time"

	"github.com/dstdfx/go-akamai/akamai/papi/v1/activation"
)

const (
	testListActivationsRawResponse = `
{
    "accountId": "act_F-AC-1519477",
    "activations": {
        "items": [
            {
                "activationId": "atv_6804950",
                "activationType": "ACTIVATE",
                "network": "PRODUCTION",
                "note": "test from go",
                "notifyEmails": [
                    "dstdfx@example.ru"
                ],
                "propertyId": "prp_511225",
                "propertyName": "1df95724-f65c-4d52-bc32-0cf20d06763c.dstdfx.space",
                "propertyVersion": 1,
                "status": "ACTIVE",
                "submitDate": "2015-03-02T15:06:13Z",
                "updateDate": "2015-03-02T15:06:13Z"
            }
        ]
    },
    "contractId": "ctr_C-10UBXF3",
    "groupId": "grp_52301"
}
`
	testGetActivationRawResponse = `
{
    "accountId": "act_F-AC-1519477",
    "activations": {
        "items": [
            {
                "activationId": "atv_6804950",
                "activationType": "ACTIVATE",
                "network": "PRODUCTION",
                "note": "test from go",
                "notifyEmails": [
                    "dstdfx@example.ru"
                ],
                "propertyId": "prp_511225",
                "propertyName": "1df95724-f65c-4d52-bc32-0cf20d06763c.dstdfx.space",
                "propertyVersion": 1,
                "status": "ACTIVE",
                "submitDate": "2015-03-02T15:06:13Z",
                "updateDate": "2015-03-02T15:06:13Z"
            }
        ]
    },
    "contractId": "ctr_C-10UBXF3",
    "groupId": "grp_52301"
}
`
	testCancelActivationRawResponse = `
{
    "activations": {
        "items": [
            {
                "activationId": "atv_6804950",
                "propertyName": "1df95724-f65c-4d52-bc32-0cf20d06763c.dstdfx.space",
                "propertyId": "prp_511225",
                "propertyVersion": 1,
                "network": "PRODUCTION",
                "activationType": "ACTIVATE",
                "status": "ABORTED",
                "submitDate": "2015-03-02T15:06:13Z",
                "updateDate": "2015-03-02T15:06:13Z",
                "note": "Sample activation cancellation",
                "accountId": "act_1-1TJZFB",
                "groupId": "grp_15225",
                "notifyEmails": [
                    "dstdfx@example.ru"
                ]
            }
        ]
    }
}
`
	testGetActivationByLinkRawResponse = `
{
    "accountId": "act_F-AC-1519477",
    "activations": {
        "items": [
            {
                "activationId": "atv_6804950",
                "activationType": "ACTIVATE",
                "network": "PRODUCTION",
                "note": "test from go",
                "notifyEmails": [
                    "dstdfx@example.ru"
                ],
                "propertyId": "prp_511225",
                "propertyName": "1df95724-f65c-4d52-bc32-0cf20d06763c.dstdfx.space",
                "propertyVersion": 1,
                "status": "ACTIVE",
                "submitDate": "2015-03-02T15:06:13Z",
                "updateDate": "2015-03-02T15:06:13Z"
            }
        ]
    },
    "contractId": "ctr_C-10UBXF3",
    "groupId": "grp_52301"
}
`
	testCreateActivationRawResponse = `
{
    "activationLink": "/papi/v1/properties/prp_173136/activations/atv_67037?contractId=ctr_1-1TJZFB&groupId=grp_15225"
}
`
)


const (
	testActivationLink = "/papi/v1/properties/prp_511225/activations/atv_6804950"
	testCreateActivationRawRequest = `
{
    "propertyVersion": 1,
    "network": "PRODUCTION",
    "note": "test from go",
    "notifyEmails": [
        "dstdfx@example.ru"
    ],
    "acknowledgeAllWarnings": true
}
`
)

var (
	testTime = time.Date(2015, 03,02, 15, 06,13, 0, time.UTC)
	testGetActivationExpected = &activation.Activation{
		ID:   "atv_6804950",
		Type: activation.TypeActivate,
		Network: activation.NetworkProduction,
		Note: "test from go",
		NotifyEmails: []string{"dstdfx@example.ru"},
		PropertyID: "prp_511225",
		PropertyName: "1df95724-f65c-4d52-bc32-0cf20d06763c.dstdfx.space",
		PropertyVersion: 1,
		Status: activation.StatusActive,
		SubmitDate: &testTime,
		UpdateDate: &testTime,
	}
	testCancelActivationExpected = &activation.Activation{
		ID:   "atv_6804950",
		Type: activation.TypeActivate,
		Network: activation.NetworkProduction,
		Note: "Sample activation cancellation",
		NotifyEmails: []string{"dstdfx@example.ru"},
		PropertyID: "prp_511225",
		PropertyName: "1df95724-f65c-4d52-bc32-0cf20d06763c.dstdfx.space",
		PropertyVersion: 1,
		Status: activation.StatusAborted,
		SubmitDate: &testTime,
		UpdateDate: &testTime,
	}
	testGetActivationByLinkExpected = &activation.Activation{
		ID:   "atv_6804950",
		Type: activation.TypeActivate,
		Network: activation.NetworkProduction,
		Note: "test from go",
		NotifyEmails: []string{"dstdfx@example.ru"},
		PropertyID: "prp_511225",
		PropertyName: "1df95724-f65c-4d52-bc32-0cf20d06763c.dstdfx.space",
		PropertyVersion: 1,
		Status: activation.StatusActive,
		SubmitDate: &testTime,
		UpdateDate: &testTime,
	}
	testListActivationsExpected = []*activation.Activation{
		{
			ID:   "atv_6804950",
			Type: activation.TypeActivate,
			Network: activation.NetworkProduction,
			Note: "test from go",
			NotifyEmails: []string{"dstdfx@example.ru"},
			PropertyID: "prp_511225",
			PropertyName: "1df95724-f65c-4d52-bc32-0cf20d06763c.dstdfx.space",
			PropertyVersion: 1,
			Status: activation.StatusActive,
			SubmitDate: &testTime,
			UpdateDate: &testTime,
		},
	}
	testCreateActivationExpected = `/papi/v1/properties/prp_173136/activations/atv_67037?contractId=ctr_1-1TJZFB&groupId=grp_15225`
)