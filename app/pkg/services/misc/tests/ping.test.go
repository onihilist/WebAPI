package tests

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/onihilist/WebAPI/pkg/routes"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	expectedJson := `{"message":"pong","status":` + strconv.Itoa(http.StatusOK) + `}`

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expectedJson, w.Body.String())
}
