# LocatorExecutionSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessedInstanceCount** | **int32** |  | 
**SuccessfulInstanceCount** | **int32** |  | 
**FailedInstanceCount** | **int32** |  | 
**FailureReasons** | [**[]LocatorFailureReason**](LocatorFailureReason.md) |  | 

## Methods

### NewLocatorExecutionSummary

`func NewLocatorExecutionSummary(processedInstanceCount int32, successfulInstanceCount int32, failedInstanceCount int32, failureReasons []LocatorFailureReason, ) *LocatorExecutionSummary`

NewLocatorExecutionSummary instantiates a new LocatorExecutionSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocatorExecutionSummaryWithDefaults

`func NewLocatorExecutionSummaryWithDefaults() *LocatorExecutionSummary`

NewLocatorExecutionSummaryWithDefaults instantiates a new LocatorExecutionSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessedInstanceCount

`func (o *LocatorExecutionSummary) GetProcessedInstanceCount() int32`

GetProcessedInstanceCount returns the ProcessedInstanceCount field if non-nil, zero value otherwise.

### GetProcessedInstanceCountOk

`func (o *LocatorExecutionSummary) GetProcessedInstanceCountOk() (*int32, bool)`

GetProcessedInstanceCountOk returns a tuple with the ProcessedInstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedInstanceCount

`func (o *LocatorExecutionSummary) SetProcessedInstanceCount(v int32)`

SetProcessedInstanceCount sets ProcessedInstanceCount field to given value.


### GetSuccessfulInstanceCount

`func (o *LocatorExecutionSummary) GetSuccessfulInstanceCount() int32`

GetSuccessfulInstanceCount returns the SuccessfulInstanceCount field if non-nil, zero value otherwise.

### GetSuccessfulInstanceCountOk

`func (o *LocatorExecutionSummary) GetSuccessfulInstanceCountOk() (*int32, bool)`

GetSuccessfulInstanceCountOk returns a tuple with the SuccessfulInstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulInstanceCount

`func (o *LocatorExecutionSummary) SetSuccessfulInstanceCount(v int32)`

SetSuccessfulInstanceCount sets SuccessfulInstanceCount field to given value.


### GetFailedInstanceCount

`func (o *LocatorExecutionSummary) GetFailedInstanceCount() int32`

GetFailedInstanceCount returns the FailedInstanceCount field if non-nil, zero value otherwise.

### GetFailedInstanceCountOk

`func (o *LocatorExecutionSummary) GetFailedInstanceCountOk() (*int32, bool)`

GetFailedInstanceCountOk returns a tuple with the FailedInstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedInstanceCount

`func (o *LocatorExecutionSummary) SetFailedInstanceCount(v int32)`

SetFailedInstanceCount sets FailedInstanceCount field to given value.


### GetFailureReasons

`func (o *LocatorExecutionSummary) GetFailureReasons() []LocatorFailureReason`

GetFailureReasons returns the FailureReasons field if non-nil, zero value otherwise.

### GetFailureReasonsOk

`func (o *LocatorExecutionSummary) GetFailureReasonsOk() (*[]LocatorFailureReason, bool)`

GetFailureReasonsOk returns a tuple with the FailureReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReasons

`func (o *LocatorExecutionSummary) SetFailureReasons(v []LocatorFailureReason)`

SetFailureReasons sets FailureReasons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


