# BulkConfigTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**EnvironmentParams** | Pointer to **interface{}** | Any &lt;a href&#x3D;\&quot;https://jinja.palletsprojects.com/en/stable/api/#jinja2.Environment\&quot;&gt;additional parameters&lt;/a&gt; to pass when constructing the Jinja environment | [optional] 
**TemplateCode** | **string** | Jinja template code. | 
**MimeType** | Pointer to **string** | Defaults to &lt;code&gt;text/plain; charset&#x3D;utf-8&lt;/code&gt; | [optional] 
**FileName** | Pointer to **string** | Filename to give to the rendered export file | [optional] 
**FileExtension** | Pointer to **string** | Extension to append to the rendered filename | [optional] 
**AsAttachment** | Pointer to **bool** | Download file as attachment | [optional] 
**Debug** | Pointer to **bool** | Enable verbose error output when rendering this template. Not recommended for production use. | [optional] 
**DataSource** | Pointer to [**BulkConfigContextProfileRequestDataSource**](BulkConfigContextProfileRequestDataSource.md) |  | [optional] 
**AutoSyncEnabled** | Pointer to **bool** | Enable automatic synchronization of data when the data file is updated | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 

## Methods

### NewBulkConfigTemplateRequest

`func NewBulkConfigTemplateRequest(id int32, name string, templateCode string, ) *BulkConfigTemplateRequest`

NewBulkConfigTemplateRequest instantiates a new BulkConfigTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkConfigTemplateRequestWithDefaults

`func NewBulkConfigTemplateRequestWithDefaults() *BulkConfigTemplateRequest`

NewBulkConfigTemplateRequestWithDefaults instantiates a new BulkConfigTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkConfigTemplateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkConfigTemplateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkConfigTemplateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BulkConfigTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkConfigTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkConfigTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkConfigTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkConfigTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkConfigTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkConfigTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentParams

`func (o *BulkConfigTemplateRequest) GetEnvironmentParams() interface{}`

GetEnvironmentParams returns the EnvironmentParams field if non-nil, zero value otherwise.

### GetEnvironmentParamsOk

`func (o *BulkConfigTemplateRequest) GetEnvironmentParamsOk() (*interface{}, bool)`

GetEnvironmentParamsOk returns a tuple with the EnvironmentParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentParams

`func (o *BulkConfigTemplateRequest) SetEnvironmentParams(v interface{})`

SetEnvironmentParams sets EnvironmentParams field to given value.

### HasEnvironmentParams

`func (o *BulkConfigTemplateRequest) HasEnvironmentParams() bool`

HasEnvironmentParams returns a boolean if a field has been set.

### SetEnvironmentParamsNil

`func (o *BulkConfigTemplateRequest) SetEnvironmentParamsNil(b bool)`

 SetEnvironmentParamsNil sets the value for EnvironmentParams to be an explicit nil

### UnsetEnvironmentParams
`func (o *BulkConfigTemplateRequest) UnsetEnvironmentParams()`

UnsetEnvironmentParams ensures that no value is present for EnvironmentParams, not even an explicit nil
### GetTemplateCode

`func (o *BulkConfigTemplateRequest) GetTemplateCode() string`

GetTemplateCode returns the TemplateCode field if non-nil, zero value otherwise.

### GetTemplateCodeOk

`func (o *BulkConfigTemplateRequest) GetTemplateCodeOk() (*string, bool)`

GetTemplateCodeOk returns a tuple with the TemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCode

`func (o *BulkConfigTemplateRequest) SetTemplateCode(v string)`

SetTemplateCode sets TemplateCode field to given value.


### GetMimeType

`func (o *BulkConfigTemplateRequest) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *BulkConfigTemplateRequest) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *BulkConfigTemplateRequest) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *BulkConfigTemplateRequest) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetFileName

`func (o *BulkConfigTemplateRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *BulkConfigTemplateRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *BulkConfigTemplateRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *BulkConfigTemplateRequest) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileExtension

`func (o *BulkConfigTemplateRequest) GetFileExtension() string`

GetFileExtension returns the FileExtension field if non-nil, zero value otherwise.

### GetFileExtensionOk

`func (o *BulkConfigTemplateRequest) GetFileExtensionOk() (*string, bool)`

GetFileExtensionOk returns a tuple with the FileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtension

`func (o *BulkConfigTemplateRequest) SetFileExtension(v string)`

SetFileExtension sets FileExtension field to given value.

### HasFileExtension

`func (o *BulkConfigTemplateRequest) HasFileExtension() bool`

HasFileExtension returns a boolean if a field has been set.

### GetAsAttachment

`func (o *BulkConfigTemplateRequest) GetAsAttachment() bool`

GetAsAttachment returns the AsAttachment field if non-nil, zero value otherwise.

### GetAsAttachmentOk

`func (o *BulkConfigTemplateRequest) GetAsAttachmentOk() (*bool, bool)`

GetAsAttachmentOk returns a tuple with the AsAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsAttachment

`func (o *BulkConfigTemplateRequest) SetAsAttachment(v bool)`

SetAsAttachment sets AsAttachment field to given value.

### HasAsAttachment

`func (o *BulkConfigTemplateRequest) HasAsAttachment() bool`

HasAsAttachment returns a boolean if a field has been set.

### GetDebug

`func (o *BulkConfigTemplateRequest) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *BulkConfigTemplateRequest) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *BulkConfigTemplateRequest) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *BulkConfigTemplateRequest) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetDataSource

`func (o *BulkConfigTemplateRequest) GetDataSource() BulkConfigContextProfileRequestDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *BulkConfigTemplateRequest) GetDataSourceOk() (*BulkConfigContextProfileRequestDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *BulkConfigTemplateRequest) SetDataSource(v BulkConfigContextProfileRequestDataSource)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *BulkConfigTemplateRequest) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetAutoSyncEnabled

`func (o *BulkConfigTemplateRequest) GetAutoSyncEnabled() bool`

GetAutoSyncEnabled returns the AutoSyncEnabled field if non-nil, zero value otherwise.

### GetAutoSyncEnabledOk

`func (o *BulkConfigTemplateRequest) GetAutoSyncEnabledOk() (*bool, bool)`

GetAutoSyncEnabledOk returns a tuple with the AutoSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncEnabled

`func (o *BulkConfigTemplateRequest) SetAutoSyncEnabled(v bool)`

SetAutoSyncEnabled sets AutoSyncEnabled field to given value.

### HasAutoSyncEnabled

`func (o *BulkConfigTemplateRequest) HasAutoSyncEnabled() bool`

HasAutoSyncEnabled returns a boolean if a field has been set.

### GetOwner

`func (o *BulkConfigTemplateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkConfigTemplateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkConfigTemplateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkConfigTemplateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkConfigTemplateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkConfigTemplateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *BulkConfigTemplateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkConfigTemplateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkConfigTemplateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkConfigTemplateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


