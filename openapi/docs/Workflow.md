# Workflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**WorkflowName** | **string** |  | 
**DemandUrl** | Pointer to **string** |  | [optional] 
**GroupId** | **int32** |  | 
**GroupName** | **string** |  | [readonly] 
**DbName** | **string** |  | 
**SyntaxType** | [**SyntaxTypeEnum**](SyntaxTypeEnum.md) |  | [readonly] 
**IsBackup** | Pointer to **bool** | * &#x60;False&#x60; - 否 * &#x60;True&#x60; - 是 | [optional] 
**Engineer** | Pointer to **string** |  | [optional] 
**EngineerDisplay** | **string** |  | [readonly] 
**Status** | [**StatusEnum**](StatusEnum.md) |  | [readonly] 
**AuditAuthGroups** | **string** |  | [readonly] 
**RunDateStart** | Pointer to **NullableTime** |  | [optional] 
**RunDateEnd** | Pointer to **NullableTime** |  | [optional] 
**CreateTime** | **time.Time** |  | [readonly] 
**FinishTime** | **NullableTime** |  | [readonly] 
**IsManual** | [**IsOfflineExportEnum**](IsOfflineExportEnum.md) |  | [readonly] 
**IsOfflineExport** | Pointer to [**IsOfflineExportEnum**](IsOfflineExportEnum.md) |  | [optional] 
**ExportFormat** | Pointer to **NullableString** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**Instance** | **int32** |  | 

## Methods

### NewWorkflow

`func NewWorkflow(id int32, workflowName string, groupId int32, groupName string, dbName string, syntaxType SyntaxTypeEnum, engineerDisplay string, status StatusEnum, auditAuthGroups string, createTime time.Time, finishTime NullableTime, isManual IsOfflineExportEnum, instance int32, ) *Workflow`

NewWorkflow instantiates a new Workflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWithDefaults

`func NewWorkflowWithDefaults() *Workflow`

NewWorkflowWithDefaults instantiates a new Workflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Workflow) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Workflow) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Workflow) SetId(v int32)`

SetId sets Id field to given value.


### GetWorkflowName

`func (o *Workflow) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *Workflow) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *Workflow) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetDemandUrl

`func (o *Workflow) GetDemandUrl() string`

GetDemandUrl returns the DemandUrl field if non-nil, zero value otherwise.

### GetDemandUrlOk

`func (o *Workflow) GetDemandUrlOk() (*string, bool)`

GetDemandUrlOk returns a tuple with the DemandUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemandUrl

`func (o *Workflow) SetDemandUrl(v string)`

SetDemandUrl sets DemandUrl field to given value.

### HasDemandUrl

`func (o *Workflow) HasDemandUrl() bool`

HasDemandUrl returns a boolean if a field has been set.

### GetGroupId

`func (o *Workflow) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Workflow) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Workflow) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetGroupName

`func (o *Workflow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *Workflow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *Workflow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetDbName

`func (o *Workflow) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *Workflow) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *Workflow) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetSyntaxType

`func (o *Workflow) GetSyntaxType() SyntaxTypeEnum`

GetSyntaxType returns the SyntaxType field if non-nil, zero value otherwise.

### GetSyntaxTypeOk

`func (o *Workflow) GetSyntaxTypeOk() (*SyntaxTypeEnum, bool)`

GetSyntaxTypeOk returns a tuple with the SyntaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntaxType

`func (o *Workflow) SetSyntaxType(v SyntaxTypeEnum)`

SetSyntaxType sets SyntaxType field to given value.


### GetIsBackup

`func (o *Workflow) GetIsBackup() bool`

GetIsBackup returns the IsBackup field if non-nil, zero value otherwise.

### GetIsBackupOk

`func (o *Workflow) GetIsBackupOk() (*bool, bool)`

GetIsBackupOk returns a tuple with the IsBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackup

`func (o *Workflow) SetIsBackup(v bool)`

SetIsBackup sets IsBackup field to given value.

### HasIsBackup

`func (o *Workflow) HasIsBackup() bool`

HasIsBackup returns a boolean if a field has been set.

### GetEngineer

`func (o *Workflow) GetEngineer() string`

GetEngineer returns the Engineer field if non-nil, zero value otherwise.

### GetEngineerOk

`func (o *Workflow) GetEngineerOk() (*string, bool)`

GetEngineerOk returns a tuple with the Engineer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineer

`func (o *Workflow) SetEngineer(v string)`

SetEngineer sets Engineer field to given value.

### HasEngineer

`func (o *Workflow) HasEngineer() bool`

HasEngineer returns a boolean if a field has been set.

### GetEngineerDisplay

`func (o *Workflow) GetEngineerDisplay() string`

GetEngineerDisplay returns the EngineerDisplay field if non-nil, zero value otherwise.

### GetEngineerDisplayOk

`func (o *Workflow) GetEngineerDisplayOk() (*string, bool)`

GetEngineerDisplayOk returns a tuple with the EngineerDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineerDisplay

`func (o *Workflow) SetEngineerDisplay(v string)`

SetEngineerDisplay sets EngineerDisplay field to given value.


### GetStatus

`func (o *Workflow) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Workflow) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Workflow) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.


### GetAuditAuthGroups

`func (o *Workflow) GetAuditAuthGroups() string`

GetAuditAuthGroups returns the AuditAuthGroups field if non-nil, zero value otherwise.

### GetAuditAuthGroupsOk

