/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PatchedVLANGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedVLANGroupRequest{}

// PatchedVLANGroupRequest Adds support for custom fields and tags.
type PatchedVLANGroupRequest struct {
	Name      *string        `json:"name,omitempty"`
	Slug      *string        `json:"slug,omitempty" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	ScopeType NullableString `json:"scope_type,omitempty"`
	ScopeId   NullableInt32  `json:"scope_id,omitempty"`
	// Lowest permissible ID of a child VLAN
	MinVid *int32 `json:"min_vid,omitempty"`
	// Highest permissible ID of a child VLAN
	MaxVid               *int32                 `json:"max_vid,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedVLANGroupRequest PatchedVLANGroupRequest

// NewPatchedVLANGroupRequest instantiates a new PatchedVLANGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedVLANGroupRequest() *PatchedVLANGroupRequest {
	this := PatchedVLANGroupRequest{}
	return &this
}

// NewPatchedVLANGroupRequestWithDefaults instantiates a new PatchedVLANGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedVLANGroupRequestWithDefaults() *PatchedVLANGroupRequest {
	this := PatchedVLANGroupRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedVLANGroupRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedVLANGroupRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedVLANGroupRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedVLANGroupRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *PatchedVLANGroupRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedVLANGroupRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PatchedVLANGroupRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PatchedVLANGroupRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetScopeType returns the ScopeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedVLANGroupRequest) GetScopeType() string {
	if o == nil || IsNil(o.ScopeType.Get()) {
		var ret string
		return ret
	}
	return *o.ScopeType.Get()
}

// GetScopeTypeOk returns a tuple with the ScopeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedVLANGroupRequest) GetScopeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScopeType.Get(), o.ScopeType.IsSet()
}

// HasScopeType returns a boolean if a field has been set.
func (o *PatchedVLANGroupRequest) HasScopeType() bool {
	if o != nil && o.ScopeType.IsSet() {
		return true
	}

	return false
}

// SetScopeType gets a reference to the given NullableString and assigns it to the ScopeType field.
func (o *PatchedVLANGroupRequest) SetScopeType(v string) {
	o.ScopeType.Set(&v)
}

// SetScopeTypeNil sets the value for ScopeType to be an explicit nil
func (o *PatchedVLANGroupRequest) SetScopeTypeNil() {
	o.ScopeType.Set(nil)
}

// UnsetScopeType ensures that no value is present for ScopeType, not even an explicit nil
func (o *PatchedVLANGroupRequest) UnsetScopeType() {
	o.ScopeType.Unset()
}

// GetScopeId returns the ScopeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedVLANGroupRequest) GetScopeId() int32 {
	if o == nil || IsNil(o.ScopeId.Get()) {
		var ret int32
		return ret
	}
	return *o.ScopeId.Get()
}

// GetScopeIdOk returns a tuple with the ScopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedVLANGroupRequest) GetScopeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScopeId.Get(), o.ScopeId.IsSet()
}

// HasScopeId returns a boolean if a field has been set.
func (o *PatchedVLANGroupRequest) HasScopeId() bool {
	if o != nil && o.ScopeId.IsSet() {
		return true
	}

	return false
}

// SetScopeId gets a reference to the given NullableInt32 and assigns it to the ScopeId field.
func (o *PatchedVLANGroupRequest) SetScopeId(v int32) {
	o.ScopeId.Set(&v)
}

// SetScopeIdNil sets the value for ScopeId to be an explicit nil
func (o *PatchedVLANGroupRequest) SetScopeIdNil() {
	o.ScopeId.Set(nil)
}

// UnsetScopeId ensures that no value is present for ScopeId, not even an explicit nil
func (o *PatchedVLANGroupRequest) UnsetScopeId() {
	o.ScopeId.Unset()
}

// GetMinVid returns the MinVid field value if set, zero value otherwise.
func (o *PatchedVLANGroupRequest) GetMinVid() int32 {
	if o == nil || IsNil(o.MinVid) {
		var ret int32
		return ret
	}
	return *o.MinVid
}

// GetMinVidOk returns a tuple with the MinVid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedVLANGroupRequest) GetMinVidOk() (*int32, bool) {
	if o == nil || IsNil(o.MinVid) {
		return nil, false
	}
	return o.MinVid, true
}

// HasMinVid returns a boolean if a field has been set.
func (o *PatchedVLANGroupRequest) HasMinVid() bool {
	if o != nil && !IsNil(o.MinVid) {
		return true
	}

	return false
}

// SetMinVid gets a reference to the given int32 and assigns it to the MinVid field.
func (o *PatchedVLANGroupRequest) SetMinVid(v int32) {
	o.MinVid = &v
}

// GetMaxVid returns the MaxVid field value if set, zero value otherwise.
func (o *PatchedVLANGroupRequest) GetMaxVid() int32 {
	if o == nil || IsNil(o.MaxVid) {
		var ret int32
		return ret
	}
	return *o.MaxVid
}

// GetMaxVidOk returns a tuple with the MaxVid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedVLANGroupRequest) GetMaxVidOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxVid) {
		return nil, false
	}
	return o.MaxVid, true
}

// HasMaxVid returns a boolean if a field has been set.
func (o *PatchedVLANGroupRequest) HasMaxVid() bool {
	if o != nil && !IsNil(o.MaxVid) {
		return true
	}

	return false
}

// SetMaxVid gets a reference to the given int32 and assigns it to the MaxVid field.
func (o *PatchedVLANGroupRequest) SetMaxVid(v int32) {
	o.MaxVid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedVLANGroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedVLANGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedVLANGroupRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedVLANGroupRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedVLANGroupRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedVLANGroupRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedVLANGroupRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedVLANGroupRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedVLANGroupRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedVLANGroupRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedVLANGroupRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedVLANGroupRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedVLANGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedVLANGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if o.ScopeType.IsSet() {
		toSerialize["scope_type"] = o.ScopeType.Get()
	}
	if o.ScopeId.IsSet() {
		toSerialize["scope_id"] = o.ScopeId.Get()
	}
	if !IsNil(o.MinVid) {
		toSerialize["min_vid"] = o.MinVid
	}
	if !IsNil(o.MaxVid) {
		toSerialize["max_vid"] = o.MaxVid
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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

func (o *PatchedVLANGroupRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedVLANGroupRequest := _PatchedVLANGroupRequest{}

	err = json.Unmarshal(data, &varPatchedVLANGroupRequest)

	if err != nil {
		return err
	}

	*o = PatchedVLANGroupRequest(varPatchedVLANGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "scope_type")
		delete(additionalProperties, "scope_id")
		delete(additionalProperties, "min_vid")
		delete(additionalProperties, "max_vid")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedVLANGroupRequest struct {
	value *PatchedVLANGroupRequest
	isSet bool
}

func (v NullablePatchedVLANGroupRequest) Get() *PatchedVLANGroupRequest {
	return v.value
}

func (v *NullablePatchedVLANGroupRequest) Set(val *PatchedVLANGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedVLANGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedVLANGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedVLANGroupRequest(val *PatchedVLANGroupRequest) *NullablePatchedVLANGroupRequest {
	return &NullablePatchedVLANGroupRequest{value: val, isSet: true}
}

func (v NullablePatchedVLANGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedVLANGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
