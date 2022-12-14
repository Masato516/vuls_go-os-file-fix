# TaskUpdateTaskResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvisoryIDs** | Pointer to **[]string** | advisoryIDs of cve | [optional] 
**ApplyingPatchOn** | Pointer to **string** | ApplyingPatchOn of task | [optional] 
**Comments** | Pointer to [**[]TaskCommentResponseBody**](TaskCommentResponseBody.md) | Comment of task | [optional] 
**CreatedAt** | **time.Time** | created time of task | 
**CveID** | **string** | CVE ID of task | 
**Cvss** | Pointer to **map[string]os.File** | Key Value of CveID and Cvss of task | [optional] 
**DetectionMethods** | Pointer to [**[]DetectionMethodResponseBody**](DetectionMethodResponseBody.md) | DetectionMethod of task | [optional] 
**DetectionTools** | Pointer to [**[]DetectionToolResponseBody**](DetectionToolResponseBody.md) | DetectionTools of task | [optional] 
**Id** | **int64** | ID of task | 
**Ignore** | **bool** | Ignore of task | 
**IgnoreUntil** | Pointer to **string** | Ignore until of task | [optional] 
**MainUserID** | Pointer to **int64** | MainUserID of task | [optional] 
**MainUserName** | Pointer to **string** | MainUserName of task | [optional] 
**PackageStatuses** | Pointer to **map[string]string** | packageStatus of task | [optional] 
**PkgCpes** | Pointer to [**[]PkgCpeChildResponseBody**](PkgCpeChildResponseBody.md) | Pcakge And Cpe list of task | [optional] 
**Priority** | **string** | Priority of task | 
**RoleID** | **int64** | ServerRoleID of task | 
**RoleName** | **string** | ServerRoleName of task | 
**Server** | [**ServerChildResponseBody**](ServerChildResponseBody.md) |  | 
**ServerID** | **int64** | ServerID of task | 
**Status** | **string** | Status of task | 
**SubUserID** | Pointer to **int64** | SubUserID of task | [optional] 
**SubUserName** | Pointer to **string** | SubUserName of task | [optional] 
**UpdatedAt** | **time.Time** | updated time of task | 

## Methods

### NewTaskUpdateTaskResponseBody

`func NewTaskUpdateTaskResponseBody(createdAt time.Time, cveID string, id int64, ignore bool, priority string, roleID int64, roleName string, server ServerChildResponseBody, serverID int64, status string, updatedAt time.Time, ) *TaskUpdateTaskResponseBody`

NewTaskUpdateTaskResponseBody instantiates a new TaskUpdateTaskResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskUpdateTaskResponseBodyWithDefaults

`func NewTaskUpdateTaskResponseBodyWithDefaults() *TaskUpdateTaskResponseBody`

NewTaskUpdateTaskResponseBodyWithDefaults instantiates a new TaskUpdateTaskResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisoryIDs

`func (o *TaskUpdateTaskResponseBody) GetAdvisoryIDs() []string`

GetAdvisoryIDs returns the AdvisoryIDs field if non-nil, zero value otherwise.

### GetAdvisoryIDsOk

`func (o *TaskUpdateTaskResponseBody) GetAdvisoryIDsOk() (*[]string, bool)`

GetAdvisoryIDsOk returns a tuple with the AdvisoryIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryIDs

`func (o *TaskUpdateTaskResponseBody) SetAdvisoryIDs(v []string)`

SetAdvisoryIDs sets AdvisoryIDs field to given value.

### HasAdvisoryIDs

`func (o *TaskUpdateTaskResponseBody) HasAdvisoryIDs() bool`

HasAdvisoryIDs returns a boolean if a field has been set.

### GetApplyingPatchOn

`func (o *TaskUpdateTaskResponseBody) GetApplyingPatchOn() string`

GetApplyingPatchOn returns the ApplyingPatchOn field if non-nil, zero value otherwise.

### GetApplyingPatchOnOk

`func (o *TaskUpdateTaskResponseBody) GetApplyingPatchOnOk() (*string, bool)`

GetApplyingPatchOnOk returns a tuple with the ApplyingPatchOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyingPatchOn

`func (o *TaskUpdateTaskResponseBody) SetApplyingPatchOn(v string)`

SetApplyingPatchOn sets ApplyingPatchOn field to given value.

### HasApplyingPatchOn

`func (o *TaskUpdateTaskResponseBody) HasApplyingPatchOn() bool`

HasApplyingPatchOn returns a boolean if a field has been set.

### GetComments

`func (o *TaskUpdateTaskResponseBody) GetComments() []TaskCommentResponseBody`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *TaskUpdateTaskResponseBody) GetCommentsOk() (*[]TaskCommentResponseBody, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *TaskUpdateTaskResponseBody) SetComments(v []TaskCommentResponseBody)`

SetComments sets Comments field to given value.

### HasComments

`func (o *TaskUpdateTaskResponseBody) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskUpdateTaskResponseBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskUpdateTaskResponseBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskUpdateTaskResponseBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCveID

