package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConnection(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ping", nil)
	responseRecorder := httptest.NewRecorder()
	Router.ServeHTTP(responseRecorder, req)
	assert.Equal(t, 200, responseRecorder.Code)
	assert.Equal(t, "pong", responseRecorder.Body.String())
}
