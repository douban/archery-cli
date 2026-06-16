# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**InstanceName** | **string** |  | 
**Type** | [**TypeA85Enum**](TypeA85Enum.md) |  | 
**DbType** | [**DbTypeEnum**](DbTypeEnum.md) |  | 
**Mode** | Pointer to **string** |  | [optional] 
**Host** | **string** |  | 
**Port** | Pointer to **int32** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**IsSsl** | Pointer to **bool** |  | [optional] 
**VerifySsl** | Pointer to **bool** |  | [optional] 
**DbName** | Pointer to **string** |  | [optional] 
**ShowDbNameRegex** | Pointer to **string** | 正则表达式。示例：^(test_db|dmp_db|za.*)$。Redis示例: ^(0|4|6|11|12|13)$ | [optional] 
**DeniedDbNameRegex** | Pointer to **string** | 正则表达式。隐藏大于显示，此规则优先。 | [optional] 
**Charset** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **NullableString** |  | [optional] 
**Sid** | Pointer to **NullableString** |  | [optional] 
**CreateTime** | **time.Time** |  | [readonly] 
**UpdateTime** | **time.Time** |  | [readonly] 
**Tunnel** | Pointer to **NullableInt32** |  | [optional] 
**ResourceGroup** | Pointer to **[]int32** |  | [optional] 
**InstanceTag** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewInstance

`func NewInstance(id int32, instanceName string, type_ TypeA85Enum, dbType DbTypeEnum, host string, createTime time.Time, updateTime time.Time, ) *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Instance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Instance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Instance) SetId(v int32)`

SetId sets Id field to given value.


### GetInstanceName

`func (o *Instance) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *Instance) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *Instance) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetType

`func (o *Instance) GetType() TypeA85Enum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Instance) GetTypeOk() (*TypeA85Enum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Instance) SetType(v TypeA85Enum)`

SetType sets Type field to given value.


### GetDbType

`func (o *Instance) GetDbType() DbTypeEnum`

GetDbType returns the DbType field if non-nil, zero value otherwise.

### GetDbTypeOk

`func (o *Instance) GetDbTypeOk() (*DbTypeEnum, bool)`

GetDbTypeOk returns a tuple with the DbType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbType

`func (o *Instance) SetDbType(v DbTypeEnum)`

SetDbType sets DbType field to given value.


### GetMode

`func (o *Instance) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Instance) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Instance) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Instance) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetHost

`func (o *Instance) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Instance) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Instance) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *Instance) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Instance) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Instance) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Instance) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUser

`func (o *Instance) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Instance) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Instance) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Instance) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPassword

`func (o *Instance) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Instance) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Instance) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Instance) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetIsSsl

`func (o *Instance) GetIsSsl() bool`

GetIsSsl returns the IsSsl field if non-nil, zero value otherwise.

### GetIsSslOk

`func (o *Instance) GetIsSslOk() (*bool, bool)`

GetIsSslOk returns a tuple with the IsSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSsl

`func (o *Instance) SetIsSsl(v bool)`

SetIsSsl sets IsSsl field to given value.

### HasIsSsl

`func (o *Instance) HasIsSsl() bool`

HasIsSsl returns a boolean if a field has been set.

### GetVerifySsl

`func (o *Instance) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *Instance) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *Instance) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *Instance) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.

### GetDbName

`func (o *Instance) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *Instance) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *Instance) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *Instance) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetShowDbNameRegex

`func (o *Instance) GetShowDbNameRegex() string`

GetShowDbNameRegex returns the ShowDbNameRegex field if non-nil, zero value otherwise.

### GetShowDbNameRegexOk

`func (o *Instance) GetShowDbNameRegexOk() (*string, bool)`

GetShowDbNameRegexOk returns a tuple with the ShowDbNameRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDbNameRegex

`func (o *Instance) SetShowDbNameRegex(v string)`

SetShowDbNameRegex sets ShowDbNameRegex field to given value.

### HasShowDbNameRegex

`func (o *Instance) HasShowDbNameRegex() bool`

HasShowDbNameRegex returns a boolean if a field has been set.

### GetDeniedDbNameRegex

`func (o *Instance) GetDeniedDbNameRegex() string`

GetDeniedDbNameRegex returns the DeniedDbNameRegex field if non-nil, zero value otherwise.

### GetDeniedDbNameRegexOk

`func (o *Instance) GetDeniedDbNameRegexOk() (*string, bool)`

GetDeniedDbNameRegexOk returns a tuple with the DeniedDbNameRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedDbNameRegex

`func (o *Instance) SetDeniedDbNameRegex(v string)`

SetDeniedDbNameRegex sets DeniedDbNameRegex field to given value.

### HasDeniedDbNameRegex

`func (o *Instance) HasDeniedDbNameRegex() bool`

HasDeniedDbNameRegex returns a boolean if a field has been set.

### GetCharset

`func (o *Instance) GetCharset() string`

GetCharset returns the Charset field if non-nil, zero value otherwise.

### GetCharsetOk

`func (o *Instance) GetCharsetOk() (*string, bool)`

GetCharsetOk returns a tuple with the Charset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharset

`func (o *Instance) SetCharset(v string)`

SetCharset sets Charset field to given value.

### HasCharset

`func (o *Instance) HasCharset() bool`

HasCharset returns a boolean if a field has been set.

### GetServiceName

`func (o *Instance) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *Instance) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *Instance) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *Instance) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *Instance) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *Instance) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil
### GetSid

`func (o *Instance) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *Instance) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *Instance) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *Instance) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *Instance) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *Instance) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetCreateTime

`func (o *Instance) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Instance) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Instance) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.


### GetUpdateTime

`func (o *Instance) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *Instance) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *Instance) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.


### GetTunnel

`func (o *Instance) GetTunnel() int32`

GetTunnel returns the Tunnel field if non-nil, zero value otherwise.

### GetTunnelOk

`func (o *Instance) GetTunnelOk() (*int32, bool)`

GetTunnelOk returns a tuple with the Tunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel

`func (o *Instance) SetTunnel(v int32)`

SetTunnel sets Tunnel field to given value.

### HasTunnel

`func (o *Instance) HasTunnel() bool`

HasTunnel returns a boolean if a field has been set.

### SetTunnelNil

`func (o *Instance) SetTunnelNil(b bool)`

 SetTunnelNil sets the value for Tunnel to be an explicit nil

### UnsetTunnel
`func (o *Instance) UnsetTunnel()`

UnsetTunnel ensures that no value is present for Tunnel, not even an explicit nil
### GetResourceGroup

`func (o *Instance) GetResourceGroup() []int32`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *Instance) GetResourceGroupOk() (*[]int32, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *Instance) SetResourceGroup(v []int32)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *Instance) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetInstanceTag

`func (o *Instance) GetInstanceTag() []int32`

GetInstanceTag returns the InstanceTag field if non-nil, zero value otherwise.

### GetInstanceTagOk

`func (o *Instance) GetInstanceTagOk() (*[]int32, bool)`

GetInstanceTagOk returns a tuple with the InstanceTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTag

`func (o *Instance) SetInstanceTag(v []int32)`

SetInstanceTag sets InstanceTag field to given value.

### HasInstanceTag

`func (o *Instance) HasInstanceTag() bool`

HasInstanceTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


