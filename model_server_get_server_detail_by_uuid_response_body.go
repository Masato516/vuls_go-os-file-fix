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

// ServerGetServerDetailByUUIDResponseBody struct for ServerGetServerDetailByUUIDResponseBody
type ServerGetServerDetailByUUIDResponseBody struct {
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

// NewServerGetServerDetailByUUIDResponseBody instantiates a new ServerGetServerDetailByUUIDResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerGetServerDetailByUUIDResponseBody(createdAt string, hostUuid string, id int64, needKernelRestart bool, osFamily string, osVersion string, platformInstanceId string, platformName string, serverIpv4 string, serverName string, serverUuid string, serverroleId int64, serverroleName string, updatedAt string) *ServerGetServerDetailByUUIDResponseBody {
	this := ServerGetServerDetailByUUIDResponseBody{}
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

// NewServerGetServerDetailByUUIDResponseBodyWithDefaults instantiates a new ServerGetServerDetailByUUIDResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerGetServerDetailByUUIDResponseBodyWithDefaults() *ServerGetServerDetailByUUIDResponseBody {
	this := ServerGetServerDetailByUUIDResponseBody{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *ServerGetServerDetailByUUIDResponseBody) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ServerGetServerDetailByUUIDResponseBody) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDefaultUserId returns the DefaultUserId field value if set, zero value otherwise.
func (o *ServerGetServerDetailByUUIDResponseBody) GetDefaultUserId() int64 {
	if o == nil || o.DefaultUserId == nil {
		var ret int64
		return ret
	}
	return *o.DefaultUserId
}

// GetDefaultUserIdOk returns a tuple with the DefaultUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetDefaultUserIdOk() (*int64, bool) {
	if o == nil || o.DefaultUserId == nil {
		return nil, false
	}
	return o.DefaultUserId, true
}

// HasDefaultUserId returns a boolean if a field has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) HasDefaultUserId() bool {
	if o != nil && o.DefaultUserId != nil {
		return true
	}

	return false
}

// SetDefaultUserId gets a reference to the given int64 and assigns it to the DefaultUserId field.
func (o *ServerGetServerDetailByUUIDResponseBody) SetDefaultUserId(v int64) {
	o.DefaultUserId = &v
}

// GetDefaultUserName returns the DefaultUserName field value if set, zero value otherwise.
func (o *ServerGetServerDetailByUUIDResponseBody) GetDefaultUserName() string {
	if o == nil || o.DefaultUserName == nil {
		var ret string
		return ret
	}
	return *o.DefaultUserName
}

// GetDefaultUserNameOk returns a tuple with the DefaultUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetDefaultUserNameOk() (*string, bool) {
	if o == nil || o.DefaultUserName == nil {
		return nil, false
	}
	return o.DefaultUserName, true
}

// HasDefaultUserName returns a boolean if a field has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) HasDefaultUserName() bool {
	if o != nil && o.DefaultUserName != nil {
		return true
	}

	return false
}

// SetDefaultUserName gets a reference to the given string and assigns it to the DefaultUserName field.
func (o *ServerGetServerDetailByUUIDResponseBody) SetDefaultUserName(v string) {
	o.DefaultUserName = &v
}

// GetHostUuid returns the HostUuid field value
func (o *ServerGetServerDetailByUUIDResponseBody) GetHostUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostUuid
}

// GetHostUuidOk returns a tuple with the HostUuid field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetHostUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostUuid, true
}

// SetHostUuid sets field value
func (o *ServerGetServerDetailByUUIDResponseBody) SetHostUuid(v string) {
	o.HostUuid = v
}

// GetId returns the Id field value
func (o *ServerGetServerDetailByUUIDResponseBody) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServerGetServerDetailByUUIDResponseBody) SetId(v int64) {
	o.Id = v
}

// GetLastScannedAt returns the LastScannedAt field value if set, zero value otherwise.
func (o *ServerGetServerDetailByUUIDResponseBody) GetLastScannedAt() string {
	if o == nil || o.LastScannedAt == nil {
		var ret string
		return ret
	}
	return *o.LastScannedAt
}

