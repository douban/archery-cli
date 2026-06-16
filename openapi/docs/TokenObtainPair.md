# TokenObtainPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Password** | **string** |  | 
**Access** | **string** |  | [readonly] 
**Refresh** | **string** |  | [readonly] 

## Methods

### NewTokenObtainPair

`func NewTokenObtainPair(username string, password string, access string, refresh string, ) *TokenObtainPair`

NewTokenObtainPair instantiates a new TokenObtainPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenObtainPairWithDefaults

`func NewTokenObtainPairWithDefaults() *TokenObtainPair`

NewTokenObtainPairWithDefaults instantiates a new TokenObtainPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *TokenObtainPair) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TokenObtainPair) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TokenObtainPair) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *TokenObtainPair) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TokenObtainPair) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TokenObtainPair) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetAccess

`func (o *TokenObtainPair) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *TokenObtainPair) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *TokenObtainPair) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetRefresh

`func (o *TokenObtainPair) GetRefresh() string`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *TokenObtainPair) GetRefreshOk() (*string, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *TokenObtainPair) SetRefresh(v string)`

SetRefresh sets Refresh field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


