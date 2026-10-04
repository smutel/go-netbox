# PatchedBulkImageAttachmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ObjectType** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Image** | Pointer to ***os.File** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedBulkImageAttachmentRequest

`func NewPatchedBulkImageAttachmentRequest(id int32, ) *PatchedBulkImageAttachmentRequest`

NewPatchedBulkImageAttachmentRequest instantiates a new PatchedBulkImageAttachmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkImageAttachmentRequestWithDefaults

`func NewPatchedBulkImageAttachmentRequestWithDefaults() *PatchedBulkImageAttachmentRequest`

NewPatchedBulkImageAttachmentRequestWithDefaults instantiates a new PatchedBulkImageAttachmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkImageAttachmentRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkImageAttachmentRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkImageAttachmentRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetObjectType

`func (o *PatchedBulkImageAttachmentRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PatchedBulkImageAttachmentRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PatchedBulkImageAttachmentRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *PatchedBulkImageAttachmentRequest) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectId

`func (o *PatchedBulkImageAttachmentRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PatchedBulkImageAttachmentRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PatchedBulkImageAttachmentRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *PatchedBulkImageAttachmentRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkImageAttachmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkImageAttachmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkImageAttachmentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkImageAttachmentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *PatchedBulkImageAttachmentRequest) GetImage() *os.File`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PatchedBulkImageAttachmentRequest) GetImageOk() (**os.File, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PatchedBulkImageAttachmentRequest) SetImage(v *os.File)`

SetImage sets Image field to given value.

### HasImage

`func (o *PatchedBulkImageAttachmentRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkImageAttachmentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkImageAttachmentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkImageAttachmentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkImageAttachmentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


