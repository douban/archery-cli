# TableInstanceLookupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** |  | 
**Msg** | **string** |  | 
**Count** | **int32** |  | 
**Data** | [**[]TableInstance**](TableInstance.md) |  | 
**Summary** | Pointer to [**NullableLocatorExecutionSummary**](LocatorExecutionSummary.md) |  | [optional] 

## Methods

### NewTableInstanceLookupResponse

`func NewTableInstanceLookupResponse(status int32, msg string, count int32, data []TableInstance, ) *TableInstanceLookupResponse`

NewTableInstanceLookupResponse instantiates a new TableInstanceLookupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableInstanceLookupResponseWithDefaults

`func NewTableInstanceLookupResponseWithDefaults() *TableInstanceLookupResponse`

NewTableInstanceLookupResponseWithDefaults instantiates a new TableInstanceLookupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TableInstanceLookupResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TableInstanceLookupResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TableInstanceLookupResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetMsg

`func (o *TableInstanceLookupResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *TableInstanceLookupResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *TableInstanceLookupResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetCount

`func (o *TableInstanceLookupResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TableInstanceLookupResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TableInstanceLookupResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetData

`func (o *TableInstanceLookupResponse) GetData() []TableInstance`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TableInstanceLookupResponse) GetDataOk() (*[]TableInstance, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TableInstanceLookupResponse) SetData(v []TableInstance)`

SetData sets Data field to given value.


### GetSummary

`func (o *TableInstanceLookupResponse) GetSummary() LocatorExecutionSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *TableInstanceLookupResponse) GetSummaryOk() (*LocatorExecutionSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *TableInstanceLookupResponse) SetSummary(v LocatorExecutionSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *TableInstanceLookupResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *TableInstanceLookupResponse) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *TableInstanceLookupResponse) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


