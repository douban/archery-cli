# CloudAccessKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Type** | Pointer to [**CloudAccessKeyTypeEnum**](CloudAccessKeyTypeEnum.md) |  | [optional] 
**KeyId** | **string** |  | 
**KeySecret** | **string** |  | 
**Remark** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudAccessKey

`func NewCloudAccessKey(id int32, keyId string, keySecret string, ) *CloudAccessKey`

NewCloudAccessKey instantiates a new CloudAccessKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAccessKeyWithDefaults

`func NewCloudAccessKeyWithDefaults() *CloudAccessKey`

NewCloudAccessKeyWithDefaults instantiates a new CloudAccessKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudAccessKey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudAccessKey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudAccessKey) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *CloudAccessKey) GetType() CloudAccessKeyTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudAccessKey) GetTypeOk() (*CloudAccessKeyTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudAccessKey) SetType(v CloudAccessKeyTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *CloudAccessKey) HasType() bool`

HasType returns a boolean if a field has been set.

### GetKeyId

`func (o *CloudAccessKey) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *CloudAccessKey) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *CloudAccessKey) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetKeySecret

`func (o *CloudAccessKey) GetKeySecret() string`

GetKeySecret returns the KeySecret field if non-nil, zero value otherwise.

### GetKeySecretOk

`func (o *CloudAccessKey) GetKeySecretOk() (*string, bool)`

GetKeySecretOk returns a tuple with the KeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySecret

`func (o *CloudAccessKey) SetKeySecret(v string)`

SetKeySecret sets KeySecret field to given value.


### GetRemark

`func (o *CloudAccessKey) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *CloudAccessKey) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *CloudAccessKey) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *CloudAccessKey) HasRemark() bool`

HasRemark returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


