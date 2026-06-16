# TwoFAVerify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engineer** | **string** |  | 
**Otp** | **int32** |  | 
**Key** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**AuthType** | **string** |  | 

## Methods

### NewTwoFAVerify

`func NewTwoFAVerify(engineer string, otp int32, authType string, ) *TwoFAVerify`

NewTwoFAVerify instantiates a new TwoFAVerify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwoFAVerifyWithDefaults

`func NewTwoFAVerifyWithDefaults() *TwoFAVerify`

NewTwoFAVerifyWithDefaults instantiates a new TwoFAVerify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineer

`func (o *TwoFAVerify) GetEngineer() string`

GetEngineer returns the Engineer field if non-nil, zero value otherwise.

### GetEngineerOk

`func (o *TwoFAVerify) GetEngineerOk() (*string, bool)`

GetEngineerOk returns a tuple with the Engineer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineer

`func (o *TwoFAVerify) SetEngineer(v string)`

SetEngineer sets Engineer field to given value.


### GetOtp

`func (o *TwoFAVerify) GetOtp() int32`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *TwoFAVerify) GetOtpOk() (*int32, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *TwoFAVerify) SetOtp(v int32)`

SetOtp sets Otp field to given value.


### GetKey

`func (o *TwoFAVerify) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TwoFAVerify) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TwoFAVerify) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TwoFAVerify) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPhone

`func (o *TwoFAVerify) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *TwoFAVerify) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *TwoFAVerify) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *TwoFAVerify) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAuthType

`func (o *TwoFAVerify) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *TwoFAVerify) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *TwoFAVerify) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