`func (o *TaskUpdateTaskResponseBody) GetCveID() string`

GetCveID returns the CveID field if non-nil, zero value otherwise.

### GetCveIDOk

`func (o *TaskUpdateTaskResponseBody) GetCveIDOk() (*string, bool)`

GetCveIDOk returns a tuple with the CveID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveID

`func (o *TaskUpdateTaskResponseBody) SetCveID(v string)`

SetCveID sets CveID field to given value.


### GetCvss

`func (o *TaskUpdateTaskResponseBody) GetCvss() map[string]os.File`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *TaskUpdateTaskResponseBody) GetCvssOk() (*map[string]os.File, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *TaskUpdateTaskResponseBody) SetCvss(v map[string]os.File)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *TaskUpdateTaskResponseBody) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetDetectionMethods

`func (o *TaskUpdateTaskResponseBody) GetDetectionMethods() []DetectionMethodResponseBody`

GetDetectionMethods returns the DetectionMethods field if non-nil, zero value otherwise.

### GetDetectionMethodsOk

`func (o *TaskUpdateTaskResponseBody) GetDetectionMethodsOk() (*[]DetectionMethodResponseBody, bool)`

GetDetectionMethodsOk returns a tuple with the DetectionMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionMethods

`func (o *TaskUpdateTaskResponseBody) SetDetectionMethods(v []DetectionMethodResponseBody)`

SetDetectionMethods sets DetectionMethods field to given value.

### HasDetectionMethods

`func (o *TaskUpdateTaskResponseBody) HasDetectionMethods() bool`

HasDetectionMethods returns a boolean if a field has been set.

### GetDetectionTools

`func (o *TaskUpdateTaskResponseBody) GetDetectionTools() []DetectionToolResponseBody`

GetDetectionTools returns the DetectionTools field if non-nil, zero value otherwise.

### GetDetectionToolsOk

`func (o *TaskUpdateTaskResponseBody) GetDetectionToolsOk() (*[]DetectionToolResponseBody, bool)`

GetDetectionToolsOk returns a tuple with the DetectionTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionTools

`func (o *TaskUpdateTaskResponseBody) SetDetectionTools(v []DetectionToolResponseBody)`

SetDetectionTools sets DetectionTools field to given value.

### HasDetectionTools

`func (o *TaskUpdateTaskResponseBody) HasDetectionTools() bool`

HasDetectionTools returns a boolean if a field has been set.

### GetId

`func (o *TaskUpdateTaskResponseBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskUpdateTaskResponseBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskUpdateTaskResponseBody) SetId(v int64)`

SetId sets Id field to given value.


### GetIgnore

`func (o *TaskUpdateTaskResponseBody) GetIgnore() bool`

GetIgnore returns the Ignore field if non-nil, zero value otherwise.

### GetIgnoreOk

`func (o *TaskUpdateTaskResponseBody) GetIgnoreOk() (*bool, bool)`

GetIgnoreOk returns a tuple with the Ignore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnore

`func (o *TaskUpdateTaskResponseBody) SetIgnore(v bool)`

SetIgnore sets Ignore field to given value.


### GetIgnoreUntil

`func (o *TaskUpdateTaskResponseBody) GetIgnoreUntil() string`

GetIgnoreUntil returns the IgnoreUntil field if non-nil, zero value otherwise.

### GetIgnoreUntilOk

`func (o *TaskUpdateTaskResponseBody) GetIgnoreUntilOk() (*string, bool)`

GetIgnoreUntilOk returns a tuple with the IgnoreUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreUntil

`func (o *TaskUpdateTaskResponseBody) SetIgnoreUntil(v string)`

SetIgnoreUntil sets IgnoreUntil field to given value.

### HasIgnoreUntil

`func (o *TaskUpdateTaskResponseBody) HasIgnoreUntil() bool`

HasIgnoreUntil returns a boolean if a field has been set.

### GetMainUserID

`func (o *TaskUpdateTaskResponseBody) GetMainUserID() int64`

GetMainUserID returns the MainUserID field if non-nil, zero value otherwise.

### GetMainUserIDOk

`func (o *TaskUpdateTaskResponseBody) GetMainUserIDOk() (*int64, bool)`

GetMainUserIDOk returns a tuple with the MainUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainUserID

`func (o *TaskUpdateTaskResponseBody) SetMainUserID(v int64)`

SetMainUserID sets MainUserID field to given value.

### HasMainUserID

`func (o *TaskUpdateTaskResponseBody) HasMainUserID() bool`

HasMainUserID returns a boolean if a field has been set.

### GetMainUserName

`func (o *TaskUpdateTaskResponseBody) GetMainUserName() string`

GetMainUserName returns the MainUserName field if non-nil, zero value otherwise.

### GetMainUserNameOk

`func (o *TaskUpdateTaskResponseBody) GetMainUserNameOk() (*string, bool)`

GetMainUserNameOk returns a tuple with the MainUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainUserName

`func (o *TaskUpdateTaskResponseBody) SetMainUserName(v string)`

