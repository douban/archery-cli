# \ApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAuthTokenCreate**](ApiAPI.md#ApiAuthTokenCreate) | **Post** /api/auth/token/ | 
[**ApiAuthTokenRefreshCreate**](ApiAPI.md#ApiAuthTokenRefreshCreate) | **Post** /api/auth/token/refresh/ | 
[**ApiAuthTokenVerifyCreate**](ApiAPI.md#ApiAuthTokenVerifyCreate) | **Post** /api/auth/token/verify/ | 
[**ApiSchemaRetrieve**](ApiAPI.md#ApiSchemaRetrieve) | **Get** /api/schema/ | 
[**ApiV1InstanceCreate**](ApiAPI.md#ApiV1InstanceCreate) | **Post** /api/v1/instance/ | 创建实例
[**ApiV1InstanceDestroy**](ApiAPI.md#ApiV1InstanceDestroy) | **Delete** /api/v1/instance/{id}/ | 删除实例
[**ApiV1InstanceList**](ApiAPI.md#ApiV1InstanceList) | **Get** /api/v1/instance/ | 实例清单
[**ApiV1InstanceRdsCreate**](ApiAPI.md#ApiV1InstanceRdsCreate) | **Post** /api/v1/instance/rds/ | 创建AliyunRDS
[**ApiV1InstanceRdsList**](ApiAPI.md#ApiV1InstanceRdsList) | **Get** /api/v1/instance/rds/ | AliyunRDS清单
[**ApiV1InstanceResourceCreate**](ApiAPI.md#ApiV1InstanceResourceCreate) | **Post** /api/v1/instance/resource/ | 实例资源
[**ApiV1InstanceTableInstancesCreate**](ApiAPI.md#ApiV1InstanceTableInstancesCreate) | **Post** /api/v1/instance/table-instances/ | 按表名查询所属实例
[**ApiV1InstanceTunnelCreate**](ApiAPI.md#ApiV1InstanceTunnelCreate) | **Post** /api/v1/instance/tunnel/ | 创建隧道
[**ApiV1InstanceTunnelList**](ApiAPI.md#ApiV1InstanceTunnelList) | **Get** /api/v1/instance/tunnel/ | 隧道清单
[**ApiV1InstanceUpdate**](ApiAPI.md#ApiV1InstanceUpdate) | **Put** /api/v1/instance/{id}/ | 更新实例
[**ApiV1SqlqueryDescribetableCreate**](ApiAPI.md#ApiV1SqlqueryDescribetableCreate) | **Post** /api/v1/sqlquery/describetable/ | SQLQuery 表结构
[**ApiV1SqlqueryExecuteCreate**](ApiAPI.md#ApiV1SqlqueryExecuteCreate) | **Post** /api/v1/sqlquery/execute/ | SQLQuery 执行查询
[**ApiV1SqlqueryFavoritesCreate**](ApiAPI.md#ApiV1SqlqueryFavoritesCreate) | **Post** /api/v1/sqlquery/favorites/ | SQLQuery 收藏/取消收藏
[**ApiV1SqlqueryInstancesRetrieve**](ApiAPI.md#ApiV1SqlqueryInstancesRetrieve) | **Get** /api/v1/sqlquery/instances/ | SQLQuery 可访问实例列表
[**ApiV1SqlqueryLogsRetrieve**](ApiAPI.md#ApiV1SqlqueryLogsRetrieve) | **Get** /api/v1/sqlquery/logs/ | SQLQuery 历史查询记录
[**ApiV1SqlqueryResourcesRetrieve**](ApiAPI.md#ApiV1SqlqueryResourcesRetrieve) | **Get** /api/v1/sqlquery/resources/ | SQLQuery 实例资源列表
[**ApiV1User2faCreate**](ApiAPI.md#ApiV1User2faCreate) | **Post** /api/v1/user/2fa/ | 配置2fa
[**ApiV1User2faSaveCreate**](ApiAPI.md#ApiV1User2faSaveCreate) | **Post** /api/v1/user/2fa/save/ | 保存2fa配置
[**ApiV1User2faStateCreate**](ApiAPI.md#ApiV1User2faStateCreate) | **Post** /api/v1/user/2fa/state/ | 查询2fa配置情况
[**ApiV1User2faVerifyCreate**](ApiAPI.md#ApiV1User2faVerifyCreate) | **Post** /api/v1/user/2fa/verify/ | 检验2fa密码
[**ApiV1UserAuthCreate**](ApiAPI.md#ApiV1UserAuthCreate) | **Post** /api/v1/user/auth/ | 用户认证校验
[**ApiV1UserCreate**](ApiAPI.md#ApiV1UserCreate) | **Post** /api/v1/user/ | 创建用户
[**ApiV1UserDestroy**](ApiAPI.md#ApiV1UserDestroy) | **Delete** /api/v1/user/{id}/ | 删除用户
[**ApiV1UserGroupCreate**](ApiAPI.md#ApiV1UserGroupCreate) | **Post** /api/v1/user/group/ | 创建用户组
[**ApiV1UserGroupDestroy**](ApiAPI.md#ApiV1UserGroupDestroy) | **Delete** /api/v1/user/group/{id}/ | 删除用户组
[**ApiV1UserGroupList**](ApiAPI.md#ApiV1UserGroupList) | **Get** /api/v1/user/group/ | 用户组清单
[**ApiV1UserGroupUpdate**](ApiAPI.md#ApiV1UserGroupUpdate) | **Put** /api/v1/user/group/{id}/ | 更新用户组
[**ApiV1UserList**](ApiAPI.md#ApiV1UserList) | **Get** /api/v1/user/ | 用户清单
[**ApiV1UserResourcegroupCreate**](ApiAPI.md#ApiV1UserResourcegroupCreate) | **Post** /api/v1/user/resourcegroup/ | 创建资源组
[**ApiV1UserResourcegroupDestroy**](ApiAPI.md#ApiV1UserResourcegroupDestroy) | **Delete** /api/v1/user/resourcegroup/{id}/ | 删除资源组
[**ApiV1UserResourcegroupList**](ApiAPI.md#ApiV1UserResourcegroupList) | **Get** /api/v1/user/resourcegroup/ | 资源组清单
[**ApiV1UserResourcegroupUpdate**](ApiAPI.md#ApiV1UserResourcegroupUpdate) | **Put** /api/v1/user/resourcegroup/{id}/ | 更新资源组
[**ApiV1UserUpdate**](ApiAPI.md#ApiV1UserUpdate) | **Put** /api/v1/user/{id}/ | 更新用户
[**ApiV1WorkflowAuditCreate**](ApiAPI.md#ApiV1WorkflowAuditCreate) | **Post** /api/v1/workflow/audit/ | 审核工单
[**ApiV1WorkflowAuditlistCreate**](ApiAPI.md#ApiV1WorkflowAuditlistCreate) | **Post** /api/v1/workflow/auditlist/ | 待审核清单
[**ApiV1WorkflowCreate**](ApiAPI.md#ApiV1WorkflowCreate) | **Post** /api/v1/workflow/ | 提交SQL上线工单
[**ApiV1WorkflowExecuteCreate**](ApiAPI.md#ApiV1WorkflowExecuteCreate) | **Post** /api/v1/workflow/execute/ | 执行工单
[**ApiV1WorkflowList**](ApiAPI.md#ApiV1WorkflowList) | **Get** /api/v1/workflow/ | SQL上线工单清单
[**ApiV1WorkflowLogCreate**](ApiAPI.md#ApiV1WorkflowLogCreate) | **Post** /api/v1/workflow/log/ | 工单日志
[**ApiV1WorkflowSqlcheckCreate**](ApiAPI.md#ApiV1WorkflowSqlcheckCreate) | **Post** /api/v1/workflow/sqlcheck/ | SQL检查



## ApiAuthTokenCreate

> TokenObtainPair ApiAuthTokenCreate(ctx).TokenObtainPair(tokenObtainPair).Execute()





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
	tokenObtainPair := *openapiclient.NewTokenObtainPair("Username_example", "Password_example", "Access_example", "Refresh_example") // TokenObtainPair | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiAuthTokenCreate(context.Background()).TokenObtainPair(tokenObtainPair).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiAuthTokenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiAuthTokenCreate`: TokenObtainPair
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiAuthTokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAuthTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenObtainPair** | [**TokenObtainPair**](TokenObtainPair.md) |  | 

### Return type

[**TokenObtainPair**](TokenObtainPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAuthTokenRefreshCreate

> TokenRefresh ApiAuthTokenRefreshCreate(ctx).TokenRefresh(tokenRefresh).Execute()





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
	tokenRefresh := *openapiclient.NewTokenRefresh("Access_example", "Refresh_example") // TokenRefresh | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiAuthTokenRefreshCreate(context.Background()).TokenRefresh(tokenRefresh).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiAuthTokenRefreshCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiAuthTokenRefreshCreate`: TokenRefresh
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiAuthTokenRefreshCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAuthTokenRefreshCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRefresh** | [**TokenRefresh**](TokenRefresh.md) |  | 

### Return type

[**TokenRefresh**](TokenRefresh.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAuthTokenVerifyCreate

> TokenVerify ApiAuthTokenVerifyCreate(ctx).TokenVerify(tokenVerify).Execute()





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
	tokenVerify := *openapiclient.NewTokenVerify("Token_example") // TokenVerify | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiAuthTokenVerifyCreate(context.Background()).TokenVerify(tokenVerify).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiAuthTokenVerifyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiAuthTokenVerifyCreate`: TokenVerify
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiAuthTokenVerifyCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAuthTokenVerifyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenVerify** | [**TokenVerify**](TokenVerify.md) |  | 

### Return type

[**TokenVerify**](TokenVerify.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSchemaRetrieve

> map[string]interface{} ApiSchemaRetrieve(ctx).Format(format).Lang(lang).Execute()





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
	format := "format_example" // string |  (optional)
	lang := "lang_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiSchemaRetrieve(context.Background()).Format(format).Lang(lang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiSchemaRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiSchemaRetrieve`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiSchemaRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSchemaRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** |  | 
 **lang** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.oai.openapi, application/yaml, application/vnd.oai.openapi+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1InstanceCreate

> Instance ApiV1InstanceCreate(ctx).Instance(instance).Execute()

创建实例



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.intra.douban.com/sa/archery-cli"
)

func main() {
	instance := *openapiclient.NewInstance(int32(123), "InstanceName_example", openapiclient.TypeA85Enum("master"), openapiclient.DbTypeEnum("mysql"), "Host_example", time.Now(), time.Now()) // Instance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1InstanceCreate(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1InstanceCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1InstanceCreate`: Instance
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1InstanceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1InstanceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**Instance**](Instance.md) |  | 

### Return type

[**Instance**](Instance.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1InstanceDestroy

> ApiV1InstanceDestroy(ctx, id).Execute()

删除实例



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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApiAPI.ApiV1InstanceDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1InstanceDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1InstanceDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ApiV1InstanceList

> PaginatedInstanceList ApiV1InstanceList(ctx).Page(page).Size(size).Execute()

实例清单



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
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	size := int32(56) // int32 | Number of results to return per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1InstanceList(context.Background()).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1InstanceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1InstanceList`: PaginatedInstanceList
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1InstanceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1InstanceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 
 **size** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedInstanceList**](PaginatedInstanceList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1InstanceRdsCreate

> AliyunRds ApiV1InstanceRdsCreate(ctx).AliyunRds(aliyunRds).Execute()

创建AliyunRDS



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
	aliyunRds := *openapiclient.NewAliyunRds(int32(123), "RdsDbinstanceid_example", int32(123), *openapiclient.NewCloudAccessKey(int32(123), "KeyId_example", "KeySecret_example")) // AliyunRds | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1InstanceRdsCreate(context.Background()).AliyunRds(aliyunRds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1InstanceRdsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1InstanceRdsCreate`: AliyunRds
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1InstanceRdsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1InstanceRdsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aliyunRds** | [**AliyunRds**](AliyunRds.md) |  | 

### Return type

[**AliyunRds**](AliyunRds.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1InstanceRdsList

> PaginatedAliyunRdsList ApiV1InstanceRdsList(ctx).Page(page).Size(size).Execute()

AliyunRDS清单



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
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	size := int32(56) // int32 | Number of results to return per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1InstanceRdsList(context.Background()).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1InstanceRdsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1InstanceRdsList`: PaginatedAliyunRdsList
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1InstanceRdsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1InstanceRdsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 
 **size** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedAliyunRdsList**](PaginatedAliyunRdsList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1InstanceResourceCreate

> InstanceResourceList ApiV1InstanceResourceCreate(ctx).InstanceResource(instanceResource).Execute()

实例资源



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
	instanceResource := *openapiclient.NewInstanceResource(int32(123), openapiclient.ResourceTypeEnum("database")) // InstanceResource | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1InstanceResourceCreate(context.Background()).InstanceResource(instanceResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1InstanceResourceCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1InstanceResourceCreate`: InstanceResourceList
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1InstanceResourceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1InstanceResourceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceResource** | [**InstanceResource**](InstanceResource.md) |  | 

### Return type

[**InstanceResourceList**](InstanceResourceList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1InstanceTableInstancesCreate

> TableInstanceLookupResponse ApiV1InstanceTableInstancesCreate(ctx).TableInstanceLookup(tableInstanceLookup).Execute()

按表名查询所属实例



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
	tableInstanceLookup := *openapiclient.NewTableInstanceLookup("TableName_example") // TableInstanceLookup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1InstanceTableInstancesCreate(context.Background()).TableInstanceLookup(tableInstanceLookup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1InstanceTableInstancesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1InstanceTableInstancesCreate`: TableInstanceLookupResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1InstanceTableInstancesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1InstanceTableInstancesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tableInstanceLookup** | [**TableInstanceLookup**](TableInstanceLookup.md) |  | 

### Return type

[**TableInstanceLookupResponse**](TableInstanceLookupResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1InstanceTunnelCreate

> Tunnel ApiV1InstanceTunnelCreate(ctx).Tunnel(tunnel).Execute()

创建隧道



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.intra.douban.com/sa/archery-cli"
)

func main() {
	tunnel := *openapiclient.NewTunnel(int32(123), "TunnelName_example", "Host_example", time.Now(), time.Now()) // Tunnel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1InstanceTunnelCreate(context.Background()).Tunnel(tunnel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1InstanceTunnelCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1InstanceTunnelCreate`: Tunnel
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1InstanceTunnelCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1InstanceTunnelCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tunnel** | [**Tunnel**](Tunnel.md) |  | 

### Return type

[**Tunnel**](Tunnel.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1InstanceTunnelList

> PaginatedTunnelList ApiV1InstanceTunnelList(ctx).Page(page).Size(size).Execute()

隧道清单



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
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	size := int32(56) // int32 | Number of results to return per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1InstanceTunnelList(context.Background()).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1InstanceTunnelList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1InstanceTunnelList`: PaginatedTunnelList
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1InstanceTunnelList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1InstanceTunnelListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 
 **size** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedTunnelList**](PaginatedTunnelList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1InstanceUpdate

> InstanceDetail ApiV1InstanceUpdate(ctx, id).InstanceDetail(instanceDetail).Execute()

更新实例



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.intra.douban.com/sa/archery-cli"
)

func main() {
	id := int32(56) // int32 | 
	instanceDetail := *openapiclient.NewInstanceDetail(int32(123), time.Now(), time.Now()) // InstanceDetail |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1InstanceUpdate(context.Background(), id).InstanceDetail(instanceDetail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1InstanceUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1InstanceUpdate`: InstanceDetail
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1InstanceUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1InstanceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceDetail** | [**InstanceDetail**](InstanceDetail.md) |  | 

### Return type

[**InstanceDetail**](InstanceDetail.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1SqlqueryDescribetableCreate

> ApiV1SqlqueryDescribetableCreate(ctx).Execute()

SQLQuery 表结构

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
	r, err := apiClient.ApiAPI.ApiV1SqlqueryDescribetableCreate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1SqlqueryDescribetableCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1SqlqueryDescribetableCreateRequest struct via the builder pattern


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


## ApiV1SqlqueryExecuteCreate

> ApiV1SqlqueryExecuteCreate(ctx).Execute()

SQLQuery 执行查询

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
	r, err := apiClient.ApiAPI.ApiV1SqlqueryExecuteCreate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1SqlqueryExecuteCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1SqlqueryExecuteCreateRequest struct via the builder pattern


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


## ApiV1SqlqueryFavoritesCreate

> ApiV1SqlqueryFavoritesCreate(ctx).Execute()

SQLQuery 收藏/取消收藏

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
	r, err := apiClient.ApiAPI.ApiV1SqlqueryFavoritesCreate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1SqlqueryFavoritesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1SqlqueryFavoritesCreateRequest struct via the builder pattern


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


## ApiV1SqlqueryInstancesRetrieve

> ApiV1SqlqueryInstancesRetrieve(ctx).Execute()

SQLQuery 可访问实例列表

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
	r, err := apiClient.ApiAPI.ApiV1SqlqueryInstancesRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1SqlqueryInstancesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1SqlqueryInstancesRetrieveRequest struct via the builder pattern


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


## ApiV1SqlqueryLogsRetrieve

> ApiV1SqlqueryLogsRetrieve(ctx).Execute()

SQLQuery 历史查询记录

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
	r, err := apiClient.ApiAPI.ApiV1SqlqueryLogsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1SqlqueryLogsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1SqlqueryLogsRetrieveRequest struct via the builder pattern


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


## ApiV1SqlqueryResourcesRetrieve

> ApiV1SqlqueryResourcesRetrieve(ctx).Execute()

SQLQuery 实例资源列表

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
	r, err := apiClient.ApiAPI.ApiV1SqlqueryResourcesRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1SqlqueryResourcesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1SqlqueryResourcesRetrieveRequest struct via the builder pattern


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


## ApiV1User2faCreate

> ApiV1User2faCreate(ctx).TwoFA(twoFA).Execute()

配置2fa



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
	twoFA := *openapiclient.NewTwoFA("Engineer_example", openapiclient.EnableEnum("true"), openapiclient.TwoFAAuthTypeEnum("totp")) // TwoFA | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApiAPI.ApiV1User2faCreate(context.Background()).TwoFA(twoFA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1User2faCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1User2faCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **twoFA** | [**TwoFA**](TwoFA.md) |  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1User2faSaveCreate

> ApiV1User2faSaveCreate(ctx).TwoFASave(twoFASave).Execute()

保存2fa配置



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
	twoFASave := *openapiclient.NewTwoFASave("Engineer_example", openapiclient.TwoFASaveAuthTypeEnum("disabled")) // TwoFASave | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApiAPI.ApiV1User2faSaveCreate(context.Background()).TwoFASave(twoFASave).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1User2faSaveCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1User2faSaveCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **twoFASave** | [**TwoFASave**](TwoFASave.md) |  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1User2faStateCreate

> ApiV1User2faStateCreate(ctx).TwoFAState(twoFAState).Execute()

查询2fa配置情况



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
	twoFAState := *openapiclient.NewTwoFAState("Engineer_example") // TwoFAState | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApiAPI.ApiV1User2faStateCreate(context.Background()).TwoFAState(twoFAState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1User2faStateCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1User2faStateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **twoFAState** | [**TwoFAState**](TwoFAState.md) |  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1User2faVerifyCreate

> ApiV1User2faVerifyCreate(ctx).TwoFAVerify(twoFAVerify).Execute()

检验2fa密码



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
	twoFAVerify := *openapiclient.NewTwoFAVerify("Engineer_example", int32(123), "AuthType_example") // TwoFAVerify | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApiAPI.ApiV1User2faVerifyCreate(context.Background()).TwoFAVerify(twoFAVerify).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1User2faVerifyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1User2faVerifyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **twoFAVerify** | [**TwoFAVerify**](TwoFAVerify.md) |  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UserAuthCreate

> ApiV1UserAuthCreate(ctx).UserAuth(userAuth).Execute()

用户认证校验



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
	userAuth := *openapiclient.NewUserAuth("Engineer_example", "Password_example") // UserAuth | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApiAPI.ApiV1UserAuthCreate(context.Background()).UserAuth(userAuth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1UserAuthCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UserAuthCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAuth** | [**UserAuth**](UserAuth.md) |  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UserCreate

> User ApiV1UserCreate(ctx).User(user).Execute()

创建用户



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
	user := *openapiclient.NewUser(int32(123), "Password_example", "Username_example", "Display_example") // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1UserCreate(context.Background()).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1UserCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UserCreate`: User
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1UserCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UserCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UserDestroy

> ApiV1UserDestroy(ctx, id).Execute()

删除用户



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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApiAPI.ApiV1UserDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1UserDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UserDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ApiV1UserGroupCreate

> Group ApiV1UserGroupCreate(ctx).Group(group).Execute()

创建用户组



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
	group := *openapiclient.NewGroup(int32(123), "Name_example") // Group | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1UserGroupCreate(context.Background()).Group(group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1UserGroupCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UserGroupCreate`: Group
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1UserGroupCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UserGroupCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | [**Group**](Group.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UserGroupDestroy

> ApiV1UserGroupDestroy(ctx, id).Execute()

删除用户组



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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApiAPI.ApiV1UserGroupDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1UserGroupDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UserGroupDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ApiV1UserGroupList

> PaginatedGroupList ApiV1UserGroupList(ctx).Page(page).Size(size).Execute()

用户组清单



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
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	size := int32(56) // int32 | Number of results to return per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1UserGroupList(context.Background()).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1UserGroupList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UserGroupList`: PaginatedGroupList
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1UserGroupList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UserGroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 
 **size** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedGroupList**](PaginatedGroupList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UserGroupUpdate

> Group ApiV1UserGroupUpdate(ctx, id).Group(group).Execute()

更新用户组



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
	id := int32(56) // int32 | 
	group := *openapiclient.NewGroup(int32(123), "Name_example") // Group | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1UserGroupUpdate(context.Background(), id).Group(group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1UserGroupUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UserGroupUpdate`: Group
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1UserGroupUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UserGroupUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **group** | [**Group**](Group.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UserList

> PaginatedUserList ApiV1UserList(ctx).Page(page).Size(size).Execute()

用户清单



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
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	size := int32(56) // int32 | Number of results to return per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1UserList(context.Background()).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1UserList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UserList`: PaginatedUserList
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1UserList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 
 **size** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedUserList**](PaginatedUserList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UserResourcegroupCreate

> ResourceGroup ApiV1UserResourcegroupCreate(ctx).ResourceGroup(resourceGroup).Execute()

创建资源组



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.intra.douban.com/sa/archery-cli"
)

func main() {
	resourceGroup := *openapiclient.NewResourceGroup(int32(123), "GroupName_example", time.Now(), time.Now()) // ResourceGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1UserResourcegroupCreate(context.Background()).ResourceGroup(resourceGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1UserResourcegroupCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UserResourcegroupCreate`: ResourceGroup
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1UserResourcegroupCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UserResourcegroupCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceGroup** | [**ResourceGroup**](ResourceGroup.md) |  | 

### Return type

[**ResourceGroup**](ResourceGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UserResourcegroupDestroy

> ApiV1UserResourcegroupDestroy(ctx, id).Execute()

删除资源组



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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApiAPI.ApiV1UserResourcegroupDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1UserResourcegroupDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UserResourcegroupDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ApiV1UserResourcegroupList

> PaginatedResourceGroupList ApiV1UserResourcegroupList(ctx).Page(page).Size(size).Execute()

资源组清单



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
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	size := int32(56) // int32 | Number of results to return per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1UserResourcegroupList(context.Background()).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1UserResourcegroupList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UserResourcegroupList`: PaginatedResourceGroupList
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1UserResourcegroupList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UserResourcegroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 
 **size** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedResourceGroupList**](PaginatedResourceGroupList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UserResourcegroupUpdate

> ResourceGroup ApiV1UserResourcegroupUpdate(ctx, id).ResourceGroup(resourceGroup).Execute()

更新资源组



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.intra.douban.com/sa/archery-cli"
)

func main() {
	id := int32(56) // int32 | 
	resourceGroup := *openapiclient.NewResourceGroup(int32(123), "GroupName_example", time.Now(), time.Now()) // ResourceGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1UserResourcegroupUpdate(context.Background(), id).ResourceGroup(resourceGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1UserResourcegroupUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UserResourcegroupUpdate`: ResourceGroup
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1UserResourcegroupUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UserResourcegroupUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceGroup** | [**ResourceGroup**](ResourceGroup.md) |  | 

### Return type

[**ResourceGroup**](ResourceGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UserUpdate

> UserDetail ApiV1UserUpdate(ctx, id).UserDetail(userDetail).Execute()

更新用户



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
	id := int32(56) // int32 | 
	userDetail := *openapiclient.NewUserDetail(int32(123)) // UserDetail |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1UserUpdate(context.Background(), id).UserDetail(userDetail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1UserUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UserUpdate`: UserDetail
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1UserUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UserUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userDetail** | [**UserDetail**](UserDetail.md) |  | 

### Return type

[**UserDetail**](UserDetail.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowAuditCreate

> ApiV1WorkflowAuditCreate(ctx).AuditWorkflow(auditWorkflow).Execute()

审核工单



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
	auditWorkflow := *openapiclient.NewAuditWorkflow("Engineer_example", int32(123), "AuditRemark_example", openapiclient.WorkflowTypeA27Enum(1), openapiclient.AuditTypeEnum("pass")) // AuditWorkflow | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApiAPI.ApiV1WorkflowAuditCreate(context.Background()).AuditWorkflow(auditWorkflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1WorkflowAuditCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowAuditCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auditWorkflow** | [**AuditWorkflow**](AuditWorkflow.md) |  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowAuditlistCreate

> WorkflowAuditList ApiV1WorkflowAuditlistCreate(ctx).WorkflowAudit(workflowAudit).Execute()

待审核清单



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
	workflowAudit := *openapiclient.NewWorkflowAudit("Engineer_example") // WorkflowAudit | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1WorkflowAuditlistCreate(context.Background()).WorkflowAudit(workflowAudit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1WorkflowAuditlistCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1WorkflowAuditlistCreate`: WorkflowAuditList
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1WorkflowAuditlistCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowAuditlistCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowAudit** | [**WorkflowAudit**](WorkflowAudit.md) |  | 

### Return type

[**WorkflowAuditList**](WorkflowAuditList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowCreate

> WorkflowContent ApiV1WorkflowCreate(ctx).WorkflowContent(workflowContent).Execute()

提交SQL上线工单



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.intra.douban.com/sa/archery-cli"
)

func main() {
	workflowContent := *openapiclient.NewWorkflowContent(int32(123), int32(123), *openapiclient.NewWorkflow(int32(123), "WorkflowName_example", int32(123), "GroupName_example", "DbName_example", openapiclient.SyntaxTypeEnum(0), "EngineerDisplay_example", openapiclient.StatusEnum("workflow_finish"), "AuditAuthGroups_example", time.Now(), time.Now(), openapiclient.IsOfflineExportEnum(0), int32(123)), "SqlContent_example", "ReviewContent_example", "ExecuteResult_example") // WorkflowContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1WorkflowCreate(context.Background()).WorkflowContent(workflowContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1WorkflowCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1WorkflowCreate`: WorkflowContent
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1WorkflowCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowContent** | [**WorkflowContent**](WorkflowContent.md) |  | 

### Return type

[**WorkflowContent**](WorkflowContent.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowExecuteCreate

> ApiV1WorkflowExecuteCreate(ctx).ExecuteWorkflow(executeWorkflow).Execute()

执行工单



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
	executeWorkflow := *openapiclient.NewExecuteWorkflow(int32(123), openapiclient.ExecuteWorkflowWorkflowTypeEnum(2)) // ExecuteWorkflow | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApiAPI.ApiV1WorkflowExecuteCreate(context.Background()).ExecuteWorkflow(executeWorkflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1WorkflowExecuteCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowExecuteCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executeWorkflow** | [**ExecuteWorkflow**](ExecuteWorkflow.md) |  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowList

> PaginatedWorkflowContentList ApiV1WorkflowList(ctx).Page(page).Size(size).Execute()

SQL上线工单清单



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
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	size := int32(56) // int32 | Number of results to return per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1WorkflowList(context.Background()).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1WorkflowList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1WorkflowList`: PaginatedWorkflowContentList
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1WorkflowList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 
 **size** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedWorkflowContentList**](PaginatedWorkflowContentList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowLogCreate

> WorkflowLogList ApiV1WorkflowLogCreate(ctx).WorkflowLog(workflowLog).Execute()

工单日志



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
	workflowLog := *openapiclient.NewWorkflowLog(int32(123), openapiclient.WorkflowLogWorkflowTypeEnum(1)) // WorkflowLog | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1WorkflowLogCreate(context.Background()).WorkflowLog(workflowLog).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1WorkflowLogCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1WorkflowLogCreate`: WorkflowLogList
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1WorkflowLogCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowLogCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowLog** | [**WorkflowLog**](WorkflowLog.md) |  | 

### Return type

[**WorkflowLogList**](WorkflowLogList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowSqlcheckCreate

> ExecuteCheckResult ApiV1WorkflowSqlcheckCreate(ctx).ExecuteCheck(executeCheck).Execute()

SQL检查



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
	executeCheck := *openapiclient.NewExecuteCheck(int32(123), "DbName_example", "FullSql_example") // ExecuteCheck | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiAPI.ApiV1WorkflowSqlcheckCreate(context.Background()).ExecuteCheck(executeCheck).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiAPI.ApiV1WorkflowSqlcheckCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1WorkflowSqlcheckCreate`: ExecuteCheckResult
	fmt.Fprintf(os.Stdout, "Response from `ApiAPI.ApiV1WorkflowSqlcheckCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowSqlcheckCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executeCheck** | [**ExecuteCheck**](ExecuteCheck.md) |  | 

### Return type

[**ExecuteCheckResult**](ExecuteCheckResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [jwtAuth](../README.md#jwtAuth), [knoxApiToken](../README.md#knoxApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

