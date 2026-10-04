# PatchedBulkExportTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ObjectTypes** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnvironmentParams** | Pointer to **interface{}** | Any &lt;a href&#x3D;\&quot;https://jinja.palletsprojects.com/en/stable/api/#jinja2.Environment\&quot;&gt;additional parameters&lt;/a&gt; to pass when constructing the Jinja environment | [optional] 
**TemplateCode** | Pointer to **string** | Jinja template code. | [optional] 
**MimeType** | Pointer to **string** | Defaults to &lt;code&gt;text/plain; charset&#x3D;utf-8&lt;/code&gt; | [optional] 
**FileName** | Pointer to **string** | Filename to give to the rendered export file | [optional] 
**FileExtension** | Pointer to **string** | Extension to append to the rendered filename | [optional] 
**AsAttachment** | Pointer to **bool** | Download file as attachment | [optional] 
**DataSource** | Pointer to [**BulkConfigContextProfileRequestDataSource**](BulkConfigContextProfileRequestDataSource.md) |  | [optional] 
**AutoSyncEnabled** | Pointer to **bool** | Enable automatic synchronization of data when the data file is updated | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 

## Methods

### NewPatchedBulkExportTemplateRequest

`func NewPatchedBulkExportTemplateRequest(id int32, ) *PatchedBulkExportTemplateRequest`

NewPatchedBulkExportTemplateRequest instantiates a new PatchedBulkExportTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkExportTemplateRequestWithDefaults

`func NewPatchedBulkExportTemplateRequestWithDefaults() *PatchedBulkExportTemplateRequest`

NewPatchedBulkExportTemplateRequestWithDefaults instantiates a new PatchedBulkExportTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkExportTemplateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkExportTemplateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkExportTemplateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetObjectTypes

`func (o *PatchedBulkExportTemplateRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *PatchedBulkExportTemplateRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *PatchedBulkExportTemplateRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.

### HasObjectTypes

`func (o *PatchedBulkExportTemplateRequest) HasObjectTypes() bool`

HasObjectTypes returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkExportTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkExportTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkExportTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkExportTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkExportTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkExportTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkExportTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkExportTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentParams

`func (o *PatchedBulkExportTemplateRequest) GetEnvironmentParams() interface{}`

GetEnvironmentParams returns the EnvironmentParams field if non-nil, zero value otherwise.

### GetEnvironmentParamsOk

`func (o *PatchedBulkExportTemplateRequest) GetEnvironmentParamsOk() (*interface{}, bool)`

GetEnvironmentParamsOk returns a tuple with the EnvironmentParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentParams

`func (o *PatchedBulkExportTemplateRequest) SetEnvironmentParams(v interface{})`

SetEnvironmentParams sets EnvironmentParams field to given value.

### HasEnvironmentParams

`func (o *PatchedBulkExportTemplateRequest) HasEnvironmentParams() bool`

HasEnvironmentParams returns a boolean if a field has been set.

### SetEnvironmentParamsNil

`func (o *PatchedBulkExportTemplateRequest) SetEnvironmentParamsNil(b bool)`

 SetEnvironmentParamsNil sets the value for EnvironmentParams to be an explicit nil

### UnsetEnvironmentParams
`func (o *PatchedBulkExportTemplateRequest) UnsetEnvironmentParams()`

UnsetEnvironmentParams ensures that no value is present for EnvironmentParams, not even an explicit nil
### GetTemplateCode

`func (o *PatchedBulkExportTemplateRequest) GetTemplateCode() string`

GetTemplateCode returns the TemplateCode field if non-nil, zero value otherwise.

### GetTemplateCodeOk

`func (o *PatchedBulkExportTemplateRequest) GetTemplateCodeOk() (*string, bool)`

GetTemplateCodeOk returns a tuple with the TemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCode

`func (o *PatchedBulkExportTemplateRequest) SetTemplateCode(v string)`

SetTemplateCode sets TemplateCode field to given value.

### HasTemplateCode

`func (o *PatchedBulkExportTemplateRequest) HasTemplateCode() bool`

HasTemplateCode returns a boolean if a field has been set.

### GetMimeType

`func (o *PatchedBulkExportTemplateRequest) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *PatchedBulkExportTemplateRequest) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *PatchedBulkExportTemplateRequest) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *PatchedBulkExportTemplateRequest) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetFileName

`func (o *PatchedBulkExportTemplateRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *PatchedBulkExportTemplateRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *PatchedBulkExportTemplateRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *PatchedBulkExportTemplateRequest) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileExtension

`func (o *PatchedBulkExportTemplateRequest) GetFileExtension() string`

GetFileExtension returns the FileExtension field if non-nil, zero value otherwise.

### GetFileExtensionOk

`func (o *PatchedBulkExportTemplateRequest) GetFileExtensionOk() (*string, bool)`

GetFileExtensionOk returns a tuple with the FileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtension

`func (o *PatchedBulkExportTemplateRequest) SetFileExtension(v string)`

SetFileExtension sets FileExtension field to given value.

### HasFileExtension

`func (o *PatchedBulkExportTemplateRequest) HasFileExtension() bool`

HasFileExtension returns a boolean if a field has been set.

### GetAsAttachment

`func (o *PatchedBulkExportTemplateRequest) GetAsAttachment() bool`

GetAsAttachment returns the AsAttachment field if non-nil, zero value otherwise.

### GetAsAttachmentOk

`func (o *PatchedBulkExportTemplateRequest) GetAsAttachmentOk() (*bool, bool)`

GetAsAttachmentOk returns a tuple with the AsAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsAttachment

`func (o *PatchedBulkExportTemplateRequest) SetAsAttachment(v bool)`

SetAsAttachment sets AsAttachment field to given value.

### HasAsAttachment

`func (o *PatchedBulkExportTemplateRequest) HasAsAttachment() bool`

HasAsAttachment returns a boolean if a field has been set.

### GetDataSource

`func (o *PatchedBulkExportTemplateRequest) GetDataSource() BulkConfigContextProfileRequestDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *PatchedBulkExportTemplateRequest) GetDataSourceOk() (*BulkConfigContextProfileRequestDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *PatchedBulkExportTemplateRequest) SetDataSource(v BulkConfigContextProfileRequestDataSource)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *PatchedBulkExportTemplateRequest) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetAutoSyncEnabled

`func (o *PatchedBulkExportTemplateRequest) GetAutoSyncEnabled() bool`

GetAutoSyncEnabled returns the AutoSyncEnabled field if non-nil, zero value otherwise.

### GetAutoSyncEnabledOk

`func (o *PatchedBulkExportTemplateRequest) GetAutoSyncEnabledOk() (*bool, bool)`

GetAutoSyncEnabledOk returns a tuple with the AutoSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncEnabled

`func (o *PatchedBulkExportTemplateRequest) SetAutoSyncEnabled(v bool)`

SetAutoSyncEnabled sets AutoSyncEnabled field to given value.

### HasAutoSyncEnabled

`func (o *PatchedBulkExportTemplateRequest) HasAutoSyncEnabled() bool`

HasAutoSyncEnabled returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkExportTemplateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkExportTemplateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkExportTemplateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkExportTemplateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkExportTemplateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkExportTemplateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


