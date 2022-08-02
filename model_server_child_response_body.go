/*
Future Vuls Public API

Future Vuls Public API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vuls

import (
	"encoding/json"
)

// ServerChildResponseBody struct for ServerChildResponseBody
type ServerChildResponseBody struct {
	// crated time of server
	CreatedAt string `json:"createdAt"`
	// default user ID of server
	DefaultUserId *int64 `json:"defaultUserId,omitempty"`
	// default user name of server
	DefaultUserName *string `json:"defaultUserName,omitempty"`
	// UUID of server
	HostUuid string `json:"hostUuid"`
	// ID of server
	Id int64 `json:"id"`
	// last scanned time of server
	LastScannedAt *string `json:"lastScannedAt,omitempty"`
	// last uploaded time of server
	LastUploadedAt *string `json:"lastUploadedAt,omitempty"`
	// Whether server needs kernel restart
	NeedKernelRestart bool `json:"needKernelRestart"`
	// OS Name of server
	OsFamily string `json:"osFamily"`
	// OS Version of server
	OsVersion string `json:"osVersion"`
	// Name of server
	ServerName string `json:"serverName"`
	// UUID of server
	ServerUuid string `json:"serverUuid"`
	// ID of server role
	ServerroleId int64 `json:"serverroleId"`
	// Name of server role
	ServerroleName string `json:"serverroleName"`
	// tags is list of server tag
	Tags []string `json:"tags,omitempty"`
	// updated time of server
	UpdatedAt string `json:"updatedAt"`
}

// NewServerChildResponseBody instantiates a new ServerChildResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerChildResponseBody(createdAt string, hostUuid string, id int64, needKernelRestart bool, osFamily string, osVersion string, serverName string, serverUuid string, serverroleId int64, serverroleName string, updatedAt string) *ServerChildResponseBody {
	this := ServerChildResponseBody{}
	this.CreatedAt = createdAt
	this.HostUuid = hostUuid
	this.Id = id
	this.NeedKernelRestart = needKernelRestart
	this.OsFamily = osFamily
	this.OsVersion = osVersion
	this.ServerName = serverName
	this.ServerUuid = serverUuid
	this.ServerroleId = serverroleId
	this.ServerroleName = serverroleName
	this.UpdatedAt = updatedAt
	return &this
}

// NewServerChildResponseBodyWithDefaults instantiates a new ServerChildResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerChildResponseBodyWithDefaults() *ServerChildResponseBody {
	this := ServerChildResponseBody{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *ServerChildResponseBody) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ServerChildResponseBody) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ServerChildResponseBody) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDefaultUserId returns the DefaultUserId field value if set, zero value otherwise.
func (o *ServerChildResponseBody) GetDefaultUserId() int64 {
	if o == nil || o.DefaultUserId == nil {
		var ret int64
		return ret
	}
	return *o.DefaultUserId
}

// GetDefaultUserIdOk returns a tuple with the DefaultUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerChildResponseBody) GetDefaultUserIdOk() (*int64, bool) {
	if o == nil || o.DefaultUserId == nil {
		return nil, false
	}
	return o.DefaultUserId, true
}

// HasDefaultUserId returns a boolean if a field has been set.
func (o *ServerChildResponseBody) HasDefaultUserId() bool {
	if o != nil && o.DefaultUserId != nil {
		return true
	}

	return false
}

// SetDefaultUserId gets a reference to the given int64 and assigns it to the DefaultUserId field.
func (o *ServerChildResponseBody) SetDefaultUserId(v int64) {
	o.DefaultUserId = &v
}

// GetDefaultUserName returns the DefaultUserName field value if set, zero value otherwise.
func (o *ServerChildResponseBody) GetDefaultUserName() string {
	if o == nil || o.DefaultUserName == nil {
		var ret string
		return ret
	}
	return *o.DefaultUserName
}

// GetDefaultUserNameOk returns a tuple with the DefaultUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerChildResponseBody) GetDefaultUserNameOk() (*string, bool) {
	if o == nil || o.DefaultUserName == nil {
		return nil, false
	}
	return o.DefaultUserName, true
}

// HasDefaultUserName returns a boolean if a field has been set.
func (o *ServerChildResponseBody) HasDefaultUserName() bool {
	if o != nil && o.DefaultUserName != nil {
		return true
	}

	return false
}

// SetDefaultUserName gets a reference to the given string and assigns it to the DefaultUserName field.
func (o *ServerChildResponseBody) SetDefaultUserName(v string) {
	o.DefaultUserName = &v
}

// GetHostUuid returns the HostUuid field value
func (o *ServerChildResponseBody) GetHostUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostUuid
}

// GetHostUuidOk returns a tuple with the HostUuid field value
// and a boolean to check if the value has been set.
func (o *ServerChildResponseBody) GetHostUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostUuid, true
}

// SetHostUuid sets field value
func (o *ServerChildResponseBody) SetHostUuid(v string) {
	o.HostUuid = v
}

// GetId returns the Id field value
func (o *ServerChildResponseBody) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServerChildResponseBody) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServerChildResponseBody) SetId(v int64) {
	o.Id = v
}

// GetLastScannedAt returns the LastScannedAt field value if set, zero value otherwise.
func (o *ServerChildResponseBody) GetLastScannedAt() string {
	if o == nil || o.LastScannedAt == nil {
		var ret string
		return ret
	}
	return *o.LastScannedAt
}

// GetLastScannedAtOk returns a tuple with the LastScannedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerChildResponseBody) GetLastScannedAtOk() (*string, bool) {
	if o == nil || o.LastScannedAt == nil {
		return nil, false
	}
	return o.LastScannedAt, true
}

// HasLastScannedAt returns a boolean if a field has been set.
func (o *ServerChildResponseBody) HasLastScannedAt() bool {
	if o != nil && o.LastScannedAt != nil {
		return true
	}

	return false
}

// SetLastScannedAt gets a reference to the given string and assigns it to the LastScannedAt field.
func (o *ServerChildResponseBody) SetLastScannedAt(v string) {
	o.LastScannedAt = &v
}

// GetLastUploadedAt returns the LastUploadedAt field value if set, zero value otherwise.
func (o *ServerChildResponseBody) GetLastUploadedAt() string {
	if o == nil || o.LastUploadedAt == nil {
		var ret string
		return ret
	}
	return *o.LastUploadedAt
}

// GetLastUploadedAtOk returns a tuple with the LastUploadedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerChildResponseBody) GetLastUploadedAtOk() (*string, bool) {
	if o == nil || o.LastUploadedAt == nil {
		return nil, false
	}
	return o.LastUploadedAt, true
}

// HasLastUploadedAt returns a boolean if a field has been set.
func (o *ServerChildResponseBody) HasLastUploadedAt() bool {
	if o != nil && o.LastUploadedAt != nil {
		return true
	}

	return false
}

// SetLastUploadedAt gets a reference to the given string and assigns it to the LastUploadedAt field.
func (o *ServerChildResponseBody) SetLastUploadedAt(v string) {
	o.LastUploadedAt = &v
}

// GetNeedKernelRestart returns the NeedKernelRestart field value
func (o *ServerChildResponseBody) GetNeedKernelRestart() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NeedKernelRestart
}

// GetNeedKernelRestartOk returns a tuple with the NeedKernelRestart field value
// and a boolean to check if the value has been set.
func (o *ServerChildResponseBody) GetNeedKernelRestartOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NeedKernelRestart, true
}

// SetNeedKernelRestart sets field value
func (o *ServerChildResponseBody) SetNeedKernelRestart(v bool) {
	o.NeedKernelRestart = v
}

// GetOsFamily returns the OsFamily field value
func (o *ServerChildResponseBody) GetOsFamily() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsFamily
}

// GetOsFamilyOk returns a tuple with the OsFamily field value
// and a boolean to check if the value has been set.
func (o *ServerChildResponseBody) GetOsFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsFamily, true
}

// SetOsFamily sets field value
func (o *ServerChildResponseBody) SetOsFamily(v string) {
	o.OsFamily = v
}

// GetOsVersion returns the OsVersion field value
func (o *ServerChildResponseBody) GetOsVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value
// and a boolean to check if the value has been set.
func (o *ServerChildResponseBody) GetOsVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsVersion, true
}

// SetOsVersion sets field value
func (o *ServerChildResponseBody) SetOsVersion(v string) {
	o.OsVersion = v
}

// GetServerName returns the ServerName field value
func (o *ServerChildResponseBody) GetServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value
// and a boolean to check if the value has been set.
func (o *ServerChildResponseBody) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerName, true
}

// SetServerName sets field value
func (o *ServerChildResponseBody) SetServerName(v string) {
	o.ServerName = v
}

// GetServerUuid returns the ServerUuid field value
func (o *ServerChildResponseBody) GetServerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerUuid
}

// GetServerUuidOk returns a tuple with the ServerUuid field value
// and a boolean to check if the value has been set.
func (o *ServerChildResponseBody) GetServerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerUuid, true
}

// SetServerUuid sets field value
func (o *ServerChildResponseBody) SetServerUuid(v string) {
	o.ServerUuid = v
}

// GetServerroleId returns the ServerroleId field value
func (o *ServerChildResponseBody) GetServerroleId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServerroleId
}

// GetServerroleIdOk returns a tuple with the ServerroleId field value
// and a boolean to check if the value has been set.
func (o *ServerChildResponseBody) GetServerroleIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerroleId, true
}

// SetServerroleId sets field value
func (o *ServerChildResponseBody) SetServerroleId(v int64) {
	o.ServerroleId = v
}

// GetServerroleName returns the ServerroleName field value
func (o *ServerChildResponseBody) GetServerroleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerroleName
}

// GetServerroleNameOk returns a tuple with the ServerroleName field value
// and a boolean to check if the value has been set.
func (o *ServerChildResponseBody) GetServerroleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerroleName, true
}

// SetServerroleName sets field value
func (o *ServerChildResponseBody) SetServerroleName(v string) {
	o.ServerroleName = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ServerChildResponseBody) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerChildResponseBody) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ServerChildResponseBody) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ServerChildResponseBody) SetTags(v []string) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ServerChildResponseBody) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ServerChildResponseBody) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ServerChildResponseBody) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o ServerChildResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.DefaultUserId != nil {
		toSerialize["defaultUserId"] = o.DefaultUserId
	}
	if o.DefaultUserName != nil {
		toSerialize["defaultUserName"] = o.DefaultUserName
	}
	if true {
		toSerialize["hostUuid"] = o.HostUuid
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.LastScannedAt != nil {
		toSerialize["lastScannedAt"] = o.LastScannedAt
	}
	if o.LastUploadedAt != nil {
		toSerialize["lastUploadedAt"] = o.LastUploadedAt
	}
	if true {
		toSerialize["needKernelRestart"] = o.NeedKernelRestart
	}
	if true {
		toSerialize["osFamily"] = o.OsFamily
	}
	if true {
		toSerialize["osVersion"] = o.OsVersion
	}
	if true {
		toSerialize["serverName"] = o.ServerName
	}
	if true {
		toSerialize["serverUuid"] = o.ServerUuid
	}
	if true {
		toSerialize["serverroleId"] = o.ServerroleId
	}
	if true {
		toSerialize["serverroleName"] = o.ServerroleName
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableServerChildResponseBody struct {
	value *ServerChildResponseBody
	isSet bool
}

func (v NullableServerChildResponseBody) Get() *ServerChildResponseBody {
	return v.value
}

func (v *NullableServerChildResponseBody) Set(val *ServerChildResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableServerChildResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableServerChildResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerChildResponseBody(val *ServerChildResponseBody) *NullableServerChildResponseBody {
	return &NullableServerChildResponseBody{value: val, isSet: true}
}

func (v NullableServerChildResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerChildResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


