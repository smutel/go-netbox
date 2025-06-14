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

// checks if the PatchedASNRangeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedASNRangeRequest{}

// PatchedASNRangeRequest Adds support for custom fields and tags.
type PatchedASNRangeRequest struct {
	Name                 *string                    `json:"name,omitempty"`
	Slug                 *string                    `json:"slug,omitempty" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Rir                  *BriefRIRRequest           `json:"rir,omitempty"`
	Start                *int64                     `json:"start,omitempty"`
	End                  *int64                     `json:"end,omitempty"`
	Tenant               NullableBriefTenantRequest `json:"tenant,omitempty"`
	Description          *string                    `json:"description,omitempty"`
	Tags                 []NestedTagRequest         `json:"tags,omitempty"`
	CustomFields         map[string]interface{}     `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedASNRangeRequest PatchedASNRangeRequest

// NewPatchedASNRangeRequest instantiates a new PatchedASNRangeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedASNRangeRequest() *PatchedASNRangeRequest {
	this := PatchedASNRangeRequest{}
	return &this
}

// NewPatchedASNRangeRequestWithDefaults instantiates a new PatchedASNRangeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedASNRangeRequestWithDefaults() *PatchedASNRangeRequest {
	this := PatchedASNRangeRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedASNRangeRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedASNRangeRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedASNRangeRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedASNRangeRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *PatchedASNRangeRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedASNRangeRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PatchedASNRangeRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PatchedASNRangeRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetRir returns the Rir field value if set, zero value otherwise.
func (o *PatchedASNRangeRequest) GetRir() BriefRIRRequest {
	if o == nil || IsNil(o.Rir) {
		var ret BriefRIRRequest
		return ret
	}
	return *o.Rir
}

// GetRirOk returns a tuple with the Rir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedASNRangeRequest) GetRirOk() (*BriefRIRRequest, bool) {
	if o == nil || IsNil(o.Rir) {
		return nil, false
	}
	return o.Rir, true
}

// HasRir returns a boolean if a field has been set.
func (o *PatchedASNRangeRequest) HasRir() bool {
	if o != nil && !IsNil(o.Rir) {
		return true
	}

	return false
}

// SetRir gets a reference to the given BriefRIRRequest and assigns it to the Rir field.
func (o *PatchedASNRangeRequest) SetRir(v BriefRIRRequest) {
	o.Rir = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *PatchedASNRangeRequest) GetStart() int64 {
	if o == nil || IsNil(o.Start) {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedASNRangeRequest) GetStartOk() (*int64, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *PatchedASNRangeRequest) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *PatchedASNRangeRequest) SetStart(v int64) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *PatchedASNRangeRequest) GetEnd() int64 {
	if o == nil || IsNil(o.End) {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedASNRangeRequest) GetEndOk() (*int64, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *PatchedASNRangeRequest) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *PatchedASNRangeRequest) SetEnd(v int64) {
	o.End = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedASNRangeRequest) GetTenant() BriefTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedASNRangeRequest) GetTenantOk() (*BriefTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedASNRangeRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenantRequest and assigns it to the Tenant field.
func (o *PatchedASNRangeRequest) SetTenant(v BriefTenantRequest) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedASNRangeRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedASNRangeRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedASNRangeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedASNRangeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedASNRangeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedASNRangeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedASNRangeRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedASNRangeRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedASNRangeRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedASNRangeRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedASNRangeRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedASNRangeRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedASNRangeRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedASNRangeRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedASNRangeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedASNRangeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Rir) {
		toSerialize["rir"] = o.Rir
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
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

func (o *PatchedASNRangeRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedASNRangeRequest := _PatchedASNRangeRequest{}

	err = json.Unmarshal(data, &varPatchedASNRangeRequest)

	if err != nil {
		return err
	}

	*o = PatchedASNRangeRequest(varPatchedASNRangeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "rir")
		delete(additionalProperties, "start")
		delete(additionalProperties, "end")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedASNRangeRequest struct {
	value *PatchedASNRangeRequest
	isSet bool
}

func (v NullablePatchedASNRangeRequest) Get() *PatchedASNRangeRequest {
	return v.value
}

func (v *NullablePatchedASNRangeRequest) Set(val *PatchedASNRangeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedASNRangeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedASNRangeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedASNRangeRequest(val *PatchedASNRangeRequest) *NullablePatchedASNRangeRequest {
	return &NullablePatchedASNRangeRequest{value: val, isSet: true}
}

func (v NullablePatchedASNRangeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedASNRangeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
