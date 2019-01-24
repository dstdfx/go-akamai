package testing

import "github.com/dstdfx/go-akamai/akamai/papi/v1/product"

const (
	testListProductsRawResponse = `
{
    "accountId": "act_1-1TJZFB",
    "contractId": "ctr_1-1TJZH5",
    "products": {
        "items": [
            {
                "productName": "Alta",
                "productId": "prd_Alta"
            }
        ]
    }
}
`
)

var testListProductsExpected = []*product.Product{
	{
		ID: "prd_Alta",
		Name: "Alta",
	},
}
