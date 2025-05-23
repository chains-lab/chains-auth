/*
Cifra SSO REST API

SSO REST API for Cifra app

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta{}

// ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta struct for ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta
type ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta struct {
	// Parameter is the name of the query parameter or request body field that caused the error
	Parameter string `json:"parameter"`
	// Pointer is a JSON Pointer [RFC6901] to the associated entity in the request document. It may be empty if no specific entity is identified.
	Pointer string `json:"pointer"`
}

type _ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta

// NewReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta instantiates a new ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta(parameter string, pointer string) *ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta {
	this := ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta{}
	this.Parameter = parameter
	this.Pointer = pointer
	return &this
}

// NewReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMetaWithDefaults instantiates a new ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMetaWithDefaults() *ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta {
	this := ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta{}
	return &this
}

// GetParameter returns the Parameter field value
func (o *ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta) GetParameter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value
// and a boolean to check if the value has been set.
func (o *ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta) GetParameterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameter, true
}

// SetParameter sets field value
func (o *ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta) SetParameter(v string) {
	o.Parameter = v
}

// GetPointer returns the Pointer field value
func (o *ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta) GetPointer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pointer
}

// GetPointerOk returns a tuple with the Pointer field value
// and a boolean to check if the value has been set.
func (o *ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta) GetPointerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pointer, true
}

// SetPointer sets field value
func (o *ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta) SetPointer(v string) {
	o.Pointer = v
}

func (o ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parameter"] = o.Parameter
	toSerialize["pointer"] = o.Pointer
	return toSerialize, nil
}

func (o *ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"parameter",
		"pointer",
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

	varReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta := _ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta)

	if err != nil {
		return err
	}

	*o = ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta(varReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta)

	return err
}

type NullableReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta struct {
	value *ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta
	isSet bool
}

func (v NullableReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta) Get() *ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta {
	return v.value
}

func (v *NullableReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta) Set(val *ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta(val *ReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta) *NullableReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta {
	return &NullableReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta{value: val, isSet: true}
}

func (v NullableReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReNewsChainsAuthV1OwnGet400ResponseErrorsInnerMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


