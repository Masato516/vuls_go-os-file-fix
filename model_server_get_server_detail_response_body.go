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

// ServerGetServerDetailResponseBody struct for ServerGetServerDetailResponseBody
type ServerGetServerDetailResponseBody struct {
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
	// platformInstanceId of server
	PlatformInstanceId string `json:"platformInstanceId"`
	// platformName of server
	PlatformName string `json:"platformName"`
	// IPv4 of server
	ServerIpv4 string `json:"serverIpv4"`
	// Name of server
	ServerName string `json:"serverName"`
	// UUID of server
	ServerUuid string `json:"serverUuid"`
	// ID of server role
	ServerroleId int64 `json:"serverroleId"`
	// Name of server role
	ServerroleName string `json:"serverroleName"`
	// tags is list of server tag
	Tags []ServerTagResponseBody `json:"tags,omitempty"`
	// tasks of server
	Tasks []ChildTaskResponseBody `json:"tasks,omitempty"`
	// updated time of server
	UpdatedAt string `json:"updatedAt"`
}

// NewServerGetServerDetailResponseBody instantiates a new ServerGetServerDetailResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerGetServerDetailResponseBody(createdAt string, hostUuid string, id int64, needKernelRestart bool, osFamily string, osVersion string, platformInstanceId string, platformName string, serverIpv4 string, serverName string, serverUuid string, serverroleId int64, serverroleName string, updatedAt string) *ServerGetServerDetailResponseBody {
	this := ServerGetServerDetailResponseBody{}
	this.CreatedAt = createdAt
	this.HostUuid = hostUuid
	this.Id = id
	this.NeedKernelRestart = needKernelRestart
	this.OsFamily = osFamily
	this.OsVersion = osVersion
	this.PlatformInstanceId = platformInstanceId
	this.PlatformName = platformName
	this.ServerIpv4 = serverIpv4
	this.ServerName = serverName
	this.ServerUuid = serverUuid
	this.ServerroleId = serverroleId
	this.ServerroleName = serverroleName
	this.UpdatedAt = updatedAt
	return &this
}

// NewServerGetServerDetailResponseBodyWithDefaults instantiates a new ServerGetServerDetailResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerGetServerDetailResponseBodyWithDefaults() *ServerGetServerDetailResponseBody {
	this := ServerGetServerDetailResponseBody{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *ServerGetServerDetailResponseBody) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ServerGetServerDetailResponseBody) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDefaultUserId returns the DefaultUserId field value if set, zero value otherwise.
func (o *ServerGetServerDetailResponseBody) GetDefaultUserId() int64 {
	if o == nil || o.DefaultUserId == nil {
		var ret int64
		return ret
	}
	return *o.DefaultUserId
}

// GetDefaultUserIdOk returns a tuple with the DefaultUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetDefaultUserIdOk() (*int64, bool) {
	if o == nil || o.DefaultUserId == nil {
		return nil, false
	}
	return o.DefaultUserId, true
}

// HasDefaultUserId returns a boolean if a field has been set.
func (o *ServerGetServerDetailResponseBody) HasDefaultUserId() bool {
	if o != nil && o.DefaultUserId != nil {
		return true
	}

	return false
}

// SetDefaultUserId gets a reference to the given int64 and assigns it to the DefaultUserId field.
func (o *ServerGetServerDetailResponseBody) SetDefaultUserId(v int64) {
	o.DefaultUserId = &v
}

// GetDefaultUserName returns the DefaultUserName field value if set, zero value otherwise.
func (o *ServerGetServerDetailResponseBody) GetDefaultUserName() string {
	if o == nil || o.DefaultUserName == nil {
		var ret string
		return ret
	}
	return *o.DefaultUserName
}

// GetDefaultUserNameOk returns a tuple with the DefaultUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetDefaultUserNameOk() (*string, bool) {
	if o == nil || o.DefaultUserName == nil {
		return nil, false
	}
	return o.DefaultUserName, true
}

// HasDefaultUserName returns a boolean if a field has been set.
func (o *ServerGetServerDetailResponseBody) HasDefaultUserName() bool {
	if o != nil && o.DefaultUserName != nil {
		return true
	}

	return false
}

// SetDefaultUserName gets a reference to the given string and assigns it to the DefaultUserName field.
func (o *ServerGetServerDetailResponseBody) SetDefaultUserName(v string) {
	o.DefaultUserName = &v
}

