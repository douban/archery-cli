# WorkflowLogList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationTypeDesc** | **string** |  | 
**OperationInfo** | **string** |  | 
**OperatorDisplay** | Pointer to **string** |  | [optional] 
**OperationTime** | **time.Time** |  | [readonly] 

## Methods

### NewWorkflowLogList

`func NewWorkflowLogList(operationTypeDesc string, operationInfo string, operationTime time.Time, ) *WorkflowLogList`

NewWorkflowLogList instantiates a new WorkflowLogList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowLogListWithDefaults

`func NewWorkflowLogListWithDefaults() *WorkflowLogList`

NewWorkflowLogListWithDefaults instantiates a new WorkflowLogList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationTypeDesc

`func (o *WorkflowLogList) GetOperationTypeDesc() string`

GetOperationTypeDesc returns the OperationTypeDesc field if non-nil, zero value otherwise.

### GetOperationTypeDescOk

`func (o *WorkflowLogList) GetOperationTypeDescOk() (*string, bool)`

GetOperationTypeDescOk returns a tuple with the OperationTypeDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationTypeDesc

`func (o *WorkflowLogList) SetOperationTypeDesc(v string)`

SetOperationTypeDesc sets OperationTypeDesc field to given value.


### GetOperationInfo

`func (o *WorkflowLogList) GetOperationInfo() string`

GetOperationInfo returns the OperationInfo field if non-nil, zero value otherwise.

### GetOperationInfoOk

`func (o *WorkflowLogList) GetOperationInfoOk() (*string, bool)`

GetOperationInfoOk returns a tuple with the OperationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationInfo

`func (o *WorkflowLogList) SetOperationInfo(v string)`

SetOperationInfo sets OperationInfo field to given value.


### GetOperatorDisplay

`func (o *WorkflowLogList) GetOperatorDisplay() string`

GetOperatorDisplay returns the OperatorDisplay field if non-nil, zero value otherwise.

### GetOperatorDisplayOk

`func (o *WorkflowLogList) GetOperatorDisplayOk() (*string, bool)`

GetOperatorDisplayOk returns a tuple with the OperatorDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorDisplay

`func (o *WorkflowLogList) SetOperatorDisplay(v string)`

SetOperatorDisplay sets OperatorDisplay field to given value.

### HasOperatorDisplay

`func (o *WorkflowLogList) HasOperatorDisplay() bool`

HasOperatorDisplay returns a boolean if a field has been set.

### GetOperationTime

`func (o *WorkflowLogList) GetOperationTime() time.Time`

GetOperationTime returns the OperationTime field if non-nil, zero value otherwise.

### GetOperationTimeOk

`func (o *WorkflowLogList) GetOperationTimeOk() (*time.Time, bool)`

GetOperationTimeOk returns a tuple with the OperationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationTime

`func (o *WorkflowLogList) SetOperationTime(v time.Time)`

SetOperationTime sets OperationTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


