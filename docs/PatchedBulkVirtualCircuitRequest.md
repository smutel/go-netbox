# PatchedBulkVirtualCircuitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Cid** | Pointer to **string** | Unique circuit ID | [optional] 
**ProviderNetwork** | Pointer to [**BriefVirtualCircuitRequestProviderNetwork**](BriefVirtualCircuitRequestProviderNetwork.md) |  | [optional] 
**ProviderAccount** | Pointer to [**NullableBulkCircuitRequestProviderAccount**](BulkCircuitRequestProviderAccount.md) |  | [optional] 
**Type** | Pointer to [**BulkVirtualCircuitRequestType**](BulkVirtualCircuitRequestType.md) |  | [optional] 
**Status** | Pointer to [**BulkCircuitRequestStatus**](BulkCircuitRequestStatus.md) |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkVirtualCircuitRequest

`func NewPatchedBulkVirtualCircuitRequest(id int32, ) *PatchedBulkVirtualCircuitRequest`

NewPatchedBulkVirtualCircuitRequest instantiates a new PatchedBulkVirtualCircuitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkVirtualCircuitRequestWithDefaults

`func NewPatchedBulkVirtualCircuitRequestWithDefaults() *PatchedBulkVirtualCircuitRequest`

NewPatchedBulkVirtualCircuitRequestWithDefaults instantiates a new PatchedBulkVirtualCircuitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkVirtualCircuitRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkVirtualCircuitRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkVirtualCircuitRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetCid

`func (o *PatchedBulkVirtualCircuitRequest) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *PatchedBulkVirtualCircuitRequest) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *PatchedBulkVirtualCircuitRequest) SetCid(v string)`

SetCid sets Cid field to given value.

### HasCid

`func (o *PatchedBulkVirtualCircuitRequest) HasCid() bool`

HasCid returns a boolean if a field has been set.

### GetProviderNetwork

`func (o *PatchedBulkVirtualCircuitRequest) GetProviderNetwork() BriefVirtualCircuitRequestProviderNetwork`

GetProviderNetwork returns the ProviderNetwork field if non-nil, zero value otherwise.

### GetProviderNetworkOk

`func (o *PatchedBulkVirtualCircuitRequest) GetProviderNetworkOk() (*BriefVirtualCircuitRequestProviderNetwork, bool)`

GetProviderNetworkOk returns a tuple with the ProviderNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNetwork

`func (o *PatchedBulkVirtualCircuitRequest) SetProviderNetwork(v BriefVirtualCircuitRequestProviderNetwork)`

SetProviderNetwork sets ProviderNetwork field to given value.

### HasProviderNetwork

`func (o *PatchedBulkVirtualCircuitRequest) HasProviderNetwork() bool`

HasProviderNetwork returns a boolean if a field has been set.

### GetProviderAccount

`func (o *PatchedBulkVirtualCircuitRequest) GetProviderAccount() BulkCircuitRequestProviderAccount`

GetProviderAccount returns the ProviderAccount field if non-nil, zero value otherwise.

### GetProviderAccountOk

`func (o *PatchedBulkVirtualCircuitRequest) GetProviderAccountOk() (*BulkCircuitRequestProviderAccount, bool)`

GetProviderAccountOk returns a tuple with the ProviderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccount

`func (o *PatchedBulkVirtualCircuitRequest) SetProviderAccount(v BulkCircuitRequestProviderAccount)`

SetProviderAccount sets ProviderAccount field to given value.

### HasProviderAccount

`func (o *PatchedBulkVirtualCircuitRequest) HasProviderAccount() bool`

HasProviderAccount returns a boolean if a field has been set.

### SetProviderAccountNil

`func (o *PatchedBulkVirtualCircuitRequest) SetProviderAccountNil(b bool)`

 SetProviderAccountNil sets the value for ProviderAccount to be an explicit nil

### UnsetProviderAccount
`func (o *PatchedBulkVirtualCircuitRequest) UnsetProviderAccount()`

UnsetProviderAccount ensures that no value is present for ProviderAccount, not even an explicit nil
### GetType

`func (o *PatchedBulkVirtualCircuitRequest) GetType() BulkVirtualCircuitRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedBulkVirtualCircuitRequest) GetTypeOk() (*BulkVirtualCircuitRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedBulkVirtualCircuitRequest) SetType(v BulkVirtualCircuitRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedBulkVirtualCircuitRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedBulkVirtualCircuitRequest) GetStatus() BulkCircuitRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkVirtualCircuitRequest) GetStatusOk() (*BulkCircuitRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkVirtualCircuitRequest) SetStatus(v BulkCircuitRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkVirtualCircuitRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedBulkVirtualCircuitRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkVirtualCircuitRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkVirtualCircuitRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkVirtualCircuitRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkVirtualCircuitRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkVirtualCircuitRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetDescription

`func (o *PatchedBulkVirtualCircuitRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkVirtualCircuitRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkVirtualCircuitRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkVirtualCircuitRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkVirtualCircuitRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkVirtualCircuitRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkVirtualCircuitRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkVirtualCircuitRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkVirtualCircuitRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkVirtualCircuitRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkVirtualCircuitRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkVirtualCircuitRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkVirtualCircuitRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkVirtualCircuitRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkVirtualCircuitRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkVirtualCircuitRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkVirtualCircuitRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkVirtualCircuitRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkVirtualCircuitRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkVirtualCircuitRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkVirtualCircuitRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkVirtualCircuitRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


