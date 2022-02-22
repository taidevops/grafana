package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

// Bind deserializes JSON payload from the request
func Bind(req *http.Request, v interface{}) error {
	if req.Body != nil {
		defer func() { _ = req.Body.Close() }()
		err := json.NewDecoder(req.Body).Decode(v)
		if err != nil && !errors.Is(err, io.EOF) {
			return err
		}
	}
	return validate(v)
}

func validate(obj interface{}) error {
	// First check if obj is nil, because we cannot validate those.
	if obj == nil {
		return nil
	}

	// Second, check if obj has a nil interface value.
	// This is to prevent panics when obj is an instance of uninitialised struct pointer / interface.
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		return nil
	}

	// If type has a Validate() method - use that
	if validator, ok := obj.(Validator); ok {

	}
}
