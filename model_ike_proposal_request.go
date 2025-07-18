/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// checks if the IKEProposalRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IKEProposalRequest{}

// IKEProposalRequest Adds support for custom fields and tags.
type IKEProposalRequest struct {
	Name                    string                                   `json:"name"`
	Description             *string                                  `json:"description,omitempty"`
	AuthenticationMethod    IKEProposalAuthenticationMethodValue     `json:"authentication_method"`
	EncryptionAlgorithm     IKEProposalEncryptionAlgorithmValue      `json:"encryption_algorithm"`
	AuthenticationAlgorithm *IKEProposalAuthenticationAlgorithmValue `json:"authentication_algorithm,omitempty"`
	Group                   IKEProposalGroupValue                    `json:"group"`
	// Security association lifetime (in seconds)
	SaLifetime           NullableInt32          `json:"sa_lifetime,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IKEProposalRequest IKEProposalRequest

// NewIKEProposalRequest instantiates a new IKEProposalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIKEProposalRequest(name string, authenticationMethod IKEProposalAuthenticationMethodValue, encryptionAlgorithm IKEProposalEncryptionAlgorithmValue, group IKEProposalGroupValue) *IKEProposalRequest {
	this := IKEProposalRequest{}
	this.Name = name
	this.AuthenticationMethod = authenticationMethod
	this.EncryptionAlgorithm = encryptionAlgorithm
	this.Group = group
	return &this
}

// NewIKEProposalRequestWithDefaults instantiates a new IKEProposalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIKEProposalRequestWithDefaults() *IKEProposalRequest {
	this := IKEProposalRequest{}
	return &this
}

// GetName returns the Name field value
func (o *IKEProposalRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IKEProposalRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IKEProposalRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IKEProposalRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEProposalRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IKEProposalRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IKEProposalRequest) SetDescription(v string) {
	o.Description = &v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value
func (o *IKEProposalRequest) GetAuthenticationMethod() IKEProposalAuthenticationMethodValue {
	if o == nil {
		var ret IKEProposalAuthenticationMethodValue
		return ret
	}

	return o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value
// and a boolean to check if the value has been set.
func (o *IKEProposalRequest) GetAuthenticationMethodOk() (*IKEProposalAuthenticationMethodValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationMethod, true
}

// SetAuthenticationMethod sets field value
func (o *IKEProposalRequest) SetAuthenticationMethod(v IKEProposalAuthenticationMethodValue) {
	o.AuthenticationMethod = v
}

// GetEncryptionAlgorithm returns the EncryptionAlgorithm field value
func (o *IKEProposalRequest) GetEncryptionAlgorithm() IKEProposalEncryptionAlgorithmValue {
	if o == nil {
		var ret IKEProposalEncryptionAlgorithmValue
		return ret
	}

	return o.EncryptionAlgorithm
}

// GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field value
// and a boolean to check if the value has been set.
func (o *IKEProposalRequest) GetEncryptionAlgorithmOk() (*IKEProposalEncryptionAlgorithmValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptionAlgorithm, true
}

// SetEncryptionAlgorithm sets field value
func (o *IKEProposalRequest) SetEncryptionAlgorithm(v IKEProposalEncryptionAlgorithmValue) {
	o.EncryptionAlgorithm = v
}

// GetAuthenticationAlgorithm returns the AuthenticationAlgorithm field value if set, zero value otherwise.
func (o *IKEProposalRequest) GetAuthenticationAlgorithm() IKEProposalAuthenticationAlgorithmValue {
	if o == nil || IsNil(o.AuthenticationAlgorithm) {
		var ret IKEProposalAuthenticationAlgorithmValue
		return ret
	}
	return *o.AuthenticationAlgorithm
}

// GetAuthenticationAlgorithmOk returns a tuple with the AuthenticationAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEProposalRequest) GetAuthenticationAlgorithmOk() (*IKEProposalAuthenticationAlgorithmValue, bool) {
	if o == nil || IsNil(o.AuthenticationAlgorithm) {
		return nil, false
	}
	return o.AuthenticationAlgorithm, true
}

// HasAuthenticationAlgorithm returns a boolean if a field has been set.
func (o *IKEProposalRequest) HasAuthenticationAlgorithm() bool {
	if o != nil && !IsNil(o.AuthenticationAlgorithm) {
		return true
	}

	return false
}

// SetAuthenticationAlgorithm gets a reference to the given IKEProposalAuthenticationAlgorithmValue and assigns it to the AuthenticationAlgorithm field.
func (o *IKEProposalRequest) SetAuthenticationAlgorithm(v IKEProposalAuthenticationAlgorithmValue) {
	o.AuthenticationAlgorithm = &v
}

// GetGroup returns the Group field value
func (o *IKEProposalRequest) GetGroup() IKEProposalGroupValue {
	if o == nil {
		var ret IKEProposalGroupValue
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *IKEProposalRequest) GetGroupOk() (*IKEProposalGroupValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *IKEProposalRequest) SetGroup(v IKEProposalGroupValue) {
	o.Group = v
}

// GetSaLifetime returns the SaLifetime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IKEProposalRequest) GetSaLifetime() int32 {
	if o == nil || IsNil(o.SaLifetime.Get()) {
		var ret int32
		return ret
	}
	return *o.SaLifetime.Get()
}

// GetSaLifetimeOk returns a tuple with the SaLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IKEProposalRequest) GetSaLifetimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaLifetime.Get(), o.SaLifetime.IsSet()
}

// HasSaLifetime returns a boolean if a field has been set.
func (o *IKEProposalRequest) HasSaLifetime() bool {
	if o != nil && o.SaLifetime.IsSet() {
		return true
	}

	return false
}

// SetSaLifetime gets a reference to the given NullableInt32 and assigns it to the SaLifetime field.
func (o *IKEProposalRequest) SetSaLifetime(v int32) {
	o.SaLifetime.Set(&v)
}

// SetSaLifetimeNil sets the value for SaLifetime to be an explicit nil
func (o *IKEProposalRequest) SetSaLifetimeNil() {
	o.SaLifetime.Set(nil)
}

// UnsetSaLifetime ensures that no value is present for SaLifetime, not even an explicit nil
func (o *IKEProposalRequest) UnsetSaLifetime() {
	o.SaLifetime.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *IKEProposalRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEProposalRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *IKEProposalRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *IKEProposalRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IKEProposalRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEProposalRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IKEProposalRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *IKEProposalRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *IKEProposalRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEProposalRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *IKEProposalRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *IKEProposalRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o IKEProposalRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IKEProposalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["authentication_method"] = o.AuthenticationMethod
	toSerialize["encryption_algorithm"] = o.EncryptionAlgorithm
	if !IsNil(o.AuthenticationAlgorithm) {
		toSerialize["authentication_algorithm"] = o.AuthenticationAlgorithm
	}
	toSerialize["group"] = o.Group
	if o.SaLifetime.IsSet() {
		toSerialize["sa_lifetime"] = o.SaLifetime.Get()
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IKEProposalRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"authentication_method",
		"encryption_algorithm",
		"group",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	varIKEProposalRequest := _IKEProposalRequest{}

	err = json.Unmarshal(data, &varIKEProposalRequest)

	if err != nil {
		return err
	}

	*o = IKEProposalRequest(varIKEProposalRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "authentication_method")
		delete(additionalProperties, "encryption_algorithm")
		delete(additionalProperties, "authentication_algorithm")
		delete(additionalProperties, "group")
		delete(additionalProperties, "sa_lifetime")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIKEProposalRequest struct {
	value *IKEProposalRequest
	isSet bool
}

func (v NullableIKEProposalRequest) Get() *IKEProposalRequest {
	return v.value
}

func (v *NullableIKEProposalRequest) Set(val *IKEProposalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIKEProposalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIKEProposalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIKEProposalRequest(val *IKEProposalRequest) *NullableIKEProposalRequest {
	return &NullableIKEProposalRequest{value: val, isSet: true}
}

func (v NullableIKEProposalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIKEProposalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
