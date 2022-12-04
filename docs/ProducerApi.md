# \ProducerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Send**](ProducerApi.md#Send) | **Post** /topics/{topicname} | 
[**SendToPartition**](ProducerApi.md#SendToPartition) | **Post** /topics/{topicname}/partitions/{partitionid} | 



## Send

> OffsetRecordSentList Send(ctx, topicname).ProducerRecordList(producerRecordList).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    topicname := "topicname_example" // string | Name of the topic to send records to or retrieve metadata from.
    producerRecordList := *openapiclient.NewProducerRecordList() // ProducerRecordList | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProducerApi.Send(context.Background(), topicname).ProducerRecordList(producerRecordList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProducerApi.Send``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Send`: OffsetRecordSentList
    fmt.Fprintf(os.Stdout, "Response from `ProducerApi.Send`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicname** | **string** | Name of the topic to send records to or retrieve metadata from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **producerRecordList** | [**ProducerRecordList**](ProducerRecordList.md) |  | 

### Return type

[**OffsetRecordSentList**](OffsetRecordSentList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.kafka.json.v2+json, application/vnd.kafka.binary.v2+json
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendToPartition

> OffsetRecordSentList SendToPartition(ctx, topicname, partitionid).ProducerRecordToPartitionList(producerRecordToPartitionList).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    topicname := "topicname_example" // string | Name of the topic to send records to or retrieve metadata from.
    partitionid := int32(56) // int32 | ID of the partition to send records to or retrieve metadata from.
    producerRecordToPartitionList := *openapiclient.NewProducerRecordToPartitionList() // ProducerRecordToPartitionList | List of records to send to a given topic partition, including a value (required) and a key (optional).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProducerApi.SendToPartition(context.Background(), topicname, partitionid).ProducerRecordToPartitionList(producerRecordToPartitionList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProducerApi.SendToPartition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendToPartition`: OffsetRecordSentList
    fmt.Fprintf(os.Stdout, "Response from `ProducerApi.SendToPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicname** | **string** | Name of the topic to send records to or retrieve metadata from. | 
**partitionid** | **int32** | ID of the partition to send records to or retrieve metadata from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendToPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **producerRecordToPartitionList** | [**ProducerRecordToPartitionList**](ProducerRecordToPartitionList.md) | List of records to send to a given topic partition, including a value (required) and a key (optional). | 

### Return type

[**OffsetRecordSentList**](OffsetRecordSentList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.kafka.json.v2+json, application/vnd.kafka.binary.v2+json
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

