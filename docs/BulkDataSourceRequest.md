# BulkDataSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Type** | [**BulkDataSourceRequestType**](BulkDataSourceRequestType.md) |  | 
**SourceUrl** | **string** |  | 
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

### NewBulkDataSourceRequest

`func NewBulkDataSourceRequest(id int32, name string, type_ BulkDataSourceRequestType, sourceUrl string, ) *BulkDataSourceRequest`

NewBulkDataSourceRequest instantiates a new BulkDataSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDataSourceRequestWithDefaults

`func NewBulkDataSourceRequestWithDefaults() *BulkDataSourceRequest`

NewBulkDataSourceRequestWithDefaults instantiates a new BulkDataSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkDataSourceRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkDataSourceRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkDataSourceRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BulkDataSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkDataSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkDataSourceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *BulkDataSourceRequest) GetType() BulkDataSourceRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkDataSourceRequest) GetTypeOk() (*BulkDataSourceRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkDataSourceRequest) SetType(v BulkDataSourceRequestType)`

SetType sets Type field to given value.


### GetSourceUrl

`func (o *BulkDataSourceRequest) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *BulkDataSourceRequest) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *BulkDataSourceRequest) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.


### GetEnabled

`func (o *BulkDataSourceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BulkDataSourceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BulkDataSourceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BulkDataSourceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *BulkDataSourceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkDataSourceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkDataSourceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkDataSourceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSyncInterval

`func (o *BulkDataSourceRequest) GetSyncInterval() BulkDataSourceRequestSyncInterval`

GetSyncInterval returns the SyncInterval field if non-nil, zero value otherwise.

### GetSyncIntervalOk

`func (o *BulkDataSourceRequest) GetSyncIntervalOk() (*BulkDataSourceRequestSyncInterval, bool)`

GetSyncIntervalOk returns a tuple with the SyncInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncInterval

`func (o *BulkDataSourceRequest) SetSyncInterval(v BulkDataSourceRequestSyncInterval)`

SetSyncInterval sets SyncInterval field to given value.

### HasSyncInterval

`func (o *BulkDataSourceRequest) HasSyncInterval() bool`

HasSyncInterval returns a boolean if a field has been set.

### SetSyncIntervalNil

`func (o *BulkDataSourceRequest) SetSyncIntervalNil(b bool)`

 SetSyncIntervalNil sets the value for SyncInterval to be an explicit nil

### UnsetSyncInterval
`func (o *BulkDataSourceRequest) UnsetSyncInterval()`

UnsetSyncInterval ensures that no value is present for SyncInterval, not even an explicit nil
### GetParameters

`func (o *BulkDataSourceRequest) GetParameters() interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *BulkDataSourceRequest) GetParametersOk() (*interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *BulkDataSourceRequest) SetParameters(v interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *BulkDataSourceRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *BulkDataSourceRequest) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *BulkDataSourceRequest) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetIgnoreRules

`func (o *BulkDataSourceRequest) GetIgnoreRules() string`

GetIgnoreRules returns the IgnoreRules field if non-nil, zero value otherwise.

### GetIgnoreRulesOk

`func (o *BulkDataSourceRequest) GetIgnoreRulesOk() (*string, bool)`

GetIgnoreRulesOk returns a tuple with the IgnoreRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreRules

`func (o *BulkDataSourceRequest) SetIgnoreRules(v string)`

SetIgnoreRules sets IgnoreRules field to given value.

### HasIgnoreRules

`func (o *BulkDataSourceRequest) HasIgnoreRules() bool`

HasIgnoreRules returns a boolean if a field has been set.

### GetOwner

`func (o *BulkDataSourceRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkDataSourceRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkDataSourceRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkDataSourceRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkDataSourceRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkDataSourceRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkDataSourceRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkDataSourceRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkDataSourceRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkDataSourceRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkDataSourceRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkDataSourceRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkDataSourceRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkDataSourceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkDataSourceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkDataSourceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkDataSourceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkDataSourceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


