# ProjectNext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** | The name of the project. | [optional] 
**RevisionId** | **string** |  | 
**Services** | [**ProjectServices**](ProjectServices.md) |  | 
**Slug** | **string** | The project&#39;s slug | [readonly] 
**State** | **string** | The state of the project. | [readonly] 

## Methods

### NewProjectNext

`func NewProjectNext(id string, revisionId string, services ProjectServices, slug string, state string, ) *ProjectNext`

NewProjectNext instantiates a new ProjectNext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectNextWithDefaults

`func NewProjectNextWithDefaults() *ProjectNext`

NewProjectNextWithDefaults instantiates a new ProjectNext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectNext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectNext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectNext) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProjectNext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectNext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectNext) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectNext) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRevisionId

`func (o *ProjectNext) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *ProjectNext) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *ProjectNext) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.


### GetServices

`func (o *ProjectNext) GetServices() ProjectServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ProjectNext) GetServicesOk() (*ProjectServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ProjectNext) SetServices(v ProjectServices)`

SetServices sets Services field to given value.


### GetSlug

`func (o *ProjectNext) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectNext) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectNext) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetState

`func (o *ProjectNext) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProjectNext) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProjectNext) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


