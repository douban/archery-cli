# UserAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engineer** | **string** |  | 
**Password** | **string** |  | 

## Methods

### NewUserAuth

`func NewUserAuth(engineer string, password string, ) *UserAuth`

NewUserAuth instantiates a new UserAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAuthWithDefaults

`func NewUserAuthWithDefaults() *UserAuth`

NewUserAuthWithDefaults instantiates a new UserAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineer

`func (o *UserAuth) GetEngineer() string`

GetEngineer returns the Engineer field if non-nil, zero value otherwise.

### GetEngineerOk

`func (o *UserAuth) GetEngineerOk() (*string, bool)`

GetEngineerOk returns a tuple with the Engineer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineer

`func (o *UserAuth) SetEngineer(v string)`

SetEngineer sets Engineer field to given value.


### GetPassword

`func (o *UserAuth) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserAuth) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserAuth) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