// GetLastScannedAtOk returns a tuple with the LastScannedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetLastScannedAtOk() (*string, bool) {
	if o == nil || o.LastScannedAt == nil {
		return nil, false
	}
	return o.LastScannedAt, true
}

// HasLastScannedAt returns a boolean if a field has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) HasLastScannedAt() bool {
	if o != nil && o.LastScannedAt != nil {
		return true
	}

	return false
}

// SetLastScannedAt gets a reference to the given string and assigns it to the LastScannedAt field.
func (o *ServerGetServerDetailByUUIDResponseBody) SetLastScannedAt(v string) {
	o.LastScannedAt = &v
}

// GetLastUploadedAt returns the LastUploadedAt field value if set, zero value otherwise.
func (o *ServerGetServerDetailByUUIDResponseBody) GetLastUploadedAt() string {
	if o == nil || o.LastUploadedAt == nil {
		var ret string
		return ret
	}
	return *o.LastUploadedAt
}

// GetLastUploadedAtOk returns a tuple with the LastUploadedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetLastUploadedAtOk() (*string, bool) {
	if o == nil || o.LastUploadedAt == nil {
		return nil, false
	}
	return o.LastUploadedAt, true
}

// HasLastUploadedAt returns a boolean if a field has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) HasLastUploadedAt() bool {
	if o != nil && o.LastUploadedAt != nil {
		return true
	}

	return false
}

// SetLastUploadedAt gets a reference to the given string and assigns it to the LastUploadedAt field.
func (o *ServerGetServerDetailByUUIDResponseBody) SetLastUploadedAt(v string) {
	o.LastUploadedAt = &v
}

// GetNeedKernelRestart returns the NeedKernelRestart field value
func (o *ServerGetServerDetailByUUIDResponseBody) GetNeedKernelRestart() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NeedKernelRestart
}

// GetNeedKernelRestartOk returns a tuple with the NeedKernelRestart field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetNeedKernelRestartOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NeedKernelRestart, true
}

// SetNeedKernelRestart sets field value
func (o *ServerGetServerDetailByUUIDResponseBody) SetNeedKernelRestart(v bool) {
	o.NeedKernelRestart = v
}

// GetOsFamily returns the OsFamily field value
func (o *ServerGetServerDetailByUUIDResponseBody) GetOsFamily() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsFamily
}

// GetOsFamilyOk returns a tuple with the OsFamily field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetOsFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsFamily, true
}

// SetOsFamily sets field value
func (o *ServerGetServerDetailByUUIDResponseBody) SetOsFamily(v string) {
	o.OsFamily = v
}

// GetOsVersion returns the OsVersion field value
func (o *ServerGetServerDetailByUUIDResponseBody) GetOsVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetOsVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsVersion, true
}

// SetOsVersion sets field value
func (o *ServerGetServerDetailByUUIDResponseBody) SetOsVersion(v string) {
	o.OsVersion = v
}

// GetPlatformInstanceId returns the PlatformInstanceId field value
func (o *ServerGetServerDetailByUUIDResponseBody) GetPlatformInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlatformInstanceId
}

// GetPlatformInstanceIdOk returns a tuple with the PlatformInstanceId field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetPlatformInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlatformInstanceId, true
}

// SetPlatformInstanceId sets field value
func (o *ServerGetServerDetailByUUIDResponseBody) SetPlatformInstanceId(v string) {
	o.PlatformInstanceId = v
}

// GetPlatformName returns the PlatformName field value
func (o *ServerGetServerDetailByUUIDResponseBody) GetPlatformName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlatformName
}

// GetPlatformNameOk returns a tuple with the PlatformName field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetPlatformNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlatformName, true
}

// SetPlatformName sets field value
func (o *ServerGetServerDetailByUUIDResponseBody) SetPlatformName(v string) {
	o.PlatformName = v
}

