# BulkIPSecProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | [**BulkIPSecProfileRequestMode**](BulkIPSecProfileRequestMode.md) |  | 
**IkePolicy** | [**BulkIPSecProfileRequestIkePolicy**](BulkIPSecProfileRequestIkePolicy.md) |  | 
**IpsecPolicy** | [**BulkIPSecProfileRequestIpsecPolicy**](BulkIPSecProfileRequestIpsecPolicy.md) |  | 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkIPSecProfileRequest

`func NewBulkIPSecProfileRequest(id int32, name string, mode BulkIPSecProfileRequestMode, ikePolicy BulkIPSecProfileRequestIkePolicy, ipsecPolicy BulkIPSecProfileRequestIpsecPolicy, ) *BulkIPSecProfileRequest`

NewBulkIPSecProfileRequest instantiates a new BulkIPSecProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkIPSecProfileRequestWithDefaults

`func NewBulkIPSecProfileRequestWithDefaults() *BulkIPSecProfileRequest`

NewBulkIPSecProfileRequestWithDefaults instantiates a new BulkIPSecProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkIPSecProfileRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkIPSecProfileRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkIPSecProfileRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BulkIPSecProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkIPSecProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkIPSecProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkIPSecProfileRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkIPSecProfileRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkIPSecProfileRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkIPSecProfileRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *BulkIPSecProfileRequest) GetMode() BulkIPSecProfileRequestMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *BulkIPSecProfileRequest) GetModeOk() (*BulkIPSecProfileRequestMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *BulkIPSecProfileRequest) SetMode(v BulkIPSecProfileRequestMode)`

SetMode sets Mode field to given value.


### GetIkePolicy

`func (o *BulkIPSecProfileRequest) GetIkePolicy() BulkIPSecProfileRequestIkePolicy`

GetIkePolicy returns the IkePolicy field if non-nil, zero value otherwise.

### GetIkePolicyOk

`func (o *BulkIPSecProfileRequest) GetIkePolicyOk() (*BulkIPSecProfileRequestIkePolicy, bool)`

GetIkePolicyOk returns a tuple with the IkePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkePolicy

`func (o *BulkIPSecProfileRequest) SetIkePolicy(v BulkIPSecProfileRequestIkePolicy)`

SetIkePolicy sets IkePolicy field to given value.


### GetIpsecPolicy

`func (o *BulkIPSecProfileRequest) GetIpsecPolicy() BulkIPSecProfileRequestIpsecPolicy`

GetIpsecPolicy returns the IpsecPolicy field if non-nil, zero value otherwise.

### GetIpsecPolicyOk

`func (o *BulkIPSecProfileRequest) GetIpsecPolicyOk() (*BulkIPSecProfileRequestIpsecPolicy, bool)`

GetIpsecPolicyOk returns a tuple with the IpsecPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPolicy

`func (o *BulkIPSecProfileRequest) SetIpsecPolicy(v BulkIPSecProfileRequestIpsecPolicy)`

SetIpsecPolicy sets IpsecPolicy field to given value.


### GetOwner

`func (o *BulkIPSecProfileRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkIPSecProfileRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkIPSecProfileRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkIPSecProfileRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkIPSecProfileRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkIPSecProfileRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkIPSecProfileRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkIPSecProfileRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkIPSecProfileRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkIPSecProfileRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkIPSecProfileRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkIPSecProfileRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkIPSecProfileRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkIPSecProfileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkIPSecProfileRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkIPSecProfileRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkIPSecProfileRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkIPSecProfileRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


