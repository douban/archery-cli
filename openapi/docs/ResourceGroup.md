# ResourceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **int32** |  | [readonly] 
**GroupName** | **string** |  | 
**GroupParentId** | Pointer to **int64** |  | [optional] 
**GroupSort** | Pointer to **int32** |  | [optional] 
**GroupLevel** | Pointer to **int32** |  | [optional] 
**DingWebhook** | Pointer to **string** |  | [optional] 
**FeishuWebhook** | Pointer to **string** |  | [optional] 
**QywxWebhook** | Pointer to **string** |  | [optional] 
**IsDeleted** | Pointer to [**IsOfflineExportEnum**](IsOfflineExportEnum.md) |  | [optional] 
**CreateTime** | **time.Time** |  | [readonly] 
**SysTime** | **time.Time** |  | [readonly] 

## Methods

### NewResourceGroup

`func NewResourceGroup(groupId int32, groupName string, createTime time.Time, sysTime time.Time, ) *ResourceGroup`

NewResourceGroup instantiates a new ResourceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceGroupWithDefaults

`func NewResourceGroupWithDefaults() *ResourceGroup`

NewResourceGroupWithDefaults instantiates a new ResourceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ResourceGroup) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ResourceGroup) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ResourceGroup) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetGroupName

`func (o *ResourceGroup) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ResourceGroup) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ResourceGroup) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetGroupParentId

`func (o *ResourceGroup) GetGroupParentId() int64`

GetGroupParentId returns the GroupParentId field if non-nil, zero value otherwise.

### GetGroupParentIdOk

`func (o *ResourceGroup) GetGroupParentIdOk() (*int64, bool)`

GetGroupParentIdOk returns a tuple with the GroupParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupParentId

`func (o *ResourceGroup) SetGroupParentId(v int64)`

SetGroupParentId sets GroupParentId field to given value.

### HasGroupParentId

`func (o *ResourceGroup) HasGroupParentId() bool`

HasGroupParentId returns a boolean if a field has been set.

### GetGroupSort

`func (o *ResourceGroup) GetGroupSort() int32`

GetGroupSort returns the GroupSort field if non-nil, zero value otherwise.

### GetGroupSortOk

`func (o *ResourceGroup) GetGroupSortOk() (*int32, bool)`

GetGroupSortOk returns a tuple with the GroupSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSort

`func (o *ResourceGroup) SetGroupSort(v int32)`

SetGroupSort sets GroupSort field to given value.

### HasGroupSort

`func (o *ResourceGroup) HasGroupSort() bool`

HasGroupSort returns a boolean if a field has been set.

### GetGroupLevel

`func (o *ResourceGroup) GetGroupLevel() int32`

GetGroupLevel returns the GroupLevel field if non-nil, zero value otherwise.

### GetGroupLevelOk

`func (o *ResourceGroup) GetGroupLevelOk() (*int32, bool)`

GetGroupLevelOk returns a tuple with the GroupLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLevel

`func (o *ResourceGroup) SetGroupLevel(v int32)`

SetGroupLevel sets GroupLevel field to given value.

### HasGroupLevel

`func (o *ResourceGroup) HasGroupLevel() bool`

HasGroupLevel returns a boolean if a field has been set.

### GetDingWebhook

`func (o *ResourceGroup) GetDingWebhook() string`

GetDingWebhook returns the DingWebhook field if non-nil, zero value otherwise.

### GetDingWebhookOk

`func (o *ResourceGroup) GetDingWebhookOk() (*string, bool)`

GetDingWebhookOk returns a tuple with the DingWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDingWebhook

`func (o *ResourceGroup) SetDingWebhook(v string)`

SetDingWebhook sets DingWebhook field to given value.

### HasDingWebhook

`func (o *ResourceGroup) HasDingWebhook() bool`

HasDingWebhook returns a boolean if a field has been set.

### GetFeishuWebhook

`func (o *ResourceGroup) GetFeishuWebhook() string`

GetFeishuWebhook returns the FeishuWebhook field if non-nil, zero value otherwise.

### GetFeishuWebhookOk

`func (o *ResourceGroup) GetFeishuWebhookOk() (*string, bool)`

GetFeishuWebhookOk returns a tuple with the FeishuWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeishuWebhook

`func (o *ResourceGroup) SetFeishuWebhook(v string)`

SetFeishuWebhook sets FeishuWebhook field to given value.

### HasFeishuWebhook

`func (o *ResourceGroup) HasFeishuWebhook() bool`

HasFeishuWebhook returns a boolean if a field has been set.

### GetQywxWebhook

`func (o *ResourceGroup) GetQywxWebhook() string`

GetQywxWebhook returns the QywxWebhook field if non-nil, zero value otherwise.

### GetQywxWebhookOk

`func (o *ResourceGroup) GetQywxWebhookOk() (*string, bool)`

GetQywxWebhookOk returns a tuple with the QywxWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQywxWebhook

`func (o *ResourceGroup) SetQywxWebhook(v string)`

SetQywxWebhook sets QywxWebhook field to given value.

### HasQywxWebhook

`func (o *ResourceGroup) HasQywxWebhook() bool`

HasQywxWebhook returns a boolean if a field has been set.

### GetIsDeleted

`func (o *ResourceGroup) GetIsDeleted() IsOfflineExportEnum`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ResourceGroup) GetIsDeletedOk() (*IsOfflineExportEnum, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ResourceGroup) SetIsDeleted(v IsOfflineExportEnum)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ResourceGroup) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetCreateTime

`func (o *ResourceGroup) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ResourceGroup) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ResourceGroup) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.


### GetSysTime

`func (o *ResourceGroup) GetSysTime() time.Time`

GetSysTime returns the SysTime field if non-nil, zero value otherwise.

### GetSysTimeOk

`func (o *ResourceGroup) GetSysTimeOk() (*time.Time, bool)`

GetSysTimeOk returns a tuple with the SysTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysTime

`func (o *ResourceGroup) SetSysTime(v time.Time)`

SetSysTime sets SysTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


