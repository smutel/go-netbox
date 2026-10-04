# RenderConfigInputRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigTemplateId** | Pointer to **int32** | Optional ID of the ConfigTemplate to render. If omitted, the object&#39;s assigned config template is used. | [optional] 

## Methods

### NewRenderConfigInputRequest

`func NewRenderConfigInputRequest() *RenderConfigInputRequest`

NewRenderConfigInputRequest instantiates a new RenderConfigInputRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderConfigInputRequestWithDefaults

`func NewRenderConfigInputRequestWithDefaults() *RenderConfigInputRequest`

NewRenderConfigInputRequestWithDefaults instantiates a new RenderConfigInputRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigTemplateId

`func (o *RenderConfigInputRequest) GetConfigTemplateId() int32`

GetConfigTemplateId returns the ConfigTemplateId field if non-nil, zero value otherwise.

### GetConfigTemplateIdOk

`func (o *RenderConfigInputRequest) GetConfigTemplateIdOk() (*int32, bool)`

GetConfigTemplateIdOk returns a tuple with the ConfigTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplateId

`func (o *RenderConfigInputRequest) SetConfigTemplateId(v int32)`

SetConfigTemplateId sets ConfigTemplateId field to given value.

### HasConfigTemplateId

`func (o *RenderConfigInputRequest) HasConfigTemplateId() bool`

HasConfigTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


