# PatchedBulkConfigContextProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to **interface{}** | A JSON schema specifying the structure of the context data for this profile | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**DataSource** | Pointer to [**BulkConfigContextProfileRequestDataSource**](BulkConfigContextProfileRequestDataSource.md) |  | [optional] 

## Methods

### NewPatchedBulkConfigContextProfileRequest

`func NewPatchedBulkConfigContextProfileRequest(id int32, ) *PatchedBulkConfigContextProfileRequest`

NewPatchedBulkConfigContextProfileRequest instantiates a new PatchedBulkConfigContextProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkConfigContextProfileRequestWithDefaults

`func NewPatchedBulkConfigContextProfileRequestWithDefaults() *PatchedBulkConfigContextProfileRequest`

NewPatchedBulkConfigContextProfileRequestWithDefaults instantiates a new PatchedBulkConfigContextProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkConfigContextProfileRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkConfigContextProfileRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkConfigContextProfileRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkConfigContextProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkConfigContextProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkConfigContextProfileRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkConfigContextProfileRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkConfigContextProfileRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkConfigContextProfileRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkConfigContextProfileRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkConfigContextProfileRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSchema

`func (o *PatchedBulkConfigContextProfileRequest) GetSchema() interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *PatchedBulkConfigContextProfileRequest) GetSchemaOk() (*interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *PatchedBulkConfigContextProfileRequest) SetSchema(v interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *PatchedBulkConfigContextProfileRequest) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *PatchedBulkConfigContextProfileRequest) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *PatchedBulkConfigContextProfileRequest) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetTags

`func (o *PatchedBulkConfigContextProfileRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkConfigContextProfileRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkConfigContextProfileRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkConfigContextProfileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkConfigContextProfileRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkConfigContextProfileRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkConfigContextProfileRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkConfigContextProfileRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkConfigContextProfileRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkConfigContextProfileRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkConfigContextProfileRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkConfigContextProfileRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkConfigContextProfileRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkConfigContextProfileRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetDataSource

`func (o *PatchedBulkConfigContextProfileRequest) GetDataSource() BulkConfigContextProfileRequestDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *PatchedBulkConfigContextProfileRequest) GetDataSourceOk() (*BulkConfigContextProfileRequestDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *PatchedBulkConfigContextProfileRequest) SetDataSource(v BulkConfigContextProfileRequestDataSource)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *PatchedBulkConfigContextProfileRequest) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


