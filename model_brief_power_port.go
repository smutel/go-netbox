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

// checks if the BriefPowerPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefPowerPort{}

// BriefPowerPort Adds support for custom fields and tags.
type BriefPowerPort struct {
	Id                   int32              `json:"id"`
	Url                  string             `json:"url"`
	Display              string             `json:"display"`
	Device               BriefDevice        `json:"device"`
	Name                 string             `json:"name"`
	Description          *string            `json:"description,omitempty"`
	Cable                NullableBriefCable `json:"cable"`
	Occupied             bool               `json:"_occupied"`
	AdditionalProperties map[string]interface{}
}

type _BriefPowerPort BriefPowerPort

// NewBriefPowerPort instantiates a new BriefPowerPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefPowerPort(id int32, url string, display string, device BriefDevice, name string, cable NullableBriefCable, occupied bool) *BriefPowerPort {
	this := BriefPowerPort{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Device = device
	this.Name = name
	this.Cable = cable
	this.Occupied = occupied
	return &this
}

// NewBriefPowerPortWithDefaults instantiates a new BriefPowerPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefPowerPortWithDefaults() *BriefPowerPort {
	this := BriefPowerPort{}
	return &this
}

// GetId returns the Id field value
func (o *BriefPowerPort) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefPowerPort) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefPowerPort) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefPowerPort) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefPowerPort) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefPowerPort) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefPowerPort) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefPowerPort) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefPowerPort) SetDisplay(v string) {
	o.Display = v
}

// GetDevice returns the Device field value
func (o *BriefPowerPort) GetDevice() BriefDevice {
	if o == nil {
		var ret BriefDevice
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *BriefPowerPort) GetDeviceOk() (*BriefDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *BriefPowerPort) SetDevice(v BriefDevice) {
	o.Device = v
}

// GetName returns the Name field value
func (o *BriefPowerPort) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefPowerPort) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefPowerPort) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefPowerPort) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefPowerPort) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefPowerPort) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefPowerPort) SetDescription(v string) {
	o.Description = &v
}

// GetCable returns the Cable field value
// If the value is explicit nil, the zero value for BriefCable will be returned
func (o *BriefPowerPort) GetCable() BriefCable {
	if o == nil || o.Cable.Get() == nil {
		var ret BriefCable
		return ret
	}

	return *o.Cable.Get()
}

// GetCableOk returns a tuple with the Cable field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BriefPowerPort) GetCableOk() (*BriefCable, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cable.Get(), o.Cable.IsSet()
}

// SetCable sets field value
func (o *BriefPowerPort) SetCable(v BriefCable) {
	o.Cable.Set(&v)
}

// GetOccupied returns the Occupied field value
func (o *BriefPowerPort) GetOccupied() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Occupied
}

// GetOccupiedOk returns a tuple with the Occupied field value
// and a boolean to check if the value has been set.
func (o *BriefPowerPort) GetOccupiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Occupied, true
}

// SetOccupied sets field value
func (o *BriefPowerPort) SetOccupied(v bool) {
	o.Occupied = v
}

func (o BriefPowerPort) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefPowerPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["device"] = o.Device
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["cable"] = o.Cable.Get()
	toSerialize["_occupied"] = o.Occupied

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefPowerPort) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"device",
		"name",
		"cable",
		"_occupied",
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
	varBriefPowerPort := _BriefPowerPort{}

	err = json.Unmarshal(data, &varBriefPowerPort)

	if err != nil {
		return err
	}

	*o = BriefPowerPort(varBriefPowerPort)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "device")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "cable")
		delete(additionalProperties, "_occupied")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefPowerPort struct {
	value *BriefPowerPort
	isSet bool
}

func (v NullableBriefPowerPort) Get() *BriefPowerPort {
	return v.value
}

func (v *NullableBriefPowerPort) Set(val *BriefPowerPort) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefPowerPort) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefPowerPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefPowerPort(val *BriefPowerPort) *NullableBriefPowerPort {
	return &NullableBriefPowerPort{value: val, isSet: true}
}

func (v NullableBriefPowerPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefPowerPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
