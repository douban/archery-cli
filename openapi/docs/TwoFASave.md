# TwoFASave

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engineer** | **string** |  | 
**Key** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**AuthType** | [**TwoFASaveAuthTypeEnum**](TwoFASaveAuthTypeEnum.md) |  | 

## Methods

### NewTwoFASave

`func NewTwoFASave(engineer string, authType TwoFASaveAuthTypeEnum, ) *TwoFASave`

NewTwoFASave instantiates a new TwoFASave object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwoFASaveWithDefaults

`func NewTwoFASaveWithDefaults() *TwoFASave`

NewTwoFASaveWithDefaults instantiates a new TwoFASave object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineer

`func (o *TwoFASave) GetEngineer() string`

GetEngineer returns the Engineer field if non-nil, zero value otherwise.

### GetEngineerOk

`func (o *TwoFASave) GetEngineerOk() (*string, bool)`

GetEngineerOk returns a tuple with the Engineer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineer

`func (o *TwoFASave) SetEngineer(v string)`

SetEngineer sets Engineer field to given value.


### GetKey

`func (o *TwoFASave) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TwoFASave) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TwoFASave) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TwoFASave) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPhone

`func (o *TwoFASave) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *TwoFASave) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *TwoFASave) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *TwoFASave) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAuthType

`func (o *TwoFASave) GetAuthType() TwoFASaveAuthTypeEnum`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *TwoFASave) GetAuthTypeOk() (*TwoFASaveAuthTypeEnum, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *TwoFASave) SetAuthType(v TwoFASaveAuthTypeEnum)`

SetAuthType sets AuthType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


