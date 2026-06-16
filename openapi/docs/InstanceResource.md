# InstanceResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **int32** |  | 
**ResourceType** | [**ResourceTypeEnum**](ResourceTypeEnum.md) |  | 
**DbName** | Pointer to **string** |  | [optional] 
**SchemaName** | Pointer to **string** |  | [optional] 
**TbName** | Pointer to **string** |  | [optional] 

## Methods

### NewInstanceResource

`func NewInstanceResource(instanceId int32, resourceType ResourceTypeEnum, ) *InstanceResource`

NewInstanceResource instantiates a new InstanceResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceResourceWithDefaults

`func NewInstanceResourceWithDefaults() *InstanceResource`

NewInstanceResourceWithDefaults instantiates a new InstanceResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *InstanceResource) GetInstanceId() int32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *InstanceResource) GetInstanceIdOk() (*int32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *InstanceResource) SetInstanceId(v int32)`

SetInstanceId sets InstanceId field to given value.


### GetResourceType

`func (o *InstanceResource) GetResourceType() ResourceTypeEnum`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *InstanceResource) GetResourceTypeOk() (*ResourceTypeEnum, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *InstanceResource) SetResourceType(v ResourceTypeEnum)`

SetResourceType sets ResourceType field to given value.


### GetDbName

`func (o *InstanceResource) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *InstanceResource) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *InstanceResource) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *InstanceResource) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetSchemaName

`func (o *InstanceResource) GetSchemaName() string`

GetSchemaName returns the SchemaName field if non-nil, zero value otherwise.

### GetSchemaNameOk

`func (o *InstanceResource) GetSchemaNameOk() (*string, bool)`

GetSchemaNameOk returns a tuple with the SchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaName

`func (o *InstanceResource) SetSchemaName(v string)`

SetSchemaName sets SchemaName field to given value.

### HasSchemaName

`func (o *InstanceResource) HasSchemaName() bool`

HasSchemaName returns a boolean if a field has been set.

### GetTbName

`func (o *InstanceResource) GetTbName() string`

GetTbName returns the TbName field if non-nil, zero value otherwise.

### GetTbNameOk

`func (o *InstanceResource) GetTbNameOk() (*string, bool)`

GetTbNameOk returns a tuple with the TbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTbName

`func (o *InstanceResource) SetTbName(v string)`

SetTbName sets TbName field to given value.

### HasTbName

`func (o *InstanceResource) HasTbName() bool`

HasTbName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