// GetHostUuid returns the HostUuid field value
func (o *ServerGetServerDetailResponseBody) GetHostUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostUuid
}

// GetHostUuidOk returns a tuple with the HostUuid field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetHostUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostUuid, true
}

// SetHostUuid sets field value
func (o *ServerGetServerDetailResponseBody) SetHostUuid(v string) {
	o.HostUuid = v
}

// GetId returns the Id field value
func (o *ServerGetServerDetailResponseBody) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServerGetServerDetailResponseBody) SetId(v int64) {
	o.Id = v
}

// GetLastScannedAt returns the LastScannedAt field value if set, zero value otherwise.
func (o *ServerGetServerDetailResponseBody) GetLastScannedAt() string {
	if o == nil || o.LastScannedAt == nil {
		var ret string
		return ret
	}
	return *o.LastScannedAt
}

// GetLastScannedAtOk returns a tuple with the LastScannedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetLastScannedAtOk() (*string, bool) {
	if o == nil || o.LastScannedAt == nil {
		return nil, false
	}
	return o.LastScannedAt, true
}

// HasLastScannedAt returns a boolean if a field has been set.
func (o *ServerGetServerDetailResponseBody) HasLastScannedAt() bool {
	if o != nil && o.LastScannedAt != nil {
		return true
	}

	return false
}

// SetLastScannedAt gets a reference to the given string and assigns it to the LastScannedAt field.
func (o *ServerGetServerDetailResponseBody) SetLastScannedAt(v string) {
	o.LastScannedAt = &v
}

// GetLastUploadedAt returns the LastUploadedAt field value if set, zero value otherwise.
func (o *ServerGetServerDetailResponseBody) GetLastUploadedAt() string {
	if o == nil || o.LastUploadedAt == nil {
		var ret string
		return ret
	}
	return *o.LastUploadedAt
}

// GetLastUploadedAtOk returns a tuple with the LastUploadedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetLastUploadedAtOk() (*string, bool) {
	if o == nil || o.LastUploadedAt == nil {
		return nil, false
	}
	return o.LastUploadedAt, true
}

// HasLastUploadedAt returns a boolean if a field has been set.
func (o *ServerGetServerDetailResponseBody) HasLastUploadedAt() bool {
	if o != nil && o.LastUploadedAt != nil {
		return true
	}

	return false
}

// SetLastUploadedAt gets a reference to the given string and assigns it to the LastUploadedAt field.
func (o *ServerGetServerDetailResponseBody) SetLastUploadedAt(v string) {
	o.LastUploadedAt = &v
}

// GetNeedKernelRestart returns the NeedKernelRestart field value
func (o *ServerGetServerDetailResponseBody) GetNeedKernelRestart() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NeedKernelRestart
}

// GetNeedKernelRestartOk returns a tuple with the NeedKernelRestart field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetNeedKernelRestartOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NeedKernelRestart, true
}

// SetNeedKernelRestart sets field value
func (o *ServerGetServerDetailResponseBody) SetNeedKernelRestart(v bool) {
	o.NeedKernelRestart = v
}

// GetOsFamily returns the OsFamily field value
func (o *ServerGetServerDetailResponseBody) GetOsFamily() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsFamily
}

// GetOsFamilyOk returns a tuple with the OsFamily field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetOsFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsFamily, true
}

// SetOsFamily sets field value
func (o *ServerGetServerDetailResponseBody) SetOsFamily(v string) {
	o.OsFamily = v
}

// GetOsVersion returns the OsVersion field value
func (o *ServerGetServerDetailResponseBody) GetOsVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetOsVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsVersion, true
}

// SetOsVersion sets field value
func (o *ServerGetServerDetailResponseBody) SetOsVersion(v string) {
	o.OsVersion = v
}

// GetPlatformInstanceId returns the PlatformInstanceId field value
func (o *ServerGetServerDetailResponseBody) GetPlatformInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlatformInstanceId
}

// GetPlatformInstanceIdOk returns a tuple with the PlatformInstanceId field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetPlatformInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlatformInstanceId, true
}

// SetPlatformInstanceId sets field value
func (o *ServerGetServerDetailResponseBody) SetPlatformInstanceId(v string) {
	o.PlatformInstanceId = v
}

