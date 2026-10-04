# BulkImageAttachmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ObjectType** | **string** |  | 
**ObjectId** | **int64** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Image** | ***os.File** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkImageAttachmentRequest

`func NewBulkImageAttachmentRequest(id int32, objectType string, objectId int64, image *os.File, ) *BulkImageAttachmentRequest`

NewBulkImageAttachmentRequest instantiates a new BulkImageAttachmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkImageAttachmentRequestWithDefaults

`func NewBulkImageAttachmentRequestWithDefaults() *BulkImageAttachmentRequest`

NewBulkImageAttachmentRequestWithDefaults instantiates a new BulkImageAttachmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkImageAttachmentRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkImageAttachmentRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkImageAttachmentRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetObjectType

`func (o *BulkImageAttachmentRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkImageAttachmentRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkImageAttachmentRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *BulkImageAttachmentRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BulkImageAttachmentRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BulkImageAttachmentRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetName

`func (o *BulkImageAttachmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkImageAttachmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkImageAttachmentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BulkImageAttachmentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *BulkImageAttachmentRequest) GetImage() *os.File`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BulkImageAttachmentRequest) GetImageOk() (**os.File, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BulkImageAttachmentRequest) SetImage(v *os.File)`

SetImage sets Image field to given value.


### GetDescription

`func (o *BulkImageAttachmentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkImageAttachmentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkImageAttachmentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkImageAttachmentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


