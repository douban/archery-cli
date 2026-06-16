# WorkflowLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **int32** |  | 
**WorkflowType** | [**WorkflowLogWorkflowTypeEnum**](WorkflowLogWorkflowTypeEnum.md) |  | 

## Methods

### NewWorkflowLog

`func NewWorkflowLog(workflowId int32, workflowType WorkflowLogWorkflowTypeEnum, ) *WorkflowLog`

NewWorkflowLog instantiates a new WorkflowLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowLogWithDefaults

`func NewWorkflowLogWithDefaults() *WorkflowLog`

NewWorkflowLogWithDefaults instantiates a new WorkflowLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *WorkflowLog) GetWorkflowId() int32`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowLog) GetWorkflowIdOk() (*int32, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowLog) SetWorkflowId(v int32)`

SetWorkflowId sets WorkflowId field to given value.


### GetWorkflowType

`func (o *WorkflowLog) GetWorkflowType() WorkflowLogWorkflowTypeEnum`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *WorkflowLog) GetWorkflowTypeOk() (*WorkflowLogWorkflowTypeEnum, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *WorkflowLog) SetWorkflowType(v WorkflowLogWorkflowTypeEnum)`

SetWorkflowType sets WorkflowType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


