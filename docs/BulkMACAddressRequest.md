# BulkMACAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MacAddress** | **string** |  | 
**AssignedObjectType** | Pointer to **NullableString** |  | [optional] 
**AssignedObjectId** | Pointer to **NullableInt64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkMACAddressRequest

`func NewBulkMACAddressRequest(id int32, macAddress string, ) *BulkMACAddressRequest`

NewBulkMACAddressRequest instantiates a new BulkMACAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkMACAddressRequestWithDefaults

`func NewBulkMACAddressRequestWithDefaults() *BulkMACAddressRequest`

NewBulkMACAddressRequestWithDefaults instantiates a new BulkMACAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkMACAddressRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkMACAddressRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkMACAddressRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetMacAddress

`func (o *BulkMACAddressRequest) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *BulkMACAddressRequest) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *BulkMACAddressRequest) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.


### GetAssignedObjectType

`func (o *BulkMACAddressRequest) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *BulkMACAddressRequest) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *BulkMACAddressRequest) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.

### HasAssignedObjectType

`func (o *BulkMACAddressRequest) HasAssignedObjectType() bool`

HasAssignedObjectType returns a boolean if a field has been set.

### SetAssignedObjectTypeNil

`func (o *BulkMACAddressRequest) SetAssignedObjectTypeNil(b bool)`

 SetAssignedObjectTypeNil sets the value for AssignedObjectType to be an explicit nil

### UnsetAssignedObjectType
`func (o *BulkMACAddressRequest) UnsetAssignedObjectType()`

UnsetAssignedObjectType ensures that no value is present for AssignedObjectType, not even an explicit nil
### GetAssignedObjectId

`func (o *BulkMACAddressRequest) GetAssignedObjectId() int64`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *BulkMACAddressRequest) GetAssignedObjectIdOk() (*int64, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *BulkMACAddressRequest) SetAssignedObjectId(v int64)`

SetAssignedObjectId sets AssignedObjectId field to given value.

### HasAssignedObjectId

`func (o *BulkMACAddressRequest) HasAssignedObjectId() bool`

HasAssignedObjectId returns a boolean if a field has been set.

### SetAssignedObjectIdNil

`func (o *BulkMACAddressRequest) SetAssignedObjectIdNil(b bool)`

 SetAssignedObjectIdNil sets the value for AssignedObjectId to be an explicit nil

### UnsetAssignedObjectId
`func (o *BulkMACAddressRequest) UnsetAssignedObjectId()`

UnsetAssignedObjectId ensures that no value is present for AssignedObjectId, not even an explicit nil
### GetDescription

`func (o *BulkMACAddressRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkMACAddressRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkMACAddressRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkMACAddressRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *BulkMACAddressRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkMACAddressRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkMACAddressRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkMACAddressRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkMACAddressRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkMACAddressRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkMACAddressRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkMACAddressRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkMACAddressRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkMACAddressRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkMACAddressRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkMACAddressRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkMACAddressRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkMACAddressRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkMACAddressRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkMACAddressRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkMACAddressRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkMACAddressRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


