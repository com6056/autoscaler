/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// NatGatewayLanProperties struct for NatGatewayLanProperties
type NatGatewayLanProperties struct {
	// Collection of gateway IP addresses of the NAT Gateway. Will be auto-generated if not provided. Should ideally be an IP belonging to the same subnet as the LAN
	GatewayIps *[]string `json:"gatewayIps,omitempty"`
	// Id for the LAN connected to the NAT Gateway
	Id *int32 `json:"id"`
}

// NewNatGatewayLanProperties instantiates a new NatGatewayLanProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNatGatewayLanProperties(id int32) *NatGatewayLanProperties {
	this := NatGatewayLanProperties{}

	this.Id = &id

	return &this
}

// NewNatGatewayLanPropertiesWithDefaults instantiates a new NatGatewayLanProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNatGatewayLanPropertiesWithDefaults() *NatGatewayLanProperties {
	this := NatGatewayLanProperties{}
	return &this
}

// GetGatewayIps returns the GatewayIps field value
// If the value is explicit nil, nil is returned
func (o *NatGatewayLanProperties) GetGatewayIps() *[]string {
	if o == nil {
		return nil
	}

	return o.GatewayIps

}

// GetGatewayIpsOk returns a tuple with the GatewayIps field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NatGatewayLanProperties) GetGatewayIpsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.GatewayIps, true
}

// SetGatewayIps sets field value
func (o *NatGatewayLanProperties) SetGatewayIps(v []string) {

	o.GatewayIps = &v

}

// HasGatewayIps returns a boolean if a field has been set.
func (o *NatGatewayLanProperties) HasGatewayIps() bool {
	if o != nil && o.GatewayIps != nil {
		return true
	}

	return false
}

// GetId returns the Id field value
// If the value is explicit nil, nil is returned
func (o *NatGatewayLanProperties) GetId() *int32 {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NatGatewayLanProperties) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *NatGatewayLanProperties) SetId(v int32) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *NatGatewayLanProperties) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

func (o NatGatewayLanProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GatewayIps != nil {
		toSerialize["gatewayIps"] = o.GatewayIps
	}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	return json.Marshal(toSerialize)
}

type NullableNatGatewayLanProperties struct {
	value *NatGatewayLanProperties
	isSet bool
}

func (v NullableNatGatewayLanProperties) Get() *NatGatewayLanProperties {
	return v.value
}

func (v *NullableNatGatewayLanProperties) Set(val *NatGatewayLanProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableNatGatewayLanProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableNatGatewayLanProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNatGatewayLanProperties(val *NatGatewayLanProperties) *NullableNatGatewayLanProperties {
	return &NullableNatGatewayLanProperties{value: val, isSet: true}
}

func (v NullableNatGatewayLanProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNatGatewayLanProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
