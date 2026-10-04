# BulkVirtualCircuitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Cid** | **string** | Unique circuit ID | 
**ProviderNetwork** | [**BriefVirtualCircuitRequestProviderNetwork**](BriefVirtualCircuitRequestProviderNetwork.md) |  | 
**ProviderAccount** | Pointer to [**NullableBulkCircuitRequestProviderAccount**](BulkCircuitRequestProviderAccount.md) |  | [optional] 
**Type** | [**BulkVirtualCircuitRequestType**](BulkVirtualCircuitRequestType.md) |  | 
**Status** | Pointer to [**BulkCircuitRequestStatus**](BulkCircuitRequestStatus.md) |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkVirtualCircuitRequest

`func NewBulkVirtualCircuitRequest(id int32, cid string, providerNetwork BriefVirtualCircuitRequestProviderNetwork, type_ BulkVirtualCircuitRequestType, ) *BulkVirtualCircuitRequest`

NewBulkVirtualCircuitRequest instantiates a new BulkVirtualCircuitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkVirtualCircuitRequestWithDefaults

`func NewBulkVirtualCircuitRequestWithDefaults() *BulkVirtualCircuitRequest`

NewBulkVirtualCircuitRequestWithDefaults instantiates a new BulkVirtualCircuitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkVirtualCircuitRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkVirtualCircuitRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkVirtualCircuitRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetCid

`func (o *BulkVirtualCircuitRequest) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *BulkVirtualCircuitRequest) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *BulkVirtualCircuitRequest) SetCid(v string)`

SetCid sets Cid field to given value.


### GetProviderNetwork

`func (o *BulkVirtualCircuitRequest) GetProviderNetwork() BriefVirtualCircuitRequestProviderNetwork`

GetProviderNetwork returns the ProviderNetwork field if non-nil, zero value otherwise.

### GetProviderNetworkOk

`func (o *BulkVirtualCircuitRequest) GetProviderNetworkOk() (*BriefVirtualCircuitRequestProviderNetwork, bool)`

GetProviderNetworkOk returns a tuple with the ProviderNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNetwork

`func (o *BulkVirtualCircuitRequest) SetProviderNetwork(v BriefVirtualCircuitRequestProviderNetwork)`

SetProviderNetwork sets ProviderNetwork field to given value.


### GetProviderAccount

`func (o *BulkVirtualCircuitRequest) GetProviderAccount() BulkCircuitRequestProviderAccount`

GetProviderAccount returns the ProviderAccount field if non-nil, zero value otherwise.

### GetProviderAccountOk

`func (o *BulkVirtualCircuitRequest) GetProviderAccountOk() (*BulkCircuitRequestProviderAccount, bool)`

GetProviderAccountOk returns a tuple with the ProviderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccount

`func (o *BulkVirtualCircuitRequest) SetProviderAccount(v BulkCircuitRequestProviderAccount)`

SetProviderAccount sets ProviderAccount field to given value.

### HasProviderAccount

`func (o *BulkVirtualCircuitRequest) HasProviderAccount() bool`

HasProviderAccount returns a boolean if a field has been set.

### SetProviderAccountNil

`func (o *BulkVirtualCircuitRequest) SetProviderAccountNil(b bool)`

 SetProviderAccountNil sets the value for ProviderAccount to be an explicit nil

### UnsetProviderAccount
`func (o *BulkVirtualCircuitRequest) UnsetProviderAccount()`

UnsetProviderAccount ensures that no value is present for ProviderAccount, not even an explicit nil
### GetType

`func (o *BulkVirtualCircuitRequest) GetType() BulkVirtualCircuitRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkVirtualCircuitRequest) GetTypeOk() (*BulkVirtualCircuitRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkVirtualCircuitRequest) SetType(v BulkVirtualCircuitRequestType)`

SetType sets Type field to given value.


### GetStatus

`func (o *BulkVirtualCircuitRequest) GetStatus() BulkCircuitRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkVirtualCircuitRequest) GetStatusOk() (*BulkCircuitRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkVirtualCircuitRequest) SetStatus(v BulkCircuitRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkVirtualCircuitRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *BulkVirtualCircuitRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkVirtualCircuitRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkVirtualCircuitRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkVirtualCircuitRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkVirtualCircuitRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkVirtualCircuitRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetDescription

`func (o *BulkVirtualCircuitRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkVirtualCircuitRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkVirtualCircuitRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkVirtualCircuitRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *BulkVirtualCircuitRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkVirtualCircuitRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkVirtualCircuitRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkVirtualCircuitRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkVirtualCircuitRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkVirtualCircuitRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkVirtualCircuitRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkVirtualCircuitRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkVirtualCircuitRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkVirtualCircuitRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkVirtualCircuitRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkVirtualCircuitRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkVirtualCircuitRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkVirtualCircuitRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkVirtualCircuitRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkVirtualCircuitRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkVirtualCircuitRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkVirtualCircuitRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


