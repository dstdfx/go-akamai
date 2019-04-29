package testing

import (
	"github.com/dstdfx/go-akamai/akamai/ccu/v3/cache"
	"github.com/dstdfx/go-akamai/akamai/internal/util"
)

const (
	testInvalidateCacheByURLRawResponse = `
{
	"title": "",
	"describedBy": "",
    "httpStatus": 201,
    "estimatedSeconds": 5,
    "purgeId": "e535071c-26b2-11e7-94d7-276f2f54d938",
    "supportId": "17PY1492793544958045-219026624",
    "detail": "Request accepted"
}
`

	testInvalidateCacheByURLWithHostnameRawResponse = `
{
	"title": "",
	"describedBy": "",
    "httpStatus": 201,
    "estimatedSeconds": 5,
    "purgeId": "e635071c-26b2-11e7-94d7-276f2f54d938",
    "supportId": "17PY1492793544958045-219026624",
    "detail": "Request accepted"
}
`
	testInvalidateCacheByCPCodeRawResponse = `
{
	"title": "",
	"describedBy": "",
    "httpStatus": 201,
    "estimatedSeconds": 5,
    "purgeId": "e735071c-26b2-11e7-94d7-276f2f54d938",
    "supportId": "17PY1492793544958045-219026624",
    "detail": "Request accepted"
}
`
	testInvalidateCacheByCacheTagRawResponse = `
{
	"title": "",
	"describedBy": "",
    "httpStatus": 201,
    "estimatedSeconds": 5,
    "purgeId": "e835071c-26b2-11e7-94d7-276f2f54d938",
    "supportId": "17PY1492793544958045-219026624",
    "detail": "Request accepted"
}
`

	testDeleteCacheByURLRawResponse = `
{
	"title": "",
	"describedBy": "",
    "httpStatus": 201,
    "estimatedSeconds": 5,
    "purgeId": "e535071c-26b2-11e7-94d7-276f2f54d938",
    "supportId": "17PY1492793544958045-219026624",
    "detail": "Request accepted"
}
`

	testDeleteCacheByURLWithHostnameRawResponse = `
{
	"title": "",
	"describedBy": "",
    "httpStatus": 201,
    "estimatedSeconds": 5,
    "purgeId": "e635071c-26b2-11e7-94d7-276f2f54d938",
    "supportId": "17PY1492793544958045-219026624",
    "detail": "Request accepted"
}
`
	testDeleteCacheByCPCodeRawResponse = `
{
	"title": "",
	"describedBy": "",
    "httpStatus": 201,
    "estimatedSeconds": 5,
    "purgeId": "e735071c-26b2-11e7-94d7-276f2f54d938",
    "supportId": "17PY1492793544958045-219026624",
    "detail": "Request accepted"
}
`
	testDeleteCacheByCacheTagRawResponse = `
{
	"title": "",
	"describedBy": "",
    "httpStatus": 201,
    "estimatedSeconds": 5,
    "purgeId": "e835071c-26b2-11e7-94d7-276f2f54d938",
    "supportId": "17PY1492793544958045-219026624",
    "detail": "Request accepted"
}
`
)


const (
	testInvalidateCacheByURLRawRequest = `
{
	"objects": [
		"https://5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net/file.jpg",
		"https://5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net/video.mp4",
		"https://5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net/video2.mp4"]
}
`
	testInvalidateCacheByURLWithHostnameRawRequest = `
{
	"hostname": "5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net",
	"objects": ["/file.jpg", "/video.mp4", "/video2.mp4"]
}
`
	testInvalidateCacheByCPCodeRawRequest = `
{
	"objects": [123, 321]
}`
	testInvalidateCacheByCacheTagRawRequest = `
{
	"objects": ["5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net"]
}
`
	testDeleteCacheByURLRawRequest = `
{
	"objects": [
		"https://5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net/file.jpg",
		"https://5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net/video.mp4",
		"https://5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net/video2.mp4"]
}
`
	testDeleteCacheByURLWithHostnameRawRequest = `
{
	"hostname": "5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net",
	"objects": ["/file.jpg", "/video.mp4", "/video2.mp4"]
}
`
	testDeleteCacheByCPCodeRawRequest = `
{
	"objects": [123, 321]
}`
	testDeleteCacheByCacheTagRawRequest = `
{
	"objects": ["5136022d-e7a4-45ec-b60c-c41913eff6db.akamaized.net"]
}
`
)


var (
	testInvalidateCacheByURLExpected = &cache.CCUResponse{
		HTTPStatus: util.IntPtr(201),
		DescribedBy: "",
		Detail: "Request accepted",
		EstimatedSeconds: util.IntPtr(5),
		PurgeID: "e535071c-26b2-11e7-94d7-276f2f54d938",
		SupportID: "17PY1492793544958045-219026624",
		Title: "",
	}
	testInvalidateCacheByURLWithHostnameExpected = &cache.CCUResponse{
		HTTPStatus: util.IntPtr(201),
		DescribedBy: "",
		Detail: "Request accepted",
		EstimatedSeconds: util.IntPtr(5),
		PurgeID: "e635071c-26b2-11e7-94d7-276f2f54d938",
		SupportID: "17PY1492793544958045-219026624",
		Title: "",
	}
	testInvalidateCacheByCpCodeExpected = &cache.CCUResponse{
		HTTPStatus: util.IntPtr(201),
		DescribedBy: "",
		Detail: "Request accepted",
		EstimatedSeconds: util.IntPtr(5),
		PurgeID: "e735071c-26b2-11e7-94d7-276f2f54d938",
		SupportID: "17PY1492793544958045-219026624",
		Title: "",
	}
	testInvalidateCacheByCacheTagExpected = &cache.CCUResponse{
		HTTPStatus: util.IntPtr(201),
		DescribedBy: "",
		Detail: "Request accepted",
		EstimatedSeconds: util.IntPtr(5),
		PurgeID: "e835071c-26b2-11e7-94d7-276f2f54d938",
		SupportID: "17PY1492793544958045-219026624",
		Title: "",
	}

	testDeleteCacheByURLExpected = &cache.CCUResponse{
		HTTPStatus: util.IntPtr(201),
		DescribedBy: "",
		Detail: "Request accepted",
		EstimatedSeconds: util.IntPtr(5),
		PurgeID: "e535071c-26b2-11e7-94d7-276f2f54d938",
		SupportID: "17PY1492793544958045-219026624",
		Title: "",
	}
	testDeleteCacheByURLWithHostnameExpected = &cache.CCUResponse{
		HTTPStatus: util.IntPtr(201),
		DescribedBy: "",
		Detail: "Request accepted",
		EstimatedSeconds: util.IntPtr(5),
		PurgeID: "e635071c-26b2-11e7-94d7-276f2f54d938",
		SupportID: "17PY1492793544958045-219026624",
		Title: "",
	}
	testDeleteCacheByCpCodeExpected = &cache.CCUResponse{
		HTTPStatus: util.IntPtr(201),
		DescribedBy: "",
		Detail: "Request accepted",
		EstimatedSeconds: util.IntPtr(5),
		PurgeID: "e735071c-26b2-11e7-94d7-276f2f54d938",
		SupportID: "17PY1492793544958045-219026624",
		Title: "",
	}
	testDeleteCacheByCacheTagExpected = &cache.CCUResponse{
		HTTPStatus: util.IntPtr(201),
		DescribedBy: "",
		Detail: "Request accepted",
		EstimatedSeconds: util.IntPtr(5),
		PurgeID: "e835071c-26b2-11e7-94d7-276f2f54d938",
		SupportID: "17PY1492793544958045-219026624",
		Title: "",
	}
)