# BulkExportTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ObjectTypes** | **[]string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**EnvironmentParams** | Pointer to **interface{}** | Any &lt;a href&#x3D;\&quot;https://jinja.palletsprojects.com/en/stable/api/#jinja2.Environment\&quot;&gt;additional parameters&lt;/a&gt; to pass when constructing the Jinja environment | [optional] 
**TemplateCode** | **string** | Jinja template code. | 
**MimeType** | Pointer to **string** | Defaults to &lt;code&gt;text/plain; charset&#x3D;utf-8&lt;/code&gt; | [optional] 
**FileName** | Pointer to **string** | Filename to give to the rendered export file | [optional] 
**FileExtension** | Pointer to **string** | Extension to append to the rendered filename | [optional] 
**AsAttachment** | Pointer to **bool** | Download file as attachment | [optional] 
**DataSource** | Pointer to [**BulkConfigContextProfileRequestDataSource**](BulkConfigContextProfileRequestDataSource.md) |  | [optional] 
**AutoSyncEnabled** | Pointer to **bool** | Enable automatic synchronization of data when the data file is updated | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 

## Methods

### NewBulkExportTemplateRequest

`func NewBulkExportTemplateRequest(id int32, objectTypes []string, name string, templateCode string, ) *BulkExportTemplateRequest`

NewBulkExportTemplateRequest instantiates a new BulkExportTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkExportTemplateRequestWithDefaults

`func NewBulkExportTemplateRequestWithDefaults() *BulkExportTemplateRequest`

NewBulkExportTemplateRequestWithDefaults instantiates a new BulkExportTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkExportTemplateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkExportTemplateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkExportTemplateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetObjectTypes

`func (o *BulkExportTemplateRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *BulkExportTemplateRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *BulkExportTemplateRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.


### GetName

`func (o *BulkExportTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkExportTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkExportTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkExportTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkExportTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkExportTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkExportTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentParams

`func (o *BulkExportTemplateRequest) GetEnvironmentParams() interface{}`

GetEnvironmentParams returns the EnvironmentParams field if non-nil, zero value otherwise.

### GetEnvironmentParamsOk

`func (o *BulkExportTemplateRequest) GetEnvironmentParamsOk() (*interface{}, bool)`

GetEnvironmentParamsOk returns a tuple with the EnvironmentParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentParams

`func (o *BulkExportTemplateRequest) SetEnvironmentParams(v interface{})`

SetEnvironmentParams sets EnvironmentParams field to given value.

### HasEnvironmentParams

`func (o *BulkExportTemplateRequest) HasEnvironmentParams() bool`

HasEnvironmentParams returns a boolean if a field has been set.

### SetEnvironmentParamsNil

`func (o *BulkExportTemplateRequest) SetEnvironmentParamsNil(b bool)`

 SetEnvironmentParamsNil sets the value for EnvironmentParams to be an explicit nil

### UnsetEnvironmentParams
`func (o *BulkExportTemplateRequest) UnsetEnvironmentParams()`

UnsetEnvironmentParams ensures that no value is present for EnvironmentParams, not even an explicit nil
### GetTemplateCode

`func (o *BulkExportTemplateRequest) GetTemplateCode() string`

GetTemplateCode returns the TemplateCode field if non-nil, zero value otherwise.

### GetTemplateCodeOk

`func (o *BulkExportTemplateRequest) GetTemplateCodeOk() (*string, bool)`

GetTemplateCodeOk returns a tuple with the TemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCode

`func (o *BulkExportTemplateRequest) SetTemplateCode(v string)`

SetTemplateCode sets TemplateCode field to given value.


### GetMimeType

`func (o *BulkExportTemplateRequest) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *BulkExportTemplateRequest) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *BulkExportTemplateRequest) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *BulkExportTemplateRequest) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetFileName

`func (o *BulkExportTemplateRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *BulkExportTemplateRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *BulkExportTemplateRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *BulkExportTemplateRequest) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileExtension

`func (o *BulkExportTemplateRequest) GetFileExtension() string`

GetFileExtension returns the FileExtension field if non-nil, zero value otherwise.

### GetFileExtensionOk

`func (o *BulkExportTemplateRequest) GetFileExtensionOk() (*string, bool)`

GetFileExtensionOk returns a tuple with the FileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtension

`func (o *BulkExportTemplateRequest) SetFileExtension(v string)`

SetFileExtension sets FileExtension field to given value.

### HasFileExtension

`func (o *BulkExportTemplateRequest) HasFileExtension() bool`

HasFileExtension returns a boolean if a field has been set.

### GetAsAttachment

`func (o *BulkExportTemplateRequest) GetAsAttachment() bool`

GetAsAttachment returns the AsAttachment field if non-nil, zero value otherwise.

### GetAsAttachmentOk

`func (o *BulkExportTemplateRequest) GetAsAttachmentOk() (*bool, bool)`

GetAsAttachmentOk returns a tuple with the AsAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsAttachment

`func (o *BulkExportTemplateRequest) SetAsAttachment(v bool)`

SetAsAttachment sets AsAttachment field to given value.

### HasAsAttachment

`func (o *BulkExportTemplateRequest) HasAsAttachment() bool`

HasAsAttachment returns a boolean if a field has been set.

### GetDataSource

`func (o *BulkExportTemplateRequest) GetDataSource() BulkConfigContextProfileRequestDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *BulkExportTemplateRequest) GetDataSourceOk() (*BulkConfigContextProfileRequestDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *BulkExportTemplateRequest) SetDataSource(v BulkConfigContextProfileRequestDataSource)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *BulkExportTemplateRequest) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetAutoSyncEnabled

`func (o *BulkExportTemplateRequest) GetAutoSyncEnabled() bool`

GetAutoSyncEnabled returns the AutoSyncEnabled field if non-nil, zero value otherwise.

### GetAutoSyncEnabledOk

`func (o *BulkExportTemplateRequest) GetAutoSyncEnabledOk() (*bool, bool)`

GetAutoSyncEnabledOk returns a tuple with the AutoSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncEnabled

`func (o *BulkExportTemplateRequest) SetAutoSyncEnabled(v bool)`

SetAutoSyncEnabled sets AutoSyncEnabled field to given value.

### HasAutoSyncEnabled

`func (o *BulkExportTemplateRequest) HasAutoSyncEnabled() bool`

HasAutoSyncEnabled returns a boolean if a field has been set.

### GetOwner

`func (o *BulkExportTemplateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkExportTemplateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkExportTemplateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkExportTemplateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkExportTemplateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkExportTemplateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


