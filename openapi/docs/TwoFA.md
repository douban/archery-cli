# TwoFA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engineer** | **string** |  | 
**Enable** | [**EnableEnum**](EnableEnum.md) |  | 
**Phone** | Pointer to **string** |  | [optional] 
**AuthType** | [**TwoFAAuthTypeEnum**](TwoFAAuthTypeEnum.md) |  | 

## Methods

### NewTwoFA

`func NewTwoFA(engineer string, enable EnableEnum, authType TwoFAAuthTypeEnum, ) *TwoFA`

NewTwoFA instantiates a new TwoFA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwoFAWithDefaults

`func NewTwoFAWithDefaults() *TwoFA`

NewTwoFAWithDefaults instantiates a new TwoFA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineer

`func (o *TwoFA) GetEngineer() string`

GetEngineer returns the Engineer field if non-nil, zero value otherwise.

### GetEngineerOk

`func (o *TwoFA) GetEngineerOk() (*string, bool)`

GetEngineerOk returns a tuple with the Engineer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineer

`func (o *TwoFA) SetEngineer(v string)`

SetEngineer sets Engineer field to given value.


### GetEnable

`func (o *TwoFA) GetEnable() EnableEnum`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *TwoFA) GetEnableOk() (*EnableEnum, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *TwoFA) SetEnable(v EnableEnum)`

SetEnable sets Enable field to given value.


### GetPhone

`func (o *TwoFA) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *TwoFA) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *TwoFA) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *TwoFA) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAuthType

`func (o *TwoFA) GetAuthType() TwoFAAuthTypeEnum`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *TwoFA) GetAuthTypeOk() (*TwoFAAuthTypeEnum, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *TwoFA) SetAuthType(v TwoFAAuthTypeEnum)`

SetAuthType sets AuthType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


