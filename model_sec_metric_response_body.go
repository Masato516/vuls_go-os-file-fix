/*
Future Vuls Public API

Future Vuls Public API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vuls

import (
	"encoding/json"
	"time"
)

// SecMetricResponseBody SecMetric describes a secMetric
type SecMetricResponseBody struct {
	// AR of secMetric
	Ar string `json:"ar"`
	// CR of secMetric
	Cr string `json:"cr"`
	// created time
	CreatedAt time.Time `json:"createdAt"`
	// IR of secMetric
	Ir string `json:"ir"`
	// ServerRoleID of secMetric
	RoleID int64 `json:"roleID"`
	// ServerRoleName of secMetric
	RoleName string `json:"roleName"`
	// updated time
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewSecMetricResponseBody instantiates a new SecMetricResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecMetricResponseBody(ar string, cr string, createdAt time.Time, ir string, roleID int64, roleName string, updatedAt time.Time) *SecMetricResponseBody {
	this := SecMetricResponseBody{}
	this.Ar = ar
	this.Cr = cr
	this.CreatedAt = createdAt
	this.Ir = ir
	this.RoleID = roleID
	this.RoleName = roleName
	this.UpdatedAt = updatedAt
	return &this
}

// NewSecMetricResponseBodyWithDefaults instantiates a new SecMetricResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecMetricResponseBodyWithDefaults() *SecMetricResponseBody {
	this := SecMetricResponseBody{}
	return &this
}

// GetAr returns the Ar field value
func (o *SecMetricResponseBody) GetAr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ar
}

// GetArOk returns a tuple with the Ar field value
// and a boolean to check if the value has been set.
func (o *SecMetricResponseBody) GetArOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ar, true
}

// SetAr sets field value
func (o *SecMetricResponseBody) SetAr(v string) {
	o.Ar = v
}

// GetCr returns the Cr field value
func (o *SecMetricResponseBody) GetCr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cr
}

// GetCrOk returns a tuple with the Cr field value
// and a boolean to check if the value has been set.
func (o *SecMetricResponseBody) GetCrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cr, true
}

// SetCr sets field value
func (o *SecMetricResponseBody) SetCr(v string) {
	o.Cr = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SecMetricResponseBody) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SecMetricResponseBody) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SecMetricResponseBody) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetIr returns the Ir field value
func (o *SecMetricResponseBody) GetIr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ir
}

// GetIrOk returns a tuple with the Ir field value
// and a boolean to check if the value has been set.
func (o *SecMetricResponseBody) GetIrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ir, true
}

// SetIr sets field value
func (o *SecMetricResponseBody) SetIr(v string) {
	o.Ir = v
}

// GetRoleID returns the RoleID field value
func (o *SecMetricResponseBody) GetRoleID() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RoleID
}

// GetRoleIDOk returns a tuple with the RoleID field value
// and a boolean to check if the value has been set.
func (o *SecMetricResponseBody) GetRoleIDOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleID, true
}

// SetRoleID sets field value
func (o *SecMetricResponseBody) SetRoleID(v int64) {
	o.RoleID = v
}

// GetRoleName returns the RoleName field value
func (o *SecMetricResponseBody) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *SecMetricResponseBody) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *SecMetricResponseBody) SetRoleName(v string) {
	o.RoleName = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *SecMetricResponseBody) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SecMetricResponseBody) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *SecMetricResponseBody) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o SecMetricResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ar"] = o.Ar
	}
	if true {
		toSerialize["cr"] = o.Cr
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["ir"] = o.Ir
	}
	if true {
		toSerialize["roleID"] = o.RoleID
	}
	if true {
		toSerialize["roleName"] = o.RoleName
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableSecMetricResponseBody struct {
	value *SecMetricResponseBody
	isSet bool
}

func (v NullableSecMetricResponseBody) Get() *SecMetricResponseBody {
	return v.value
}

func (v *NullableSecMetricResponseBody) Set(val *SecMetricResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSecMetricResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSecMetricResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecMetricResponseBody(val *SecMetricResponseBody) *NullableSecMetricResponseBody {
	return &NullableSecMetricResponseBody{value: val, isSet: true}
}

func (v NullableSecMetricResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecMetricResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


