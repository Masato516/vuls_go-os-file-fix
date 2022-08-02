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

// PkgCpeGetPkgDetailResponseBody struct for PkgCpeGetPkgDetailResponseBody
type PkgCpeGetPkgDetailResponseBody struct {
	// AffectedProcess list of package
	AffectedProcs []AffectedProcResponseBody `json:"affectedProcs,omitempty"`
	// crated time of package or cpe
	CreatedAt time.Time `json:"createdAt"`
	// ID of package
	Id *int64 `json:"id,omitempty"`
	// Name of package or cpe
	Name string `json:"name"`
	// NeedRestartProcess list of package
	NeedRestartProcs []NeedRestartProcResponseBody `json:"needRestartProcs,omitempty"`
	// New Release of package
	NewRelease *string `json:"newRelease,omitempty"`
	// New Version of package
	NewVersion *string `json:"newVersion,omitempty"`
	// package status of package
	PackageStatuses *map[string]string `json:"packageStatuses,omitempty"`
	// Release of package
	Release string `json:"release"`
	// Repository of package
	Repository *string `json:"repository,omitempty"`
	Server ServerChildResponseBody `json:"server"`
	// updated time of server
	Tasks []ChildTaskResponseBody `json:"tasks,omitempty"`
	// updated time of package or cpe
	UpdatedAt time.Time `json:"updatedAt"`
	// Version of package or cpe
	Version string `json:"version"`
}

// NewPkgCpeGetPkgDetailResponseBody instantiates a new PkgCpeGetPkgDetailResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkgCpeGetPkgDetailResponseBody(createdAt time.Time, name string, release string, server ServerChildResponseBody, updatedAt time.Time, version string) *PkgCpeGetPkgDetailResponseBody {
	this := PkgCpeGetPkgDetailResponseBody{}
	this.CreatedAt = createdAt
	this.Name = name
	this.Release = release
	this.Server = server
	this.UpdatedAt = updatedAt
	this.Version = version
	return &this
}

// NewPkgCpeGetPkgDetailResponseBodyWithDefaults instantiates a new PkgCpeGetPkgDetailResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkgCpeGetPkgDetailResponseBodyWithDefaults() *PkgCpeGetPkgDetailResponseBody {
	this := PkgCpeGetPkgDetailResponseBody{}
	return &this
}

// GetAffectedProcs returns the AffectedProcs field value if set, zero value otherwise.
func (o *PkgCpeGetPkgDetailResponseBody) GetAffectedProcs() []AffectedProcResponseBody {
	if o == nil || o.AffectedProcs == nil {
		var ret []AffectedProcResponseBody
		return ret
	}
	return o.AffectedProcs
}

// GetAffectedProcsOk returns a tuple with the AffectedProcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeGetPkgDetailResponseBody) GetAffectedProcsOk() ([]AffectedProcResponseBody, bool) {
	if o == nil || o.AffectedProcs == nil {
		return nil, false
	}
	return o.AffectedProcs, true
}

// HasAffectedProcs returns a boolean if a field has been set.
func (o *PkgCpeGetPkgDetailResponseBody) HasAffectedProcs() bool {
	if o != nil && o.AffectedProcs != nil {
		return true
	}

	return false
}

// SetAffectedProcs gets a reference to the given []AffectedProcResponseBody and assigns it to the AffectedProcs field.
func (o *PkgCpeGetPkgDetailResponseBody) SetAffectedProcs(v []AffectedProcResponseBody) {
	o.AffectedProcs = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PkgCpeGetPkgDetailResponseBody) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PkgCpeGetPkgDetailResponseBody) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PkgCpeGetPkgDetailResponseBody) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PkgCpeGetPkgDetailResponseBody) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeGetPkgDetailResponseBody) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PkgCpeGetPkgDetailResponseBody) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *PkgCpeGetPkgDetailResponseBody) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *PkgCpeGetPkgDetailResponseBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PkgCpeGetPkgDetailResponseBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PkgCpeGetPkgDetailResponseBody) SetName(v string) {
	o.Name = v
}

