package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"server/cmd/server/routers"
	"testing"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	_ = os.Setenv("TESTING", "LOCAL")
	_ = os.Setenv("ENV_PATH", "../.env")
	router = routers.RunGinRouter(gin.Default())
	code := m.Run()
	os.Exit(code)
}

func TestConnection(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ping", nil)
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(t, 200, responseRecorder.Code)
	assert.Equal(t, "pong", responseRecorder.Body.String())
}
