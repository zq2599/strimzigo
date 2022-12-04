# \SeekApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Seek**](SeekApi.md#Seek) | **Post** /consumers/{groupid}/instances/{name}/positions | 
[**SeekToBeginning**](SeekApi.md#SeekToBeginning) | **Post** /consumers/{groupid}/instances/{name}/positions/beginning | 
[**SeekToEnd**](SeekApi.md#SeekToEnd) | **Post** /consumers/{groupid}/instances/{name}/positions/end | 



## Seek

> Seek(ctx, groupid, name).OffsetCommitSeekList(offsetCommitSeekList).Execute()





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
    groupid := "groupid_example" // string | ID of the consumer group to which the consumer belongs.
    name := "name_example" // string | Name of the subscribed consumer.
    offsetCommitSeekList := *openapiclient.NewOffsetCommitSeekList() // OffsetCommitSeekList | List of partition offsets from which the subscribed consumer will next fetch records.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeekApi.Seek(context.Background(), groupid, name).OffsetCommitSeekList(offsetCommitSeekList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeekApi.Seek``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupid** | **string** | ID of the consumer group to which the consumer belongs. | 
**name** | **string** | Name of the subscribed consumer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSeekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offsetCommitSeekList** | [**OffsetCommitSeekList**](OffsetCommitSeekList.md) | List of partition offsets from which the subscribed consumer will next fetch records. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.kafka.v2+json
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeekToBeginning

> SeekToBeginning(ctx, groupid, name).Partitions(partitions).Execute()





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
    groupid := "groupid_example" // string | ID of the consumer group to which the subscribed consumer belongs.
    name := "name_example" // string | Name of the subscribed consumer.
    partitions := *openapiclient.NewPartitions() // Partitions | List of topic partitions to which the consumer is subscribed. The consumer will seek the first offset in the specified partitions.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeekApi.SeekToBeginning(context.Background(), groupid, name).Partitions(partitions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeekApi.SeekToBeginning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupid** | **string** | ID of the consumer group to which the subscribed consumer belongs. | 
**name** | **string** | Name of the subscribed consumer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSeekToBeginningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **partitions** | [**Partitions**](Partitions.md) | List of topic partitions to which the consumer is subscribed. The consumer will seek the first offset in the specified partitions. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.kafka.v2+json
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeekToEnd

> SeekToEnd(ctx, groupid, name).Partitions(partitions).Execute()





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
    groupid := "groupid_example" // string | ID of the consumer group to which the subscribed consumer belongs.
    name := "name_example" // string | Name of the subscribed consumer.
    partitions := *openapiclient.NewPartitions() // Partitions | List of topic partitions to which the consumer is subscribed. The consumer will seek the last offset in the specified partitions.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeekApi.SeekToEnd(context.Background(), groupid, name).Partitions(partitions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeekApi.SeekToEnd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupid** | **string** | ID of the consumer group to which the subscribed consumer belongs. | 
**name** | **string** | Name of the subscribed consumer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSeekToEndRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **partitions** | [**Partitions**](Partitions.md) | List of topic partitions to which the consumer is subscribed. The consumer will seek the last offset in the specified partitions. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.kafka.v2+json
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

