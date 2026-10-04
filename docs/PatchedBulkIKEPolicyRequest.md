# PatchedBulkIKEPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Version** | Pointer to [**BulkIKEPolicyRequestVersion**](BulkIKEPolicyRequestVersion.md) |  | [optional] 
**Mode** | Pointer to [**BulkIKEPolicyRequestMode**](BulkIKEPolicyRequestMode.md) |  | [optional] 
**Proposals** | Pointer to **[]int32** |  | [optional] 
**PresharedKey** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkIKEPolicyRequest

`func NewPatchedBulkIKEPolicyRequest(id int32, ) *PatchedBulkIKEPolicyRequest`

NewPatchedBulkIKEPolicyRequest instantiates a new PatchedBulkIKEPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkIKEPolicyRequestWithDefaults

`func NewPatchedBulkIKEPolicyRequestWithDefaults() *PatchedBulkIKEPolicyRequest`

NewPatchedBulkIKEPolicyRequestWithDefaults instantiates a new PatchedBulkIKEPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkIKEPolicyRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkIKEPolicyRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkIKEPolicyRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkIKEPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkIKEPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkIKEPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkIKEPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkIKEPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkIKEPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkIKEPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkIKEPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *PatchedBulkIKEPolicyRequest) GetVersion() BulkIKEPolicyRequestVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PatchedBulkIKEPolicyRequest) GetVersionOk() (*BulkIKEPolicyRequestVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PatchedBulkIKEPolicyRequest) SetVersion(v BulkIKEPolicyRequestVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PatchedBulkIKEPolicyRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetMode

`func (o *PatchedBulkIKEPolicyRequest) GetMode() BulkIKEPolicyRequestMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedBulkIKEPolicyRequest) GetModeOk() (*BulkIKEPolicyRequestMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedBulkIKEPolicyRequest) SetMode(v BulkIKEPolicyRequestMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedBulkIKEPolicyRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetProposals

`func (o *PatchedBulkIKEPolicyRequest) GetProposals() []int32`

GetProposals returns the Proposals field if non-nil, zero value otherwise.

### GetProposalsOk

`func (o *PatchedBulkIKEPolicyRequest) GetProposalsOk() (*[]int32, bool)`

GetProposalsOk returns a tuple with the Proposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposals

`func (o *PatchedBulkIKEPolicyRequest) SetProposals(v []int32)`

SetProposals sets Proposals field to given value.

### HasProposals

`func (o *PatchedBulkIKEPolicyRequest) HasProposals() bool`

HasProposals returns a boolean if a field has been set.

### GetPresharedKey

`func (o *PatchedBulkIKEPolicyRequest) GetPresharedKey() string`

GetPresharedKey returns the PresharedKey field if non-nil, zero value otherwise.

### GetPresharedKeyOk

`func (o *PatchedBulkIKEPolicyRequest) GetPresharedKeyOk() (*string, bool)`

GetPresharedKeyOk returns a tuple with the PresharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresharedKey

`func (o *PatchedBulkIKEPolicyRequest) SetPresharedKey(v string)`

SetPresharedKey sets PresharedKey field to given value.

### HasPresharedKey

`func (o *PatchedBulkIKEPolicyRequest) HasPresharedKey() bool`

HasPresharedKey returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkIKEPolicyRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkIKEPolicyRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkIKEPolicyRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkIKEPolicyRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkIKEPolicyRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkIKEPolicyRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkIKEPolicyRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkIKEPolicyRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkIKEPolicyRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkIKEPolicyRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkIKEPolicyRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkIKEPolicyRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkIKEPolicyRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkIKEPolicyRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkIKEPolicyRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkIKEPolicyRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkIKEPolicyRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkIKEPolicyRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


