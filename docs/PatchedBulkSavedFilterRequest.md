# PatchedBulkSavedFilterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ObjectTypes** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**User** | Pointer to **NullableInt32** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Parameters** | Pointer to **interface{}** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 

## Methods

### NewPatchedBulkSavedFilterRequest

`func NewPatchedBulkSavedFilterRequest(id int32, ) *PatchedBulkSavedFilterRequest`

NewPatchedBulkSavedFilterRequest instantiates a new PatchedBulkSavedFilterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkSavedFilterRequestWithDefaults

`func NewPatchedBulkSavedFilterRequestWithDefaults() *PatchedBulkSavedFilterRequest`

NewPatchedBulkSavedFilterRequestWithDefaults instantiates a new PatchedBulkSavedFilterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkSavedFilterRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkSavedFilterRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkSavedFilterRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetObjectTypes

`func (o *PatchedBulkSavedFilterRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *PatchedBulkSavedFilterRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *PatchedBulkSavedFilterRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.

### HasObjectTypes

`func (o *PatchedBulkSavedFilterRequest) HasObjectTypes() bool`

HasObjectTypes returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkSavedFilterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkSavedFilterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkSavedFilterRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkSavedFilterRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedBulkSavedFilterRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedBulkSavedFilterRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedBulkSavedFilterRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedBulkSavedFilterRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkSavedFilterRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkSavedFilterRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkSavedFilterRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkSavedFilterRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUser

`func (o *PatchedBulkSavedFilterRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedBulkSavedFilterRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedBulkSavedFilterRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedBulkSavedFilterRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *PatchedBulkSavedFilterRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PatchedBulkSavedFilterRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetWeight

`func (o *PatchedBulkSavedFilterRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedBulkSavedFilterRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedBulkSavedFilterRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedBulkSavedFilterRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedBulkSavedFilterRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedBulkSavedFilterRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedBulkSavedFilterRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedBulkSavedFilterRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetShared

`func (o *PatchedBulkSavedFilterRequest) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *PatchedBulkSavedFilterRequest) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *PatchedBulkSavedFilterRequest) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *PatchedBulkSavedFilterRequest) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetParameters

`func (o *PatchedBulkSavedFilterRequest) GetParameters() interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PatchedBulkSavedFilterRequest) GetParametersOk() (*interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PatchedBulkSavedFilterRequest) SetParameters(v interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *PatchedBulkSavedFilterRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *PatchedBulkSavedFilterRequest) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *PatchedBulkSavedFilterRequest) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetOwner

`func (o *PatchedBulkSavedFilterRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkSavedFilterRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkSavedFilterRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkSavedFilterRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkSavedFilterRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkSavedFilterRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


