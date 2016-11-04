package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/labstack/echo"
	"net/http"
)

func SecretKeyCreate(c echo.Context) error {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return err
	}

	secret := hex.EncodeToString(b)
	response := map[string]interface{}{"secret_key": secret}

	return c.JSON(http.StatusOK, response)
}
