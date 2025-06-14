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

// checks if the PatchedSavedFilterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedSavedFilterRequest{}

// PatchedSavedFilterRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedSavedFilterRequest struct {
	ObjectTypes          []string      `json:"object_types,omitempty"`
	Name                 *string       `json:"name,omitempty"`
	Slug                 *string       `json:"slug,omitempty" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Description          *string       `json:"description,omitempty"`
	User                 NullableInt32 `json:"user,omitempty"`
	Weight               *int32        `json:"weight,omitempty"`
	Enabled              *bool         `json:"enabled,omitempty"`
	Shared               *bool         `json:"shared,omitempty"`
	Parameters           interface{}   `json:"parameters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedSavedFilterRequest PatchedSavedFilterRequest

// NewPatchedSavedFilterRequest instantiates a new PatchedSavedFilterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedSavedFilterRequest() *PatchedSavedFilterRequest {
	this := PatchedSavedFilterRequest{}
	return &this
}

// NewPatchedSavedFilterRequestWithDefaults instantiates a new PatchedSavedFilterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedSavedFilterRequestWithDefaults() *PatchedSavedFilterRequest {
	this := PatchedSavedFilterRequest{}
	return &this
}

// GetObjectTypes returns the ObjectTypes field value if set, zero value otherwise.
func (o *PatchedSavedFilterRequest) GetObjectTypes() []string {
	if o == nil || IsNil(o.ObjectTypes) {
		var ret []string
		return ret
	}
	return o.ObjectTypes
}

// GetObjectTypesOk returns a tuple with the ObjectTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSavedFilterRequest) GetObjectTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ObjectTypes) {
		return nil, false
	}
	return o.ObjectTypes, true
}

// HasObjectTypes returns a boolean if a field has been set.
func (o *PatchedSavedFilterRequest) HasObjectTypes() bool {
	if o != nil && !IsNil(o.ObjectTypes) {
		return true
	}

	return false
}

// SetObjectTypes gets a reference to the given []string and assigns it to the ObjectTypes field.
func (o *PatchedSavedFilterRequest) SetObjectTypes(v []string) {
	o.ObjectTypes = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedSavedFilterRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSavedFilterRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedSavedFilterRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedSavedFilterRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *PatchedSavedFilterRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSavedFilterRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PatchedSavedFilterRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PatchedSavedFilterRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedSavedFilterRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSavedFilterRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedSavedFilterRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedSavedFilterRequest) SetDescription(v string) {
	o.Description = &v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedSavedFilterRequest) GetUser() int32 {
	if o == nil || IsNil(o.User.Get()) {
		var ret int32
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedSavedFilterRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *PatchedSavedFilterRequest) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableInt32 and assigns it to the User field.
func (o *PatchedSavedFilterRequest) SetUser(v int32) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *PatchedSavedFilterRequest) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *PatchedSavedFilterRequest) UnsetUser() {
	o.User.Unset()
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *PatchedSavedFilterRequest) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSavedFilterRequest) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *PatchedSavedFilterRequest) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *PatchedSavedFilterRequest) SetWeight(v int32) {
	o.Weight = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchedSavedFilterRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSavedFilterRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchedSavedFilterRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchedSavedFilterRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *PatchedSavedFilterRequest) GetShared() bool {
	if o == nil || IsNil(o.Shared) {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSavedFilterRequest) GetSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *PatchedSavedFilterRequest) HasShared() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *PatchedSavedFilterRequest) SetShared(v bool) {
	o.Shared = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedSavedFilterRequest) GetParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedSavedFilterRequest) GetParametersOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *PatchedSavedFilterRequest) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given interface{} and assigns it to the Parameters field.
func (o *PatchedSavedFilterRequest) SetParameters(v interface{}) {
	o.Parameters = v
}

func (o PatchedSavedFilterRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedSavedFilterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjectTypes) {
		toSerialize["object_types"] = o.ObjectTypes
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Shared) {
		toSerialize["shared"] = o.Shared
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedSavedFilterRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedSavedFilterRequest := _PatchedSavedFilterRequest{}

	err = json.Unmarshal(data, &varPatchedSavedFilterRequest)

	if err != nil {
		return err
	}

	*o = PatchedSavedFilterRequest(varPatchedSavedFilterRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		delete(additionalProperties, "user")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "shared")
		delete(additionalProperties, "parameters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedSavedFilterRequest struct {
	value *PatchedSavedFilterRequest
	isSet bool
}

func (v NullablePatchedSavedFilterRequest) Get() *PatchedSavedFilterRequest {
	return v.value
}

func (v *NullablePatchedSavedFilterRequest) Set(val *PatchedSavedFilterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedSavedFilterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedSavedFilterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedSavedFilterRequest(val *PatchedSavedFilterRequest) *NullablePatchedSavedFilterRequest {
	return &NullablePatchedSavedFilterRequest{value: val, isSet: true}
}

func (v NullablePatchedSavedFilterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedSavedFilterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
