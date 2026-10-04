# BulkTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Version** | Pointer to [**BulkTokenRequestVersion**](BulkTokenRequestVersion.md) |  | [optional] 
**User** | [**BookmarkRequestUser**](BookmarkRequestUser.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **NullableTime** |  | [optional] 
**LastUsed** | Pointer to **NullableTime** |  | [optional] 
**Enabled** | Pointer to **bool** | Disable to temporarily revoke this token without deleting it. | [optional] 
**WriteEnabled** | Pointer to **bool** | Permit create/update/delete operations using this token | [optional] 
**PepperId** | Pointer to **NullableInt32** | ID of the cryptographic pepper used to hash the token (v2 only) | [optional] 

## Methods

### NewBulkTokenRequest

`func NewBulkTokenRequest(id int32, user BookmarkRequestUser, ) *BulkTokenRequest`

NewBulkTokenRequest instantiates a new BulkTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkTokenRequestWithDefaults

`func NewBulkTokenRequestWithDefaults() *BulkTokenRequest`

NewBulkTokenRequestWithDefaults instantiates a new BulkTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkTokenRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkTokenRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkTokenRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetVersion

`func (o *BulkTokenRequest) GetVersion() BulkTokenRequestVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BulkTokenRequest) GetVersionOk() (*BulkTokenRequestVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BulkTokenRequest) SetVersion(v BulkTokenRequestVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BulkTokenRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUser

`func (o *BulkTokenRequest) GetUser() BookmarkRequestUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BulkTokenRequest) GetUserOk() (*BookmarkRequestUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BulkTokenRequest) SetUser(v BookmarkRequestUser)`

SetUser sets User field to given value.


### GetDescription

`func (o *BulkTokenRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkTokenRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkTokenRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkTokenRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpires

`func (o *BulkTokenRequest) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *BulkTokenRequest) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *BulkTokenRequest) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *BulkTokenRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *BulkTokenRequest) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *BulkTokenRequest) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetLastUsed

`func (o *BulkTokenRequest) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *BulkTokenRequest) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *BulkTokenRequest) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *BulkTokenRequest) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### SetLastUsedNil

`func (o *BulkTokenRequest) SetLastUsedNil(b bool)`

 SetLastUsedNil sets the value for LastUsed to be an explicit nil

### UnsetLastUsed
`func (o *BulkTokenRequest) UnsetLastUsed()`

UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
### GetEnabled

`func (o *BulkTokenRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BulkTokenRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BulkTokenRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BulkTokenRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetWriteEnabled

`func (o *BulkTokenRequest) GetWriteEnabled() bool`

GetWriteEnabled returns the WriteEnabled field if non-nil, zero value otherwise.

### GetWriteEnabledOk

`func (o *BulkTokenRequest) GetWriteEnabledOk() (*bool, bool)`

GetWriteEnabledOk returns a tuple with the WriteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteEnabled

`func (o *BulkTokenRequest) SetWriteEnabled(v bool)`

SetWriteEnabled sets WriteEnabled field to given value.

### HasWriteEnabled

`func (o *BulkTokenRequest) HasWriteEnabled() bool`

HasWriteEnabled returns a boolean if a field has been set.

### GetPepperId

`func (o *BulkTokenRequest) GetPepperId() int32`

GetPepperId returns the PepperId field if non-nil, zero value otherwise.

### GetPepperIdOk

`func (o *BulkTokenRequest) GetPepperIdOk() (*int32, bool)`

GetPepperIdOk returns a tuple with the PepperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPepperId

`func (o *BulkTokenRequest) SetPepperId(v int32)`

SetPepperId sets PepperId field to given value.

### HasPepperId

`func (o *BulkTokenRequest) HasPepperId() bool`

HasPepperId returns a boolean if a field has been set.

### SetPepperIdNil

`func (o *BulkTokenRequest) SetPepperIdNil(b bool)`

 SetPepperIdNil sets the value for PepperId to be an explicit nil

### UnsetPepperId
`func (o *BulkTokenRequest) UnsetPepperId()`

UnsetPepperId ensures that no value is present for PepperId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