`func (o *Workflow) GetAuditAuthGroupsOk() (*string, bool)`

GetAuditAuthGroupsOk returns a tuple with the AuditAuthGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditAuthGroups

`func (o *Workflow) SetAuditAuthGroups(v string)`

SetAuditAuthGroups sets AuditAuthGroups field to given value.


### GetRunDateStart

`func (o *Workflow) GetRunDateStart() time.Time`

GetRunDateStart returns the RunDateStart field if non-nil, zero value otherwise.

### GetRunDateStartOk

`func (o *Workflow) GetRunDateStartOk() (*time.Time, bool)`

GetRunDateStartOk returns a tuple with the RunDateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunDateStart

`func (o *Workflow) SetRunDateStart(v time.Time)`

SetRunDateStart sets RunDateStart field to given value.

### HasRunDateStart

`func (o *Workflow) HasRunDateStart() bool`

HasRunDateStart returns a boolean if a field has been set.

### SetRunDateStartNil

`func (o *Workflow) SetRunDateStartNil(b bool)`

 SetRunDateStartNil sets the value for RunDateStart to be an explicit nil

### UnsetRunDateStart
`func (o *Workflow) UnsetRunDateStart()`

UnsetRunDateStart ensures that no value is present for RunDateStart, not even an explicit nil
### GetRunDateEnd

`func (o *Workflow) GetRunDateEnd() time.Time`

GetRunDateEnd returns the RunDateEnd field if non-nil, zero value otherwise.

### GetRunDateEndOk

`func (o *Workflow) GetRunDateEndOk() (*time.Time, bool)`

GetRunDateEndOk returns a tuple with the RunDateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunDateEnd

`func (o *Workflow) SetRunDateEnd(v time.Time)`

SetRunDateEnd sets RunDateEnd field to given value.

### HasRunDateEnd

`func (o *Workflow) HasRunDateEnd() bool`

HasRunDateEnd returns a boolean if a field has been set.

### SetRunDateEndNil

`func (o *Workflow) SetRunDateEndNil(b bool)`

 SetRunDateEndNil sets the value for RunDateEnd to be an explicit nil

### UnsetRunDateEnd
`func (o *Workflow) UnsetRunDateEnd()`

UnsetRunDateEnd ensures that no value is present for RunDateEnd, not even an explicit nil
### GetCreateTime

`func (o *Workflow) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Workflow) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Workflow) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.


### GetFinishTime

`func (o *Workflow) GetFinishTime() time.Time`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *Workflow) GetFinishTimeOk() (*time.Time, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *Workflow) SetFinishTime(v time.Time)`

SetFinishTime sets FinishTime field to given value.


### SetFinishTimeNil

`func (o *Workflow) SetFinishTimeNil(b bool)`

 SetFinishTimeNil sets the value for FinishTime to be an explicit nil

### UnsetFinishTime
`func (o *Workflow) UnsetFinishTime()`

UnsetFinishTime ensures that no value is present for FinishTime, not even an explicit nil
### GetIsManual

`func (o *Workflow) GetIsManual() IsOfflineExportEnum`

GetIsManual returns the IsManual field if non-nil, zero value otherwise.

### GetIsManualOk

`func (o *Workflow) GetIsManualOk() (*IsOfflineExportEnum, bool)`

GetIsManualOk returns a tuple with the IsManual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManual

`func (o *Workflow) SetIsManual(v IsOfflineExportEnum)`

SetIsManual sets IsManual field to given value.


### GetIsOfflineExport

`func (o *Workflow) GetIsOfflineExport() IsOfflineExportEnum`

GetIsOfflineExport returns the IsOfflineExport field if non-nil, zero value otherwise.

### GetIsOfflineExportOk

`func (o *Workflow) GetIsOfflineExportOk() (*IsOfflineExportEnum, bool)`

GetIsOfflineExportOk returns a tuple with the IsOfflineExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOfflineExport

`func (o *Workflow) SetIsOfflineExport(v IsOfflineExportEnum)`

SetIsOfflineExport sets IsOfflineExport field to given value.

### HasIsOfflineExport

`func (o *Workflow) HasIsOfflineExport() bool`

HasIsOfflineExport returns a boolean if a field has been set.

### GetExportFormat

`func (o *Workflow) GetExportFormat() string`

GetExportFormat returns the ExportFormat field if non-nil, zero value otherwise.

### GetExportFormatOk

`func (o *Workflow) GetExportFormatOk() (*string, bool)`

GetExportFormatOk returns a tuple with the ExportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportFormat

`func (o *Workflow) SetExportFormat(v string)`

SetExportFormat sets ExportFormat field to given value.

### HasExportFormat

`func (o *Workflow) HasExportFormat() bool`

HasExportFormat returns a boolean if a field has been set.

### SetExportFormatNil

`func (o *Workflow) SetExportFormatNil(b bool)`

 SetExportFormatNil sets the value for ExportFormat to be an explicit nil

### UnsetExportFormat
`func (o *Workflow) UnsetExportFormat()`

UnsetExportFormat ensures that no value is present for ExportFormat, not even an explicit nil
### GetFileName

`func (o *Workflow) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Workflow) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Workflow) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *Workflow) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *Workflow) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *Workflow) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetInstance

`func (o *Workflow) GetInstance() int32`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *Workflow) GetInstanceOk() (*int32, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *Workflow) SetInstance(v int32)`

SetInstance sets Instance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


