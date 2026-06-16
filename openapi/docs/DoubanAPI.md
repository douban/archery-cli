# \DoubanAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DoubanKnoxGenerateTokenCreate**](DoubanAPI.md#DoubanKnoxGenerateTokenCreate) | **Post** /douban/knox/generate-token/ | 
[**DoubanKnoxRevokeAllTokenCreate**](DoubanAPI.md#DoubanKnoxRevokeAllTokenCreate) | **Post** /douban/knox/revoke-all-token/ | 



## DoubanKnoxGenerateTokenCreate

> DoubanKnoxGenerateTokenCreate(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.intra.douban.com/sa/archery-cli"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DoubanAPI.DoubanKnoxGenerateTokenCreate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DoubanAPI.DoubanKnoxGenerateTokenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDoubanKnoxGenerateTokenCreateRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DoubanKnoxRevokeAllTokenCreate

> DoubanKnoxRevokeAllTokenCreate(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.intra.douban.com/sa/archery-cli"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DoubanAPI.DoubanKnoxRevokeAllTokenCreate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DoubanAPI.DoubanKnoxRevokeAllTokenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDoubanKnoxRevokeAllTokenCreateRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

