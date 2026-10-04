# PatchedBulkWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**PayloadUrl** | Pointer to **string** | This URL will be called using the HTTP method defined when the webhook is called. Must be http:// or https://. Jinja2 template processing is supported (with the same context as the request body) for part or all of the URL. | [optional] 
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

### NewPatchedBulkWebhookRequest

`func NewPatchedBulkWebhookRequest(id int32, ) *PatchedBulkWebhookRequest`

NewPatchedBulkWebhookRequest instantiates a new PatchedBulkWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWebhookRequestWithDefaults

`func NewPatchedBulkWebhookRequestWithDefaults() *PatchedBulkWebhookRequest`

NewPatchedBulkWebhookRequestWithDefaults instantiates a new PatchedBulkWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWebhookRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWebhookRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWebhookRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWebhookRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWebhookRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWebhookRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWebhookRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWebhookRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWebhookRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPayloadUrl

`func (o *PatchedBulkWebhookRequest) GetPayloadUrl() string`

GetPayloadUrl returns the PayloadUrl field if non-nil, zero value otherwise.

### GetPayloadUrlOk

`func (o *PatchedBulkWebhookRequest) GetPayloadUrlOk() (*string, bool)`

GetPayloadUrlOk returns a tuple with the PayloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadUrl

`func (o *PatchedBulkWebhookRequest) SetPayloadUrl(v string)`

SetPayloadUrl sets PayloadUrl field to given value.

### HasPayloadUrl

`func (o *PatchedBulkWebhookRequest) HasPayloadUrl() bool`

HasPayloadUrl returns a boolean if a field has been set.

### GetHttpMethod

`func (o *PatchedBulkWebhookRequest) GetHttpMethod() BulkWebhookRequestHttpMethod`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *PatchedBulkWebhookRequest) GetHttpMethodOk() (*BulkWebhookRequestHttpMethod, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *PatchedBulkWebhookRequest) SetHttpMethod(v BulkWebhookRequestHttpMethod)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *PatchedBulkWebhookRequest) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetHttpContentType

`func (o *PatchedBulkWebhookRequest) GetHttpContentType() string`

GetHttpContentType returns the HttpContentType field if non-nil, zero value otherwise.

### GetHttpContentTypeOk

`func (o *PatchedBulkWebhookRequest) GetHttpContentTypeOk() (*string, bool)`

GetHttpContentTypeOk returns a tuple with the HttpContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpContentType

`func (o *PatchedBulkWebhookRequest) SetHttpContentType(v string)`

SetHttpContentType sets HttpContentType field to given value.

### HasHttpContentType

`func (o *PatchedBulkWebhookRequest) HasHttpContentType() bool`

HasHttpContentType returns a boolean if a field has been set.

### GetAdditionalHeaders

`func (o *PatchedBulkWebhookRequest) GetAdditionalHeaders() string`

GetAdditionalHeaders returns the AdditionalHeaders field if non-nil, zero value otherwise.

### GetAdditionalHeadersOk

`func (o *PatchedBulkWebhookRequest) GetAdditionalHeadersOk() (*string, bool)`

GetAdditionalHeadersOk returns a tuple with the AdditionalHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalHeaders

`func (o *PatchedBulkWebhookRequest) SetAdditionalHeaders(v string)`

SetAdditionalHeaders sets AdditionalHeaders field to given value.

### HasAdditionalHeaders

`func (o *PatchedBulkWebhookRequest) HasAdditionalHeaders() bool`

HasAdditionalHeaders returns a boolean if a field has been set.

### GetBodyTemplate

`func (o *PatchedBulkWebhookRequest) GetBodyTemplate() string`

GetBodyTemplate returns the BodyTemplate field if non-nil, zero value otherwise.

### GetBodyTemplateOk

`func (o *PatchedBulkWebhookRequest) GetBodyTemplateOk() (*string, bool)`

GetBodyTemplateOk returns a tuple with the BodyTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyTemplate

`func (o *PatchedBulkWebhookRequest) SetBodyTemplate(v string)`

SetBodyTemplate sets BodyTemplate field to given value.

### HasBodyTemplate

`func (o *PatchedBulkWebhookRequest) HasBodyTemplate() bool`

HasBodyTemplate returns a boolean if a field has been set.

### GetSecret

`func (o *PatchedBulkWebhookRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PatchedBulkWebhookRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PatchedBulkWebhookRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PatchedBulkWebhookRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSslVerification

`func (o *PatchedBulkWebhookRequest) GetSslVerification() bool`

GetSslVerification returns the SslVerification field if non-nil, zero value otherwise.

### GetSslVerificationOk

`func (o *PatchedBulkWebhookRequest) GetSslVerificationOk() (*bool, bool)`

GetSslVerificationOk returns a tuple with the SslVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerification

`func (o *PatchedBulkWebhookRequest) SetSslVerification(v bool)`

SetSslVerification sets SslVerification field to given value.

### HasSslVerification

`func (o *PatchedBulkWebhookRequest) HasSslVerification() bool`

HasSslVerification returns a boolean if a field has been set.

### GetCaFilePath

`func (o *PatchedBulkWebhookRequest) GetCaFilePath() string`

GetCaFilePath returns the CaFilePath field if non-nil, zero value otherwise.

### GetCaFilePathOk

`func (o *PatchedBulkWebhookRequest) GetCaFilePathOk() (*string, bool)`

GetCaFilePathOk returns a tuple with the CaFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaFilePath

`func (o *PatchedBulkWebhookRequest) SetCaFilePath(v string)`

SetCaFilePath sets CaFilePath field to given value.

### HasCaFilePath

`func (o *PatchedBulkWebhookRequest) HasCaFilePath() bool`

HasCaFilePath returns a boolean if a field has been set.

### SetCaFilePathNil

`func (o *PatchedBulkWebhookRequest) SetCaFilePathNil(b bool)`

 SetCaFilePathNil sets the value for CaFilePath to be an explicit nil

### UnsetCaFilePath
`func (o *PatchedBulkWebhookRequest) UnsetCaFilePath()`

UnsetCaFilePath ensures that no value is present for CaFilePath, not even an explicit nil
### GetTimeout

`func (o *PatchedBulkWebhookRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PatchedBulkWebhookRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PatchedBulkWebhookRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PatchedBulkWebhookRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *PatchedBulkWebhookRequest) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *PatchedBulkWebhookRequest) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetCustomFields

`func (o *PatchedBulkWebhookRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWebhookRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWebhookRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWebhookRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkWebhookRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkWebhookRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkWebhookRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkWebhookRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkWebhookRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkWebhookRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *PatchedBulkWebhookRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWebhookRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWebhookRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWebhookRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


