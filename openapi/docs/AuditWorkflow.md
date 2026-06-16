# AuditWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engineer** | **string** |  | 
**WorkflowId** | **int32** |  | 
**AuditRemark** | **string** |  | 
**WorkflowType** | [**WorkflowTypeA27Enum**](WorkflowTypeA27Enum.md) |  | 
**AuditType** | [**AuditTypeEnum**](AuditTypeEnum.md) |  | 

## Methods

### NewAuditWorkflow

`func NewAuditWorkflow(engineer string, workflowId int32, auditRemark string, workflowType WorkflowTypeA27Enum, auditType AuditTypeEnum, ) *AuditWorkflow`

NewAuditWorkflow instantiates a new AuditWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditWorkflowWithDefaults

`func NewAuditWorkflowWithDefaults() *AuditWorkflow`

NewAuditWorkflowWithDefaults instantiates a new AuditWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineer

`func (o *AuditWorkflow) GetEngineer() string`

GetEngineer returns the Engineer field if non-nil, zero value otherwise.

### GetEngineerOk

`func (o *AuditWorkflow) GetEngineerOk() (*string, bool)`

GetEngineerOk returns a tuple with the Engineer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineer

`func (o *AuditWorkflow) SetEngineer(v string)`

SetEngineer sets Engineer field to given value.


### GetWorkflowId

`func (o *AuditWorkflow) GetWorkflowId() int32`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *AuditWorkflow) GetWorkflowIdOk() (*int32, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *AuditWorkflow) SetWorkflowId(v int32)`

SetWorkflowId sets WorkflowId field to given value.


### GetAuditRemark

`func (o *AuditWorkflow) GetAuditRemark() string`

GetAuditRemark returns the AuditRemark field if non-nil, zero value otherwise.

### GetAuditRemarkOk

`func (o *AuditWorkflow) GetAuditRemarkOk() (*string, bool)`

GetAuditRemarkOk returns a tuple with the AuditRemark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditRemark

`func (o *AuditWorkflow) SetAuditRemark(v string)`

SetAuditRemark sets AuditRemark field to given value.


### GetWorkflowType

`func (o *AuditWorkflow) GetWorkflowType() WorkflowTypeA27Enum`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *AuditWorkflow) GetWorkflowTypeOk() (*WorkflowTypeA27Enum, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *AuditWorkflow) SetWorkflowType(v WorkflowTypeA27Enum)`

SetWorkflowType sets WorkflowType field to given value.


### GetAuditType

`func (o *AuditWorkflow) GetAuditType() AuditTypeEnum`

GetAuditType returns the AuditType field if non-nil, zero value otherwise.

### GetAuditTypeOk

`func (o *AuditWorkflow) GetAuditTypeOk() (*AuditTypeEnum, bool)`

GetAuditTypeOk returns a tuple with the AuditType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditType

`func (o *AuditWorkflow) SetAuditType(v AuditTypeEnum)`

SetAuditType sets AuditType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


