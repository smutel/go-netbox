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

// checks if the NestedModuleBayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedModuleBayRequest{}

// NestedModuleBayRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedModuleBayRequest struct {
	InstalledModule      NullableModuleBayNestedModuleRequest `json:"installed_module,omitempty"`
	Name                 string                               `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _NestedModuleBayRequest NestedModuleBayRequest

// NewNestedModuleBayRequest instantiates a new NestedModuleBayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedModuleBayRequest(name string) *NestedModuleBayRequest {
	this := NestedModuleBayRequest{}
	this.Name = name
	return &this
}

// NewNestedModuleBayRequestWithDefaults instantiates a new NestedModuleBayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedModuleBayRequestWithDefaults() *NestedModuleBayRequest {
	this := NestedModuleBayRequest{}
	return &this
}

// GetInstalledModule returns the InstalledModule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NestedModuleBayRequest) GetInstalledModule() ModuleBayNestedModuleRequest {
	if o == nil || IsNil(o.InstalledModule.Get()) {
		var ret ModuleBayNestedModuleRequest
		return ret
	}
	return *o.InstalledModule.Get()
}

// GetInstalledModuleOk returns a tuple with the InstalledModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NestedModuleBayRequest) GetInstalledModuleOk() (*ModuleBayNestedModuleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstalledModule.Get(), o.InstalledModule.IsSet()
}

// HasInstalledModule returns a boolean if a field has been set.
func (o *NestedModuleBayRequest) HasInstalledModule() bool {
	if o != nil && o.InstalledModule.IsSet() {
		return true
	}

	return false
}

// SetInstalledModule gets a reference to the given NullableModuleBayNestedModuleRequest and assigns it to the InstalledModule field.
func (o *NestedModuleBayRequest) SetInstalledModule(v ModuleBayNestedModuleRequest) {
	o.InstalledModule.Set(&v)
}

// SetInstalledModuleNil sets the value for InstalledModule to be an explicit nil
func (o *NestedModuleBayRequest) SetInstalledModuleNil() {
	o.InstalledModule.Set(nil)
}

// UnsetInstalledModule ensures that no value is present for InstalledModule, not even an explicit nil
func (o *NestedModuleBayRequest) UnsetInstalledModule() {
	o.InstalledModule.Unset()
}

// GetName returns the Name field value
func (o *NestedModuleBayRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedModuleBayRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedModuleBayRequest) SetName(v string) {
	o.Name = v
}

func (o NestedModuleBayRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedModuleBayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.InstalledModule.IsSet() {
		toSerialize["installed_module"] = o.InstalledModule.Get()
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedModuleBayRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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
	varNestedModuleBayRequest := _NestedModuleBayRequest{}

	err = json.Unmarshal(data, &varNestedModuleBayRequest)

	if err != nil {
		return err
	}

	*o = NestedModuleBayRequest(varNestedModuleBayRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "installed_module")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedModuleBayRequest struct {
	value *NestedModuleBayRequest
	isSet bool
}

func (v NullableNestedModuleBayRequest) Get() *NestedModuleBayRequest {
	return v.value
}

func (v *NullableNestedModuleBayRequest) Set(val *NestedModuleBayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedModuleBayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedModuleBayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedModuleBayRequest(val *NestedModuleBayRequest) *NullableNestedModuleBayRequest {
	return &NullableNestedModuleBayRequest{value: val, isSet: true}
}

func (v NullableNestedModuleBayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedModuleBayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
