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

// TmpMetricResponseBody TmpMetric describes a tmpMetric
type TmpMetricResponseBody struct {
	// created time
	CreatedAt time.Time `json:"createdAt"`
	// E of tmpMetric
	E string `json:"e"`
	// RC of tmpMetric
	Rc string `json:"rc"`
	// RL of tmpMetric
	Rl string `json:"rl"`
	// updated time
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewTmpMetricResponseBody instantiates a new TmpMetricResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTmpMetricResponseBody(createdAt time.Time, e string, rc string, rl string, updatedAt time.Time) *TmpMetricResponseBody {
	this := TmpMetricResponseBody{}
	this.CreatedAt = createdAt
	this.E = e
	this.Rc = rc
	this.Rl = rl
	this.UpdatedAt = updatedAt
	return &this
}

// NewTmpMetricResponseBodyWithDefaults instantiates a new TmpMetricResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTmpMetricResponseBodyWithDefaults() *TmpMetricResponseBody {
	this := TmpMetricResponseBody{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *TmpMetricResponseBody) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TmpMetricResponseBody) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TmpMetricResponseBody) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetE returns the E field value
func (o *TmpMetricResponseBody) GetE() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.E
}

// GetEOk returns a tuple with the E field value
// and a boolean to check if the value has been set.
func (o *TmpMetricResponseBody) GetEOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.E, true
}

// SetE sets field value
func (o *TmpMetricResponseBody) SetE(v string) {
	o.E = v
}

// GetRc returns the Rc field value
func (o *TmpMetricResponseBody) GetRc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rc
}

// GetRcOk returns a tuple with the Rc field value
// and a boolean to check if the value has been set.
func (o *TmpMetricResponseBody) GetRcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rc, true
}

// SetRc sets field value
func (o *TmpMetricResponseBody) SetRc(v string) {
	o.Rc = v
}

// GetRl returns the Rl field value
func (o *TmpMetricResponseBody) GetRl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rl
}

// GetRlOk returns a tuple with the Rl field value
// and a boolean to check if the value has been set.
func (o *TmpMetricResponseBody) GetRlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rl, true
}

// SetRl sets field value
func (o *TmpMetricResponseBody) SetRl(v string) {
	o.Rl = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TmpMetricResponseBody) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TmpMetricResponseBody) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TmpMetricResponseBody) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o TmpMetricResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["e"] = o.E
	}
	if true {
		toSerialize["rc"] = o.Rc
	}
	if true {
		toSerialize["rl"] = o.Rl
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableTmpMetricResponseBody struct {
	value *TmpMetricResponseBody
	isSet bool
}

func (v NullableTmpMetricResponseBody) Get() *TmpMetricResponseBody {
	return v.value
}

func (v *NullableTmpMetricResponseBody) Set(val *TmpMetricResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTmpMetricResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTmpMetricResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTmpMetricResponseBody(val *TmpMetricResponseBody) *NullableTmpMetricResponseBody {
	return &NullableTmpMetricResponseBody{value: val, isSet: true}
}

func (v NullableTmpMetricResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTmpMetricResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


