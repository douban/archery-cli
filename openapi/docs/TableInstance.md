# TableInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**DbType** | **string** |  | 
**DbName** | **string** |  | 
**TableName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTableInstance

`func NewTableInstance(id int32, name string, dbType string, dbName string, ) *TableInstance`

NewTableInstance instantiates a new TableInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableInstanceWithDefaults

`func NewTableInstanceWithDefaults() *TableInstance`

NewTableInstanceWithDefaults instantiates a new TableInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TableInstance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TableInstance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TableInstance) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *TableInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TableInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TableInstance) SetName(v string)`

SetName sets Name field to given value.


### GetDbType

`func (o *TableInstance) GetDbType() string`

GetDbType returns the DbType field if non-nil, zero value otherwise.

### GetDbTypeOk

`func (o *TableInstance) GetDbTypeOk() (*string, bool)`

GetDbTypeOk returns a tuple with the DbType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbType

`func (o *TableInstance) SetDbType(v string)`

SetDbType sets DbType field to given value.


### GetDbName

`func (o *TableInstance) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *TableInstance) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *TableInstance) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetTableName

`func (o *TableInstance) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *TableInstance) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *TableInstance) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *TableInstance) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### SetTableNameNil

`func (o *TableInstance) SetTableNameNil(b bool)`

 SetTableNameNil sets the value for TableName to be an explicit nil

### UnsetTableName
`func (o *TableInstance) UnsetTableName()`

UnsetTableName ensures that no value is present for TableName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