// GetNeedRestartProcs returns the NeedRestartProcs field value if set, zero value otherwise.
func (o *PkgCpeGetPkgDetailResponseBody) GetNeedRestartProcs() []NeedRestartProcResponseBody {
	if o == nil || o.NeedRestartProcs == nil {
		var ret []NeedRestartProcResponseBody
		return ret
	}
	return o.NeedRestartProcs
}

// GetNeedRestartProcsOk returns a tuple with the NeedRestartProcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeGetPkgDetailResponseBody) GetNeedRestartProcsOk() ([]NeedRestartProcResponseBody, bool) {
	if o == nil || o.NeedRestartProcs == nil {
		return nil, false
	}
	return o.NeedRestartProcs, true
}

// HasNeedRestartProcs returns a boolean if a field has been set.
func (o *PkgCpeGetPkgDetailResponseBody) HasNeedRestartProcs() bool {
	if o != nil && o.NeedRestartProcs != nil {
		return true
	}

	return false
}

// SetNeedRestartProcs gets a reference to the given []NeedRestartProcResponseBody and assigns it to the NeedRestartProcs field.
func (o *PkgCpeGetPkgDetailResponseBody) SetNeedRestartProcs(v []NeedRestartProcResponseBody) {
	o.NeedRestartProcs = v
}

// GetNewRelease returns the NewRelease field value if set, zero value otherwise.
func (o *PkgCpeGetPkgDetailResponseBody) GetNewRelease() string {
	if o == nil || o.NewRelease == nil {
		var ret string
		return ret
	}
	return *o.NewRelease
}

// GetNewReleaseOk returns a tuple with the NewRelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeGetPkgDetailResponseBody) GetNewReleaseOk() (*string, bool) {
	if o == nil || o.NewRelease == nil {
		return nil, false
	}
	return o.NewRelease, true
}

// HasNewRelease returns a boolean if a field has been set.
func (o *PkgCpeGetPkgDetailResponseBody) HasNewRelease() bool {
	if o != nil && o.NewRelease != nil {
		return true
	}

	return false
}

// SetNewRelease gets a reference to the given string and assigns it to the NewRelease field.
func (o *PkgCpeGetPkgDetailResponseBody) SetNewRelease(v string) {
	o.NewRelease = &v
}

// GetNewVersion returns the NewVersion field value if set, zero value otherwise.
func (o *PkgCpeGetPkgDetailResponseBody) GetNewVersion() string {
	if o == nil || o.NewVersion == nil {
		var ret string
		return ret
	}
	return *o.NewVersion
}

// GetNewVersionOk returns a tuple with the NewVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeGetPkgDetailResponseBody) GetNewVersionOk() (*string, bool) {
	if o == nil || o.NewVersion == nil {
		return nil, false
	}
	return o.NewVersion, true
}

// HasNewVersion returns a boolean if a field has been set.
func (o *PkgCpeGetPkgDetailResponseBody) HasNewVersion() bool {
	if o != nil && o.NewVersion != nil {
		return true
	}

	return false
}

// SetNewVersion gets a reference to the given string and assigns it to the NewVersion field.
func (o *PkgCpeGetPkgDetailResponseBody) SetNewVersion(v string) {
	o.NewVersion = &v
}

// GetPackageStatuses returns the PackageStatuses field value if set, zero value otherwise.
func (o *PkgCpeGetPkgDetailResponseBody) GetPackageStatuses() map[string]string {
	if o == nil || o.PackageStatuses == nil {
		var ret map[string]string
		return ret
	}
	return *o.PackageStatuses
}

// GetPackageStatusesOk returns a tuple with the PackageStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeGetPkgDetailResponseBody) GetPackageStatusesOk() (*map[string]string, bool) {
	if o == nil || o.PackageStatuses == nil {
		return nil, false
	}
	return o.PackageStatuses, true
}

// HasPackageStatuses returns a boolean if a field has been set.
func (o *PkgCpeGetPkgDetailResponseBody) HasPackageStatuses() bool {
	if o != nil && o.PackageStatuses != nil {
		return true
	}

	return false
}