// GetPlatformName returns the PlatformName field value
func (o *ServerGetServerDetailResponseBody) GetPlatformName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlatformName
}

// GetPlatformNameOk returns a tuple with the PlatformName field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetPlatformNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlatformName, true
}

// SetPlatformName sets field value
func (o *ServerGetServerDetailResponseBody) SetPlatformName(v string) {
	o.PlatformName = v
}

// GetServerIpv4 returns the ServerIpv4 field value
func (o *ServerGetServerDetailResponseBody) GetServerIpv4() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerIpv4
}

// GetServerIpv4Ok returns a tuple with the ServerIpv4 field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetServerIpv4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerIpv4, true
}

// SetServerIpv4 sets field value
func (o *ServerGetServerDetailResponseBody) SetServerIpv4(v string) {
	o.ServerIpv4 = v
}

// GetServerName returns the ServerName field value
func (o *ServerGetServerDetailResponseBody) GetServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerName, true
}

// SetServerName sets field value
func (o *ServerGetServerDetailResponseBody) SetServerName(v string) {
	o.ServerName = v
}

// GetServerUuid returns the ServerUuid field value
func (o *ServerGetServerDetailResponseBody) GetServerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerUuid
}

// GetServerUuidOk returns a tuple with the ServerUuid field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetServerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerUuid, true
}

// SetServerUuid sets field value
func (o *ServerGetServerDetailResponseBody) SetServerUuid(v string) {
	o.ServerUuid = v
}

// GetServerroleId returns the ServerroleId field value
func (o *ServerGetServerDetailResponseBody) GetServerroleId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServerroleId
}

// GetServerroleIdOk returns a tuple with the ServerroleId field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetServerroleIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerroleId, true
}

// SetServerroleId sets field value
func (o *ServerGetServerDetailResponseBody) SetServerroleId(v int64) {
	o.ServerroleId = v
}

// GetServerroleName returns the ServerroleName field value
func (o *ServerGetServerDetailResponseBody) GetServerroleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerroleName
}

// GetServerroleNameOk returns a tuple with the ServerroleName field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetServerroleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerroleName, true
}

// SetServerroleName sets field value
func (o *ServerGetServerDetailResponseBody) SetServerroleName(v string) {
	o.ServerroleName = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ServerGetServerDetailResponseBody) GetTags() []ServerTagResponseBody {
	if o == nil || o.Tags == nil {
		var ret []ServerTagResponseBody
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetTagsOk() ([]ServerTagResponseBody, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ServerGetServerDetailResponseBody) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ServerTagResponseBody and assigns it to the Tags field.
func (o *ServerGetServerDetailResponseBody) SetTags(v []ServerTagResponseBody) {
	o.Tags = v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *ServerGetServerDetailResponseBody) GetTasks() []ChildTaskResponseBody {
	if o == nil || o.Tasks == nil {
		var ret []ChildTaskResponseBody
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetTasksOk() ([]ChildTaskResponseBody, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *ServerGetServerDetailResponseBody) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []ChildTaskResponseBody and assigns it to the Tasks field.
func (o *ServerGetServerDetailResponseBody) SetTasks(v []ChildTaskResponseBody) {
	o.Tasks = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ServerGetServerDetailResponseBody) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailResponseBody) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ServerGetServerDetailResponseBody) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o ServerGetServerDetailResponseBody) MarshalJSON() ([]byte, error) {
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
		toSerialize["platformInstanceId"] = o.PlatformInstanceId
	}
	if true {
		toSerialize["platformName"] = o.PlatformName
	}
	if true {
		toSerialize["serverIpv4"] = o.ServerIpv4
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
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableServerGetServerDetailResponseBody struct {
	value *ServerGetServerDetailResponseBody
	isSet bool
}

func (v NullableServerGetServerDetailResponseBody) Get() *ServerGetServerDetailResponseBody {
	return v.value
}

func (v *NullableServerGetServerDetailResponseBody) Set(val *ServerGetServerDetailResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableServerGetServerDetailResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableServerGetServerDetailResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerGetServerDetailResponseBody(val *ServerGetServerDetailResponseBody) *NullableServerGetServerDetailResponseBody {
	return &NullableServerGetServerDetailResponseBody{value: val, isSet: true}
}

func (v NullableServerGetServerDetailResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerGetServerDetailResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


