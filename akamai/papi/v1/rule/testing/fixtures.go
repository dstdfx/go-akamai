package testing

import (
	"github.com/dstdfx/go-akamai/akamai/papi/v1/rule"
)

const (
	testGetRuleTreeRawResponse = `
{
    "accountId": "act_F-AC-1519477",
    "contractId": "ctr_C-10UBXF3",
    "etag": "d96f6e437110f686a3a58e6ba45e575f0577440d",
    "groupId": "grp_52301",
    "propertyId": "prp_511225",
    "propertyName": "1df95724-f65c-4d52-bc32-0cf20d06763c.dstdfx.space",
    "propertyVersion": 1,
    "ruleFormat": "v2018-09-12",
    "rules": {
        "behaviors": [
            {
                "name": "origin",
                "options": {
                    "cacheKeyHostname": "REQUEST_HOST_HEADER",
                    "compress": true,
                    "enableTrueClientIp": true,
                    "forwardHostHeader": "REQUEST_HOST_HEADER",
                    "hostname": "test.dstdfx.space",
                    "httpPort": 80,
                    "httpsPort": 443,
                    "originCertificate": "",
                    "originSni": true,
                    "originType": "CUSTOMER",
                    "ports": "",
                    "trueClientIpClientSetting": false,
                    "trueClientIpHeader": "True-Client-IP",
                    "verificationMode": "PLATFORM_SETTINGS"
                }
            },
            {
                "name": "cpCode",
                "options": {
                    "value": {
                        "id": 688480
                    }
                }
            },
            {
                "name": "report",
                "options": {
                    "logAcceptLanguage": false,
                    "logCookies": "OFF",
                    "logCustomLogField": false,
                    "logHost": false,
                    "logReferer": false,
                    "logUserAgent": true
                }
            },
            {
                "name": "caching",
                "options": {
                    "behavior": "MAX_AGE",
                    "mustRevalidate": false,
                    "ttl": "2d"
                }
            },
            {
                "name": "tieredDistribution",
                "options": {
                    "enabled": true,
                    "tieredDistributionMap": "CH2"
                }
            },
            {
                "name": "cacheKeyQueryParams",
                "options": {
                    "behavior": "IGNORE_ALL"
                }
            },
            {
                "name": "netSession",
                "options": {
                    "enabled": false
                }
            }
        ],
        "children": [
            {
                "behaviors": [
                    {
                        "name": "largeFileOptimization",
                        "options": {
                            "enablePartialObjectCaching": "PARTIAL_OBJECT_CACHING",
                            "enabled": true,
                            "maximumSize": "16GB",
                            "minimumSize": "100MB",
                            "useVersioning": true
                        }
                    }
                ],
                "children": [],
                "criteria": [
                    {
                        "name": "fileExtension",
                        "options": {
                            "matchCaseSensitive": false,
                            "matchOperator": "IS_ONE_OF",
                            "values": [
                                "exe",
                                "bz2",
                                "dmg",
                                "gz",
                                "iso",
                                "mov",
                                "pkg",
                                "tar",
                                "tgz",
                                "wmv",
                                "wma",
                                "zip",
                                "webp",
                                "jxr",
                                "hdp",
                                "wdp"
                            ]
                        }
                    }
                ],
                "criteriaMustSatisfy": "all",
                "name": "Large File Optimization"
            },
            {
                "behaviors": [
                    {
                        "name": "gzipResponse",
                        "options": {
                            "behavior": "ALWAYS"
                        }
                    }
                ],
                "children": [],
                "criteria": [
                    {
                        "name": "contentType",
                        "options": {
                            "matchCaseSensitive": false,
                            "matchOperator": "IS_ONE_OF",
                            "matchWildcard": true,
                            "values": [
                                "text/html*",
                                "text/css*",
                                "application/x-javascript*"
                            ]
                        }
                    }
                ],
                "criteriaMustSatisfy": "all",
                "name": "Content Compression"
            }
        ],
        "name": "default",
        "options": {}
    },
    "warnings": [
        {
            "detail": "If your 'Origin Server' uses HTTPS, make sure to follow <a href=\"/dl/property-manager/property-manager-help/csh_lookup.html?id=PM_0034\" target=\"_blank\">this procedure</a> to avoid a service outage or a security breach when you rotate your origin's certificate.",
            "errorLocation": "#/rules/behaviors/0",
            "type": "https://problems.luna.akamaiapis.net/papi/v0/validation/validation_message.ssl_delegate_warning_rotate"
        },
        {
            "detail": "'Large File Optimization' does not work correctly if 'Origin Server' does not support etags properly, in which case content corruption may occur.",
            "errorLocation": "#/rules/children/0/behaviors/0",
            "type": "https://problems.luna.akamaiapis.net/papi/v0/validation/incompatible_features"
        }
    ]
}
`
	testUpdateRuleTreeRawResponse = `
{
    "accountId": "act_F-AC-1519477",
    "contractId": "ctr_C-10UBXF3",
    "etag": "d96f6e437110f686a3a58e6ba45e575f0577440d",
    "groupId": "grp_52301",
    "propertyId": "prp_511225",
    "propertyName": "1df95724-f65c-4d52-bc32-0cf20d06763c.dstdfx.space",
    "propertyVersion": 1,
    "ruleFormat": "v2018-09-12",
    "rules": {
        "behaviors": [
            {
                "name": "origin",
                "options": {
                    "cacheKeyHostname": "REQUEST_HOST_HEADER",
                    "compress": true,
                    "enableTrueClientIp": true,
                    "forwardHostHeader": "REQUEST_HOST_HEADER",
                    "hostname": "test.dstdfx.space",
                    "httpPort": 80,
                    "httpsPort": 443,
                    "originCertificate": "",
                    "originSni": true,
                    "originType": "CUSTOMER",
                    "ports": "",
                    "trueClientIpClientSetting": false,
                    "trueClientIpHeader": "True-Client-IP",
                    "verificationMode": "PLATFORM_SETTINGS"
                }
            },
            {
                "name": "cpCode",
                "options": {
                    "value": {
                        "id": 688480
                    }
                }
            },
            {
                "name": "report",
                "options": {
                    "logAcceptLanguage": false,
                    "logCookies": "OFF",
                    "logCustomLogField": false,
                    "logHost": false,
                    "logReferer": false,
                    "logUserAgent": true
                }
            },
            {
                "name": "caching",
                "options": {
                    "behavior": "MAX_AGE",
                    "mustRevalidate": false,
                    "ttl": "7d"
                }
            },
            {
                "name": "tieredDistribution",
                "options": {
                    "enabled": true,
                    "tieredDistributionMap": "CH2"
                }
            },
            {
                "name": "cacheKeyQueryParams",
                "options": {
                    "behavior": "IGNORE_ALL"
                }
            },
            {
                "name": "netSession",
                "options": {
                    "enabled": false
                }
            }
        ],
        "children": [
            {
                "behaviors": [
                    {
                        "name": "largeFileOptimization",
                        "options": {
                            "enablePartialObjectCaching": "PARTIAL_OBJECT_CACHING",
                            "enabled": true,
                            "maximumSize": "16GB",
                            "minimumSize": "100MB",
                            "useVersioning": true
                        }
                    }
                ],
                "children": [],
                "criteria": [
                    {
                        "name": "fileExtension",
                        "options": {
                            "matchCaseSensitive": false,
                            "matchOperator": "IS_ONE_OF",
                            "values": [
                                "exe",
                                "bz2",
                                "dmg",
                                "gz",
                                "iso",
                                "mov",
                                "pkg",
                                "tar",
                                "tgz",
                                "wmv",
                                "wma",
                                "zip",
                                "webp",
                                "jxr",
                                "hdp",
                                "wdp"
                            ]
                        }
                    }
                ],
                "criteriaMustSatisfy": "all",
                "name": "Large File Optimization"
            },
            {
                "behaviors": [
                    {
                        "name": "gzipResponse",
                        "options": {
                            "behavior": "ALWAYS"
                        }
                    }
                ],
                "children": [],
                "criteria": [
                    {
                        "name": "contentType",
                        "options": {
                            "matchCaseSensitive": false,
                            "matchOperator": "IS_ONE_OF",
                            "matchWildcard": true,
                            "values": [
                                "text/html*",
                                "text/css*",
                                "application/x-javascript*"
                            ]
                        }
                    }
                ],
                "criteriaMustSatisfy": "all",
                "name": "Content Compression"
            }
        ],
        "name": "default",
        "options": {}
    },
    "warnings": [
        {
            "detail": "If your 'Origin Server' uses HTTPS, make sure to follow <a href=\"/dl/property-manager/property-manager-help/csh_lookup.html?id=PM_0034\" target=\"_blank\">this procedure</a> to avoid a service outage or a security breach when you rotate your origin's certificate.",
            "errorLocation": "#/rules/behaviors/0",
            "type": "https://problems.luna.akamaiapis.net/papi/v0/validation/validation_message.ssl_delegate_warning_rotate"
        },
        {
            "detail": "'Large File Optimization' does not work correctly if 'Origin Server' does not support etags properly, in which case content corruption may occur.",
            "errorLocation": "#/rules/children/0/behaviors/0",
            "type": "https://problems.luna.akamaiapis.net/papi/v0/validation/incompatible_features"
        }
    ]
}
`

	testUpdateRuleTreeWithErrorsRawResponse = `
{
    "accountId": "act_F-AC-1519477",
    "contractId": "ctr_C-10UBXF3",
    "etag": "d96f6e437110f686a3a58e6ba45e575f0577440d",
    "groupId": "grp_52301",
    "propertyId": "prp_511225",
    "propertyName": "1df95724-f65c-4d52-bc32-0cf20d06763c.dstdfx.space",
    "propertyVersion": 1,
    "ruleFormat": "v2018-09-12",
    "rules": {
        "behaviors": [
            {
                "name": "report",
                "options": {
                    "logAcceptLanguage": false,
                    "logCookies": "OFF",
                    "logCustomLogField": false,
                    "logHost": false,
                    "logReferer": false,
                    "logUserAgent": true
                }
            },
            {
                "name": "caching",
                "options": {
                    "behavior": "MAX_AGE",
                    "mustRevalidate": false,
                    "ttl": "7d"
                }
            },
            {
                "name": "tieredDistribution",
                "options": {
                    "enabled": true,
                    "tieredDistributionMap": "CH2"
                }
            },
            {
                "name": "cacheKeyQueryParams",
                "options": {
                    "behavior": "IGNORE_ALL"
                }
            },
            {
                "name": "netSession",
                "options": {
                    "enabled": false
                }
            }
        ],
        "children": [
            {
                "behaviors": [
                    {
                        "name": "largeFileOptimization",
                        "options": {
                            "enablePartialObjectCaching": "PARTIAL_OBJECT_CACHING",
                            "enabled": true,
                            "maximumSize": "16GB",
                            "minimumSize": "100MB",
                            "useVersioning": true
                        }
                    }
                ],
                "children": [],
                "criteria": [
                    {
                        "name": "fileExtension",
                        "options": {
                            "matchCaseSensitive": false,
                            "matchOperator": "IS_ONE_OF",
                            "values": [
                                "exe",
                                "bz2",
                                "dmg",
                                "gz",
                                "iso",
                                "mov",
                                "pkg",
                                "tar",
                                "tgz",
                                "wmv",
                                "wma",
                                "zip",
                                "webp",
                                "jxr",
                                "hdp",
                                "wdp"
                            ]
                        }
                    }
                ],
                "criteriaMustSatisfy": "all",
                "name": "Large File Optimization"
            },
            {
                "behaviors": [
                    {
                        "name": "gzipResponse",
                        "options": {
                            "behavior": "ALWAYS"
                        }
                    }
                ],
                "children": [],
                "criteria": [
                    {
                        "name": "contentType",
                        "options": {
                            "matchCaseSensitive": false,
                            "matchOperator": "IS_ONE_OF",
                            "matchWildcard": true,
                            "values": [
                                "text/html*",
                                "text/css*",
                                "application/x-javascript*"
                            ]
                        }
                    }
                ],
                "criteriaMustSatisfy": "all",
                "name": "Content Compression"
            }
        ],
        "name": "default",
        "options": {}
    },
    "warnings": [
        {
            "detail": "If your 'Origin Server' uses HTTPS, make sure to follow <a href=\"/dl/property-manager/property-manager-help/csh_lookup.html?id=PM_0034\" target=\"_blank\">this procedure</a> to avoid a service outage or a security breach when you rotate your origin's certificate.",
            "errorLocation": "#/rules/behaviors/0",
            "type": "https://problems.luna.akamaiapis.net/papi/v0/validation/validation_message.ssl_delegate_warning_rotate"
        },
        {
            "detail": "'Large File Optimization' does not work correctly if 'Origin Server' does not support etags properly, in which case content corruption may occur.",
            "errorLocation": "#/rules/children/0/behaviors/0",
            "type": "https://problems.luna.akamaiapis.net/papi/v0/validation/incompatible_features"
        }
    ],
    "errors": [
        {
            "type": "/papi/v1/errors/validation.required_behavior",
            "title": "Missing required behavior in default rule",
            "detail": "In order for this property to work correctly behavior Content Provider Code needs to be present in the default section",
            "instance": "/papi/v1/properties/prp_173136/versions/3/rules#err_100",
            "behaviorName": "cpCode"
        },
        {
            "type": "/papi/v1/errors/validation.required_behavior",
            "title": "Missing required behavior in default rule",
            "detail": "In order for this property to work correctly behavior Origin needs to be present in the default section",
            "instance": "/papi/v1/properties/prp_173136/versions/3/rules#err_101",
            "behaviorName": "origin"
        }
    ]
}
`
)


