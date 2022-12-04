/*
Strimzi Kafka Bridge API Reference

Testing SeekApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/zq2599/strimzigo"
)

func Test_openapi_SeekApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SeekApiService Seek", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupid string
		var name string

		httpRes, err := apiClient.SeekApi.Seek(context.Background(), groupid, name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SeekApiService SeekToBeginning", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupid string
		var name string

		httpRes, err := apiClient.SeekApi.SeekToBeginning(context.Background(), groupid, name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SeekApiService SeekToEnd", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupid string
		var name string

		httpRes, err := apiClient.SeekApi.SeekToEnd(context.Background(), groupid, name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
