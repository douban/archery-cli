# ExecuteWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engineer** | Pointer to **string** |  | [optional] 
**WorkflowId** | **int32** |  | 
**WorkflowType** | [**ExecuteWorkflowWorkflowTypeEnum**](ExecuteWorkflowWorkflowTypeEnum.md) |  | 
**Mode** | Pointer to [**ExecuteWorkflowModeEnum**](ExecuteWorkflowModeEnum.md) |  | [optional] 

## Methods

### NewExecuteWorkflow

`func NewExecuteWorkflow(workflowId int32, workflowType ExecuteWorkflowWorkflowTypeEnum, ) *ExecuteWorkflow`

NewExecuteWorkflow instantiates a new ExecuteWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteWorkflowWithDefaults

`func NewExecuteWorkflowWithDefaults() *ExecuteWorkflow`

NewExecuteWorkflowWithDefaults instantiates a new ExecuteWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineer

`func (o *ExecuteWorkflow) GetEngineer() string`

GetEngineer returns the Engineer field if non-nil, zero value otherwise.

### GetEngineerOk

`func (o *ExecuteWorkflow) GetEngineerOk() (*string, bool)`

GetEngineerOk returns a tuple with the Engineer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineer

`func (o *ExecuteWorkflow) SetEngineer(v string)`

SetEngineer sets Engineer field to given value.

### HasEngineer

`func (o *ExecuteWorkflow) HasEngineer() bool`

HasEngineer returns a boolean if a field has been set.

### GetWorkflowId

`func (o *ExecuteWorkflow) GetWorkflowId() int32`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ExecuteWorkflow) GetWorkflowIdOk() (*int32, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ExecuteWorkflow) SetWorkflowId(v int32)`

SetWorkflowId sets WorkflowId field to given value.


### GetWorkflowType

`func (o *ExecuteWorkflow) GetWorkflowType() ExecuteWorkflowWorkflowTypeEnum`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *ExecuteWorkflow) GetWorkflowTypeOk() (*ExecuteWorkflowWorkflowTypeEnum, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *ExecuteWorkflow) SetWorkflowType(v ExecuteWorkflowWorkflowTypeEnum)`

SetWorkflowType sets WorkflowType field to given value.


### GetMode

`func (o *ExecuteWorkflow) GetMode() ExecuteWorkflowModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ExecuteWorkflow) GetModeOk() (*ExecuteWorkflowModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ExecuteWorkflow) SetMode(v ExecuteWorkflowModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ExecuteWorkflow) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


