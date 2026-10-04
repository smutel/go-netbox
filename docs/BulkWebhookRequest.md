# BulkWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**PayloadUrl** | **string** | This URL will be called using the HTTP method defined when the webhook is called. Must be http:// or https://. Jinja2 template processing is supported (with the same context as the request body) for part or all of the URL. | 
**HttpMethod** | Pointer to [**BulkWebhookRequestHttpMethod**](BulkWebhookRequestHttpMethod.md) |  | [optional] 
**HttpContentType** | Pointer to **string** | The complete list of official content types is available &lt;a href&#x3D;\&quot;https://www.iana.org/assignments/media-types/media-types.xhtml\&quot;&gt;here&lt;/a&gt;. | [optional] 
**AdditionalHeaders** | Pointer to **string** | User-supplied HTTP headers to be sent with the request in addition to the HTTP content type. Headers should be defined in the format &lt;code&gt;Name: Value&lt;/code&gt;. Jinja2 template processing is supported with the same context as the request body (below). When interpolating untrusted data (such as object attributes) into a header value, apply the &lt;code&gt;header_safe&lt;/code&gt; filter to guard against HTTP header injection, e.g. &lt;code&gt;X-Object: {{ data.name | header_safe }}&lt;/code&gt;. | [optional] 
**BodyTemplate** | Pointer to **string** | Jinja2 template for a custom request body. If blank, a JSON object representing the change will be included. Available context data includes: &lt;code&gt;event&lt;/code&gt;, &lt;code&gt;model&lt;/code&gt;, &lt;code&gt;timestamp&lt;/code&gt;, &lt;code&gt;request&lt;/code&gt;, and &lt;code&gt;data&lt;/code&gt;. | [optional] 
**Secret** | Pointer to **string** | When provided, the request will include a &lt;code&gt;X-Hook-Signature&lt;/code&gt; header containing a HMAC hex digest of the payload body using the secret as the key. The secret is not transmitted in the request. | [optional] 
**SslVerification** | Pointer to **bool** | Enable SSL certificate verification. Disable with caution! | [optional] 
**CaFilePath** | Pointer to **NullableString** | The specific CA certificate file to use for SSL verification. Leave blank to use the system defaults. | [optional] 
**Timeout** | Pointer to **NullableInt32** | The maximum time (in seconds) to wait for a response before failing the request. Leave blank to use the system default (60 seconds). | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 

## Methods

### NewBulkWebhookRequest

`func NewBulkWebhookRequest(id int32, name string, payloadUrl string, ) *BulkWebhookRequest`

NewBulkWebhookRequest instantiates a new BulkWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWebhookRequestWithDefaults

`func NewBulkWebhookRequestWithDefaults() *BulkWebhookRequest`

NewBulkWebhookRequestWithDefaults instantiates a new BulkWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWebhookRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWebhookRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWebhookRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BulkWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWebhookRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkWebhookRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWebhookRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWebhookRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWebhookRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPayloadUrl

`func (o *BulkWebhookRequest) GetPayloadUrl() string`

GetPayloadUrl returns the PayloadUrl field if non-nil, zero value otherwise.

### GetPayloadUrlOk

`func (o *BulkWebhookRequest) GetPayloadUrlOk() (*string, bool)`

GetPayloadUrlOk returns a tuple with the PayloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadUrl

`func (o *BulkWebhookRequest) SetPayloadUrl(v string)`

SetPayloadUrl sets PayloadUrl field to given value.


### GetHttpMethod

`func (o *BulkWebhookRequest) GetHttpMethod() BulkWebhookRequestHttpMethod`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *BulkWebhookRequest) GetHttpMethodOk() (*BulkWebhookRequestHttpMethod, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *BulkWebhookRequest) SetHttpMethod(v BulkWebhookRequestHttpMethod)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *BulkWebhookRequest) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetHttpContentType

`func (o *BulkWebhookRequest) GetHttpContentType() string`

GetHttpContentType returns the HttpContentType field if non-nil, zero value otherwise.

### GetHttpContentTypeOk

`func (o *BulkWebhookRequest) GetHttpContentTypeOk() (*string, bool)`

