# BulkDeviceRequestPrimaryIp4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**NatInside** | Pointer to [**NullableNestedIPAddressRequest**](NestedIPAddressRequest.md) |  | [optional] 
**DnsName** | Pointer to [**BriefIPAddressDnsName**](BriefIPAddressDnsName.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkDeviceRequestPrimaryIp4

`func NewBulkDeviceRequestPrimaryIp4(address string, ) *BulkDeviceRequestPrimaryIp4`

NewBulkDeviceRequestPrimaryIp4 instantiates a new BulkDeviceRequestPrimaryIp4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDeviceRequestPrimaryIp4WithDefaults

`func NewBulkDeviceRequestPrimaryIp4WithDefaults() *BulkDeviceRequestPrimaryIp4`

NewBulkDeviceRequestPrimaryIp4WithDefaults instantiates a new BulkDeviceRequestPrimaryIp4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *BulkDeviceRequestPrimaryIp4) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BulkDeviceRequestPrimaryIp4) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BulkDeviceRequestPrimaryIp4) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetNatInside

`func (o *BulkDeviceRequestPrimaryIp4) GetNatInside() NestedIPAddressRequest`

GetNatInside returns the NatInside field if non-nil, zero value otherwise.

### GetNatInsideOk

`func (o *BulkDeviceRequestPrimaryIp4) GetNatInsideOk() (*NestedIPAddressRequest, bool)`

GetNatInsideOk returns a tuple with the NatInside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatInside

`func (o *BulkDeviceRequestPrimaryIp4) SetNatInside(v NestedIPAddressRequest)`

SetNatInside sets NatInside field to given value.

### HasNatInside

`func (o *BulkDeviceRequestPrimaryIp4) HasNatInside() bool`

HasNatInside returns a boolean if a field has been set.

### SetNatInsideNil

`func (o *BulkDeviceRequestPrimaryIp4) SetNatInsideNil(b bool)`

 SetNatInsideNil sets the value for NatInside to be an explicit nil

### UnsetNatInside
`func (o *BulkDeviceRequestPrimaryIp4) UnsetNatInside()`

UnsetNatInside ensures that no value is present for NatInside, not even an explicit nil
### GetDnsName

`func (o *BulkDeviceRequestPrimaryIp4) GetDnsName() BriefIPAddressDnsName`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *BulkDeviceRequestPrimaryIp4) GetDnsNameOk() (*BriefIPAddressDnsName, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *BulkDeviceRequestPrimaryIp4) SetDnsName(v BriefIPAddressDnsName)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *BulkDeviceRequestPrimaryIp4) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDescription

`func (o *BulkDeviceRequestPrimaryIp4) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkDeviceRequestPrimaryIp4) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkDeviceRequestPrimaryIp4) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkDeviceRequestPrimaryIp4) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


