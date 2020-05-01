package util

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

// DecodeJSONBody helper function to decode json body
func DecodeJSONBody(r *http.Request, dst interface{}) error {
	if r.Header.Get("Content-Type") != "" {
		value := r.Header.Get("Content-Type")
		if value != "application/json" {
			return NewBadRequestError().WithContext("Content-Type header is not application/json")
		}
	}

	decoder := jsoniter.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		e := NewBadRequestError(err)

		switch {
		case errors.As(err, &syntaxError):
			return e.WithContext("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return e.WithContext("Request body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			return e.WithContext("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return e.WithContext("Request body contains unknown field %s", fieldName)

		case errors.Is(err, io.EOF):
			return e.WithContext("Request body must not be empty")

		default:
			return WrapError(err, "Error while parsing json body")
		}
	}

	if decoder.More() {
		// TODO: attach bad request error handling
		return NewBadRequestError().WithContext("Request body must only contain a single JSON object")
	}

	return nil
}

type Response struct {
	Code    int         `json:"code,omitempty"`
	Success bool        `json:"success, omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// OkResponse is a Helper function to generate OK return
func OkResponse(w http.ResponseWriter, data interface{}) error {
	json, err := jsoniter.Marshal(Response{
		Code:    http.StatusOK,
		Success: true,
		Message: "success",
		Data:    data,
	})

	if err != nil {
		return err
	}

	_, err = w.Write(json)

	if err != nil {
		return err
	}

	return nil
}

func ErrorResponse(w http.ResponseWriter, statusCode int, err error) error {
	json, err := jsoniter.Marshal(Response{
		Code:    statusCode,
		Message: err.Error(),
		Success: false,
	})

	if err != nil {
		return err
	}

	_, err = w.Write(json)

	if err != nil {
		return err
	}

	return nil
}
