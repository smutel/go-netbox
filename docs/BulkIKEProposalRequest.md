# BulkIKEProposalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**AuthenticationMethod** | [**BulkIKEProposalRequestAuthenticationMethod**](BulkIKEProposalRequestAuthenticationMethod.md) |  | 
**EncryptionAlgorithm** | [**BulkIKEProposalRequestEncryptionAlgorithm**](BulkIKEProposalRequestEncryptionAlgorithm.md) |  | 
**AuthenticationAlgorithm** | Pointer to [**BulkIKEProposalRequestAuthenticationAlgorithm**](BulkIKEProposalRequestAuthenticationAlgorithm.md) |  | [optional] 
**Group** | [**BulkIKEProposalRequestGroup**](BulkIKEProposalRequestGroup.md) |  | 
**SaLifetime** | Pointer to **NullableInt32** | Security association lifetime (in seconds) | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkIKEProposalRequest

`func NewBulkIKEProposalRequest(id int32, name string, authenticationMethod BulkIKEProposalRequestAuthenticationMethod, encryptionAlgorithm BulkIKEProposalRequestEncryptionAlgorithm, group BulkIKEProposalRequestGroup, ) *BulkIKEProposalRequest`

NewBulkIKEProposalRequest instantiates a new BulkIKEProposalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkIKEProposalRequestWithDefaults

`func NewBulkIKEProposalRequestWithDefaults() *BulkIKEProposalRequest`

NewBulkIKEProposalRequestWithDefaults instantiates a new BulkIKEProposalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkIKEProposalRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkIKEProposalRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkIKEProposalRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BulkIKEProposalRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkIKEProposalRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkIKEProposalRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkIKEProposalRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkIKEProposalRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkIKEProposalRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkIKEProposalRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *BulkIKEProposalRequest) GetAuthenticationMethod() BulkIKEProposalRequestAuthenticationMethod`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *BulkIKEProposalRequest) GetAuthenticationMethodOk() (*BulkIKEProposalRequestAuthenticationMethod, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *BulkIKEProposalRequest) SetAuthenticationMethod(v BulkIKEProposalRequestAuthenticationMethod)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.


### GetEncryptionAlgorithm

`func (o *BulkIKEProposalRequest) GetEncryptionAlgorithm() BulkIKEProposalRequestEncryptionAlgorithm`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *BulkIKEProposalRequest) GetEncryptionAlgorithmOk() (*BulkIKEProposalRequestEncryptionAlgorithm, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *BulkIKEProposalRequest) SetEncryptionAlgorithm(v BulkIKEProposalRequestEncryptionAlgorithm)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.


### GetAuthenticationAlgorithm

`func (o *BulkIKEProposalRequest) GetAuthenticationAlgorithm() BulkIKEProposalRequestAuthenticationAlgorithm`

GetAuthenticationAlgorithm returns the AuthenticationAlgorithm field if non-nil, zero value otherwise.

### GetAuthenticationAlgorithmOk

`func (o *BulkIKEProposalRequest) GetAuthenticationAlgorithmOk() (*BulkIKEProposalRequestAuthenticationAlgorithm, bool)`

GetAuthenticationAlgorithmOk returns a tuple with the AuthenticationAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAlgorithm

`func (o *BulkIKEProposalRequest) SetAuthenticationAlgorithm(v BulkIKEProposalRequestAuthenticationAlgorithm)`

SetAuthenticationAlgorithm sets AuthenticationAlgorithm field to given value.

### HasAuthenticationAlgorithm

`func (o *BulkIKEProposalRequest) HasAuthenticationAlgorithm() bool`

HasAuthenticationAlgorithm returns a boolean if a field has been set.

### GetGroup

`func (o *BulkIKEProposalRequest) GetGroup() BulkIKEProposalRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BulkIKEProposalRequest) GetGroupOk() (*BulkIKEProposalRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BulkIKEProposalRequest) SetGroup(v BulkIKEProposalRequestGroup)`

SetGroup sets Group field to given value.


### GetSaLifetime

`func (o *BulkIKEProposalRequest) GetSaLifetime() int32`

GetSaLifetime returns the SaLifetime field if non-nil, zero value otherwise.

### GetSaLifetimeOk

`func (o *BulkIKEProposalRequest) GetSaLifetimeOk() (*int32, bool)`

GetSaLifetimeOk returns a tuple with the SaLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaLifetime

`func (o *BulkIKEProposalRequest) SetSaLifetime(v int32)`

SetSaLifetime sets SaLifetime field to given value.

### HasSaLifetime

`func (o *BulkIKEProposalRequest) HasSaLifetime() bool`

HasSaLifetime returns a boolean if a field has been set.

### SetSaLifetimeNil

`func (o *BulkIKEProposalRequest) SetSaLifetimeNil(b bool)`

 SetSaLifetimeNil sets the value for SaLifetime to be an explicit nil

### UnsetSaLifetime
`func (o *BulkIKEProposalRequest) UnsetSaLifetime()`

UnsetSaLifetime ensures that no value is present for SaLifetime, not even an explicit nil
### GetOwner

`func (o *BulkIKEProposalRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkIKEProposalRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkIKEProposalRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkIKEProposalRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkIKEProposalRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkIKEProposalRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkIKEProposalRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkIKEProposalRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkIKEProposalRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkIKEProposalRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkIKEProposalRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkIKEProposalRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkIKEProposalRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkIKEProposalRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkIKEProposalRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkIKEProposalRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkIKEProposalRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkIKEProposalRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


