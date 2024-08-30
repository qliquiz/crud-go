package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/qliquiz/crud-go/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	s := NewServer(teststore.New())
	s.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, http.StatusOK)
}