GetHttpContentTypeOk returns a tuple with the HttpContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpContentType

`func (o *BulkWebhookRequest) SetHttpContentType(v string)`

SetHttpContentType sets HttpContentType field to given value.

### HasHttpContentType

`func (o *BulkWebhookRequest) HasHttpContentType() bool`

HasHttpContentType returns a boolean if a field has been set.

### GetAdditionalHeaders

`func (o *BulkWebhookRequest) GetAdditionalHeaders() string`

GetAdditionalHeaders returns the AdditionalHeaders field if non-nil, zero value otherwise.

### GetAdditionalHeadersOk

`func (o *BulkWebhookRequest) GetAdditionalHeadersOk() (*string, bool)`

GetAdditionalHeadersOk returns a tuple with the AdditionalHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalHeaders

`func (o *BulkWebhookRequest) SetAdditionalHeaders(v string)`

SetAdditionalHeaders sets AdditionalHeaders field to given value.

### HasAdditionalHeaders

`func (o *BulkWebhookRequest) HasAdditionalHeaders() bool`

HasAdditionalHeaders returns a boolean if a field has been set.

### GetBodyTemplate

`func (o *BulkWebhookRequest) GetBodyTemplate() string`

GetBodyTemplate returns the BodyTemplate field if non-nil, zero value otherwise.

### GetBodyTemplateOk

`func (o *BulkWebhookRequest) GetBodyTemplateOk() (*string, bool)`

GetBodyTemplateOk returns a tuple with the BodyTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyTemplate

`func (o *BulkWebhookRequest) SetBodyTemplate(v string)`

SetBodyTemplate sets BodyTemplate field to given value.

### HasBodyTemplate

`func (o *BulkWebhookRequest) HasBodyTemplate() bool`

HasBodyTemplate returns a boolean if a field has been set.

### GetSecret

`func (o *BulkWebhookRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *BulkWebhookRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *BulkWebhookRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *BulkWebhookRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSslVerification

`func (o *BulkWebhookRequest) GetSslVerification() bool`

GetSslVerification returns the SslVerification field if non-nil, zero value otherwise.

### GetSslVerificationOk

`func (o *BulkWebhookRequest) GetSslVerificationOk() (*bool, bool)`

GetSslVerificationOk returns a tuple with the SslVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerification

`func (o *BulkWebhookRequest) SetSslVerification(v bool)`

SetSslVerification sets SslVerification field to given value.

### HasSslVerification

`func (o *BulkWebhookRequest) HasSslVerification() bool`

HasSslVerification returns a boolean if a field has been set.

### GetCaFilePath

`func (o *BulkWebhookRequest) GetCaFilePath() string`

GetCaFilePath returns the CaFilePath field if non-nil, zero value otherwise.

### GetCaFilePathOk

`func (o *BulkWebhookRequest) GetCaFilePathOk() (*string, bool)`

GetCaFilePathOk returns a tuple with the CaFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaFilePath

`func (o *BulkWebhookRequest) SetCaFilePath(v string)`

SetCaFilePath sets CaFilePath field to given value.

### HasCaFilePath

`func (o *BulkWebhookRequest) HasCaFilePath() bool`

HasCaFilePath returns a boolean if a field has been set.

### SetCaFilePathNil

`func (o *BulkWebhookRequest) SetCaFilePathNil(b bool)`

 SetCaFilePathNil sets the value for CaFilePath to be an explicit nil

### UnsetCaFilePath
`func (o *BulkWebhookRequest) UnsetCaFilePath()`

UnsetCaFilePath ensures that no value is present for CaFilePath, not even an explicit nil
### GetTimeout

`func (o *BulkWebhookRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *BulkWebhookRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *BulkWebhookRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *BulkWebhookRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *BulkWebhookRequest) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *BulkWebhookRequest) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetCustomFields

`func (o *BulkWebhookRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWebhookRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWebhookRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWebhookRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetOwner

`func (o *BulkWebhookRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkWebhookRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkWebhookRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkWebhookRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkWebhookRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkWebhookRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *BulkWebhookRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWebhookRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWebhookRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWebhookRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


