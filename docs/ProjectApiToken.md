# ProjectApiToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** | The Token&#39;s Name  Set this to help you remember, for example, where you use the token. | 
**OwnerId** | **string** |  | 
**ProjectId** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** | The token&#39;s value | [optional] [readonly] 

## Methods

### NewProjectApiToken

`func NewProjectApiToken(id string, name string, ownerId string, ) *ProjectApiToken`

NewProjectApiToken instantiates a new ProjectApiToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectApiTokenWithDefaults

`func NewProjectApiTokenWithDefaults() *ProjectApiToken`

NewProjectApiTokenWithDefaults instantiates a new ProjectApiToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectApiToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectApiToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectApiToken) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProjectApiToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectApiToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectApiToken) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *ProjectApiToken) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ProjectApiToken) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ProjectApiToken) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetProjectId

`func (o *ProjectApiToken) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectApiToken) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectApiToken) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectApiToken) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetValue

`func (o *ProjectApiToken) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProjectApiToken) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProjectApiToken) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProjectApiToken) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


