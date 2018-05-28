package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	rec := httptest.NewRecorder()
	msg := "Error message"

	Error(rec, http.StatusBadRequest, msg)
	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	b, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, msg, string(bytes.TrimSpace(b)))
}

func TestJSON(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		rec := httptest.NewRecorder()
		expected := "Some message"

		JSON(rec, http.StatusOK, expected)
		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
		b, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		var received string
		err = json.Unmarshal(b, &received)
		assert.NoError(t, err)
		assert.Equal(t, expected, received)
	})

	t.Run("MarshalError", func(t *testing.T) {
		rec := httptest.NewRecorder()
		// json.Marshal return error for function values
		JSON(rec, http.StatusOK, func() {})
		res := rec.Result()
		defer res.Body.Close()
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})

}
