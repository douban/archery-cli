# AliyunRds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**RdsDbinstanceid** | **string** |  | 
**IsEnable** | Pointer to **bool** |  | [optional] 
**Instance** | **int32** |  | 
**Ak** | [**CloudAccessKey**](CloudAccessKey.md) |  | 

## Methods

### NewAliyunRds

`func NewAliyunRds(id int32, rdsDbinstanceid string, instance int32, ak CloudAccessKey, ) *AliyunRds`

NewAliyunRds instantiates a new AliyunRds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAliyunRdsWithDefaults

`func NewAliyunRdsWithDefaults() *AliyunRds`

NewAliyunRdsWithDefaults instantiates a new AliyunRds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AliyunRds) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AliyunRds) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AliyunRds) SetId(v int32)`

SetId sets Id field to given value.


### GetRdsDbinstanceid

`func (o *AliyunRds) GetRdsDbinstanceid() string`

GetRdsDbinstanceid returns the RdsDbinstanceid field if non-nil, zero value otherwise.

### GetRdsDbinstanceidOk

`func (o *AliyunRds) GetRdsDbinstanceidOk() (*string, bool)`

GetRdsDbinstanceidOk returns a tuple with the RdsDbinstanceid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsDbinstanceid

`func (o *AliyunRds) SetRdsDbinstanceid(v string)`

SetRdsDbinstanceid sets RdsDbinstanceid field to given value.


### GetIsEnable

`func (o *AliyunRds) GetIsEnable() bool`

GetIsEnable returns the IsEnable field if non-nil, zero value otherwise.

### GetIsEnableOk

`func (o *AliyunRds) GetIsEnableOk() (*bool, bool)`

GetIsEnableOk returns a tuple with the IsEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnable

`func (o *AliyunRds) SetIsEnable(v bool)`

SetIsEnable sets IsEnable field to given value.

### HasIsEnable

`func (o *AliyunRds) HasIsEnable() bool`

HasIsEnable returns a boolean if a field has been set.

### GetInstance

`func (o *AliyunRds) GetInstance() int32`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *AliyunRds) GetInstanceOk() (*int32, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *AliyunRds) SetInstance(v int32)`

SetInstance sets Instance field to given value.


### GetAk

`func (o *AliyunRds) GetAk() CloudAccessKey`

GetAk returns the Ak field if non-nil, zero value otherwise.

### GetAkOk

`func (o *AliyunRds) GetAkOk() (*CloudAccessKey, bool)`

GetAkOk returns a tuple with the Ak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAk

`func (o *AliyunRds) SetAk(v CloudAccessKey)`

SetAk sets Ak field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


