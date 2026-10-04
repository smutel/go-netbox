# PatchedBulkIPSecPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Proposals** | Pointer to **[]int32** |  | [optional] 
**PfsGroup** | Pointer to [**BulkIKEProposalRequestGroup**](BulkIKEProposalRequestGroup.md) |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkIPSecPolicyRequest

`func NewPatchedBulkIPSecPolicyRequest(id int32, ) *PatchedBulkIPSecPolicyRequest`

NewPatchedBulkIPSecPolicyRequest instantiates a new PatchedBulkIPSecPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkIPSecPolicyRequestWithDefaults

`func NewPatchedBulkIPSecPolicyRequestWithDefaults() *PatchedBulkIPSecPolicyRequest`

NewPatchedBulkIPSecPolicyRequestWithDefaults instantiates a new PatchedBulkIPSecPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkIPSecPolicyRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkIPSecPolicyRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkIPSecPolicyRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkIPSecPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkIPSecPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkIPSecPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkIPSecPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkIPSecPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkIPSecPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkIPSecPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkIPSecPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProposals

`func (o *PatchedBulkIPSecPolicyRequest) GetProposals() []int32`

GetProposals returns the Proposals field if non-nil, zero value otherwise.

### GetProposalsOk

`func (o *PatchedBulkIPSecPolicyRequest) GetProposalsOk() (*[]int32, bool)`

GetProposalsOk returns a tuple with the Proposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposals

`func (o *PatchedBulkIPSecPolicyRequest) SetProposals(v []int32)`

SetProposals sets Proposals field to given value.

### HasProposals

`func (o *PatchedBulkIPSecPolicyRequest) HasProposals() bool`

HasProposals returns a boolean if a field has been set.

### GetPfsGroup

`func (o *PatchedBulkIPSecPolicyRequest) GetPfsGroup() BulkIKEProposalRequestGroup`

GetPfsGroup returns the PfsGroup field if non-nil, zero value otherwise.

### GetPfsGroupOk

`func (o *PatchedBulkIPSecPolicyRequest) GetPfsGroupOk() (*BulkIKEProposalRequestGroup, bool)`

GetPfsGroupOk returns a tuple with the PfsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfsGroup

`func (o *PatchedBulkIPSecPolicyRequest) SetPfsGroup(v BulkIKEProposalRequestGroup)`

SetPfsGroup sets PfsGroup field to given value.

### HasPfsGroup

`func (o *PatchedBulkIPSecPolicyRequest) HasPfsGroup() bool`

HasPfsGroup returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkIPSecPolicyRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkIPSecPolicyRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkIPSecPolicyRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkIPSecPolicyRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkIPSecPolicyRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkIPSecPolicyRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkIPSecPolicyRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkIPSecPolicyRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkIPSecPolicyRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkIPSecPolicyRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkIPSecPolicyRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkIPSecPolicyRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkIPSecPolicyRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkIPSecPolicyRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkIPSecPolicyRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkIPSecPolicyRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkIPSecPolicyRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkIPSecPolicyRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