// GetServerIpv4 returns the ServerIpv4 field value
func (o *ServerGetServerDetailByUUIDResponseBody) GetServerIpv4() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerIpv4
}

// GetServerIpv4Ok returns a tuple with the ServerIpv4 field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetServerIpv4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerIpv4, true
}

// SetServerIpv4 sets field value
func (o *ServerGetServerDetailByUUIDResponseBody) SetServerIpv4(v string) {
	o.ServerIpv4 = v
}

// GetServerName returns the ServerName field value
func (o *ServerGetServerDetailByUUIDResponseBody) GetServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerName, true
}

// SetServerName sets field value
func (o *ServerGetServerDetailByUUIDResponseBody) SetServerName(v string) {
	o.ServerName = v
}

// GetServerUuid returns the ServerUuid field value
func (o *ServerGetServerDetailByUUIDResponseBody) GetServerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerUuid
}

// GetServerUuidOk returns a tuple with the ServerUuid field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetServerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerUuid, true
}

// SetServerUuid sets field value
func (o *ServerGetServerDetailByUUIDResponseBody) SetServerUuid(v string) {
	o.ServerUuid = v
}

// GetServerroleId returns the ServerroleId field value
func (o *ServerGetServerDetailByUUIDResponseBody) GetServerroleId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServerroleId
}

// GetServerroleIdOk returns a tuple with the ServerroleId field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetServerroleIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerroleId, true
}

// SetServerroleId sets field value
func (o *ServerGetServerDetailByUUIDResponseBody) SetServerroleId(v int64) {
	o.ServerroleId = v
}

// GetServerroleName returns the ServerroleName field value
func (o *ServerGetServerDetailByUUIDResponseBody) GetServerroleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerroleName
}

// GetServerroleNameOk returns a tuple with the ServerroleName field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetServerroleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerroleName, true
}

// SetServerroleName sets field value
func (o *ServerGetServerDetailByUUIDResponseBody) SetServerroleName(v string) {
	o.ServerroleName = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ServerGetServerDetailByUUIDResponseBody) GetTags() []ServerTagResponseBody {
	if o == nil || o.Tags == nil {
		var ret []ServerTagResponseBody
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetTagsOk() ([]ServerTagResponseBody, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ServerTagResponseBody and assigns it to the Tags field.
func (o *ServerGetServerDetailByUUIDResponseBody) SetTags(v []ServerTagResponseBody) {
	o.Tags = v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *ServerGetServerDetailByUUIDResponseBody) GetTasks() []ChildTaskResponseBody {
	if o == nil || o.Tasks == nil {
		var ret []ChildTaskResponseBody
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetTasksOk() ([]ChildTaskResponseBody, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []ChildTaskResponseBody and assigns it to the Tasks field.
func (o *ServerGetServerDetailByUUIDResponseBody) SetTasks(v []ChildTaskResponseBody) {
	o.Tasks = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ServerGetServerDetailByUUIDResponseBody) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ServerGetServerDetailByUUIDResponseBody) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ServerGetServerDetailByUUIDResponseBody) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o ServerGetServerDetailByUUIDResponseBody) MarshalJSON() ([]byte, error) {
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

type NullableServerGetServerDetailByUUIDResponseBody struct {
	value *ServerGetServerDetailByUUIDResponseBody
	isSet bool
}

func (v NullableServerGetServerDetailByUUIDResponseBody) Get() *ServerGetServerDetailByUUIDResponseBody {
	return v.value
}

func (v *NullableServerGetServerDetailByUUIDResponseBody) Set(val *ServerGetServerDetailByUUIDResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableServerGetServerDetailByUUIDResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableServerGetServerDetailByUUIDResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerGetServerDetailByUUIDResponseBody(val *ServerGetServerDetailByUUIDResponseBody) *NullableServerGetServerDetailByUUIDResponseBody {
	return &NullableServerGetServerDetailByUUIDResponseBody{value: val, isSet: true}
}

func (v NullableServerGetServerDetailByUUIDResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerGetServerDetailByUUIDResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


