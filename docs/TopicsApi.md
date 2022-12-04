# \TopicsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOffsets**](TopicsApi.md#GetOffsets) | **Get** /topics/{topicname}/partitions/{partitionid}/offsets | 
[**GetPartition**](TopicsApi.md#GetPartition) | **Get** /topics/{topicname}/partitions/{partitionid} | 
[**GetTopic**](TopicsApi.md#GetTopic) | **Get** /topics/{topicname} | 
[**ListPartitions**](TopicsApi.md#ListPartitions) | **Get** /topics/{topicname}/partitions | 
[**ListTopics**](TopicsApi.md#ListTopics) | **Get** /topics | 
[**Send**](TopicsApi.md#Send) | **Post** /topics/{topicname} | 
[**SendToPartition**](TopicsApi.md#SendToPartition) | **Post** /topics/{topicname}/partitions/{partitionid} | 



## GetOffsets

> OffsetsSummary GetOffsets(ctx, topicname, partitionid).Execute()





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
    topicname := "topicname_example" // string | Name of the topic containing the partition.
    partitionid := int32(56) // int32 | ID of the partition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicsApi.GetOffsets(context.Background(), topicname, partitionid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsApi.GetOffsets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOffsets`: OffsetsSummary
    fmt.Fprintf(os.Stdout, "Response from `TopicsApi.GetOffsets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicname** | **string** | Name of the topic containing the partition. | 
**partitionid** | **int32** | ID of the partition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOffsetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OffsetsSummary**](OffsetsSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartition

> PartitionMetadata GetPartition(ctx, topicname, partitionid).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicsApi.GetPartition(context.Background(), topicname, partitionid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsApi.GetPartition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartition`: PartitionMetadata
    fmt.Fprintf(os.Stdout, "Response from `TopicsApi.GetPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicname** | **string** | Name of the topic to send records to or retrieve metadata from. | 
**partitionid** | **int32** | ID of the partition to send records to or retrieve metadata from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PartitionMetadata**](PartitionMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopic

> TopicMetadata GetTopic(ctx, topicname).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicsApi.GetTopic(context.Background(), topicname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsApi.GetTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopic`: TopicMetadata
    fmt.Fprintf(os.Stdout, "Response from `TopicsApi.GetTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicname** | **string** | Name of the topic to send records to or retrieve metadata from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TopicMetadata**](TopicMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPartitions

> []PartitionMetadata ListPartitions(ctx, topicname).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicsApi.ListPartitions(context.Background(), topicname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsApi.ListPartitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPartitions`: []PartitionMetadata
    fmt.Fprintf(os.Stdout, "Response from `TopicsApi.ListPartitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicname** | **string** | Name of the topic to send records to or retrieve metadata from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PartitionMetadata**](PartitionMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTopics

> []string ListTopics(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicsApi.ListTopics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsApi.ListTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTopics`: []string
    fmt.Fprintf(os.Stdout, "Response from `TopicsApi.ListTopics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTopicsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := apiClient.TopicsApi.Send(context.Background(), topicname).ProducerRecordList(producerRecordList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsApi.Send``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Send`: OffsetRecordSentList
    fmt.Fprintf(os.Stdout, "Response from `TopicsApi.Send`: %v\n", resp)
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
    resp, r, err := apiClient.TopicsApi.SendToPartition(context.Background(), topicname, partitionid).ProducerRecordToPartitionList(producerRecordToPartitionList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsApi.SendToPartition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendToPartition`: OffsetRecordSentList
    fmt.Fprintf(os.Stdout, "Response from `TopicsApi.SendToPartition`: %v\n", resp)
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

