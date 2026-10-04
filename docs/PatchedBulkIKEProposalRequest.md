# PatchedBulkIKEProposalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AuthenticationMethod** | Pointer to [**BulkIKEProposalRequestAuthenticationMethod**](BulkIKEProposalRequestAuthenticationMethod.md) |  | [optional] 
**EncryptionAlgorithm** | Pointer to [**BulkIKEProposalRequestEncryptionAlgorithm**](BulkIKEProposalRequestEncryptionAlgorithm.md) |  | [optional] 
**AuthenticationAlgorithm** | Pointer to [**BulkIKEProposalRequestAuthenticationAlgorithm**](BulkIKEProposalRequestAuthenticationAlgorithm.md) |  | [optional] 
**Group** | Pointer to [**BulkIKEProposalRequestGroup**](BulkIKEProposalRequestGroup.md) |  | [optional] 
**SaLifetime** | Pointer to **NullableInt32** | Security association lifetime (in seconds) | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkIKEProposalRequest

`func NewPatchedBulkIKEProposalRequest(id int32, ) *PatchedBulkIKEProposalRequest`

NewPatchedBulkIKEProposalRequest instantiates a new PatchedBulkIKEProposalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkIKEProposalRequestWithDefaults

`func NewPatchedBulkIKEProposalRequestWithDefaults() *PatchedBulkIKEProposalRequest`

NewPatchedBulkIKEProposalRequestWithDefaults instantiates a new PatchedBulkIKEProposalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkIKEProposalRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkIKEProposalRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkIKEProposalRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkIKEProposalRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkIKEProposalRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkIKEProposalRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkIKEProposalRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkIKEProposalRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkIKEProposalRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkIKEProposalRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkIKEProposalRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *PatchedBulkIKEProposalRequest) GetAuthenticationMethod() BulkIKEProposalRequestAuthenticationMethod`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *PatchedBulkIKEProposalRequest) GetAuthenticationMethodOk() (*BulkIKEProposalRequestAuthenticationMethod, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *PatchedBulkIKEProposalRequest) SetAuthenticationMethod(v BulkIKEProposalRequestAuthenticationMethod)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *PatchedBulkIKEProposalRequest) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### GetEncryptionAlgorithm

`func (o *PatchedBulkIKEProposalRequest) GetEncryptionAlgorithm() BulkIKEProposalRequestEncryptionAlgorithm`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *PatchedBulkIKEProposalRequest) GetEncryptionAlgorithmOk() (*BulkIKEProposalRequestEncryptionAlgorithm, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *PatchedBulkIKEProposalRequest) SetEncryptionAlgorithm(v BulkIKEProposalRequestEncryptionAlgorithm)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *PatchedBulkIKEProposalRequest) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### GetAuthenticationAlgorithm

`func (o *PatchedBulkIKEProposalRequest) GetAuthenticationAlgorithm() BulkIKEProposalRequestAuthenticationAlgorithm`

GetAuthenticationAlgorithm returns the AuthenticationAlgorithm field if non-nil, zero value otherwise.

### GetAuthenticationAlgorithmOk

`func (o *PatchedBulkIKEProposalRequest) GetAuthenticationAlgorithmOk() (*BulkIKEProposalRequestAuthenticationAlgorithm, bool)`

GetAuthenticationAlgorithmOk returns a tuple with the AuthenticationAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAlgorithm

`func (o *PatchedBulkIKEProposalRequest) SetAuthenticationAlgorithm(v BulkIKEProposalRequestAuthenticationAlgorithm)`

SetAuthenticationAlgorithm sets AuthenticationAlgorithm field to given value.

### HasAuthenticationAlgorithm

`func (o *PatchedBulkIKEProposalRequest) HasAuthenticationAlgorithm() bool`

HasAuthenticationAlgorithm returns a boolean if a field has been set.

### GetGroup

`func (o *PatchedBulkIKEProposalRequest) GetGroup() BulkIKEProposalRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedBulkIKEProposalRequest) GetGroupOk() (*BulkIKEProposalRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedBulkIKEProposalRequest) SetGroup(v BulkIKEProposalRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedBulkIKEProposalRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetSaLifetime

`func (o *PatchedBulkIKEProposalRequest) GetSaLifetime() int32`

GetSaLifetime returns the SaLifetime field if non-nil, zero value otherwise.

### GetSaLifetimeOk

`func (o *PatchedBulkIKEProposalRequest) GetSaLifetimeOk() (*int32, bool)`

GetSaLifetimeOk returns a tuple with the SaLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaLifetime

`func (o *PatchedBulkIKEProposalRequest) SetSaLifetime(v int32)`

SetSaLifetime sets SaLifetime field to given value.

### HasSaLifetime

`func (o *PatchedBulkIKEProposalRequest) HasSaLifetime() bool`

HasSaLifetime returns a boolean if a field has been set.

### SetSaLifetimeNil

`func (o *PatchedBulkIKEProposalRequest) SetSaLifetimeNil(b bool)`

 SetSaLifetimeNil sets the value for SaLifetime to be an explicit nil

### UnsetSaLifetime
`func (o *PatchedBulkIKEProposalRequest) UnsetSaLifetime()`

UnsetSaLifetime ensures that no value is present for SaLifetime, not even an explicit nil
### GetOwner

`func (o *PatchedBulkIKEProposalRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkIKEProposalRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkIKEProposalRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkIKEProposalRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkIKEProposalRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkIKEProposalRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkIKEProposalRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkIKEProposalRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkIKEProposalRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkIKEProposalRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkIKEProposalRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkIKEProposalRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkIKEProposalRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkIKEProposalRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkIKEProposalRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkIKEProposalRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkIKEProposalRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkIKEProposalRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


