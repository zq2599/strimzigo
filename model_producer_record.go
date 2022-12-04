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

// checks if the ProducerRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProducerRecord{}

// ProducerRecord struct for ProducerRecord
type ProducerRecord struct {
	Partition *int32 `json:"partition,omitempty"`
	Value NullableProducerRecordValue `json:"value"`
	Key *ProducerRecordKey `json:"key,omitempty"`
	Headers []KafkaHeader `json:"headers,omitempty"`
}

// NewProducerRecord instantiates a new ProducerRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProducerRecord(value NullableProducerRecordValue) *ProducerRecord {
	this := ProducerRecord{}
	this.Value = value
	return &this
}

// NewProducerRecordWithDefaults instantiates a new ProducerRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProducerRecordWithDefaults() *ProducerRecord {
	this := ProducerRecord{}
	return &this
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *ProducerRecord) GetPartition() int32 {
	if o == nil || isNil(o.Partition) {
		var ret int32
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProducerRecord) GetPartitionOk() (*int32, bool) {
	if o == nil || isNil(o.Partition) {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *ProducerRecord) HasPartition() bool {
	if o != nil && !isNil(o.Partition) {
		return true
	}

	return false
}

// SetPartition gets a reference to the given int32 and assigns it to the Partition field.
func (o *ProducerRecord) SetPartition(v int32) {
	o.Partition = &v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for ProducerRecordValue will be returned
func (o *ProducerRecord) GetValue() ProducerRecordValue {
	if o == nil || o.Value.Get() == nil {
		var ret ProducerRecordValue
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProducerRecord) GetValueOk() (*ProducerRecordValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *ProducerRecord) SetValue(v ProducerRecordValue) {
	o.Value.Set(&v)
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ProducerRecord) GetKey() ProducerRecordKey {
	if o == nil || isNil(o.Key) {
		var ret ProducerRecordKey
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProducerRecord) GetKeyOk() (*ProducerRecordKey, bool) {
	if o == nil || isNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ProducerRecord) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given ProducerRecordKey and assigns it to the Key field.
func (o *ProducerRecord) SetKey(v ProducerRecordKey) {
	o.Key = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *ProducerRecord) GetHeaders() []KafkaHeader {
	if o == nil || isNil(o.Headers) {
		var ret []KafkaHeader
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProducerRecord) GetHeadersOk() ([]KafkaHeader, bool) {
	if o == nil || isNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *ProducerRecord) HasHeaders() bool {
	if o != nil && !isNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []KafkaHeader and assigns it to the Headers field.
func (o *ProducerRecord) SetHeaders(v []KafkaHeader) {
	o.Headers = v
}

func (o ProducerRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProducerRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Partition) {
		toSerialize["partition"] = o.Partition
	}
	toSerialize["value"] = o.Value.Get()
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	return toSerialize, nil
}

type NullableProducerRecord struct {
	value *ProducerRecord
	isSet bool
}

func (v NullableProducerRecord) Get() *ProducerRecord {
	return v.value
}

func (v *NullableProducerRecord) Set(val *ProducerRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableProducerRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableProducerRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProducerRecord(val *ProducerRecord) *NullableProducerRecord {
	return &NullableProducerRecord{value: val, isSet: true}
}

func (v NullableProducerRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProducerRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


