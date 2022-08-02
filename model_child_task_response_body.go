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

// ChildTaskResponseBody struct for ChildTaskResponseBody
type ChildTaskResponseBody struct {
	// ApplyingPatchOn of task
	ApplyingPatchOn *string `json:"applyingPatchOn,omitempty"`
	// created time of task
	CreatedAt time.Time `json:"createdAt"`
	// CVE ID of task
	CveID string `json:"cveID"`
	// ID of task
	Id int64 `json:"id"`
	// Ignore of task
	Ignore bool `json:"ignore"`
	// Ignore until of task
	IgnoreUntil *string `json:"ignoreUntil,omitempty"`
	// MainUserID of task
	MainUserID *int64 `json:"mainUserID,omitempty"`
	// MainUserName of task
	MainUserName *string `json:"mainUserName,omitempty"`
	// Priority of task
	Priority string `json:"priority"`
	// ServerID of task
	ServerID int64 `json:"serverID"`
	// Status of task
	Status string `json:"status"`
	// SubUserID of task
	SubUserID *int64 `json:"subUserID,omitempty"`
	// SubUserName of task
	SubUserName *string `json:"subUserName,omitempty"`
	// updated time of task
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewChildTaskResponseBody instantiates a new ChildTaskResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChildTaskResponseBody(createdAt time.Time, cveID string, id int64, ignore bool, priority string, serverID int64, status string, updatedAt time.Time) *ChildTaskResponseBody {
	this := ChildTaskResponseBody{}
	this.CreatedAt = createdAt
	this.CveID = cveID
	this.Id = id
	this.Ignore = ignore
	this.Priority = priority
	this.ServerID = serverID
	this.Status = status
	this.UpdatedAt = updatedAt
	return &this
}

// NewChildTaskResponseBodyWithDefaults instantiates a new ChildTaskResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChildTaskResponseBodyWithDefaults() *ChildTaskResponseBody {
	this := ChildTaskResponseBody{}
	return &this
}

// GetApplyingPatchOn returns the ApplyingPatchOn field value if set, zero value otherwise.
func (o *ChildTaskResponseBody) GetApplyingPatchOn() string {
	if o == nil || o.ApplyingPatchOn == nil {
		var ret string
		return ret
	}
	return *o.ApplyingPatchOn
}

// GetApplyingPatchOnOk returns a tuple with the ApplyingPatchOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildTaskResponseBody) GetApplyingPatchOnOk() (*string, bool) {
	if o == nil || o.ApplyingPatchOn == nil {
		return nil, false
	}
	return o.ApplyingPatchOn, true
}

// HasApplyingPatchOn returns a boolean if a field has been set.
func (o *ChildTaskResponseBody) HasApplyingPatchOn() bool {
	if o != nil && o.ApplyingPatchOn != nil {
		return true
	}

	return false
}

// SetApplyingPatchOn gets a reference to the given string and assigns it to the ApplyingPatchOn field.
func (o *ChildTaskResponseBody) SetApplyingPatchOn(v string) {
	o.ApplyingPatchOn = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ChildTaskResponseBody) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ChildTaskResponseBody) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ChildTaskResponseBody) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCveID returns the CveID field value
func (o *ChildTaskResponseBody) GetCveID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CveID
}

// GetCveIDOk returns a tuple with the CveID field value
// and a boolean to check if the value has been set.
func (o *ChildTaskResponseBody) GetCveIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CveID, true
}

// SetCveID sets field value
func (o *ChildTaskResponseBody) SetCveID(v string) {
	o.CveID = v
}

// GetId returns the Id field value
func (o *ChildTaskResponseBody) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChildTaskResponseBody) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChildTaskResponseBody) SetId(v int64) {
	o.Id = v
}

// GetIgnore returns the Ignore field value
func (o *ChildTaskResponseBody) GetIgnore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ignore
}

// GetIgnoreOk returns a tuple with the Ignore field value
// and a boolean to check if the value has been set.
func (o *ChildTaskResponseBody) GetIgnoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ignore, true
}

// SetIgnore sets field value
func (o *ChildTaskResponseBody) SetIgnore(v bool) {
	o.Ignore = v
}

// GetIgnoreUntil returns the IgnoreUntil field value if set, zero value otherwise.
func (o *ChildTaskResponseBody) GetIgnoreUntil() string {
	if o == nil || o.IgnoreUntil == nil {
		var ret string
		return ret
	}
	return *o.IgnoreUntil
}

// GetIgnoreUntilOk returns a tuple with the IgnoreUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildTaskResponseBody) GetIgnoreUntilOk() (*string, bool) {
	if o == nil || o.IgnoreUntil == nil {
		return nil, false
	}
	return o.IgnoreUntil, true
}

// HasIgnoreUntil returns a boolean if a field has been set.
func (o *ChildTaskResponseBody) HasIgnoreUntil() bool {
	if o != nil && o.IgnoreUntil != nil {
		return true
	}

	return false
}

// SetIgnoreUntil gets a reference to the given string and assigns it to the IgnoreUntil field.
func (o *ChildTaskResponseBody) SetIgnoreUntil(v string) {
	o.IgnoreUntil = &v
}

// GetMainUserID returns the MainUserID field value if set, zero value otherwise.
func (o *ChildTaskResponseBody) GetMainUserID() int64 {
	if o == nil || o.MainUserID == nil {
		var ret int64
		return ret
	}
	return *o.MainUserID
}

// GetMainUserIDOk returns a tuple with the MainUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildTaskResponseBody) GetMainUserIDOk() (*int64, bool) {
	if o == nil || o.MainUserID == nil {
		return nil, false
	}
	return o.MainUserID, true
}

