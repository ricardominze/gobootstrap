package restapi

import (
	"errors"
	"net/http"
)

func RestrictMethod(r *http.Request, method string) error {

	if r.Method != method {
		return errors.New("Method Not Allowed")
	}
	return nil
}
