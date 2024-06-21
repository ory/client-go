/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.11.11
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the AddProjectToWorkspaceBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddProjectToWorkspaceBody{}

// AddProjectToWorkspaceBody struct for AddProjectToWorkspaceBody
type AddProjectToWorkspaceBody struct {
	// The environment of the project in the workspace. Can be one of \"prod\" or \"dev\". Note that the number of projects in the \"prod\" environment is limited depending on the subscription. prod Production stage Staging dev Development
	Environment string `json:"environment"`
	// The action to take with the project subscription. Can be one of \"migrate\", and \"ignore\". \"migrate\" will migrate the project subscription to the workspace. \"ignore\" will ignore the project subscription. migrate ProjectSubscriptionActionMigrate  ProjectSubscriptionActionMigrate will migrate the project subscription to the  workspace. ignore ProjectSubscriptionActionIgnore  ProjectSubscriptionActionIgnore will ignore the project subscription.
	ProjectSubscription string `json:"project_subscription"`
	AdditionalProperties map[string]interface{}
}

type _AddProjectToWorkspaceBody AddProjectToWorkspaceBody

// NewAddProjectToWorkspaceBody instantiates a new AddProjectToWorkspaceBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddProjectToWorkspaceBody(environment string, projectSubscription string) *AddProjectToWorkspaceBody {
	this := AddProjectToWorkspaceBody{}
	this.Environment = environment
	this.ProjectSubscription = projectSubscription
	return &this
}

// NewAddProjectToWorkspaceBodyWithDefaults instantiates a new AddProjectToWorkspaceBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddProjectToWorkspaceBodyWithDefaults() *AddProjectToWorkspaceBody {
	this := AddProjectToWorkspaceBody{}
	return &this
}

// GetEnvironment returns the Environment field value
func (o *AddProjectToWorkspaceBody) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *AddProjectToWorkspaceBody) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *AddProjectToWorkspaceBody) SetEnvironment(v string) {
	o.Environment = v
}

// GetProjectSubscription returns the ProjectSubscription field value
func (o *AddProjectToWorkspaceBody) GetProjectSubscription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectSubscription
}

// GetProjectSubscriptionOk returns a tuple with the ProjectSubscription field value
// and a boolean to check if the value has been set.
func (o *AddProjectToWorkspaceBody) GetProjectSubscriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectSubscription, true
}

// SetProjectSubscription sets field value
func (o *AddProjectToWorkspaceBody) SetProjectSubscription(v string) {
	o.ProjectSubscription = v
}

func (o AddProjectToWorkspaceBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddProjectToWorkspaceBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environment"] = o.Environment
	toSerialize["project_subscription"] = o.ProjectSubscription

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AddProjectToWorkspaceBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environment",
		"project_subscription",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAddProjectToWorkspaceBody := _AddProjectToWorkspaceBody{}

	err = json.Unmarshal(data, &varAddProjectToWorkspaceBody)

	if err != nil {
		return err
	}

	*o = AddProjectToWorkspaceBody(varAddProjectToWorkspaceBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "environment")
		delete(additionalProperties, "project_subscription")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAddProjectToWorkspaceBody struct {
	value *AddProjectToWorkspaceBody
	isSet bool
}

func (v NullableAddProjectToWorkspaceBody) Get() *AddProjectToWorkspaceBody {
	return v.value
}

func (v *NullableAddProjectToWorkspaceBody) Set(val *AddProjectToWorkspaceBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAddProjectToWorkspaceBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAddProjectToWorkspaceBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddProjectToWorkspaceBody(val *AddProjectToWorkspaceBody) *NullableAddProjectToWorkspaceBody {
	return &NullableAddProjectToWorkspaceBody{value: val, isSet: true}
}

func (v NullableAddProjectToWorkspaceBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddProjectToWorkspaceBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

