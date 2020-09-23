/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// UserCreateData Object to create a user.
type UserCreateData struct {
	Attributes    UserCreateAttributes `json:"attributes"`
	Relationships *UserRelationships   `json:"relationships,omitempty"`
	Type          UsersType            `json:"type"`
}

// NewUserCreateData instantiates a new UserCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCreateData(attributes UserCreateAttributes, type_ UsersType) *UserCreateData {
	this := UserCreateData{}
	this.Attributes = attributes
	this.Type = type_
	return &this
}

// NewUserCreateDataWithDefaults instantiates a new UserCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCreateDataWithDefaults() *UserCreateData {
	this := UserCreateData{}
	var type_ UsersType = "users"
	this.Type = type_
	return &this
}

// GetAttributes returns the Attributes field value
func (o *UserCreateData) GetAttributes() UserCreateAttributes {
	if o == nil {
		var ret UserCreateAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *UserCreateData) GetAttributesOk() (*UserCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *UserCreateData) SetAttributes(v UserCreateAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *UserCreateData) GetRelationships() UserRelationships {
	if o == nil || o.Relationships == nil {
		var ret UserRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreateData) GetRelationshipsOk() (*UserRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *UserCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given UserRelationships and assigns it to the Relationships field.
func (o *UserCreateData) SetRelationships(v UserRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value
func (o *UserCreateData) GetType() UsersType {
	if o == nil {
		var ret UsersType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserCreateData) GetTypeOk() (*UsersType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserCreateData) SetType(v UsersType) {
	o.Type = v
}

func (o UserCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableUserCreateData struct {
	value *UserCreateData
	isSet bool
}

func (v NullableUserCreateData) Get() *UserCreateData {
	return v.value
}

func (v *NullableUserCreateData) Set(val *UserCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCreateData(val *UserCreateData) *NullableUserCreateData {
	return &NullableUserCreateData{value: val, isSet: true}
}

func (v NullableUserCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
