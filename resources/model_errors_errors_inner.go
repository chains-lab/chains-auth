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

// checks if the ErrorsErrorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorsErrorsInner{}

// ErrorsErrorsInner struct for ErrorsErrorsInner
type ErrorsErrorsInner struct {
	// Status is the HTTP status code applicable to this problem
	Status int32 `json:"status"`
	// Title is a short, human-readable summary of the problem
	Title string `json:"title"`
	// Code is an application-specific error code, expressed as a string
	Code string `json:"code"`
	// Detail is a human-readable explanation specific to this occurrence of the problem
	Detail string `json:"detail"`
	Meta ErrorsErrorsInnerMeta `json:"meta"`
}

type _ErrorsErrorsInner ErrorsErrorsInner

// NewErrorsErrorsInner instantiates a new ErrorsErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorsErrorsInner(status int32, title string, code string, detail string, meta ErrorsErrorsInnerMeta) *ErrorsErrorsInner {
	this := ErrorsErrorsInner{}
	this.Status = status
	this.Title = title
	this.Code = code
	this.Detail = detail
	this.Meta = meta
	return &this
}

// NewErrorsErrorsInnerWithDefaults instantiates a new ErrorsErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorsErrorsInnerWithDefaults() *ErrorsErrorsInner {
	this := ErrorsErrorsInner{}
	return &this
}

// GetStatus returns the Status field value
func (o *ErrorsErrorsInner) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ErrorsErrorsInner) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ErrorsErrorsInner) SetStatus(v int32) {
	o.Status = v
}

// GetTitle returns the Title field value
func (o *ErrorsErrorsInner) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ErrorsErrorsInner) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ErrorsErrorsInner) SetTitle(v string) {
	o.Title = v
}

// GetCode returns the Code field value
func (o *ErrorsErrorsInner) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ErrorsErrorsInner) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ErrorsErrorsInner) SetCode(v string) {
	o.Code = v
}

// GetDetail returns the Detail field value
func (o *ErrorsErrorsInner) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *ErrorsErrorsInner) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *ErrorsErrorsInner) SetDetail(v string) {
	o.Detail = v
}

// GetMeta returns the Meta field value
func (o *ErrorsErrorsInner) GetMeta() ErrorsErrorsInnerMeta {
	if o == nil {
		var ret ErrorsErrorsInnerMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ErrorsErrorsInner) GetMetaOk() (*ErrorsErrorsInnerMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *ErrorsErrorsInner) SetMeta(v ErrorsErrorsInnerMeta) {
	o.Meta = v
}

func (o ErrorsErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorsErrorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["title"] = o.Title
	toSerialize["code"] = o.Code
	toSerialize["detail"] = o.Detail
	toSerialize["meta"] = o.Meta
	return toSerialize, nil
}

func (o *ErrorsErrorsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"title",
		"code",
		"detail",
		"meta",
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

	varErrorsErrorsInner := _ErrorsErrorsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorsErrorsInner)

	if err != nil {
		return err
	}

	*o = ErrorsErrorsInner(varErrorsErrorsInner)

	return err
}

type NullableErrorsErrorsInner struct {
	value *ErrorsErrorsInner
	isSet bool
}

func (v NullableErrorsErrorsInner) Get() *ErrorsErrorsInner {
	return v.value
}

func (v *NullableErrorsErrorsInner) Set(val *ErrorsErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorsErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorsErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorsErrorsInner(val *ErrorsErrorsInner) *NullableErrorsErrorsInner {
	return &NullableErrorsErrorsInner{value: val, isSet: true}
}

func (v NullableErrorsErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorsErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


