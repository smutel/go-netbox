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

// checks if the ConsoleServerPortRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsoleServerPortRequest{}

// ConsoleServerPortRequest Adds support for custom fields and tags.
type ConsoleServerPortRequest struct {
	Device BriefDeviceRequest         `json:"device"`
	Module NullableBriefModuleRequest `json:"module,omitempty"`
	Name   string                     `json:"name"`
	// Physical label
	Label       *string                         `json:"label,omitempty"`
	Type        *ConsolePortTypeValue           `json:"type,omitempty"`
	Speed       NullableConsolePortRequestSpeed `json:"speed,omitempty"`
	Description *string                         `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected        *bool                  `json:"mark_connected,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConsoleServerPortRequest ConsoleServerPortRequest

// NewConsoleServerPortRequest instantiates a new ConsoleServerPortRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsoleServerPortRequest(device BriefDeviceRequest, name string) *ConsoleServerPortRequest {
	this := ConsoleServerPortRequest{}
	this.Device = device
	this.Name = name
	return &this
}

// NewConsoleServerPortRequestWithDefaults instantiates a new ConsoleServerPortRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsoleServerPortRequestWithDefaults() *ConsoleServerPortRequest {
	this := ConsoleServerPortRequest{}
	return &this
}

// GetDevice returns the Device field value
func (o *ConsoleServerPortRequest) GetDevice() BriefDeviceRequest {
	if o == nil {
		var ret BriefDeviceRequest
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *ConsoleServerPortRequest) GetDeviceOk() (*BriefDeviceRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *ConsoleServerPortRequest) SetDevice(v BriefDeviceRequest) {
	o.Device = v
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsoleServerPortRequest) GetModule() BriefModuleRequest {
	if o == nil || IsNil(o.Module.Get()) {
		var ret BriefModuleRequest
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsoleServerPortRequest) GetModuleOk() (*BriefModuleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *ConsoleServerPortRequest) HasModule() bool {
	if o != nil && o.Module.IsSet() {
		return true
	}

	return false
}

// SetModule gets a reference to the given NullableBriefModuleRequest and assigns it to the Module field.
func (o *ConsoleServerPortRequest) SetModule(v BriefModuleRequest) {
	o.Module.Set(&v)
}

// SetModuleNil sets the value for Module to be an explicit nil
func (o *ConsoleServerPortRequest) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil
func (o *ConsoleServerPortRequest) UnsetModule() {
	o.Module.Unset()
}

// GetName returns the Name field value
func (o *ConsoleServerPortRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConsoleServerPortRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConsoleServerPortRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ConsoleServerPortRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleServerPortRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ConsoleServerPortRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ConsoleServerPortRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConsoleServerPortRequest) GetType() ConsolePortTypeValue {
	if o == nil || IsNil(o.Type) {
		var ret ConsolePortTypeValue
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleServerPortRequest) GetTypeOk() (*ConsolePortTypeValue, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConsoleServerPortRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ConsolePortTypeValue and assigns it to the Type field.
func (o *ConsoleServerPortRequest) SetType(v ConsolePortTypeValue) {
	o.Type = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsoleServerPortRequest) GetSpeed() ConsolePortRequestSpeed {
	if o == nil || IsNil(o.Speed.Get()) {
		var ret ConsolePortRequestSpeed
		return ret
	}
	return *o.Speed.Get()
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsoleServerPortRequest) GetSpeedOk() (*ConsolePortRequestSpeed, bool) {
	if o == nil {
		return nil, false
	}
	return o.Speed.Get(), o.Speed.IsSet()
}

// HasSpeed returns a boolean if a field has been set.
func (o *ConsoleServerPortRequest) HasSpeed() bool {
	if o != nil && o.Speed.IsSet() {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given NullableConsolePortRequestSpeed and assigns it to the Speed field.
func (o *ConsoleServerPortRequest) SetSpeed(v ConsolePortRequestSpeed) {
	o.Speed.Set(&v)
}

// SetSpeedNil sets the value for Speed to be an explicit nil
func (o *ConsoleServerPortRequest) SetSpeedNil() {
	o.Speed.Set(nil)
}

// UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
func (o *ConsoleServerPortRequest) UnsetSpeed() {
	o.Speed.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConsoleServerPortRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleServerPortRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConsoleServerPortRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConsoleServerPortRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *ConsoleServerPortRequest) GetMarkConnected() bool {
	if o == nil || IsNil(o.MarkConnected) {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleServerPortRequest) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkConnected) {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *ConsoleServerPortRequest) HasMarkConnected() bool {
	if o != nil && !IsNil(o.MarkConnected) {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *ConsoleServerPortRequest) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ConsoleServerPortRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleServerPortRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ConsoleServerPortRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *ConsoleServerPortRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ConsoleServerPortRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleServerPortRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ConsoleServerPortRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ConsoleServerPortRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ConsoleServerPortRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsoleServerPortRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device"] = o.Device
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Speed.IsSet() {
		toSerialize["speed"] = o.Speed.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MarkConnected) {
		toSerialize["mark_connected"] = o.MarkConnected
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

func (o *ConsoleServerPortRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"device",
		"name",
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
	varConsoleServerPortRequest := _ConsoleServerPortRequest{}

	err = json.Unmarshal(data, &varConsoleServerPortRequest)

	if err != nil {
		return err
	}

	*o = ConsoleServerPortRequest(varConsoleServerPortRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mark_connected")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConsoleServerPortRequest struct {
	value *ConsoleServerPortRequest
	isSet bool
}

func (v NullableConsoleServerPortRequest) Get() *ConsoleServerPortRequest {
	return v.value
}

func (v *NullableConsoleServerPortRequest) Set(val *ConsoleServerPortRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConsoleServerPortRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConsoleServerPortRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsoleServerPortRequest(val *ConsoleServerPortRequest) *NullableConsoleServerPortRequest {
	return &NullableConsoleServerPortRequest{value: val, isSet: true}
}

func (v NullableConsoleServerPortRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsoleServerPortRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
