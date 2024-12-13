package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/v3/assert"
)

type Response struct {
	Message string `json:"message"`
}

func TestMainHanlder(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	expectedResponse := Response{Message: "handler one response"}

	router := gin.Default()

	router.GET("/one", HandlerOne)

	req, _ := http.NewRequest("GET", "/one", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	var responseJson Response

	err := json.Unmarshal(responseData, &responseJson)
	assert.NilError(t, err)

	assert.Equal(t, expectedResponse.Message, responseJson.Message)
	assert.Equal(t, http.StatusOK, w.Code)

}
