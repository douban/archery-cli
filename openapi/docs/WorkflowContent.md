# WorkflowContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**WorkflowId** | **int32** |  | [readonly] 
**Workflow** | [**Workflow**](Workflow.md) |  | 
**SqlContent** | **string** |  | 
**ReviewContent** | **string** |  | [readonly] 
**ExecuteResult** | **string** |  | [readonly] 

## Methods

### NewWorkflowContent

`func NewWorkflowContent(id int32, workflowId int32, workflow Workflow, sqlContent string, reviewContent string, executeResult string, ) *WorkflowContent`

NewWorkflowContent instantiates a new WorkflowContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowContentWithDefaults

`func NewWorkflowContentWithDefaults() *WorkflowContent`

NewWorkflowContentWithDefaults instantiates a new WorkflowContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowContent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowContent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowContent) SetId(v int32)`

SetId sets Id field to given value.


### GetWorkflowId

`func (o *WorkflowContent) GetWorkflowId() int32`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowContent) GetWorkflowIdOk() (*int32, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowContent) SetWorkflowId(v int32)`

SetWorkflowId sets WorkflowId field to given value.


### GetWorkflow

`func (o *WorkflowContent) GetWorkflow() Workflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WorkflowContent) GetWorkflowOk() (*Workflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WorkflowContent) SetWorkflow(v Workflow)`

SetWorkflow sets Workflow field to given value.


### GetSqlContent

`func (o *WorkflowContent) GetSqlContent() string`

GetSqlContent returns the SqlContent field if non-nil, zero value otherwise.

### GetSqlContentOk

`func (o *WorkflowContent) GetSqlContentOk() (*string, bool)`

GetSqlContentOk returns a tuple with the SqlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlContent

`func (o *WorkflowContent) SetSqlContent(v string)`

SetSqlContent sets SqlContent field to given value.


### GetReviewContent

`func (o *WorkflowContent) GetReviewContent() string`

GetReviewContent returns the ReviewContent field if non-nil, zero value otherwise.

### GetReviewContentOk

`func (o *WorkflowContent) GetReviewContentOk() (*string, bool)`

GetReviewContentOk returns a tuple with the ReviewContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewContent

`func (o *WorkflowContent) SetReviewContent(v string)`

SetReviewContent sets ReviewContent field to given value.


### GetExecuteResult

`func (o *WorkflowContent) GetExecuteResult() string`

GetExecuteResult returns the ExecuteResult field if non-nil, zero value otherwise.

### GetExecuteResultOk

`func (o *WorkflowContent) GetExecuteResultOk() (*string, bool)`

GetExecuteResultOk returns a tuple with the ExecuteResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteResult

`func (o *WorkflowContent) SetExecuteResult(v string)`

SetExecuteResult sets ExecuteResult field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


