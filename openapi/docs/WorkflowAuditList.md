# WorkflowAuditList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditId** | **int32** |  | [readonly] 
**GroupName** | **string** |  | 
**WorkflowType** | [**WorkflowTypeA27Enum**](WorkflowTypeA27Enum.md) |  | 
**WorkflowTitle** | **string** |  | 
**AuditAuthGroups** | **string** |  | 
**CurrentAudit** | **string** |  | 
**CurrentStatus** | [**CurrentStatusEnum**](CurrentStatusEnum.md) |  | 
**CreateUserDisplay** | Pointer to **string** |  | [optional] 
**CreateTime** | **time.Time** |  | [readonly] 

## Methods

### NewWorkflowAuditList

`func NewWorkflowAuditList(auditId int32, groupName string, workflowType WorkflowTypeA27Enum, workflowTitle string, auditAuthGroups string, currentAudit string, currentStatus CurrentStatusEnum, createTime time.Time, ) *WorkflowAuditList`

NewWorkflowAuditList instantiates a new WorkflowAuditList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowAuditListWithDefaults

`func NewWorkflowAuditListWithDefaults() *WorkflowAuditList`

NewWorkflowAuditListWithDefaults instantiates a new WorkflowAuditList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditId

`func (o *WorkflowAuditList) GetAuditId() int32`

GetAuditId returns the AuditId field if non-nil, zero value otherwise.

### GetAuditIdOk

`func (o *WorkflowAuditList) GetAuditIdOk() (*int32, bool)`

GetAuditIdOk returns a tuple with the AuditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditId

`func (o *WorkflowAuditList) SetAuditId(v int32)`

SetAuditId sets AuditId field to given value.


### GetGroupName

`func (o *WorkflowAuditList) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *WorkflowAuditList) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *WorkflowAuditList) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetWorkflowType

`func (o *WorkflowAuditList) GetWorkflowType() WorkflowTypeA27Enum`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *WorkflowAuditList) GetWorkflowTypeOk() (*WorkflowTypeA27Enum, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *WorkflowAuditList) SetWorkflowType(v WorkflowTypeA27Enum)`

SetWorkflowType sets WorkflowType field to given value.


### GetWorkflowTitle

`func (o *WorkflowAuditList) GetWorkflowTitle() string`

GetWorkflowTitle returns the WorkflowTitle field if non-nil, zero value otherwise.

### GetWorkflowTitleOk

`func (o *WorkflowAuditList) GetWorkflowTitleOk() (*string, bool)`

GetWorkflowTitleOk returns a tuple with the WorkflowTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTitle

`func (o *WorkflowAuditList) SetWorkflowTitle(v string)`

SetWorkflowTitle sets WorkflowTitle field to given value.


### GetAuditAuthGroups

`func (o *WorkflowAuditList) GetAuditAuthGroups() string`

GetAuditAuthGroups returns the AuditAuthGroups field if non-nil, zero value otherwise.

### GetAuditAuthGroupsOk

`func (o *WorkflowAuditList) GetAuditAuthGroupsOk() (*string, bool)`

GetAuditAuthGroupsOk returns a tuple with the AuditAuthGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditAuthGroups

`func (o *WorkflowAuditList) SetAuditAuthGroups(v string)`

SetAuditAuthGroups sets AuditAuthGroups field to given value.


### GetCurrentAudit

`func (o *WorkflowAuditList) GetCurrentAudit() string`

GetCurrentAudit returns the CurrentAudit field if non-nil, zero value otherwise.

### GetCurrentAuditOk

`func (o *WorkflowAuditList) GetCurrentAuditOk() (*string, bool)`

GetCurrentAuditOk returns a tuple with the CurrentAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAudit

`func (o *WorkflowAuditList) SetCurrentAudit(v string)`

SetCurrentAudit sets CurrentAudit field to given value.


### GetCurrentStatus

`func (o *WorkflowAuditList) GetCurrentStatus() CurrentStatusEnum`

GetCurrentStatus returns the CurrentStatus field if non-nil, zero value otherwise.

### GetCurrentStatusOk

`func (o *WorkflowAuditList) GetCurrentStatusOk() (*CurrentStatusEnum, bool)`

GetCurrentStatusOk returns a tuple with the CurrentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStatus

`func (o *WorkflowAuditList) SetCurrentStatus(v CurrentStatusEnum)`

SetCurrentStatus sets CurrentStatus field to given value.


### GetCreateUserDisplay

`func (o *WorkflowAuditList) GetCreateUserDisplay() string`

GetCreateUserDisplay returns the CreateUserDisplay field if non-nil, zero value otherwise.

### GetCreateUserDisplayOk

`func (o *WorkflowAuditList) GetCreateUserDisplayOk() (*string, bool)`

GetCreateUserDisplayOk returns a tuple with the CreateUserDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUserDisplay

`func (o *WorkflowAuditList) SetCreateUserDisplay(v string)`

SetCreateUserDisplay sets CreateUserDisplay field to given value.

### HasCreateUserDisplay

`func (o *WorkflowAuditList) HasCreateUserDisplay() bool`

HasCreateUserDisplay returns a boolean if a field has been set.

### GetCreateTime

`func (o *WorkflowAuditList) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *WorkflowAuditList) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *WorkflowAuditList) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