SetMainUserName sets MainUserName field to given value.

### HasMainUserName

`func (o *TaskUpdateTaskResponseBody) HasMainUserName() bool`

HasMainUserName returns a boolean if a field has been set.

### GetPackageStatuses

`func (o *TaskUpdateTaskResponseBody) GetPackageStatuses() map[string]string`

GetPackageStatuses returns the PackageStatuses field if non-nil, zero value otherwise.

### GetPackageStatusesOk

`func (o *TaskUpdateTaskResponseBody) GetPackageStatusesOk() (*map[string]string, bool)`

GetPackageStatusesOk returns a tuple with the PackageStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageStatuses

`func (o *TaskUpdateTaskResponseBody) SetPackageStatuses(v map[string]string)`

SetPackageStatuses sets PackageStatuses field to given value.

### HasPackageStatuses

`func (o *TaskUpdateTaskResponseBody) HasPackageStatuses() bool`

HasPackageStatuses returns a boolean if a field has been set.

### GetPkgCpes

`func (o *TaskUpdateTaskResponseBody) GetPkgCpes() []PkgCpeChildResponseBody`

GetPkgCpes returns the PkgCpes field if non-nil, zero value otherwise.

### GetPkgCpesOk

`func (o *TaskUpdateTaskResponseBody) GetPkgCpesOk() (*[]PkgCpeChildResponseBody, bool)`

GetPkgCpesOk returns a tuple with the PkgCpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgCpes

`func (o *TaskUpdateTaskResponseBody) SetPkgCpes(v []PkgCpeChildResponseBody)`

SetPkgCpes sets PkgCpes field to given value.

### HasPkgCpes

`func (o *TaskUpdateTaskResponseBody) HasPkgCpes() bool`

HasPkgCpes returns a boolean if a field has been set.

### GetPriority

`func (o *TaskUpdateTaskResponseBody) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TaskUpdateTaskResponseBody) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TaskUpdateTaskResponseBody) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetRoleID

`func (o *TaskUpdateTaskResponseBody) GetRoleID() int64`

GetRoleID returns the RoleID field if non-nil, zero value otherwise.

### GetRoleIDOk

`func (o *TaskUpdateTaskResponseBody) GetRoleIDOk() (*int64, bool)`

GetRoleIDOk returns a tuple with the RoleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleID

`func (o *TaskUpdateTaskResponseBody) SetRoleID(v int64)`

SetRoleID sets RoleID field to given value.


### GetRoleName

`func (o *TaskUpdateTaskResponseBody) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *TaskUpdateTaskResponseBody) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *TaskUpdateTaskResponseBody) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetServer

`func (o *TaskUpdateTaskResponseBody) GetServer() ServerChildResponseBody`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *TaskUpdateTaskResponseBody) GetServerOk() (*ServerChildResponseBody, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *TaskUpdateTaskResponseBody) SetServer(v ServerChildResponseBody)`

SetServer sets Server field to given value.


### GetServerID

`func (o *TaskUpdateTaskResponseBody) GetServerID() int64`

GetServerID returns the ServerID field if non-nil, zero value otherwise.

### GetServerIDOk

`func (o *TaskUpdateTaskResponseBody) GetServerIDOk() (*int64, bool)`

GetServerIDOk returns a tuple with the ServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerID

`func (o *TaskUpdateTaskResponseBody) SetServerID(v int64)`

SetServerID sets ServerID field to given value.


### GetStatus

`func (o *TaskUpdateTaskResponseBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskUpdateTaskResponseBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskUpdateTaskResponseBody) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubUserID

`func (o *TaskUpdateTaskResponseBody) GetSubUserID() int64`

GetSubUserID returns the SubUserID field if non-nil, zero value otherwise.

### GetSubUserIDOk

`func (o *TaskUpdateTaskResponseBody) GetSubUserIDOk() (*int64, bool)`

GetSubUserIDOk returns a tuple with the SubUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubUserID

`func (o *TaskUpdateTaskResponseBody) SetSubUserID(v int64)`

SetSubUserID sets SubUserID field to given value.

### HasSubUserID

`func (o *TaskUpdateTaskResponseBody) HasSubUserID() bool`

HasSubUserID returns a boolean if a field has been set.

### GetSubUserName

`func (o *TaskUpdateTaskResponseBody) GetSubUserName() string`

GetSubUserName returns the SubUserName field if non-nil, zero value otherwise.

### GetSubUserNameOk

`func (o *TaskUpdateTaskResponseBody) GetSubUserNameOk() (*string, bool)`

GetSubUserNameOk returns a tuple with the SubUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubUserName

`func (o *TaskUpdateTaskResponseBody) SetSubUserName(v string)`

SetSubUserName sets SubUserName field to given value.

### HasSubUserName

`func (o *TaskUpdateTaskResponseBody) HasSubUserName() bool`

HasSubUserName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TaskUpdateTaskResponseBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskUpdateTaskResponseBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskUpdateTaskResponseBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


