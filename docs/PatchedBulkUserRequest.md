# PatchedBulkUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Username** | Pointer to **string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to [**EmailAddress**](EmailAddress.md) |  | [optional] 
**IsActive** | Pointer to **bool** | Designates whether this user should be treated as active. Unselect this instead of deleting accounts. | [optional] 
**DateJoined** | Pointer to **time.Time** |  | [optional] 
**LastLogin** | Pointer to **NullableTime** |  | [optional] 
**Groups** | Pointer to **[]int32** |  | [optional] 
**Permissions** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewPatchedBulkUserRequest

`func NewPatchedBulkUserRequest(id int32, ) *PatchedBulkUserRequest`

NewPatchedBulkUserRequest instantiates a new PatchedBulkUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkUserRequestWithDefaults

`func NewPatchedBulkUserRequestWithDefaults() *PatchedBulkUserRequest`

NewPatchedBulkUserRequestWithDefaults instantiates a new PatchedBulkUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkUserRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkUserRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkUserRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetUsername

`func (o *PatchedBulkUserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatchedBulkUserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatchedBulkUserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PatchedBulkUserRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *PatchedBulkUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchedBulkUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchedBulkUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchedBulkUserRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetFirstName

`func (o *PatchedBulkUserRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PatchedBulkUserRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PatchedBulkUserRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PatchedBulkUserRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PatchedBulkUserRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PatchedBulkUserRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PatchedBulkUserRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PatchedBulkUserRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *PatchedBulkUserRequest) GetEmail() EmailAddress`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchedBulkUserRequest) GetEmailOk() (*EmailAddress, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchedBulkUserRequest) SetEmail(v EmailAddress)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatchedBulkUserRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIsActive

`func (o *PatchedBulkUserRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PatchedBulkUserRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PatchedBulkUserRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PatchedBulkUserRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetDateJoined

`func (o *PatchedBulkUserRequest) GetDateJoined() time.Time`

GetDateJoined returns the DateJoined field if non-nil, zero value otherwise.

### GetDateJoinedOk

`func (o *PatchedBulkUserRequest) GetDateJoinedOk() (*time.Time, bool)`

GetDateJoinedOk returns a tuple with the DateJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateJoined

`func (o *PatchedBulkUserRequest) SetDateJoined(v time.Time)`

SetDateJoined sets DateJoined field to given value.

### HasDateJoined

`func (o *PatchedBulkUserRequest) HasDateJoined() bool`

HasDateJoined returns a boolean if a field has been set.

### GetLastLogin

`func (o *PatchedBulkUserRequest) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *PatchedBulkUserRequest) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *PatchedBulkUserRequest) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *PatchedBulkUserRequest) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *PatchedBulkUserRequest) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *PatchedBulkUserRequest) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetGroups

`func (o *PatchedBulkUserRequest) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PatchedBulkUserRequest) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PatchedBulkUserRequest) SetGroups(v []int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *PatchedBulkUserRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetPermissions

`func (o *PatchedBulkUserRequest) GetPermissions() []int32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PatchedBulkUserRequest) GetPermissionsOk() (*[]int32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PatchedBulkUserRequest) SetPermissions(v []int32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PatchedBulkUserRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


