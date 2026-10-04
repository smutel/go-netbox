# PatchedBulkIPSecProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**BulkIPSecProfileRequestMode**](BulkIPSecProfileRequestMode.md) |  | [optional] 
**IkePolicy** | Pointer to [**BulkIPSecProfileRequestIkePolicy**](BulkIPSecProfileRequestIkePolicy.md) |  | [optional] 
**IpsecPolicy** | Pointer to [**BulkIPSecProfileRequestIpsecPolicy**](BulkIPSecProfileRequestIpsecPolicy.md) |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkIPSecProfileRequest

`func NewPatchedBulkIPSecProfileRequest(id int32, ) *PatchedBulkIPSecProfileRequest`

NewPatchedBulkIPSecProfileRequest instantiates a new PatchedBulkIPSecProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkIPSecProfileRequestWithDefaults

`func NewPatchedBulkIPSecProfileRequestWithDefaults() *PatchedBulkIPSecProfileRequest`

NewPatchedBulkIPSecProfileRequestWithDefaults instantiates a new PatchedBulkIPSecProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkIPSecProfileRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkIPSecProfileRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkIPSecProfileRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkIPSecProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkIPSecProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkIPSecProfileRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkIPSecProfileRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkIPSecProfileRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkIPSecProfileRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkIPSecProfileRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkIPSecProfileRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *PatchedBulkIPSecProfileRequest) GetMode() BulkIPSecProfileRequestMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedBulkIPSecProfileRequest) GetModeOk() (*BulkIPSecProfileRequestMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedBulkIPSecProfileRequest) SetMode(v BulkIPSecProfileRequestMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedBulkIPSecProfileRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetIkePolicy

`func (o *PatchedBulkIPSecProfileRequest) GetIkePolicy() BulkIPSecProfileRequestIkePolicy`

GetIkePolicy returns the IkePolicy field if non-nil, zero value otherwise.

### GetIkePolicyOk

`func (o *PatchedBulkIPSecProfileRequest) GetIkePolicyOk() (*BulkIPSecProfileRequestIkePolicy, bool)`

GetIkePolicyOk returns a tuple with the IkePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkePolicy

`func (o *PatchedBulkIPSecProfileRequest) SetIkePolicy(v BulkIPSecProfileRequestIkePolicy)`

SetIkePolicy sets IkePolicy field to given value.

### HasIkePolicy

`func (o *PatchedBulkIPSecProfileRequest) HasIkePolicy() bool`

HasIkePolicy returns a boolean if a field has been set.

### GetIpsecPolicy

`func (o *PatchedBulkIPSecProfileRequest) GetIpsecPolicy() BulkIPSecProfileRequestIpsecPolicy`

GetIpsecPolicy returns the IpsecPolicy field if non-nil, zero value otherwise.

### GetIpsecPolicyOk

`func (o *PatchedBulkIPSecProfileRequest) GetIpsecPolicyOk() (*BulkIPSecProfileRequestIpsecPolicy, bool)`

GetIpsecPolicyOk returns a tuple with the IpsecPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPolicy

`func (o *PatchedBulkIPSecProfileRequest) SetIpsecPolicy(v BulkIPSecProfileRequestIpsecPolicy)`

SetIpsecPolicy sets IpsecPolicy field to given value.

### HasIpsecPolicy

`func (o *PatchedBulkIPSecProfileRequest) HasIpsecPolicy() bool`

HasIpsecPolicy returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkIPSecProfileRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkIPSecProfileRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkIPSecProfileRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkIPSecProfileRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkIPSecProfileRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkIPSecProfileRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkIPSecProfileRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkIPSecProfileRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkIPSecProfileRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkIPSecProfileRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkIPSecProfileRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkIPSecProfileRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkIPSecProfileRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkIPSecProfileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkIPSecProfileRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkIPSecProfileRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkIPSecProfileRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkIPSecProfileRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


