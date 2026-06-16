# ExecuteCheckResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsExecute** | **bool** |  | [readonly] [default to false]
**Checked** | **string** |  | [readonly] 
**Warning** | **string** |  | [readonly] 
**Error** | **string** |  | [readonly] 
**WarningCount** | **int32** |  | [readonly] 
**ErrorCount** | **int32** |  | [readonly] 
**IsCritical** | **bool** |  | [readonly] [default to false]
**SyntaxType** | **int32** |  | [readonly] 
**Rows** | **interface{}** |  | [readonly] 
**ColumnList** | **interface{}** |  | [readonly] 
**Status** | **string** |  | [readonly] 
**AffectedRows** | **int32** |  | [readonly] 

## Methods

### NewExecuteCheckResult

`func NewExecuteCheckResult(isExecute bool, checked string, warning string, error_ string, warningCount int32, errorCount int32, isCritical bool, syntaxType int32, rows interface{}, columnList interface{}, status string, affectedRows int32, ) *ExecuteCheckResult`

NewExecuteCheckResult instantiates a new ExecuteCheckResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteCheckResultWithDefaults

`func NewExecuteCheckResultWithDefaults() *ExecuteCheckResult`

NewExecuteCheckResultWithDefaults instantiates a new ExecuteCheckResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsExecute

`func (o *ExecuteCheckResult) GetIsExecute() bool`

GetIsExecute returns the IsExecute field if non-nil, zero value otherwise.

### GetIsExecuteOk

`func (o *ExecuteCheckResult) GetIsExecuteOk() (*bool, bool)`

GetIsExecuteOk returns a tuple with the IsExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExecute

`func (o *ExecuteCheckResult) SetIsExecute(v bool)`

SetIsExecute sets IsExecute field to given value.


### GetChecked

`func (o *ExecuteCheckResult) GetChecked() string`

GetChecked returns the Checked field if non-nil, zero value otherwise.

### GetCheckedOk

`func (o *ExecuteCheckResult) GetCheckedOk() (*string, bool)`

GetCheckedOk returns a tuple with the Checked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecked

`func (o *ExecuteCheckResult) SetChecked(v string)`

SetChecked sets Checked field to given value.


### GetWarning

`func (o *ExecuteCheckResult) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *ExecuteCheckResult) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *ExecuteCheckResult) SetWarning(v string)`

SetWarning sets Warning field to given value.


### GetError

`func (o *ExecuteCheckResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ExecuteCheckResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ExecuteCheckResult) SetError(v string)`

SetError sets Error field to given value.


### GetWarningCount

`func (o *ExecuteCheckResult) GetWarningCount() int32`

GetWarningCount returns the WarningCount field if non-nil, zero value otherwise.

### GetWarningCountOk

`func (o *ExecuteCheckResult) GetWarningCountOk() (*int32, bool)`

GetWarningCountOk returns a tuple with the WarningCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningCount

`func (o *ExecuteCheckResult) SetWarningCount(v int32)`

SetWarningCount sets WarningCount field to given value.


### GetErrorCount

`func (o *ExecuteCheckResult) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *ExecuteCheckResult) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *ExecuteCheckResult) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.


### GetIsCritical

`func (o *ExecuteCheckResult) GetIsCritical() bool`

GetIsCritical returns the IsCritical field if non-nil, zero value otherwise.

### GetIsCriticalOk

`func (o *ExecuteCheckResult) GetIsCriticalOk() (*bool, bool)`

GetIsCriticalOk returns a tuple with the IsCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCritical

`func (o *ExecuteCheckResult) SetIsCritical(v bool)`

SetIsCritical sets IsCritical field to given value.


### GetSyntaxType

`func (o *ExecuteCheckResult) GetSyntaxType() int32`

GetSyntaxType returns the SyntaxType field if non-nil, zero value otherwise.

### GetSyntaxTypeOk

`func (o *ExecuteCheckResult) GetSyntaxTypeOk() (*int32, bool)`

GetSyntaxTypeOk returns a tuple with the SyntaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntaxType

`func (o *ExecuteCheckResult) SetSyntaxType(v int32)`

SetSyntaxType sets SyntaxType field to given value.


### GetRows

`func (o *ExecuteCheckResult) GetRows() interface{}`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ExecuteCheckResult) GetRowsOk() (*interface{}, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ExecuteCheckResult) SetRows(v interface{})`

SetRows sets Rows field to given value.


### SetRowsNil

`func (o *ExecuteCheckResult) SetRowsNil(b bool)`

 SetRowsNil sets the value for Rows to be an explicit nil

### UnsetRows
`func (o *ExecuteCheckResult) UnsetRows()`

UnsetRows ensures that no value is present for Rows, not even an explicit nil
### GetColumnList

`func (o *ExecuteCheckResult) GetColumnList() interface{}`

GetColumnList returns the ColumnList field if non-nil, zero value otherwise.

### GetColumnListOk

`func (o *ExecuteCheckResult) GetColumnListOk() (*interface{}, bool)`

GetColumnListOk returns a tuple with the ColumnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnList

`func (o *ExecuteCheckResult) SetColumnList(v interface{})`

SetColumnList sets ColumnList field to given value.


### SetColumnListNil

`func (o *ExecuteCheckResult) SetColumnListNil(b bool)`

 SetColumnListNil sets the value for ColumnList to be an explicit nil

### UnsetColumnList
`func (o *ExecuteCheckResult) UnsetColumnList()`

UnsetColumnList ensures that no value is present for ColumnList, not even an explicit nil
### GetStatus

`func (o *ExecuteCheckResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecuteCheckResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecuteCheckResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAffectedRows

`func (o *ExecuteCheckResult) GetAffectedRows() int32`

GetAffectedRows returns the AffectedRows field if non-nil, zero value otherwise.

### GetAffectedRowsOk

`func (o *ExecuteCheckResult) GetAffectedRowsOk() (*int32, bool)`

GetAffectedRowsOk returns a tuple with the AffectedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedRows

`func (o *ExecuteCheckResult) SetAffectedRows(v int32)`

SetAffectedRows sets AffectedRows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


