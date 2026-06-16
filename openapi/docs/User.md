# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Password** | **string** |  | 
**LastLogin** | Pointer to **NullableTime** |  | [optional] 
**IsSuperuser** | Pointer to **bool** | 指明该用户缺省拥有所有权限。 | [optional] 
**Username** | **string** | 必填；长度为150个字符或以下；只能包含字母、数字、特殊字符“@”、“.”、“-”和“_”。 | 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**IsStaff** | Pointer to **bool** | 指明用户是否可以登录到这个管理站点。 | [optional] 
**IsActive** | Pointer to **bool** | 指明用户是否被认为是活跃的。以反选代替删除帐号。 | [optional] 
**DateJoined** | Pointer to **time.Time** |  | [optional] 
**Display** | **string** |  | 
**DingUserId** | Pointer to **string** |  | [optional] 
**WxUserId** | Pointer to **string** |  | [optional] 
**FeishuOpenId** | Pointer to **string** |  | [optional] 
**FailedLoginCount** | Pointer to **int32** |  | [optional] 
**LastLoginFailedAt** | Pointer to **NullableTime** |  | [optional] 
**Groups** | Pointer to **[]int32** | 该用户归属的组。一个用户将得到其归属的组的所有权限。 | [optional] 
**UserPermissions** | Pointer to **[]int32** | 这个用户的特定权限。 | [optional] 
**ResourceGroup** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewUser

`func NewUser(id int32, password string, username string, display string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId sets Id field to given value.


### GetPassword

`func (o *User) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *User) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *User) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetLastLogin

`func (o *User) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *User) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *User) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *User) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *User) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *User) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetIsSuperuser

`func (o *User) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *User) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *User) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.

### HasIsSuperuser

`func (o *User) HasIsSuperuser() bool`

HasIsSuperuser returns a boolean if a field has been set.

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *User) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *User) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIsStaff

`func (o *User) GetIsStaff() bool`

GetIsStaff returns the IsStaff field if non-nil, zero value otherwise.

### GetIsStaffOk

`func (o *User) GetIsStaffOk() (*bool, bool)`

GetIsStaffOk returns a tuple with the IsStaff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStaff

`func (o *User) SetIsStaff(v bool)`

SetIsStaff sets IsStaff field to given value.

### HasIsStaff

`func (o *User) HasIsStaff() bool`

HasIsStaff returns a boolean if a field has been set.

### GetIsActive

`func (o *User) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *User) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *User) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *User) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetDateJoined

`func (o *User) GetDateJoined() time.Time`

GetDateJoined returns the DateJoined field if non-nil, zero value otherwise.

### GetDateJoinedOk

`func (o *User) GetDateJoinedOk() (*time.Time, bool)`

GetDateJoinedOk returns a tuple with the DateJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateJoined

`func (o *User) SetDateJoined(v time.Time)`

SetDateJoined sets DateJoined field to given value.

### HasDateJoined

`func (o *User) HasDateJoined() bool`

HasDateJoined returns a boolean if a field has been set.

### GetDisplay

`func (o *User) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *User) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *User) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetDingUserId

`func (o *User) GetDingUserId() string`

GetDingUserId returns the DingUserId field if non-nil, zero value otherwise.

### GetDingUserIdOk

`func (o *User) GetDingUserIdOk() (*string, bool)`

GetDingUserIdOk returns a tuple with the DingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDingUserId

`func (o *User) SetDingUserId(v string)`

SetDingUserId sets DingUserId field to given value.

### HasDingUserId

`func (o *User) HasDingUserId() bool`

HasDingUserId returns a boolean if a field has been set.

### GetWxUserId

`func (o *User) GetWxUserId() string`

GetWxUserId returns the WxUserId field if non-nil, zero value otherwise.

### GetWxUserIdOk

`func (o *User) GetWxUserIdOk() (*string, bool)`

GetWxUserIdOk returns a tuple with the WxUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWxUserId

`func (o *User) SetWxUserId(v string)`

SetWxUserId sets WxUserId field to given value.

### HasWxUserId

`func (o *User) HasWxUserId() bool`

HasWxUserId returns a boolean if a field has been set.

### GetFeishuOpenId

`func (o *User) GetFeishuOpenId() string`

GetFeishuOpenId returns the FeishuOpenId field if non-nil, zero value otherwise.

### GetFeishuOpenIdOk

`func (o *User) GetFeishuOpenIdOk() (*string, bool)`

GetFeishuOpenIdOk returns a tuple with the FeishuOpenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeishuOpenId

`func (o *User) SetFeishuOpenId(v string)`

SetFeishuOpenId sets FeishuOpenId field to given value.

### HasFeishuOpenId

`func (o *User) HasFeishuOpenId() bool`

HasFeishuOpenId returns a boolean if a field has been set.

### GetFailedLoginCount

`func (o *User) GetFailedLoginCount() int32`

GetFailedLoginCount returns the FailedLoginCount field if non-nil, zero value otherwise.

### GetFailedLoginCountOk

`func (o *User) GetFailedLoginCountOk() (*int32, bool)`

GetFailedLoginCountOk returns a tuple with the FailedLoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedLoginCount

`func (o *User) SetFailedLoginCount(v int32)`

SetFailedLoginCount sets FailedLoginCount field to given value.

### HasFailedLoginCount

`func (o *User) HasFailedLoginCount() bool`

HasFailedLoginCount returns a boolean if a field has been set.

### GetLastLoginFailedAt

`func (o *User) GetLastLoginFailedAt() time.Time`

GetLastLoginFailedAt returns the LastLoginFailedAt field if non-nil, zero value otherwise.

### GetLastLoginFailedAtOk

`func (o *User) GetLastLoginFailedAtOk() (*time.Time, bool)`

GetLastLoginFailedAtOk returns a tuple with the LastLoginFailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginFailedAt

`func (o *User) SetLastLoginFailedAt(v time.Time)`

SetLastLoginFailedAt sets LastLoginFailedAt field to given value.

### HasLastLoginFailedAt

`func (o *User) HasLastLoginFailedAt() bool`

HasLastLoginFailedAt returns a boolean if a field has been set.

### SetLastLoginFailedAtNil

`func (o *User) SetLastLoginFailedAtNil(b bool)`

 SetLastLoginFailedAtNil sets the value for LastLoginFailedAt to be an explicit nil

### UnsetLastLoginFailedAt
`func (o *User) UnsetLastLoginFailedAt()`

UnsetLastLoginFailedAt ensures that no value is present for LastLoginFailedAt, not even an explicit nil
### GetGroups

`func (o *User) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *User) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *User) SetGroups(v []int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *User) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUserPermissions

`func (o *User) GetUserPermissions() []int32`

GetUserPermissions returns the UserPermissions field if non-nil, zero value otherwise.

### GetUserPermissionsOk

`func (o *User) GetUserPermissionsOk() (*[]int32, bool)`

GetUserPermissionsOk returns a tuple with the UserPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPermissions

`func (o *User) SetUserPermissions(v []int32)`

SetUserPermissions sets UserPermissions field to given value.

### HasUserPermissions

`func (o *User) HasUserPermissions() bool`

HasUserPermissions returns a boolean if a field has been set.

### GetResourceGroup

`func (o *User) GetResourceGroup() []int32`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *User) GetResourceGroupOk() (*[]int32, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *User) SetResourceGroup(v []int32)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *User) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


