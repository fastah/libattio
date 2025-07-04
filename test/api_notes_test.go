/*
Attio API

Testing NotesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package libattio

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	libattio "github.com/fastah/libattio"
)

func Test_libattio_NotesAPIService(t *testing.T) {

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)

	t.Run("Test NotesAPIService V2NotesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NotesAPI.V2NotesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotesAPIService V2NotesNoteIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var noteId string

		resp, httpRes, err := apiClient.NotesAPI.V2NotesNoteIdDelete(context.Background(), noteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotesAPIService V2NotesNoteIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var noteId string

		resp, httpRes, err := apiClient.NotesAPI.V2NotesNoteIdGet(context.Background(), noteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotesAPIService V2NotesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NotesAPI.V2NotesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
