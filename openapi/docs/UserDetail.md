# UserDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Password** | Pointer to **string** |  | [optional] 
**LastLogin** | Pointer to **NullableTime** |  | [optional] 
**IsSuperuser** | Pointer to **bool** | 指明该用户缺省拥有所有权限。 | [optional] 
**Username** | Pointer to **string** | 必填；长度为150个字符或以下；只能包含字母、数字、特殊字符“@”、“.”、“-”和“_”。 | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**IsStaff** | Pointer to **bool** | 指明用户是否可以登录到这个管理站点。 | [optional] 
**IsActive** | Pointer to **bool** | 指明用户是否被认为是活跃的。以反选代替删除帐号。 | [optional] 
**DateJoined** | Pointer to **time.Time** |  | [optional] 
**Display** | Pointer to **string** |  | [optional] 
**DingUserId** | Pointer to **string** |  | [optional] 
**WxUserId** | Pointer to **string** |  | [optional] 
**FeishuOpenId** | Pointer to **string** |  | [optional] 
**FailedLoginCount** | Pointer to **int32** |  | [optional] 
**LastLoginFailedAt** | Pointer to **NullableTime** |  | [optional] 
**Groups** | Pointer to **[]int32** | 该用户归属的组。一个用户将得到其归属的组的所有权限。 | [optional] 
**UserPermissions** | Pointer to **[]int32** | 这个用户的特定权限。 | [optional] 
**ResourceGroup** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewUserDetail

`func NewUserDetail(id int32, ) *UserDetail`

NewUserDetail instantiates a new UserDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDetailWithDefaults

`func NewUserDetailWithDefaults() *UserDetail`

NewUserDetailWithDefaults instantiates a new UserDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDetail) SetId(v int32)`

SetId sets Id field to given value.


### GetPassword

`func (o *UserDetail) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserDetail) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserDetail) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserDetail) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetLastLogin

`func (o *UserDetail) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *UserDetail) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *UserDetail) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *UserDetail) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *UserDetail) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *UserDetail) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetIsSuperuser

`func (o *UserDetail) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *UserDetail) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *UserDetail) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.

### HasIsSuperuser

`func (o *UserDetail) HasIsSuperuser() bool`

HasIsSuperuser returns a boolean if a field has been set.

### GetUsername

`func (o *UserDetail) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserDetail) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserDetail) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserDetail) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFirstName

`func (o *UserDetail) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserDetail) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserDetail) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserDetail) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserDetail) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserDetail) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserDetail) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserDetail) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *UserDetail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserDetail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserDetail) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserDetail) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIsStaff

`func (o *UserDetail) GetIsStaff() bool`

GetIsStaff returns the IsStaff field if non-nil, zero value otherwise.

### GetIsStaffOk

`func (o *UserDetail) GetIsStaffOk() (*bool, bool)`

GetIsStaffOk returns a tuple with the IsStaff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStaff

`func (o *UserDetail) SetIsStaff(v bool)`

SetIsStaff sets IsStaff field to given value.

### HasIsStaff

`func (o *UserDetail) HasIsStaff() bool`

HasIsStaff returns a boolean if a field has been set.

### GetIsActive

`func (o *UserDetail) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UserDetail) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UserDetail) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UserDetail) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetDateJoined

`func (o *UserDetail) GetDateJoined() time.Time`

GetDateJoined returns the DateJoined field if non-nil, zero value otherwise.

### GetDateJoinedOk

`func (o *UserDetail) GetDateJoinedOk() (*time.Time, bool)`

GetDateJoinedOk returns a tuple with the DateJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateJoined

`func (o *UserDetail) SetDateJoined(v time.Time)`

SetDateJoined sets DateJoined field to given value.

### HasDateJoined

`func (o *UserDetail) HasDateJoined() bool`

HasDateJoined returns a boolean if a field has been set.

### GetDisplay

`func (o *UserDetail) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *UserDetail) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *UserDetail) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *UserDetail) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetDingUserId

`func (o *UserDetail) GetDingUserId() string`

GetDingUserId returns the DingUserId field if non-nil, zero value otherwise.

### GetDingUserIdOk

`func (o *UserDetail) GetDingUserIdOk() (*string, bool)`

GetDingUserIdOk returns a tuple with the DingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDingUserId

`func (o *UserDetail) SetDingUserId(v string)`

