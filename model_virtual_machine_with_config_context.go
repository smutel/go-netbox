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

// checks if the VirtualMachineWithConfigContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualMachineWithConfigContext{}

// VirtualMachineWithConfigContext Adds support for custom fields and tags.
type VirtualMachineWithConfigContext struct {
	Id             int32                       `json:"id"`
	Url            string                      `json:"url"`
	Display        string                      `json:"display"`
	Name           string                      `json:"name"`
	Status         *ModuleStatus               `json:"status,omitempty"`
	Site           NullableBriefSite           `json:"site,omitempty"`
	Cluster        NullableBriefCluster        `json:"cluster,omitempty"`
	Device         NullableBriefDevice         `json:"device,omitempty"`
	Role           NullableBriefDeviceRole     `json:"role,omitempty"`
	Tenant         NullableBriefTenant         `json:"tenant,omitempty"`
	Platform       NullableBriefPlatform       `json:"platform,omitempty"`
	PrimaryIp      NullableBriefIPAddress      `json:"primary_ip"`
	PrimaryIp4     NullableBriefIPAddress      `json:"primary_ip4,omitempty"`
	PrimaryIp6     NullableBriefIPAddress      `json:"primary_ip6,omitempty"`
	Vcpus          NullableFloat64             `json:"vcpus,omitempty"`
	Memory         NullableInt32               `json:"memory,omitempty"`
	Disk           NullableInt32               `json:"disk,omitempty"`
	Description    *string                     `json:"description,omitempty"`
	Comments       *string                     `json:"comments,omitempty"`
	ConfigTemplate NullableBriefConfigTemplate `json:"config_template,omitempty"`
	// Local config context data takes precedence over source contexts in the final rendered config context
	LocalContextData     interface{}            `json:"local_context_data,omitempty"`
	Tags                 []NestedTag            `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	ConfigContext        interface{}            `json:"config_context"`
	Created              NullableTime           `json:"created"`
	LastUpdated          NullableTime           `json:"last_updated"`
	InterfaceCount       int32                  `json:"interface_count"`
	VirtualDiskCount     int32                  `json:"virtual_disk_count"`
	AdditionalProperties map[string]interface{}
}

type _VirtualMachineWithConfigContext VirtualMachineWithConfigContext

// NewVirtualMachineWithConfigContext instantiates a new VirtualMachineWithConfigContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualMachineWithConfigContext(id int32, url string, display string, name string, primaryIp NullableBriefIPAddress, configContext interface{}, created NullableTime, lastUpdated NullableTime, interfaceCount int32, virtualDiskCount int32) *VirtualMachineWithConfigContext {
	this := VirtualMachineWithConfigContext{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.PrimaryIp = primaryIp
	this.ConfigContext = configContext
	this.Created = created
	this.LastUpdated = lastUpdated
	this.InterfaceCount = interfaceCount
	this.VirtualDiskCount = virtualDiskCount
	return &this
}

// NewVirtualMachineWithConfigContextWithDefaults instantiates a new VirtualMachineWithConfigContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualMachineWithConfigContextWithDefaults() *VirtualMachineWithConfigContext {
	this := VirtualMachineWithConfigContext{}
	return &this
}

// GetId returns the Id field value
func (o *VirtualMachineWithConfigContext) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VirtualMachineWithConfigContext) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VirtualMachineWithConfigContext) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *VirtualMachineWithConfigContext) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VirtualMachineWithConfigContext) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VirtualMachineWithConfigContext) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *VirtualMachineWithConfigContext) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *VirtualMachineWithConfigContext) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *VirtualMachineWithConfigContext) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *VirtualMachineWithConfigContext) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VirtualMachineWithConfigContext) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VirtualMachineWithConfigContext) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VirtualMachineWithConfigContext) GetStatus() ModuleStatus {
	if o == nil || IsNil(o.Status) {
		var ret ModuleStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineWithConfigContext) GetStatusOk() (*ModuleStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContext) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ModuleStatus and assigns it to the Status field.
func (o *VirtualMachineWithConfigContext) SetStatus(v ModuleStatus) {
	o.Status = &v
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContext) GetSite() BriefSite {
	if o == nil || IsNil(o.Site.Get()) {
		var ret BriefSite
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContext) GetSiteOk() (*BriefSite, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContext) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableBriefSite and assigns it to the Site field.
func (o *VirtualMachineWithConfigContext) SetSite(v BriefSite) {
	o.Site.Set(&v)
}

// SetSiteNil sets the value for Site to be an explicit nil
func (o *VirtualMachineWithConfigContext) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *VirtualMachineWithConfigContext) UnsetSite() {
	o.Site.Unset()
}

// GetCluster returns the Cluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContext) GetCluster() BriefCluster {
	if o == nil || IsNil(o.Cluster.Get()) {
		var ret BriefCluster
		return ret
	}
	return *o.Cluster.Get()
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContext) GetClusterOk() (*BriefCluster, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cluster.Get(), o.Cluster.IsSet()
}

// HasCluster returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContext) HasCluster() bool {
	if o != nil && o.Cluster.IsSet() {
		return true
	}

	return false
}

// SetCluster gets a reference to the given NullableBriefCluster and assigns it to the Cluster field.
func (o *VirtualMachineWithConfigContext) SetCluster(v BriefCluster) {
	o.Cluster.Set(&v)
}

// SetClusterNil sets the value for Cluster to be an explicit nil
func (o *VirtualMachineWithConfigContext) SetClusterNil() {
	o.Cluster.Set(nil)
}

// UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
func (o *VirtualMachineWithConfigContext) UnsetCluster() {
	o.Cluster.Unset()
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContext) GetDevice() BriefDevice {
	if o == nil || IsNil(o.Device.Get()) {
		var ret BriefDevice
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContext) GetDeviceOk() (*BriefDevice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContext) HasDevice() bool {
	if o != nil && o.Device.IsSet() {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableBriefDevice and assigns it to the Device field.
func (o *VirtualMachineWithConfigContext) SetDevice(v BriefDevice) {
	o.Device.Set(&v)
}

// SetDeviceNil sets the value for Device to be an explicit nil
func (o *VirtualMachineWithConfigContext) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *VirtualMachineWithConfigContext) UnsetDevice() {
	o.Device.Unset()
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContext) GetRole() BriefDeviceRole {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BriefDeviceRole
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContext) GetRoleOk() (*BriefDeviceRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContext) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBriefDeviceRole and assigns it to the Role field.
func (o *VirtualMachineWithConfigContext) SetRole(v BriefDeviceRole) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil
func (o *VirtualMachineWithConfigContext) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *VirtualMachineWithConfigContext) UnsetRole() {
	o.Role.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContext) GetTenant() BriefTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContext) GetTenantOk() (*BriefTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContext) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenant and assigns it to the Tenant field.
func (o *VirtualMachineWithConfigContext) SetTenant(v BriefTenant) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *VirtualMachineWithConfigContext) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *VirtualMachineWithConfigContext) UnsetTenant() {
	o.Tenant.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContext) GetPlatform() BriefPlatform {
	if o == nil || IsNil(o.Platform.Get()) {
		var ret BriefPlatform
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContext) GetPlatformOk() (*BriefPlatform, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContext) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableBriefPlatform and assigns it to the Platform field.
func (o *VirtualMachineWithConfigContext) SetPlatform(v BriefPlatform) {
	o.Platform.Set(&v)
}

// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *VirtualMachineWithConfigContext) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *VirtualMachineWithConfigContext) UnsetPlatform() {
	o.Platform.Unset()
}

// GetPrimaryIp returns the PrimaryIp field value
// If the value is explicit nil, the zero value for BriefIPAddress will be returned
func (o *VirtualMachineWithConfigContext) GetPrimaryIp() BriefIPAddress {
	if o == nil || o.PrimaryIp.Get() == nil {
		var ret BriefIPAddress
		return ret
	}

	return *o.PrimaryIp.Get()
}

// GetPrimaryIpOk returns a tuple with the PrimaryIp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContext) GetPrimaryIpOk() (*BriefIPAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp.Get(), o.PrimaryIp.IsSet()
}

// SetPrimaryIp sets field value
func (o *VirtualMachineWithConfigContext) SetPrimaryIp(v BriefIPAddress) {
	o.PrimaryIp.Set(&v)
}

// GetPrimaryIp4 returns the PrimaryIp4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContext) GetPrimaryIp4() BriefIPAddress {
	if o == nil || IsNil(o.PrimaryIp4.Get()) {
		var ret BriefIPAddress
		return ret
	}
	return *o.PrimaryIp4.Get()
}

// GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContext) GetPrimaryIp4Ok() (*BriefIPAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp4.Get(), o.PrimaryIp4.IsSet()
}

// HasPrimaryIp4 returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContext) HasPrimaryIp4() bool {
	if o != nil && o.PrimaryIp4.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp4 gets a reference to the given NullableBriefIPAddress and assigns it to the PrimaryIp4 field.
func (o *VirtualMachineWithConfigContext) SetPrimaryIp4(v BriefIPAddress) {
	o.PrimaryIp4.Set(&v)
}

// SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil
func (o *VirtualMachineWithConfigContext) SetPrimaryIp4Nil() {
	o.PrimaryIp4.Set(nil)
}

// UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
func (o *VirtualMachineWithConfigContext) UnsetPrimaryIp4() {
	o.PrimaryIp4.Unset()
}

// GetPrimaryIp6 returns the PrimaryIp6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContext) GetPrimaryIp6() BriefIPAddress {
	if o == nil || IsNil(o.PrimaryIp6.Get()) {
		var ret BriefIPAddress
		return ret
	}
	return *o.PrimaryIp6.Get()
}

// GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContext) GetPrimaryIp6Ok() (*BriefIPAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp6.Get(), o.PrimaryIp6.IsSet()
}

// HasPrimaryIp6 returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContext) HasPrimaryIp6() bool {
	if o != nil && o.PrimaryIp6.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp6 gets a reference to the given NullableBriefIPAddress and assigns it to the PrimaryIp6 field.
func (o *VirtualMachineWithConfigContext) SetPrimaryIp6(v BriefIPAddress) {
	o.PrimaryIp6.Set(&v)
}

// SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil
func (o *VirtualMachineWithConfigContext) SetPrimaryIp6Nil() {
	o.PrimaryIp6.Set(nil)
}

// UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
func (o *VirtualMachineWithConfigContext) UnsetPrimaryIp6() {
	o.PrimaryIp6.Unset()
}

// GetVcpus returns the Vcpus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContext) GetVcpus() float64 {
	if o == nil || IsNil(o.Vcpus.Get()) {
		var ret float64
		return ret
	}
	return *o.Vcpus.Get()
}

// GetVcpusOk returns a tuple with the Vcpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContext) GetVcpusOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vcpus.Get(), o.Vcpus.IsSet()
}

// HasVcpus returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContext) HasVcpus() bool {
	if o != nil && o.Vcpus.IsSet() {
		return true
	}

	return false
}

// SetVcpus gets a reference to the given NullableFloat64 and assigns it to the Vcpus field.
func (o *VirtualMachineWithConfigContext) SetVcpus(v float64) {
	o.Vcpus.Set(&v)
}

// SetVcpusNil sets the value for Vcpus to be an explicit nil
func (o *VirtualMachineWithConfigContext) SetVcpusNil() {
	o.Vcpus.Set(nil)
}

// UnsetVcpus ensures that no value is present for Vcpus, not even an explicit nil
func (o *VirtualMachineWithConfigContext) UnsetVcpus() {
	o.Vcpus.Unset()
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContext) GetMemory() int32 {
	if o == nil || IsNil(o.Memory.Get()) {
		var ret int32
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContext) GetMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContext) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableInt32 and assigns it to the Memory field.
func (o *VirtualMachineWithConfigContext) SetMemory(v int32) {
	o.Memory.Set(&v)
}

// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *VirtualMachineWithConfigContext) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *VirtualMachineWithConfigContext) UnsetMemory() {
	o.Memory.Unset()
}

// GetDisk returns the Disk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContext) GetDisk() int32 {
	if o == nil || IsNil(o.Disk.Get()) {
		var ret int32
		return ret
	}
	return *o.Disk.Get()
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContext) GetDiskOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disk.Get(), o.Disk.IsSet()
}

// HasDisk returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContext) HasDisk() bool {
	if o != nil && o.Disk.IsSet() {
		return true
	}

	return false
}

// SetDisk gets a reference to the given NullableInt32 and assigns it to the Disk field.
func (o *VirtualMachineWithConfigContext) SetDisk(v int32) {
	o.Disk.Set(&v)
}

// SetDiskNil sets the value for Disk to be an explicit nil
func (o *VirtualMachineWithConfigContext) SetDiskNil() {
	o.Disk.Set(nil)
}

// UnsetDisk ensures that no value is present for Disk, not even an explicit nil
func (o *VirtualMachineWithConfigContext) UnsetDisk() {
	o.Disk.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VirtualMachineWithConfigContext) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineWithConfigContext) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContext) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VirtualMachineWithConfigContext) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *VirtualMachineWithConfigContext) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineWithConfigContext) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContext) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *VirtualMachineWithConfigContext) SetComments(v string) {
	o.Comments = &v
}

// GetConfigTemplate returns the ConfigTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContext) GetConfigTemplate() BriefConfigTemplate {
	if o == nil || IsNil(o.ConfigTemplate.Get()) {
		var ret BriefConfigTemplate
		return ret
	}
	return *o.ConfigTemplate.Get()
}

// GetConfigTemplateOk returns a tuple with the ConfigTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContext) GetConfigTemplateOk() (*BriefConfigTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigTemplate.Get(), o.ConfigTemplate.IsSet()
}

// HasConfigTemplate returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContext) HasConfigTemplate() bool {
	if o != nil && o.ConfigTemplate.IsSet() {
		return true
	}

	return false
}

// SetConfigTemplate gets a reference to the given NullableBriefConfigTemplate and assigns it to the ConfigTemplate field.
func (o *VirtualMachineWithConfigContext) SetConfigTemplate(v BriefConfigTemplate) {
	o.ConfigTemplate.Set(&v)
}

// SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil
func (o *VirtualMachineWithConfigContext) SetConfigTemplateNil() {
	o.ConfigTemplate.Set(nil)
}

// UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
func (o *VirtualMachineWithConfigContext) UnsetConfigTemplate() {
	o.ConfigTemplate.Unset()
}

// GetLocalContextData returns the LocalContextData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContext) GetLocalContextData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.LocalContextData
}

// GetLocalContextDataOk returns a tuple with the LocalContextData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContext) GetLocalContextDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.LocalContextData) {
		return nil, false
	}
	return &o.LocalContextData, true
}

// HasLocalContextData returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContext) HasLocalContextData() bool {
	if o != nil && !IsNil(o.LocalContextData) {
		return true
	}

	return false
}

// SetLocalContextData gets a reference to the given interface{} and assigns it to the LocalContextData field.
func (o *VirtualMachineWithConfigContext) SetLocalContextData(v interface{}) {
	o.LocalContextData = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VirtualMachineWithConfigContext) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineWithConfigContext) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContext) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *VirtualMachineWithConfigContext) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *VirtualMachineWithConfigContext) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineWithConfigContext) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContext) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *VirtualMachineWithConfigContext) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetConfigContext returns the ConfigContext field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *VirtualMachineWithConfigContext) GetConfigContext() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ConfigContext
}

// GetConfigContextOk returns a tuple with the ConfigContext field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContext) GetConfigContextOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ConfigContext) {
		return nil, false
	}
	return &o.ConfigContext, true
}

// SetConfigContext sets field value
func (o *VirtualMachineWithConfigContext) SetConfigContext(v interface{}) {
	o.ConfigContext = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *VirtualMachineWithConfigContext) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContext) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *VirtualMachineWithConfigContext) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *VirtualMachineWithConfigContext) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContext) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *VirtualMachineWithConfigContext) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetInterfaceCount returns the InterfaceCount field value
func (o *VirtualMachineWithConfigContext) GetInterfaceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InterfaceCount
}

// GetInterfaceCountOk returns a tuple with the InterfaceCount field value
// and a boolean to check if the value has been set.
func (o *VirtualMachineWithConfigContext) GetInterfaceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterfaceCount, true
}

// SetInterfaceCount sets field value
func (o *VirtualMachineWithConfigContext) SetInterfaceCount(v int32) {
	o.InterfaceCount = v
}

// GetVirtualDiskCount returns the VirtualDiskCount field value
func (o *VirtualMachineWithConfigContext) GetVirtualDiskCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VirtualDiskCount
}

// GetVirtualDiskCountOk returns a tuple with the VirtualDiskCount field value
// and a boolean to check if the value has been set.
func (o *VirtualMachineWithConfigContext) GetVirtualDiskCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualDiskCount, true
}

// SetVirtualDiskCount sets field value
func (o *VirtualMachineWithConfigContext) SetVirtualDiskCount(v int32) {
	o.VirtualDiskCount = v
}

func (o VirtualMachineWithConfigContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualMachineWithConfigContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	if o.Cluster.IsSet() {
		toSerialize["cluster"] = o.Cluster.Get()
	}
	if o.Device.IsSet() {
		toSerialize["device"] = o.Device.Get()
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Platform.IsSet() {
		toSerialize["platform"] = o.Platform.Get()
	}
	toSerialize["primary_ip"] = o.PrimaryIp.Get()
	if o.PrimaryIp4.IsSet() {
		toSerialize["primary_ip4"] = o.PrimaryIp4.Get()
	}
	if o.PrimaryIp6.IsSet() {
		toSerialize["primary_ip6"] = o.PrimaryIp6.Get()
	}
	if o.Vcpus.IsSet() {
		toSerialize["vcpus"] = o.Vcpus.Get()
	}
	if o.Memory.IsSet() {
		toSerialize["memory"] = o.Memory.Get()
	}
	if o.Disk.IsSet() {
		toSerialize["disk"] = o.Disk.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if o.ConfigTemplate.IsSet() {
		toSerialize["config_template"] = o.ConfigTemplate.Get()
	}
	if o.LocalContextData != nil {
		toSerialize["local_context_data"] = o.LocalContextData
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if o.ConfigContext != nil {
		toSerialize["config_context"] = o.ConfigContext
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["interface_count"] = o.InterfaceCount
	toSerialize["virtual_disk_count"] = o.VirtualDiskCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualMachineWithConfigContext) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"primary_ip",
		"config_context",
		"created",
		"last_updated",
		"interface_count",
		"virtual_disk_count",
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
	varVirtualMachineWithConfigContext := _VirtualMachineWithConfigContext{}

	err = json.Unmarshal(data, &varVirtualMachineWithConfigContext)

	if err != nil {
		return err
	}

	*o = VirtualMachineWithConfigContext(varVirtualMachineWithConfigContext)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "site")
		delete(additionalProperties, "cluster")
		delete(additionalProperties, "device")
		delete(additionalProperties, "role")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "primary_ip")
		delete(additionalProperties, "primary_ip4")
		delete(additionalProperties, "primary_ip6")
		delete(additionalProperties, "vcpus")
		delete(additionalProperties, "memory")
		delete(additionalProperties, "disk")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "config_template")
		delete(additionalProperties, "local_context_data")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "config_context")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "interface_count")
		delete(additionalProperties, "virtual_disk_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualMachineWithConfigContext struct {
	value *VirtualMachineWithConfigContext
	isSet bool
}

func (v NullableVirtualMachineWithConfigContext) Get() *VirtualMachineWithConfigContext {
	return v.value
}

func (v *NullableVirtualMachineWithConfigContext) Set(val *VirtualMachineWithConfigContext) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMachineWithConfigContext) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMachineWithConfigContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMachineWithConfigContext(val *VirtualMachineWithConfigContext) *NullableVirtualMachineWithConfigContext {
	return &NullableVirtualMachineWithConfigContext{value: val, isSet: true}
}

func (v NullableVirtualMachineWithConfigContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMachineWithConfigContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