// HasMainUserID returns a boolean if a field has been set.
func (o *ChildTaskResponseBody) HasMainUserID() bool {
	if o != nil && o.MainUserID != nil {
		return true
	}

	return false
}

// SetMainUserID gets a reference to the given int64 and assigns it to the MainUserID field.
func (o *ChildTaskResponseBody) SetMainUserID(v int64) {
	o.MainUserID = &v
}

// GetMainUserName returns the MainUserName field value if set, zero value otherwise.
func (o *ChildTaskResponseBody) GetMainUserName() string {
	if o == nil || o.MainUserName == nil {
		var ret string
		return ret
	}
	return *o.MainUserName
}

// GetMainUserNameOk returns a tuple with the MainUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildTaskResponseBody) GetMainUserNameOk() (*string, bool) {
	if o == nil || o.MainUserName == nil {
		return nil, false
	}
	return o.MainUserName, true
}

// HasMainUserName returns a boolean if a field has been set.
func (o *ChildTaskResponseBody) HasMainUserName() bool {
	if o != nil && o.MainUserName != nil {
		return true
	}

	return false
}

// SetMainUserName gets a reference to the given string and assigns it to the MainUserName field.
func (o *ChildTaskResponseBody) SetMainUserName(v string) {
	o.MainUserName = &v
}

// GetPriority returns the Priority field value
func (o *ChildTaskResponseBody) GetPriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *ChildTaskResponseBody) GetPriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *ChildTaskResponseBody) SetPriority(v string) {
	o.Priority = v
}

// GetServerID returns the ServerID field value
func (o *ChildTaskResponseBody) GetServerID() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServerID
}

// GetServerIDOk returns a tuple with the ServerID field value
// and a boolean to check if the value has been set.
func (o *ChildTaskResponseBody) GetServerIDOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerID, true
}

// SetServerID sets field value
func (o *ChildTaskResponseBody) SetServerID(v int64) {
	o.ServerID = v
}

// GetStatus returns the Status field value
func (o *ChildTaskResponseBody) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ChildTaskResponseBody) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ChildTaskResponseBody) SetStatus(v string) {
	o.Status = v
}

// GetSubUserID returns the SubUserID field value if set, zero value otherwise.
func (o *ChildTaskResponseBody) GetSubUserID() int64 {
	if o == nil || o.SubUserID == nil {
		var ret int64
		return ret
	}
	return *o.SubUserID
}

// GetSubUserIDOk returns a tuple with the SubUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildTaskResponseBody) GetSubUserIDOk() (*int64, bool) {
	if o == nil || o.SubUserID == nil {
		return nil, false
	}
	return o.SubUserID, true
}

// HasSubUserID returns a boolean if a field has been set.
func (o *ChildTaskResponseBody) HasSubUserID() bool {
	if o != nil && o.SubUserID != nil {
		return true
	}

	return false
}

// SetSubUserID gets a reference to the given int64 and assigns it to the SubUserID field.
func (o *ChildTaskResponseBody) SetSubUserID(v int64) {
	o.SubUserID = &v
}

// GetSubUserName returns the SubUserName field value if set, zero value otherwise.
func (o *ChildTaskResponseBody) GetSubUserName() string {
	if o == nil || o.SubUserName == nil {
		var ret string
		return ret
	}
	return *o.SubUserName
}

// GetSubUserNameOk returns a tuple with the SubUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildTaskResponseBody) GetSubUserNameOk() (*string, bool) {
	if o == nil || o.SubUserName == nil {
		return nil, false
	}
	return o.SubUserName, true
}

// HasSubUserName returns a boolean if a field has been set.
func (o *ChildTaskResponseBody) HasSubUserName() bool {
	if o != nil && o.SubUserName != nil {
		return true
	}

	return false
}

// SetSubUserName gets a reference to the given string and assigns it to the SubUserName field.
func (o *ChildTaskResponseBody) SetSubUserName(v string) {
	o.SubUserName = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ChildTaskResponseBody) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ChildTaskResponseBody) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ChildTaskResponseBody) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o ChildTaskResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplyingPatchOn != nil {
		toSerialize["applyingPatchOn"] = o.ApplyingPatchOn
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["cveID"] = o.CveID
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["ignore"] = o.Ignore
	}
	if o.IgnoreUntil != nil {
		toSerialize["ignoreUntil"] = o.IgnoreUntil
	}
	if o.MainUserID != nil {
		toSerialize["mainUserID"] = o.MainUserID
	}
	if o.MainUserName != nil {
		toSerialize["mainUserName"] = o.MainUserName
	}
	if true {
		toSerialize["priority"] = o.Priority
	}
	if true {
		toSerialize["serverID"] = o.ServerID
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.SubUserID != nil {
		toSerialize["subUserID"] = o.SubUserID
	}
	if o.SubUserName != nil {
		toSerialize["subUserName"] = o.SubUserName
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableChildTaskResponseBody struct {
	value *ChildTaskResponseBody
	isSet bool
}

func (v NullableChildTaskResponseBody) Get() *ChildTaskResponseBody {
	return v.value
}

func (v *NullableChildTaskResponseBody) Set(val *ChildTaskResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableChildTaskResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableChildTaskResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChildTaskResponseBody(val *ChildTaskResponseBody) *NullableChildTaskResponseBody {
	return &NullableChildTaskResponseBody{value: val, isSet: true}
}

func (v NullableChildTaskResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChildTaskResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


