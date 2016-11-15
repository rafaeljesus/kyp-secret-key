package handlers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecreyKeyCreate(t *testing.T) {
	e := echo.New()
	req, _ := http.NewRequest(echo.POST, "/v1/secret-key", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	if assert.NoError(t, SecretKeyCreate(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