SetDingUserId sets DingUserId field to given value.

### HasDingUserId

`func (o *UserDetail) HasDingUserId() bool`

HasDingUserId returns a boolean if a field has been set.

### GetWxUserId

`func (o *UserDetail) GetWxUserId() string`

GetWxUserId returns the WxUserId field if non-nil, zero value otherwise.

### GetWxUserIdOk

`func (o *UserDetail) GetWxUserIdOk() (*string, bool)`

GetWxUserIdOk returns a tuple with the WxUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWxUserId

`func (o *UserDetail) SetWxUserId(v string)`

SetWxUserId sets WxUserId field to given value.

### HasWxUserId

`func (o *UserDetail) HasWxUserId() bool`

HasWxUserId returns a boolean if a field has been set.

### GetFeishuOpenId

`func (o *UserDetail) GetFeishuOpenId() string`

GetFeishuOpenId returns the FeishuOpenId field if non-nil, zero value otherwise.

### GetFeishuOpenIdOk

`func (o *UserDetail) GetFeishuOpenIdOk() (*string, bool)`

GetFeishuOpenIdOk returns a tuple with the FeishuOpenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeishuOpenId

`func (o *UserDetail) SetFeishuOpenId(v string)`

SetFeishuOpenId sets FeishuOpenId field to given value.

### HasFeishuOpenId

`func (o *UserDetail) HasFeishuOpenId() bool`

HasFeishuOpenId returns a boolean if a field has been set.

### GetFailedLoginCount

`func (o *UserDetail) GetFailedLoginCount() int32`

GetFailedLoginCount returns the FailedLoginCount field if non-nil, zero value otherwise.

### GetFailedLoginCountOk

`func (o *UserDetail) GetFailedLoginCountOk() (*int32, bool)`

GetFailedLoginCountOk returns a tuple with the FailedLoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedLoginCount

`func (o *UserDetail) SetFailedLoginCount(v int32)`

SetFailedLoginCount sets FailedLoginCount field to given value.

### HasFailedLoginCount

`func (o *UserDetail) HasFailedLoginCount() bool`

HasFailedLoginCount returns a boolean if a field has been set.

### GetLastLoginFailedAt

`func (o *UserDetail) GetLastLoginFailedAt() time.Time`

GetLastLoginFailedAt returns the LastLoginFailedAt field if non-nil, zero value otherwise.

### GetLastLoginFailedAtOk

`func (o *UserDetail) GetLastLoginFailedAtOk() (*time.Time, bool)`

GetLastLoginFailedAtOk returns a tuple with the LastLoginFailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginFailedAt

`func (o *UserDetail) SetLastLoginFailedAt(v time.Time)`

SetLastLoginFailedAt sets LastLoginFailedAt field to given value.

### HasLastLoginFailedAt

`func (o *UserDetail) HasLastLoginFailedAt() bool`

HasLastLoginFailedAt returns a boolean if a field has been set.

### SetLastLoginFailedAtNil

`func (o *UserDetail) SetLastLoginFailedAtNil(b bool)`

 SetLastLoginFailedAtNil sets the value for LastLoginFailedAt to be an explicit nil

### UnsetLastLoginFailedAt
`func (o *UserDetail) UnsetLastLoginFailedAt()`

UnsetLastLoginFailedAt ensures that no value is present for LastLoginFailedAt, not even an explicit nil
### GetGroups

`func (o *UserDetail) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserDetail) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserDetail) SetGroups(v []int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserDetail) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUserPermissions

`func (o *UserDetail) GetUserPermissions() []int32`

GetUserPermissions returns the UserPermissions field if non-nil, zero value otherwise.

### GetUserPermissionsOk

`func (o *UserDetail) GetUserPermissionsOk() (*[]int32, bool)`

GetUserPermissionsOk returns a tuple with the UserPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPermissions

`func (o *UserDetail) SetUserPermissions(v []int32)`

SetUserPermissions sets UserPermissions field to given value.

### HasUserPermissions

`func (o *UserDetail) HasUserPermissions() bool`

HasUserPermissions returns a boolean if a field has been set.

### GetResourceGroup

`func (o *UserDetail) GetResourceGroup() []int32`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *UserDetail) GetResourceGroupOk() (*[]int32, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *UserDetail) SetResourceGroup(v []int32)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *UserDetail) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


