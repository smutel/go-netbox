# BulkIPSecPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Proposals** | Pointer to **[]int32** |  | [optional] 
**PfsGroup** | Pointer to [**BulkIKEProposalRequestGroup**](BulkIKEProposalRequestGroup.md) |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkIPSecPolicyRequest

`func NewBulkIPSecPolicyRequest(id int32, name string, ) *BulkIPSecPolicyRequest`

NewBulkIPSecPolicyRequest instantiates a new BulkIPSecPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkIPSecPolicyRequestWithDefaults

`func NewBulkIPSecPolicyRequestWithDefaults() *BulkIPSecPolicyRequest`

NewBulkIPSecPolicyRequestWithDefaults instantiates a new BulkIPSecPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkIPSecPolicyRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkIPSecPolicyRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkIPSecPolicyRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BulkIPSecPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkIPSecPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkIPSecPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkIPSecPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkIPSecPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkIPSecPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkIPSecPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProposals

`func (o *BulkIPSecPolicyRequest) GetProposals() []int32`

GetProposals returns the Proposals field if non-nil, zero value otherwise.

### GetProposalsOk

`func (o *BulkIPSecPolicyRequest) GetProposalsOk() (*[]int32, bool)`

GetProposalsOk returns a tuple with the Proposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposals

`func (o *BulkIPSecPolicyRequest) SetProposals(v []int32)`

SetProposals sets Proposals field to given value.

### HasProposals

`func (o *BulkIPSecPolicyRequest) HasProposals() bool`

HasProposals returns a boolean if a field has been set.

### GetPfsGroup

`func (o *BulkIPSecPolicyRequest) GetPfsGroup() BulkIKEProposalRequestGroup`

GetPfsGroup returns the PfsGroup field if non-nil, zero value otherwise.

### GetPfsGroupOk

`func (o *BulkIPSecPolicyRequest) GetPfsGroupOk() (*BulkIKEProposalRequestGroup, bool)`

GetPfsGroupOk returns a tuple with the PfsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfsGroup

`func (o *BulkIPSecPolicyRequest) SetPfsGroup(v BulkIKEProposalRequestGroup)`

SetPfsGroup sets PfsGroup field to given value.

### HasPfsGroup

`func (o *BulkIPSecPolicyRequest) HasPfsGroup() bool`

HasPfsGroup returns a boolean if a field has been set.

### GetOwner

`func (o *BulkIPSecPolicyRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkIPSecPolicyRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkIPSecPolicyRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkIPSecPolicyRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkIPSecPolicyRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkIPSecPolicyRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkIPSecPolicyRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkIPSecPolicyRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkIPSecPolicyRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkIPSecPolicyRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkIPSecPolicyRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkIPSecPolicyRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkIPSecPolicyRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkIPSecPolicyRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkIPSecPolicyRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkIPSecPolicyRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkIPSecPolicyRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkIPSecPolicyRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


