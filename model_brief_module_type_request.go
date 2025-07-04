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

// checks if the BriefModuleTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefModuleTypeRequest{}

// BriefModuleTypeRequest Adds support for custom fields and tags.
type BriefModuleTypeRequest struct {
	Manufacturer         BriefManufacturerRequest `json:"manufacturer"`
	Model                string                   `json:"model"`
	Description          *string                  `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefModuleTypeRequest BriefModuleTypeRequest

// NewBriefModuleTypeRequest instantiates a new BriefModuleTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefModuleTypeRequest(manufacturer BriefManufacturerRequest, model string) *BriefModuleTypeRequest {
	this := BriefModuleTypeRequest{}
	this.Manufacturer = manufacturer
	this.Model = model
	return &this
}

// NewBriefModuleTypeRequestWithDefaults instantiates a new BriefModuleTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefModuleTypeRequestWithDefaults() *BriefModuleTypeRequest {
	this := BriefModuleTypeRequest{}
	return &this
}

// GetManufacturer returns the Manufacturer field value
func (o *BriefModuleTypeRequest) GetManufacturer() BriefManufacturerRequest {
	if o == nil {
		var ret BriefManufacturerRequest
		return ret
	}

	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value
// and a boolean to check if the value has been set.
func (o *BriefModuleTypeRequest) GetManufacturerOk() (*BriefManufacturerRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manufacturer, true
}

// SetManufacturer sets field value
func (o *BriefModuleTypeRequest) SetManufacturer(v BriefManufacturerRequest) {
	o.Manufacturer = v
}

// GetModel returns the Model field value
func (o *BriefModuleTypeRequest) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *BriefModuleTypeRequest) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *BriefModuleTypeRequest) SetModel(v string) {
	o.Model = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefModuleTypeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefModuleTypeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefModuleTypeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefModuleTypeRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefModuleTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefModuleTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["manufacturer"] = o.Manufacturer
	toSerialize["model"] = o.Model
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefModuleTypeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"manufacturer",
		"model",
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
	varBriefModuleTypeRequest := _BriefModuleTypeRequest{}

	err = json.Unmarshal(data, &varBriefModuleTypeRequest)

	if err != nil {
		return err
	}

	*o = BriefModuleTypeRequest(varBriefModuleTypeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "model")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefModuleTypeRequest struct {
	value *BriefModuleTypeRequest
	isSet bool
}

func (v NullableBriefModuleTypeRequest) Get() *BriefModuleTypeRequest {
	return v.value
}

func (v *NullableBriefModuleTypeRequest) Set(val *BriefModuleTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefModuleTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefModuleTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefModuleTypeRequest(val *BriefModuleTypeRequest) *NullableBriefModuleTypeRequest {
	return &NullableBriefModuleTypeRequest{value: val, isSet: true}
}

func (v NullableBriefModuleTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefModuleTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
