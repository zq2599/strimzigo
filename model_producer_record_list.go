/*
Strimzi Kafka Bridge API Reference

The Strimzi Kafka Bridge provides a REST API for integrating HTTP based client applications with a Kafka cluster. You can use the API to create and manage consumers and send and receive records over HTTP rather than the native Kafka protocol. 

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ProducerRecordList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProducerRecordList{}

// ProducerRecordList struct for ProducerRecordList
type ProducerRecordList struct {
	Records []ProducerRecord `json:"records,omitempty"`
}

// NewProducerRecordList instantiates a new ProducerRecordList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProducerRecordList() *ProducerRecordList {
	this := ProducerRecordList{}
	return &this
}

// NewProducerRecordListWithDefaults instantiates a new ProducerRecordList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProducerRecordListWithDefaults() *ProducerRecordList {
	this := ProducerRecordList{}
	return &this
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *ProducerRecordList) GetRecords() []ProducerRecord {
	if o == nil || isNil(o.Records) {
		var ret []ProducerRecord
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProducerRecordList) GetRecordsOk() ([]ProducerRecord, bool) {
	if o == nil || isNil(o.Records) {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *ProducerRecordList) HasRecords() bool {
	if o != nil && !isNil(o.Records) {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []ProducerRecord and assigns it to the Records field.
func (o *ProducerRecordList) SetRecords(v []ProducerRecord) {
	o.Records = v
}

func (o ProducerRecordList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProducerRecordList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Records) {
		toSerialize["records"] = o.Records
	}
	return toSerialize, nil
}

type NullableProducerRecordList struct {
	value *ProducerRecordList
	isSet bool
}

func (v NullableProducerRecordList) Get() *ProducerRecordList {
	return v.value
}

func (v *NullableProducerRecordList) Set(val *ProducerRecordList) {
	v.value = val
	v.isSet = true
}

func (v NullableProducerRecordList) IsSet() bool {
	return v.isSet
}

func (v *NullableProducerRecordList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProducerRecordList(val *ProducerRecordList) *NullableProducerRecordList {
	return &NullableProducerRecordList{value: val, isSet: true}
}

func (v NullableProducerRecordList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProducerRecordList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


