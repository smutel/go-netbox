# PatchedBulkDataSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**BulkDataSourceRequestType**](BulkDataSourceRequestType.md) |  | [optional] 
**SourceUrl** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**SyncInterval** | Pointer to [**NullableBulkDataSourceRequestSyncInterval**](BulkDataSourceRequestSyncInterval.md) |  | [optional] 
**Parameters** | Pointer to **interface{}** |  | [optional] 
**IgnoreRules** | Pointer to **string** | Patterns (one per line) matching files or paths to ignore when syncing | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkDataSourceRequest

`func NewPatchedBulkDataSourceRequest(id int32, ) *PatchedBulkDataSourceRequest`

NewPatchedBulkDataSourceRequest instantiates a new PatchedBulkDataSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkDataSourceRequestWithDefaults

`func NewPatchedBulkDataSourceRequestWithDefaults() *PatchedBulkDataSourceRequest`

NewPatchedBulkDataSourceRequestWithDefaults instantiates a new PatchedBulkDataSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkDataSourceRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkDataSourceRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkDataSourceRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkDataSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkDataSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkDataSourceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkDataSourceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PatchedBulkDataSourceRequest) GetType() BulkDataSourceRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedBulkDataSourceRequest) GetTypeOk() (*BulkDataSourceRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedBulkDataSourceRequest) SetType(v BulkDataSourceRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedBulkDataSourceRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSourceUrl

`func (o *PatchedBulkDataSourceRequest) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *PatchedBulkDataSourceRequest) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *PatchedBulkDataSourceRequest) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *PatchedBulkDataSourceRequest) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedBulkDataSourceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedBulkDataSourceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedBulkDataSourceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedBulkDataSourceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkDataSourceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkDataSourceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkDataSourceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkDataSourceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSyncInterval

`func (o *PatchedBulkDataSourceRequest) GetSyncInterval() BulkDataSourceRequestSyncInterval`

GetSyncInterval returns the SyncInterval field if non-nil, zero value otherwise.

### GetSyncIntervalOk

`func (o *PatchedBulkDataSourceRequest) GetSyncIntervalOk() (*BulkDataSourceRequestSyncInterval, bool)`

GetSyncIntervalOk returns a tuple with the SyncInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncInterval

`func (o *PatchedBulkDataSourceRequest) SetSyncInterval(v BulkDataSourceRequestSyncInterval)`

SetSyncInterval sets SyncInterval field to given value.

### HasSyncInterval

`func (o *PatchedBulkDataSourceRequest) HasSyncInterval() bool`

HasSyncInterval returns a boolean if a field has been set.

### SetSyncIntervalNil

`func (o *PatchedBulkDataSourceRequest) SetSyncIntervalNil(b bool)`

 SetSyncIntervalNil sets the value for SyncInterval to be an explicit nil

### UnsetSyncInterval
`func (o *PatchedBulkDataSourceRequest) UnsetSyncInterval()`

UnsetSyncInterval ensures that no value is present for SyncInterval, not even an explicit nil
### GetParameters

`func (o *PatchedBulkDataSourceRequest) GetParameters() interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PatchedBulkDataSourceRequest) GetParametersOk() (*interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PatchedBulkDataSourceRequest) SetParameters(v interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *PatchedBulkDataSourceRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *PatchedBulkDataSourceRequest) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *PatchedBulkDataSourceRequest) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetIgnoreRules

`func (o *PatchedBulkDataSourceRequest) GetIgnoreRules() string`

GetIgnoreRules returns the IgnoreRules field if non-nil, zero value otherwise.

### GetIgnoreRulesOk

`func (o *PatchedBulkDataSourceRequest) GetIgnoreRulesOk() (*string, bool)`

GetIgnoreRulesOk returns a tuple with the IgnoreRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreRules

`func (o *PatchedBulkDataSourceRequest) SetIgnoreRules(v string)`

SetIgnoreRules sets IgnoreRules field to given value.

### HasIgnoreRules

`func (o *PatchedBulkDataSourceRequest) HasIgnoreRules() bool`

HasIgnoreRules returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkDataSourceRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkDataSourceRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkDataSourceRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkDataSourceRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkDataSourceRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkDataSourceRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkDataSourceRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkDataSourceRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkDataSourceRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkDataSourceRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkDataSourceRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkDataSourceRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkDataSourceRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkDataSourceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkDataSourceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkDataSourceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkDataSourceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkDataSourceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