// SetPackageStatuses gets a reference to the given map[string]string and assigns it to the PackageStatuses field.
func (o *PkgCpeGetPkgDetailResponseBody) SetPackageStatuses(v map[string]string) {
	o.PackageStatuses = &v
}

// GetRelease returns the Release field value
func (o *PkgCpeGetPkgDetailResponseBody) GetRelease() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Release
}

// GetReleaseOk returns a tuple with the Release field value
// and a boolean to check if the value has been set.
func (o *PkgCpeGetPkgDetailResponseBody) GetReleaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Release, true
}

// SetRelease sets field value
func (o *PkgCpeGetPkgDetailResponseBody) SetRelease(v string) {
	o.Release = v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *PkgCpeGetPkgDetailResponseBody) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeGetPkgDetailResponseBody) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *PkgCpeGetPkgDetailResponseBody) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *PkgCpeGetPkgDetailResponseBody) SetRepository(v string) {
	o.Repository = &v
}

// GetServer returns the Server field value
func (o *PkgCpeGetPkgDetailResponseBody) GetServer() ServerChildResponseBody {
	if o == nil {
		var ret ServerChildResponseBody
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *PkgCpeGetPkgDetailResponseBody) GetServerOk() (*ServerChildResponseBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *PkgCpeGetPkgDetailResponseBody) SetServer(v ServerChildResponseBody) {
	o.Server = v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *PkgCpeGetPkgDetailResponseBody) GetTasks() []ChildTaskResponseBody {
	if o == nil || o.Tasks == nil {
		var ret []ChildTaskResponseBody
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeGetPkgDetailResponseBody) GetTasksOk() ([]ChildTaskResponseBody, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *PkgCpeGetPkgDetailResponseBody) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []ChildTaskResponseBody and assigns it to the Tasks field.
func (o *PkgCpeGetPkgDetailResponseBody) SetTasks(v []ChildTaskResponseBody) {
	o.Tasks = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PkgCpeGetPkgDetailResponseBody) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PkgCpeGetPkgDetailResponseBody) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PkgCpeGetPkgDetailResponseBody) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetVersion returns the Version field value
func (o *PkgCpeGetPkgDetailResponseBody) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *PkgCpeGetPkgDetailResponseBody) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *PkgCpeGetPkgDetailResponseBody) SetVersion(v string) {
	o.Version = v
}

func (o PkgCpeGetPkgDetailResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AffectedProcs != nil {
		toSerialize["affectedProcs"] = o.AffectedProcs
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NeedRestartProcs != nil {
		toSerialize["needRestartProcs"] = o.NeedRestartProcs
	}
	if o.NewRelease != nil {
		toSerialize["newRelease"] = o.NewRelease
	}
	if o.NewVersion != nil {
		toSerialize["newVersion"] = o.NewVersion
	}
	if o.PackageStatuses != nil {
		toSerialize["packageStatuses"] = o.PackageStatuses
	}
	if true {
		toSerialize["release"] = o.Release
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if true {
		toSerialize["server"] = o.Server
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullablePkgCpeGetPkgDetailResponseBody struct {
	value *PkgCpeGetPkgDetailResponseBody
	isSet bool
}

func (v NullablePkgCpeGetPkgDetailResponseBody) Get() *PkgCpeGetPkgDetailResponseBody {
	return v.value
}

func (v *NullablePkgCpeGetPkgDetailResponseBody) Set(val *PkgCpeGetPkgDetailResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullablePkgCpeGetPkgDetailResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullablePkgCpeGetPkgDetailResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkgCpeGetPkgDetailResponseBody(val *PkgCpeGetPkgDetailResponseBody) *NullablePkgCpeGetPkgDetailResponseBody {
	return &NullablePkgCpeGetPkgDetailResponseBody{value: val, isSet: true}
}

func (v NullablePkgCpeGetPkgDetailResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkgCpeGetPkgDetailResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