const (
	testUpdateRuleTreeRawRequest = `
{
   "rules":{
      "behaviors":[
         {
            "name":"origin",
            "options":{
               "cacheKeyHostname":"REQUEST_HOST_HEADER",
               "compress":true,
               "enableTrueClientIp":true,
               "forwardHostHeader":"REQUEST_HOST_HEADER",
               "hostname":"test.dstdfx.space",
               "httpPort":80,
               "httpsPort":443,
               "originCertificate":"",
               "originSni":true,
               "originType":"CUSTOMER",
               "ports":"",
               "trueClientIpClientSetting":false,
               "trueClientIpHeader":"True-Client-IP",
               "verificationMode":"PLATFORM_SETTINGS"
            }
         },
         {
            "name":"cpCode",
            "options":{
               "value":{
                  "id":688480
               }
            }
         },
         {
            "name":"report",
            "options":{
               "logAcceptLanguage":false,
               "logCookies":"OFF",
               "logCustomLogField":false,
               "logHost":false,
               "logReferer":false,
               "logUserAgent":true
            }
         },
         {
            "name":"caching",
            "options":{
               "behavior":"MAX_AGE",
               "mustRevalidate":false,
               "ttl":"7d"
            }
         },
         {
            "name":"tieredDistribution",
            "options":{
               "enabled":true,
               "tieredDistributionMap":"CH2"
            }
         },
         {
            "name":"cacheKeyQueryParams",
            "options":{
               "behavior":"IGNORE_ALL"
            }
         },
         {
            "name":"netSession",
            "options":{
               "enabled":false
            }
         }
      ],
      "children":[
         {
            "behaviors":[
               {
                  "name":"largeFileOptimization",
                  "options":{
                     "enablePartialObjectCaching":"PARTIAL_OBJECT_CACHING",
                     "enabled":true,
                     "maximumSize":"16GB",
                     "minimumSize":"100MB",
                     "useVersioning":true
                  }
               }
            ],
            "criteria":[
               {
                  "name":"fileExtension",
                  "options":{
                     "matchCaseSensitive":false,
                     "matchOperator":"IS_ONE_OF",
                     "values":[
                        "exe",
                        "bz2",
                        "dmg",
                        "gz",
                        "iso",
                        "mov",
                        "pkg",
                        "tar",
                        "tgz",
                        "wmv",
                        "wma",
                        "zip",
                        "webp",
                        "jxr",
                        "hdp",
                        "wdp"
                     ]
                  }
               }
            ],
            "criteriaMustSatisfy":"all",
            "name":"Large File Optimization"
         },
         {
            "behaviors":[
               {
                  "name":"gzipResponse",
                  "options":{
                     "behavior":"ALWAYS"
                  }
               }
            ],
            "criteria":[
               {
                  "name":"contentType",
                  "options":{
                     "matchCaseSensitive":false,
                     "matchOperator":"IS_ONE_OF",
                     "matchWildcard":true,
                     "values":[
                        "text/html*",
                        "text/css*",
                        "application/x-javascript*"
                     ]
                  }
               }
            ],
            "criteriaMustSatisfy":"all",
            "name":"Content Compression"
         }
      ],
      "name":"default",
      "options":{}
   }
}
`
	testUpdateRuleTreeWithErrorsRawRequest = `
{
   "rules":{
      "behaviors":[
         {
            "name":"report",
            "options":{
               "logAcceptLanguage":false,
               "logCookies":"OFF",
               "logCustomLogField":false,
               "logHost":false,
               "logReferer":false,
               "logUserAgent":true
            }
         },
         {
            "name":"caching",
            "options":{
               "behavior":"MAX_AGE",
               "mustRevalidate":false,
               "ttl":"7d"
            }
         },
         {
            "name":"tieredDistribution",
            "options":{
               "enabled":true,
               "tieredDistributionMap":"CH2"
            }
         },
         {
            "name":"cacheKeyQueryParams",
            "options":{
               "behavior":"IGNORE_ALL"
            }
         },
         {
            "name":"netSession",
            "options":{
               "enabled":false
            }
         }
      ],
      "children":[
         {
            "behaviors":[
               {
                  "name":"largeFileOptimization",
                  "options":{
                     "enablePartialObjectCaching":"PARTIAL_OBJECT_CACHING",
                     "enabled":true,
                     "maximumSize":"16GB",
                     "minimumSize":"100MB",
                     "useVersioning":true
                  }
               }
            ],
            "criteria":[
               {
                  "name":"fileExtension",
                  "options":{
                     "matchCaseSensitive":false,
                     "matchOperator":"IS_ONE_OF",
                     "values":[
                        "exe",
                        "bz2",
                        "dmg",
                        "gz",
                        "iso",
                        "mov",
                        "pkg",
                        "tar",
                        "tgz",
                        "wmv",
                        "wma",
                        "zip",
                        "webp",
                        "jxr",
                        "hdp",
                        "wdp"
                     ]
                  }
               }
            ],
            "criteriaMustSatisfy":"all",
            "name":"Large File Optimization"
         },
         {
            "behaviors":[
               {
                  "name":"gzipResponse",
                  "options":{
                     "behavior":"ALWAYS"
                  }
               }
            ],
            "criteria":[
               {
                  "name":"contentType",
                  "options":{
                     "matchCaseSensitive":false,
                     "matchOperator":"IS_ONE_OF",
                     "matchWildcard":true,
                     "values":[
                        "text/html*",
                        "text/css*",
                        "application/x-javascript*"
                     ]
                  }
               }
            ],
            "criteriaMustSatisfy":"all",
            "name":"Content Compression"
         }
      ],
      "name":"default",
      "options":{}
   }
}
`
)

