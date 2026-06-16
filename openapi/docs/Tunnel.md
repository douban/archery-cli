# Tunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**TunnelName** | **string** |  | 
**Host** | **string** |  | 
**Port** | Pointer to **int32** |  | [optional] 
**User** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**Pkey** | Pointer to **NullableString** |  | [optional] 
**PkeyPath** | Pointer to **NullableString** |  | [optional] 
**PkeyPassword** | Pointer to **NullableString** |  | [optional] 
**CreateTime** | **time.Time** |  | [readonly] 
**UpdateTime** | **time.Time** |  | [readonly] 

## Methods

### NewTunnel

`func NewTunnel(id int32, tunnelName string, host string, createTime time.Time, updateTime time.Time, ) *Tunnel`

NewTunnel instantiates a new Tunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelWithDefaults

`func NewTunnelWithDefaults() *Tunnel`

NewTunnelWithDefaults instantiates a new Tunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Tunnel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tunnel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tunnel) SetId(v int32)`

SetId sets Id field to given value.


### GetTunnelName

`func (o *Tunnel) GetTunnelName() string`

GetTunnelName returns the TunnelName field if non-nil, zero value otherwise.

### GetTunnelNameOk

`func (o *Tunnel) GetTunnelNameOk() (*string, bool)`

GetTunnelNameOk returns a tuple with the TunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelName

`func (o *Tunnel) SetTunnelName(v string)`

SetTunnelName sets TunnelName field to given value.


### GetHost

`func (o *Tunnel) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Tunnel) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Tunnel) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *Tunnel) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Tunnel) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Tunnel) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Tunnel) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUser

`func (o *Tunnel) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Tunnel) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Tunnel) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Tunnel) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *Tunnel) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *Tunnel) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPassword

`func (o *Tunnel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Tunnel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Tunnel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Tunnel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *Tunnel) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *Tunnel) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPkey

`func (o *Tunnel) GetPkey() string`

GetPkey returns the Pkey field if non-nil, zero value otherwise.

### GetPkeyOk

`func (o *Tunnel) GetPkeyOk() (*string, bool)`

GetPkeyOk returns a tuple with the Pkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkey

`func (o *Tunnel) SetPkey(v string)`

SetPkey sets Pkey field to given value.

### HasPkey

`func (o *Tunnel) HasPkey() bool`

HasPkey returns a boolean if a field has been set.

### SetPkeyNil

`func (o *Tunnel) SetPkeyNil(b bool)`

 SetPkeyNil sets the value for Pkey to be an explicit nil

### UnsetPkey
`func (o *Tunnel) UnsetPkey()`

UnsetPkey ensures that no value is present for Pkey, not even an explicit nil
### GetPkeyPath

`func (o *Tunnel) GetPkeyPath() string`

GetPkeyPath returns the PkeyPath field if non-nil, zero value otherwise.

### GetPkeyPathOk

`func (o *Tunnel) GetPkeyPathOk() (*string, bool)`

GetPkeyPathOk returns a tuple with the PkeyPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkeyPath

`func (o *Tunnel) SetPkeyPath(v string)`

SetPkeyPath sets PkeyPath field to given value.

### HasPkeyPath

`func (o *Tunnel) HasPkeyPath() bool`

HasPkeyPath returns a boolean if a field has been set.

### SetPkeyPathNil

`func (o *Tunnel) SetPkeyPathNil(b bool)`

 SetPkeyPathNil sets the value for PkeyPath to be an explicit nil

### UnsetPkeyPath
`func (o *Tunnel) UnsetPkeyPath()`

UnsetPkeyPath ensures that no value is present for PkeyPath, not even an explicit nil
### GetPkeyPassword

`func (o *Tunnel) GetPkeyPassword() string`

GetPkeyPassword returns the PkeyPassword field if non-nil, zero value otherwise.

### GetPkeyPasswordOk

`func (o *Tunnel) GetPkeyPasswordOk() (*string, bool)`

GetPkeyPasswordOk returns a tuple with the PkeyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkeyPassword

`func (o *Tunnel) SetPkeyPassword(v string)`

SetPkeyPassword sets PkeyPassword field to given value.

### HasPkeyPassword

`func (o *Tunnel) HasPkeyPassword() bool`

HasPkeyPassword returns a boolean if a field has been set.

### SetPkeyPasswordNil

`func (o *Tunnel) SetPkeyPasswordNil(b bool)`

 SetPkeyPasswordNil sets the value for PkeyPassword to be an explicit nil

### UnsetPkeyPassword
`func (o *Tunnel) UnsetPkeyPassword()`

UnsetPkeyPassword ensures that no value is present for PkeyPassword, not even an explicit nil
### GetCreateTime

`func (o *Tunnel) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Tunnel) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Tunnel) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.


### GetUpdateTime

`func (o *Tunnel) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *Tunnel) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *Tunnel) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


