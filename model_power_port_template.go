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
	"time"
)

// checks if the PowerPortTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerPortTemplate{}

// PowerPortTemplate Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PowerPortTemplate struct {
	Id         int32                   `json:"id"`
	Url        string                  `json:"url"`
	Display    string                  `json:"display"`
	DeviceType NullableBriefDeviceType `json:"device_type,omitempty"`
	ModuleType NullableBriefModuleType `json:"module_type,omitempty"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name"`
	// Physical label
	Label *string               `json:"label,omitempty"`
	Type  NullablePowerPortType `json:"type,omitempty"`
	// Maximum power draw (watts)
	MaximumDraw NullableInt32 `json:"maximum_draw,omitempty"`
	// Allocated power draw (watts)
	AllocatedDraw        NullableInt32 `json:"allocated_draw,omitempty"`
	Description          *string       `json:"description,omitempty"`
	Created              NullableTime  `json:"created"`
	LastUpdated          NullableTime  `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _PowerPortTemplate PowerPortTemplate

// NewPowerPortTemplate instantiates a new PowerPortTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerPortTemplate(id int32, url string, display string, name string, created NullableTime, lastUpdated NullableTime) *PowerPortTemplate {
	this := PowerPortTemplate{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewPowerPortTemplateWithDefaults instantiates a new PowerPortTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerPortTemplateWithDefaults() *PowerPortTemplate {
	this := PowerPortTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *PowerPortTemplate) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PowerPortTemplate) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PowerPortTemplate) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *PowerPortTemplate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PowerPortTemplate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PowerPortTemplate) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *PowerPortTemplate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *PowerPortTemplate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *PowerPortTemplate) SetDisplay(v string) {
	o.Display = v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerPortTemplate) GetDeviceType() BriefDeviceType {
	if o == nil || IsNil(o.DeviceType.Get()) {
		var ret BriefDeviceType
		return ret
	}
	return *o.DeviceType.Get()
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerPortTemplate) GetDeviceTypeOk() (*BriefDeviceType, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceType.Get(), o.DeviceType.IsSet()
}

// HasDeviceType returns a boolean if a field has been set.
func (o *PowerPortTemplate) HasDeviceType() bool {
	if o != nil && o.DeviceType.IsSet() {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given NullableBriefDeviceType and assigns it to the DeviceType field.
func (o *PowerPortTemplate) SetDeviceType(v BriefDeviceType) {
	o.DeviceType.Set(&v)
}

// SetDeviceTypeNil sets the value for DeviceType to be an explicit nil
func (o *PowerPortTemplate) SetDeviceTypeNil() {
	o.DeviceType.Set(nil)
}

// UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
func (o *PowerPortTemplate) UnsetDeviceType() {
	o.DeviceType.Unset()
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerPortTemplate) GetModuleType() BriefModuleType {
	if o == nil || IsNil(o.ModuleType.Get()) {
		var ret BriefModuleType
		return ret
	}
	return *o.ModuleType.Get()
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerPortTemplate) GetModuleTypeOk() (*BriefModuleType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleType.Get(), o.ModuleType.IsSet()
}

// HasModuleType returns a boolean if a field has been set.
func (o *PowerPortTemplate) HasModuleType() bool {
	if o != nil && o.ModuleType.IsSet() {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given NullableBriefModuleType and assigns it to the ModuleType field.
func (o *PowerPortTemplate) SetModuleType(v BriefModuleType) {
	o.ModuleType.Set(&v)
}

// SetModuleTypeNil sets the value for ModuleType to be an explicit nil
func (o *PowerPortTemplate) SetModuleTypeNil() {
	o.ModuleType.Set(nil)
}

// UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
func (o *PowerPortTemplate) UnsetModuleType() {
	o.ModuleType.Unset()
}

// GetName returns the Name field value
func (o *PowerPortTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PowerPortTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PowerPortTemplate) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PowerPortTemplate) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPortTemplate) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PowerPortTemplate) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PowerPortTemplate) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerPortTemplate) GetType() PowerPortType {
	if o == nil || IsNil(o.Type.Get()) {
		var ret PowerPortType
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerPortTemplate) GetTypeOk() (*PowerPortType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *PowerPortTemplate) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullablePowerPortType and assigns it to the Type field.
func (o *PowerPortTemplate) SetType(v PowerPortType) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *PowerPortTemplate) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *PowerPortTemplate) UnsetType() {
	o.Type.Unset()
}

// GetMaximumDraw returns the MaximumDraw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerPortTemplate) GetMaximumDraw() int32 {
	if o == nil || IsNil(o.MaximumDraw.Get()) {
		var ret int32
		return ret
	}
	return *o.MaximumDraw.Get()
}

// GetMaximumDrawOk returns a tuple with the MaximumDraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerPortTemplate) GetMaximumDrawOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaximumDraw.Get(), o.MaximumDraw.IsSet()
}

// HasMaximumDraw returns a boolean if a field has been set.
func (o *PowerPortTemplate) HasMaximumDraw() bool {
	if o != nil && o.MaximumDraw.IsSet() {
		return true
	}

	return false
}

// SetMaximumDraw gets a reference to the given NullableInt32 and assigns it to the MaximumDraw field.
func (o *PowerPortTemplate) SetMaximumDraw(v int32) {
	o.MaximumDraw.Set(&v)
}

// SetMaximumDrawNil sets the value for MaximumDraw to be an explicit nil
func (o *PowerPortTemplate) SetMaximumDrawNil() {
	o.MaximumDraw.Set(nil)
}

// UnsetMaximumDraw ensures that no value is present for MaximumDraw, not even an explicit nil
func (o *PowerPortTemplate) UnsetMaximumDraw() {
	o.MaximumDraw.Unset()
}

// GetAllocatedDraw returns the AllocatedDraw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerPortTemplate) GetAllocatedDraw() int32 {
	if o == nil || IsNil(o.AllocatedDraw.Get()) {
		var ret int32
		return ret
	}
	return *o.AllocatedDraw.Get()
}

// GetAllocatedDrawOk returns a tuple with the AllocatedDraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerPortTemplate) GetAllocatedDrawOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllocatedDraw.Get(), o.AllocatedDraw.IsSet()
}

// HasAllocatedDraw returns a boolean if a field has been set.
func (o *PowerPortTemplate) HasAllocatedDraw() bool {
	if o != nil && o.AllocatedDraw.IsSet() {
		return true
	}

	return false
}

// SetAllocatedDraw gets a reference to the given NullableInt32 and assigns it to the AllocatedDraw field.
func (o *PowerPortTemplate) SetAllocatedDraw(v int32) {
	o.AllocatedDraw.Set(&v)
}

// SetAllocatedDrawNil sets the value for AllocatedDraw to be an explicit nil
func (o *PowerPortTemplate) SetAllocatedDrawNil() {
	o.AllocatedDraw.Set(nil)
}

// UnsetAllocatedDraw ensures that no value is present for AllocatedDraw, not even an explicit nil
func (o *PowerPortTemplate) UnsetAllocatedDraw() {
	o.AllocatedDraw.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PowerPortTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPortTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PowerPortTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PowerPortTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *PowerPortTemplate) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerPortTemplate) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *PowerPortTemplate) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *PowerPortTemplate) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerPortTemplate) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *PowerPortTemplate) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o PowerPortTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerPortTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	if o.DeviceType.IsSet() {
		toSerialize["device_type"] = o.DeviceType.Get()
	}
	if o.ModuleType.IsSet() {
		toSerialize["module_type"] = o.ModuleType.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.MaximumDraw.IsSet() {
		toSerialize["maximum_draw"] = o.MaximumDraw.Get()
	}
	if o.AllocatedDraw.IsSet() {
		toSerialize["allocated_draw"] = o.AllocatedDraw.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PowerPortTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"created",
		"last_updated",
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
	varPowerPortTemplate := _PowerPortTemplate{}

	err = json.Unmarshal(data, &varPowerPortTemplate)

	if err != nil {
		return err
	}

	*o = PowerPortTemplate(varPowerPortTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "maximum_draw")
		delete(additionalProperties, "allocated_draw")
		delete(additionalProperties, "description")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePowerPortTemplate struct {
	value *PowerPortTemplate
	isSet bool
}

func (v NullablePowerPortTemplate) Get() *PowerPortTemplate {
	return v.value
}

func (v *NullablePowerPortTemplate) Set(val *PowerPortTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerPortTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerPortTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerPortTemplate(val *PowerPortTemplate) *NullablePowerPortTemplate {
	return &NullablePowerPortTemplate{value: val, isSet: true}
}

func (v NullablePowerPortTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerPortTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
