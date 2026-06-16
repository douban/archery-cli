# InstanceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**InstanceName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**TypeA85Enum**](TypeA85Enum.md) |  | [optional] 
**DbType** | Pointer to [**DbTypeEnum**](DbTypeEnum.md) |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
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

### NewInstanceDetail

`func NewInstanceDetail(id int32, createTime time.Time, updateTime time.Time, ) *InstanceDetail`

NewInstanceDetail instantiates a new InstanceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceDetailWithDefaults

`func NewInstanceDetailWithDefaults() *InstanceDetail`

NewInstanceDetailWithDefaults instantiates a new InstanceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceDetail) SetId(v int32)`

SetId sets Id field to given value.


### GetInstanceName

`func (o *InstanceDetail) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *InstanceDetail) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *InstanceDetail) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *InstanceDetail) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetType

`func (o *InstanceDetail) GetType() TypeA85Enum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceDetail) GetTypeOk() (*TypeA85Enum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceDetail) SetType(v TypeA85Enum)`

SetType sets Type field to given value.

### HasType

`func (o *InstanceDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDbType

`func (o *InstanceDetail) GetDbType() DbTypeEnum`

GetDbType returns the DbType field if non-nil, zero value otherwise.

### GetDbTypeOk

`func (o *InstanceDetail) GetDbTypeOk() (*DbTypeEnum, bool)`

GetDbTypeOk returns a tuple with the DbType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbType

`func (o *InstanceDetail) SetDbType(v DbTypeEnum)`

SetDbType sets DbType field to given value.

### HasDbType

`func (o *InstanceDetail) HasDbType() bool`

HasDbType returns a boolean if a field has been set.

### GetMode

`func (o *InstanceDetail) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InstanceDetail) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InstanceDetail) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InstanceDetail) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetHost

`func (o *InstanceDetail) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *InstanceDetail) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *InstanceDetail) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *InstanceDetail) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *InstanceDetail) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InstanceDetail) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InstanceDetail) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *InstanceDetail) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUser

`func (o *InstanceDetail) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *InstanceDetail) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *InstanceDetail) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *InstanceDetail) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPassword

`func (o *InstanceDetail) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InstanceDetail) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InstanceDetail) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InstanceDetail) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetIsSsl

`func (o *InstanceDetail) GetIsSsl() bool`

GetIsSsl returns the IsSsl field if non-nil, zero value otherwise.

### GetIsSslOk

`func (o *InstanceDetail) GetIsSslOk() (*bool, bool)`

GetIsSslOk returns a tuple with the IsSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSsl

`func (o *InstanceDetail) SetIsSsl(v bool)`

SetIsSsl sets IsSsl field to given value.

### HasIsSsl

`func (o *InstanceDetail) HasIsSsl() bool`

HasIsSsl returns a boolean if a field has been set.

### GetVerifySsl

`func (o *InstanceDetail) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *InstanceDetail) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *InstanceDetail) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *InstanceDetail) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.

### GetDbName

`func (o *InstanceDetail) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *InstanceDetail) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *InstanceDetail) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *InstanceDetail) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetShowDbNameRegex

`func (o *InstanceDetail) GetShowDbNameRegex() string`

GetShowDbNameRegex returns the ShowDbNameRegex field if non-nil, zero value otherwise.

### GetShowDbNameRegexOk

`func (o *InstanceDetail) GetShowDbNameRegexOk() (*string, bool)`

GetShowDbNameRegexOk returns a tuple with the ShowDbNameRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDbNameRegex

`func (o *InstanceDetail) SetShowDbNameRegex(v string)`

SetShowDbNameRegex sets ShowDbNameRegex field to given value.

### HasShowDbNameRegex

`func (o *InstanceDetail) HasShowDbNameRegex() bool`

HasShowDbNameRegex returns a boolean if a field has been set.

### GetDeniedDbNameRegex

`func (o *InstanceDetail) GetDeniedDbNameRegex() string`

GetDeniedDbNameRegex returns the DeniedDbNameRegex field if non-nil, zero value otherwise.

### GetDeniedDbNameRegexOk

`func (o *InstanceDetail) GetDeniedDbNameRegexOk() (*string, bool)`

GetDeniedDbNameRegexOk returns a tuple with the DeniedDbNameRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedDbNameRegex

`func (o *InstanceDetail) SetDeniedDbNameRegex(v string)`

SetDeniedDbNameRegex sets DeniedDbNameRegex field to given value.

### HasDeniedDbNameRegex

`func (o *InstanceDetail) HasDeniedDbNameRegex() bool`

HasDeniedDbNameRegex returns a boolean if a field has been set.

### GetCharset

`func (o *InstanceDetail) GetCharset() string`

GetCharset returns the Charset field if non-nil, zero value otherwise.

### GetCharsetOk

`func (o *InstanceDetail) GetCharsetOk() (*string, bool)`

GetCharsetOk returns a tuple with the Charset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharset

`func (o *InstanceDetail) SetCharset(v string)`

SetCharset sets Charset field to given value.

### HasCharset

`func (o *InstanceDetail) HasCharset() bool`

HasCharset returns a boolean if a field has been set.

### GetServiceName

`func (o *InstanceDetail) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *InstanceDetail) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *InstanceDetail) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *InstanceDetail) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *InstanceDetail) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *InstanceDetail) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil
### GetSid

`func (o *InstanceDetail) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *InstanceDetail) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *InstanceDetail) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *InstanceDetail) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *InstanceDetail) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *InstanceDetail) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetCreateTime

`func (o *InstanceDetail) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *InstanceDetail) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *InstanceDetail) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.


### GetUpdateTime

`func (o *InstanceDetail) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *InstanceDetail) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *InstanceDetail) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.


### GetTunnel

`func (o *InstanceDetail) GetTunnel() int32`

GetTunnel returns the Tunnel field if non-nil, zero value otherwise.

### GetTunnelOk

`func (o *InstanceDetail) GetTunnelOk() (*int32, bool)`

GetTunnelOk returns a tuple with the Tunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel

`func (o *InstanceDetail) SetTunnel(v int32)`

SetTunnel sets Tunnel field to given value.

### HasTunnel

`func (o *InstanceDetail) HasTunnel() bool`

HasTunnel returns a boolean if a field has been set.

### SetTunnelNil

`func (o *InstanceDetail) SetTunnelNil(b bool)`

 SetTunnelNil sets the value for Tunnel to be an explicit nil

### UnsetTunnel
`func (o *InstanceDetail) UnsetTunnel()`

UnsetTunnel ensures that no value is present for Tunnel, not even an explicit nil
### GetResourceGroup

`func (o *InstanceDetail) GetResourceGroup() []int32`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *InstanceDetail) GetResourceGroupOk() (*[]int32, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *InstanceDetail) SetResourceGroup(v []int32)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *InstanceDetail) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetInstanceTag

`func (o *InstanceDetail) GetInstanceTag() []int32`

GetInstanceTag returns the InstanceTag field if non-nil, zero value otherwise.

### GetInstanceTagOk

`func (o *InstanceDetail) GetInstanceTagOk() (*[]int32, bool)`

GetInstanceTagOk returns a tuple with the InstanceTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTag

`func (o *InstanceDetail) SetInstanceTag(v []int32)`

SetInstanceTag sets InstanceTag field to given value.

### HasInstanceTag

`func (o *InstanceDetail) HasInstanceTag() bool`

HasInstanceTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