var (
	testGetRuleTreeExpected = &rule.Tree{
		RuleFormat: "v2018-09-12",
		Rule: &rule.Rule{
			Name: "default",
			Options: &rule.TopLevelRuleOptions{},
			Behaviors: []*rule.Behavior{
				{
					Name: "origin",
					Options: rule.OptionValue{
						"cacheKeyHostname": "REQUEST_HOST_HEADER",
						"compress": true,
						"enableTrueClientIp": true,
						"forwardHostHeader": "REQUEST_HOST_HEADER",
						"hostname": "test.dstdfx.space",
						"httpPort": 80,
						"httpsPort": 443,
						"originCertificate": "",
						"originSni": true,
						"originType": "CUSTOMER",
						"ports": "",
						"trueClientIpClientSetting": false,
						"trueClientIpHeader": "True-Client-IP",
						"verificationMode": "PLATFORM_SETTINGS",
					},
				},
				{
					Name: "cpCode",
					Options: rule.OptionValue{
						"value": rule.OptionValue{
							"id": 688480,
						},
					},
				},
				{
					Name: "report",
					Options: rule.OptionValue{
						"logAcceptLanguage": false,
						"logCookies": "OFF",
						"logCustomLogField": false,
						"logHost": false,
						"logReferer": false,
						"logUserAgent": true,
					},
				},
				{
					Name: "caching",
					Options: rule.OptionValue{
						"behavior": "MAX_AGE",
						"mustRevalidate": false,
						"ttl": "2d",
					},
				},
				{
					Name: "tieredDistribution",
					Options: rule.OptionValue{
						"enabled": true,
						"tieredDistributionMap": "CH2",
					},
				},
				{
					Name: "cacheKeyQueryParams",
					Options: rule.OptionValue{
						"behavior": "IGNORE_ALL",
					},
				},
				{
					Name: "netSession",
					Options: rule.OptionValue{
						"enabled": false,
					},
				},
			},
			Children: []*rule.Rule{
				{
					Name: "Large File Optimization",
					CriteriaMustSatisfy: rule.CriteriaMustSatisfyAll,
					Behaviors: []*rule.Behavior{
						{
							Name: "largeFileOptimization",
							Options: rule.OptionValue{
								"enablePartialObjectCaching": "PARTIAL_OBJECT_CACHING",
								"enabled": true,
								"maximumSize": "16GB",
								"minimumSize": "100MB",
								"useVersioning": true,
							},
						},
					},
					Criteria: []*rule.Criteria{
						{
							Name: "fileExtension",
							Options: rule.OptionValue{
								"matchCaseSensitive": false,
								"matchOperator": "IS_ONE_OF",
								"values": []string{
									"exe",
									"bz2",
									"dmg",
									"gz",
									"iso",
									"mov",
									"pkg",
									"tar",
									"tgz",
									"wmv",
									"wma",
									"zip",
									"webp",
									"jxr",
									"hdp",
									"wdp",
								},
							},

						},
					},
					Children: []*rule.Rule{},
				},
				{
					Name: "Content Compression",
					CriteriaMustSatisfy: rule.CriteriaMustSatisfyAll,
					Behaviors: []*rule.Behavior{
						{
							Name: "gzipResponse",
							Options: rule.OptionValue{
								"behavior": "ALWAYS",
							},
						},
					},
					Criteria: []*rule.Criteria{
						{
							Name: "contentType",
							Options: rule.OptionValue{
								"matchCaseSensitive": false,
								"matchOperator": "IS_ONE_OF",
								"matchWildcard": true,
								"values": []string{
									"text/html*",
									"text/css*",
									"application/x-javascript*",
								},
							},
						},
					},
					Children: []*rule.Rule{},
				},
			},
		},
		Warnings: []*rule.Warning{
			{
				Type: "https://problems.luna.akamaiapis.net/papi/v0/validation/validation_message.ssl_delegate_warning_rotate",
				ErrorLocation: "#/rules/behaviors/0",
				Detail: "If your 'Origin Server' uses HTTPS, make sure to follow <a href=\"/dl/property-manager/property-manager-help/csh_lookup.html?id=PM_0034\" target=\"_blank\">this procedure</a> to avoid a service outage or a security breach when you rotate your origin's certificate.",
			},
			{
				Type: "https://problems.luna.akamaiapis.net/papi/v0/validation/incompatible_features",
				ErrorLocation: "#/rules/children/0/behaviors/0",
				Detail: "'Large File Optimization' does not work correctly if 'Origin Server' does not support etags properly, in which case content corruption may occur.",
			},
		},
	}

	testUpdateRuleTreeExpected = &rule.Tree{
		RuleFormat: "v2018-09-12",
		Rule: &rule.Rule{
			Name: "default",
			Options: &rule.TopLevelRuleOptions{},
			Behaviors: []*rule.Behavior{
				{
					Name: "origin",
					Options: rule.OptionValue{
						"cacheKeyHostname": "REQUEST_HOST_HEADER",
						"compress": true,
						"enableTrueClientIp": true,
						"forwardHostHeader": "REQUEST_HOST_HEADER",
						"hostname": "test.dstdfx.space",
						"httpPort": 80,
						"httpsPort": 443,
						"originCertificate": "",
						"originSni": true,
						"originType": "CUSTOMER",
						"ports": "",
						"trueClientIpClientSetting": false,
						"trueClientIpHeader": "True-Client-IP",
						"verificationMode": "PLATFORM_SETTINGS",
					},
				},
				{
					Name: "cpCode",
					Options: rule.OptionValue{
						"value": rule.OptionValue{
							"id": 688480,
						},
					},
				},
				{
					Name: "report",
					Options: rule.OptionValue{
						"logAcceptLanguage": false,
						"logCookies": "OFF",
						"logCustomLogField": false,
						"logHost": false,
						"logReferer": false,
						"logUserAgent": true,
					},
				},
				{
					Name: "caching",
					Options: rule.OptionValue{
						"behavior": "MAX_AGE",
						"mustRevalidate": false,
						"ttl": "7d",
					},
				},
				{
					Name: "tieredDistribution",
					Options: rule.OptionValue{
						"enabled": true,
						"tieredDistributionMap": "CH2",
					},
				},
				{
					Name: "cacheKeyQueryParams",
					Options: rule.OptionValue{
						"behavior": "IGNORE_ALL",
					},
				},
				{
					Name: "netSession",
					Options: rule.OptionValue{
						"enabled": false,
					},
				},
			},
			Children: []*rule.Rule{
				{
					Name: "Large File Optimization",
					CriteriaMustSatisfy: rule.CriteriaMustSatisfyAll,
					Behaviors: []*rule.Behavior{
						{
							Name: "largeFileOptimization",
							Options: rule.OptionValue{
								"enablePartialObjectCaching": "PARTIAL_OBJECT_CACHING",
								"enabled": true,
								"maximumSize": "16GB",
								"minimumSize": "100MB",
								"useVersioning": true,
							},
						},
					},
					Criteria: []*rule.Criteria{
						{
							Name: "fileExtension",
							Options: rule.OptionValue{
								"matchCaseSensitive": false,
								"matchOperator": "IS_ONE_OF",
								"values": []string{
									"exe",
									"bz2",
									"dmg",
									"gz",
									"iso",
									"mov",
									"pkg",
									"tar",
									"tgz",
									"wmv",
									"wma",
									"zip",
									"webp",
									"jxr",
									"hdp",
									"wdp",
								},
							},

						},
					},
					Children: []*rule.Rule{},
				},
				{
					Name: "Content Compression",
					CriteriaMustSatisfy: rule.CriteriaMustSatisfyAll,
					Behaviors: []*rule.Behavior{
						{
							Name: "gzipResponse",
							Options: rule.OptionValue{
								"behavior": "ALWAYS",
							},
						},
					},
					Criteria: []*rule.Criteria{
						{
							Name: "contentType",
							Options: rule.OptionValue{
								"matchCaseSensitive": false,
								"matchOperator": "IS_ONE_OF",
								"matchWildcard": true,
								"values": []string{
									"text/html*",
									"text/css*",
									"application/x-javascript*",
								},
							},
						},
					},
					Children: []*rule.Rule{},
				},
			},
		},
		Warnings: []*rule.Warning{
			{
				Type: "https://problems.luna.akamaiapis.net/papi/v0/validation/validation_message.ssl_delegate_warning_rotate",
				ErrorLocation: "#/rules/behaviors/0",
				Detail: "If your 'Origin Server' uses HTTPS, make sure to follow <a href=\"/dl/property-manager/property-manager-help/csh_lookup.html?id=PM_0034\" target=\"_blank\">this procedure</a> to avoid a service outage or a security breach when you rotate your origin's certificate.",
			},
			{
				Type: "https://problems.luna.akamaiapis.net/papi/v0/validation/incompatible_features",
				ErrorLocation: "#/rules/children/0/behaviors/0",
				Detail: "'Large File Optimization' does not work correctly if 'Origin Server' does not support etags properly, in which case content corruption may occur.",
			},
		},
	}
	testUpdateRuleTreeWithErrorsExpected = &rule.Tree{
		RuleFormat: "v2018-09-12",
		Rule: &rule.Rule{
			Name: "default",
			Options: &rule.TopLevelRuleOptions{},
			Behaviors: []*rule.Behavior{
				{
					Name: "report",
					Options: rule.OptionValue{
						"logAcceptLanguage": false,
						"logCookies": "OFF",
						"logCustomLogField": false,
						"logHost": false,
						"logReferer": false,
						"logUserAgent": true,
					},
				},
				{
					Name: "caching",
					Options: rule.OptionValue{
						"behavior": "MAX_AGE",
						"mustRevalidate": false,
						"ttl": "7d",
					},
				},
				{
					Name: "tieredDistribution",
					Options: rule.OptionValue{
						"enabled": true,
						"tieredDistributionMap": "CH2",
					},
				},
				{
					Name: "cacheKeyQueryParams",
					Options: rule.OptionValue{
						"behavior": "IGNORE_ALL",
					},
				},
				{
					Name: "netSession",
					Options: rule.OptionValue{
						"enabled": false,
					},
				},
			},
			Children: []*rule.Rule{
				{
					Name: "Large File Optimization",
					CriteriaMustSatisfy: rule.CriteriaMustSatisfyAll,
					Behaviors: []*rule.Behavior{
						{
							Name: "largeFileOptimization",
							Options: rule.OptionValue{
								"enablePartialObjectCaching": "PARTIAL_OBJECT_CACHING",
								"enabled": true,
								"maximumSize": "16GB",
								"minimumSize": "100MB",
								"useVersioning": true,
							},
						},
					},
					Criteria: []*rule.Criteria{
						{
							Name: "fileExtension",
							Options: rule.OptionValue{
								"matchCaseSensitive": false,
								"matchOperator": "IS_ONE_OF",
								"values": []string{
									"exe",
									"bz2",
									"dmg",
									"gz",
									"iso",
									"mov",
									"pkg",
									"tar",
									"tgz",
									"wmv",
									"wma",
									"zip",
									"webp",
									"jxr",
									"hdp",
									"wdp",
								},
							},

						},
					},
					Children: []*rule.Rule{},
				},
				{
					Name: "Content Compression",
					CriteriaMustSatisfy: rule.CriteriaMustSatisfyAll,
					Behaviors: []*rule.Behavior{
						{
							Name: "gzipResponse",
							Options: rule.OptionValue{
								"behavior": "ALWAYS",
							},
						},
					},
					Criteria: []*rule.Criteria{
						{
							Name: "contentType",
							Options: rule.OptionValue{
								"matchCaseSensitive": false,
								"matchOperator": "IS_ONE_OF",
								"matchWildcard": true,
								"values": []string{
									"text/html*",
									"text/css*",
									"application/x-javascript*",
								},
							},
						},
					},
					Children: []*rule.Rule{},
				},
			},
		},
		Warnings: []*rule.Warning{
			{
				Type: "https://problems.luna.akamaiapis.net/papi/v0/validation/validation_message.ssl_delegate_warning_rotate",
				ErrorLocation: "#/rules/behaviors/0",
				Detail: "If your 'Origin Server' uses HTTPS, make sure to follow <a href=\"/dl/property-manager/property-manager-help/csh_lookup.html?id=PM_0034\" target=\"_blank\">this procedure</a> to avoid a service outage or a security breach when you rotate your origin's certificate.",
			},
			{
				Type: "https://problems.luna.akamaiapis.net/papi/v0/validation/incompatible_features",
				ErrorLocation: "#/rules/children/0/behaviors/0",
				Detail: "'Large File Optimization' does not work correctly if 'Origin Server' does not support etags properly, in which case content corruption may occur.",
			},
		},
		Errors: []*rule.Error{
			{
				Type: "/papi/v1/errors/validation.required_behavior",
				Title: "Missing required behavior in default rule",
				Detail: "In order for this property to work correctly behavior Content Provider Code needs to be present in the default section",
				Instance: "/papi/v1/properties/prp_173136/versions/3/rules#err_100",
				BehaviorName: "cpCode",
			},
			{
				Type: "/papi/v1/errors/validation.required_behavior",
				Title: "Missing required behavior in default rule",
				Detail: "In order for this property to work correctly behavior Origin needs to be present in the default section",
				Instance: "/papi/v1/properties/prp_173136/versions/3/rules#err_101",
				BehaviorName: "origin",
			},
		},
	}
)
