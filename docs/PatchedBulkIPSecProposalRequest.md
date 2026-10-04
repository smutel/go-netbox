# PatchedBulkIPSecProposalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EncryptionAlgorithm** | Pointer to [**BulkIKEProposalRequestEncryptionAlgorithm**](BulkIKEProposalRequestEncryptionAlgorithm.md) |  | [optional] 
**AuthenticationAlgorithm** | Pointer to [**BulkIKEProposalRequestAuthenticationAlgorithm**](BulkIKEProposalRequestAuthenticationAlgorithm.md) |  | [optional] 
**SaLifetimeSeconds** | Pointer to **NullableInt32** | Security association lifetime (seconds) | [optional] 
**SaLifetimeData** | Pointer to **NullableInt32** | Security association lifetime (in kilobytes) | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkIPSecProposalRequest

`func NewPatchedBulkIPSecProposalRequest(id int32, ) *PatchedBulkIPSecProposalRequest`

NewPatchedBulkIPSecProposalRequest instantiates a new PatchedBulkIPSecProposalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkIPSecProposalRequestWithDefaults

`func NewPatchedBulkIPSecProposalRequestWithDefaults() *PatchedBulkIPSecProposalRequest`

NewPatchedBulkIPSecProposalRequestWithDefaults instantiates a new PatchedBulkIPSecProposalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkIPSecProposalRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkIPSecProposalRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkIPSecProposalRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkIPSecProposalRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkIPSecProposalRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkIPSecProposalRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkIPSecProposalRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkIPSecProposalRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkIPSecProposalRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkIPSecProposalRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkIPSecProposalRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEncryptionAlgorithm

`func (o *PatchedBulkIPSecProposalRequest) GetEncryptionAlgorithm() BulkIKEProposalRequestEncryptionAlgorithm`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *PatchedBulkIPSecProposalRequest) GetEncryptionAlgorithmOk() (*BulkIKEProposalRequestEncryptionAlgorithm, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *PatchedBulkIPSecProposalRequest) SetEncryptionAlgorithm(v BulkIKEProposalRequestEncryptionAlgorithm)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *PatchedBulkIPSecProposalRequest) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### GetAuthenticationAlgorithm

`func (o *PatchedBulkIPSecProposalRequest) GetAuthenticationAlgorithm() BulkIKEProposalRequestAuthenticationAlgorithm`

GetAuthenticationAlgorithm returns the AuthenticationAlgorithm field if non-nil, zero value otherwise.

### GetAuthenticationAlgorithmOk

`func (o *PatchedBulkIPSecProposalRequest) GetAuthenticationAlgorithmOk() (*BulkIKEProposalRequestAuthenticationAlgorithm, bool)`

GetAuthenticationAlgorithmOk returns a tuple with the AuthenticationAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAlgorithm

`func (o *PatchedBulkIPSecProposalRequest) SetAuthenticationAlgorithm(v BulkIKEProposalRequestAuthenticationAlgorithm)`

SetAuthenticationAlgorithm sets AuthenticationAlgorithm field to given value.

### HasAuthenticationAlgorithm

`func (o *PatchedBulkIPSecProposalRequest) HasAuthenticationAlgorithm() bool`

HasAuthenticationAlgorithm returns a boolean if a field has been set.

### GetSaLifetimeSeconds

`func (o *PatchedBulkIPSecProposalRequest) GetSaLifetimeSeconds() int32`

GetSaLifetimeSeconds returns the SaLifetimeSeconds field if non-nil, zero value otherwise.

### GetSaLifetimeSecondsOk

`func (o *PatchedBulkIPSecProposalRequest) GetSaLifetimeSecondsOk() (*int32, bool)`

GetSaLifetimeSecondsOk returns a tuple with the SaLifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaLifetimeSeconds

`func (o *PatchedBulkIPSecProposalRequest) SetSaLifetimeSeconds(v int32)`

SetSaLifetimeSeconds sets SaLifetimeSeconds field to given value.

### HasSaLifetimeSeconds

`func (o *PatchedBulkIPSecProposalRequest) HasSaLifetimeSeconds() bool`

HasSaLifetimeSeconds returns a boolean if a field has been set.

### SetSaLifetimeSecondsNil

`func (o *PatchedBulkIPSecProposalRequest) SetSaLifetimeSecondsNil(b bool)`

 SetSaLifetimeSecondsNil sets the value for SaLifetimeSeconds to be an explicit nil

### UnsetSaLifetimeSeconds
`func (o *PatchedBulkIPSecProposalRequest) UnsetSaLifetimeSeconds()`

UnsetSaLifetimeSeconds ensures that no value is present for SaLifetimeSeconds, not even an explicit nil
### GetSaLifetimeData

`func (o *PatchedBulkIPSecProposalRequest) GetSaLifetimeData() int32`

GetSaLifetimeData returns the SaLifetimeData field if non-nil, zero value otherwise.

### GetSaLifetimeDataOk

`func (o *PatchedBulkIPSecProposalRequest) GetSaLifetimeDataOk() (*int32, bool)`

GetSaLifetimeDataOk returns a tuple with the SaLifetimeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaLifetimeData

`func (o *PatchedBulkIPSecProposalRequest) SetSaLifetimeData(v int32)`

SetSaLifetimeData sets SaLifetimeData field to given value.

### HasSaLifetimeData

`func (o *PatchedBulkIPSecProposalRequest) HasSaLifetimeData() bool`

HasSaLifetimeData returns a boolean if a field has been set.

### SetSaLifetimeDataNil

`func (o *PatchedBulkIPSecProposalRequest) SetSaLifetimeDataNil(b bool)`

 SetSaLifetimeDataNil sets the value for SaLifetimeData to be an explicit nil

### UnsetSaLifetimeData
`func (o *PatchedBulkIPSecProposalRequest) UnsetSaLifetimeData()`

UnsetSaLifetimeData ensures that no value is present for SaLifetimeData, not even an explicit nil
### GetOwner

`func (o *PatchedBulkIPSecProposalRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkIPSecProposalRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkIPSecProposalRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkIPSecProposalRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkIPSecProposalRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkIPSecProposalRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkIPSecProposalRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkIPSecProposalRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkIPSecProposalRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkIPSecProposalRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkIPSecProposalRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkIPSecProposalRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkIPSecProposalRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkIPSecProposalRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkIPSecProposalRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkIPSecProposalRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkIPSecProposalRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkIPSecProposalRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


