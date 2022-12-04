# ProducerRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partition** | Pointer to **int32** |  | [optional] 
**Value** | [**NullableProducerRecordValue**](ProducerRecordValue.md) |  | 
**Key** | Pointer to [**ProducerRecordKey**](ProducerRecordKey.md) |  | [optional] 
**Headers** | Pointer to [**[]KafkaHeader**](KafkaHeader.md) |  | [optional] 

## Methods

### NewProducerRecord

`func NewProducerRecord(value NullableProducerRecordValue, ) *ProducerRecord`

NewProducerRecord instantiates a new ProducerRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProducerRecordWithDefaults

`func NewProducerRecordWithDefaults() *ProducerRecord`

NewProducerRecordWithDefaults instantiates a new ProducerRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartition

`func (o *ProducerRecord) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *ProducerRecord) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *ProducerRecord) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *ProducerRecord) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetValue

`func (o *ProducerRecord) GetValue() ProducerRecordValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProducerRecord) GetValueOk() (*ProducerRecordValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProducerRecord) SetValue(v ProducerRecordValue)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *ProducerRecord) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ProducerRecord) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetKey

`func (o *ProducerRecord) GetKey() ProducerRecordKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProducerRecord) GetKeyOk() (*ProducerRecordKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProducerRecord) SetKey(v ProducerRecordKey)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProducerRecord) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetHeaders

`func (o *ProducerRecord) GetHeaders() []KafkaHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ProducerRecord) GetHeadersOk() (*[]KafkaHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ProducerRecord) SetHeaders(v []KafkaHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ProducerRecord) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


