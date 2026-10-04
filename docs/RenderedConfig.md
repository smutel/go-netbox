# RenderedConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configtemplate** | [**BriefConfigTemplate**](BriefConfigTemplate.md) |  | [readonly] 
**Content** | **string** | The rendered template output. | [readonly] 

## Methods

### NewRenderedConfig

`func NewRenderedConfig(configtemplate BriefConfigTemplate, content string, ) *RenderedConfig`

NewRenderedConfig instantiates a new RenderedConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderedConfigWithDefaults

`func NewRenderedConfigWithDefaults() *RenderedConfig`

NewRenderedConfigWithDefaults instantiates a new RenderedConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigtemplate

`func (o *RenderedConfig) GetConfigtemplate() BriefConfigTemplate`

GetConfigtemplate returns the Configtemplate field if non-nil, zero value otherwise.

### GetConfigtemplateOk

`func (o *RenderedConfig) GetConfigtemplateOk() (*BriefConfigTemplate, bool)`

GetConfigtemplateOk returns a tuple with the Configtemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigtemplate

`func (o *RenderedConfig) SetConfigtemplate(v BriefConfigTemplate)`

SetConfigtemplate sets Configtemplate field to given value.


### GetContent

`func (o *RenderedConfig) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *RenderedConfig) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *RenderedConfig) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


