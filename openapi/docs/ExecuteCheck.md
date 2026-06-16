# ExecuteCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **int32** |  | 
**DbName** | **string** |  | 
**FullSql** | **string** |  | 

## Methods

### NewExecuteCheck

`func NewExecuteCheck(instanceId int32, dbName string, fullSql string, ) *ExecuteCheck`

NewExecuteCheck instantiates a new ExecuteCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteCheckWithDefaults

`func NewExecuteCheckWithDefaults() *ExecuteCheck`

NewExecuteCheckWithDefaults instantiates a new ExecuteCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *ExecuteCheck) GetInstanceId() int32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ExecuteCheck) GetInstanceIdOk() (*int32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ExecuteCheck) SetInstanceId(v int32)`

SetInstanceId sets InstanceId field to given value.


### GetDbName

`func (o *ExecuteCheck) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *ExecuteCheck) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *ExecuteCheck) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetFullSql

`func (o *ExecuteCheck) GetFullSql() string`

GetFullSql returns the FullSql field if non-nil, zero value otherwise.

### GetFullSqlOk

`func (o *ExecuteCheck) GetFullSqlOk() (*string, bool)`

GetFullSqlOk returns a tuple with the FullSql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullSql

`func (o *ExecuteCheck) SetFullSql(v string)`

SetFullSql sets FullSql field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


