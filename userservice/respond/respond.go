// Package respond provides utility functions to send response to clients
package respond

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"userservice/userservice/errors"

	log "github.com/sirupsen/logrus"
)

// Response struct contains all the fields needed to respond
// to a particular request
type Response struct {
	StatusCode int
	Data       interface{}
	Headers    map[string]string
}

// SendResponse is a helper function which sends a response with the passed data
func SendResponse(w http.ResponseWriter, statusCode int, data interface{}, headers map[string]string) error {
	return NewResponse(statusCode, data, headers).Send(w)
}

// NewResponse returns a new response object.
func NewResponse(statusCode int, data interface{}, headers map[string]string) *Response {
	return &Response{
		StatusCode: statusCode,
		Data:       data,
		Headers:    headers,
	}
}

// Send sends data encoded to JSON
func (res *Response) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	if res.Headers != nil {
		for key, value := range res.Headers {
			w.Header().Set(key, value)
		}
	}
	w.WriteHeader(res.StatusCode)

	if res.StatusCode != http.StatusNoContent {
		if err := json.NewEncoder(w).Encode(res.Data); err != nil {
			log.Error("respond.send.error: ", err)
			// TODO: handle err, notify developers. Configurable
			// http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	// FIXME: remove errors
	return nil
}

// FIXME: only 2XX responses will be handled from here.
// ALL 4XX & 5XX will be handled by errors thrown by the api handler

// TODO: FIXME: no need to return from the respond pkg
// handle gracefully, nofify developers

// 2xx -------------------------------------------------------------------------

// OK is a helper function used to send response data
// with StatusOK status code (200)
func OK(w http.ResponseWriter, data, meta interface{}) error {
	return SendResponse(w, http.StatusOK, WrapPayload(data, meta), nil)
}

// Created is a helper function used to send response data
// with StatusCreated status code (201)
func Created(w http.ResponseWriter, data, meta interface{}) error {
	return SendResponse(w, http.StatusCreated, WrapPayload(data, meta), nil)
}

// NoContent is a helper function used to send a NoContent Header (204)
// Note : the sent data and meta are ignored.
func NoContent(w http.ResponseWriter, data, meta interface{}) error {
	return SendResponse(w, http.StatusNoContent, nil, nil)
}

// 4xx & 5XX -------------------------------------------------------------------

// Fail write the error response
// Common func to send all the error response
func Fail(w http.ResponseWriter, e *errors.AppError) {
	log.Errorf("StatusCode: %d, Error: %s\n DEBUG: %s\n",
		e.Status, e.Error(), e.Debug)
	SendResponse(w, e.Status, WrapPayload(nil, e), nil)
}

// TODO: remove all the rest

// Unauthorized is a helper function used to send response data
// with StatusUnauthorized status code (401)
func Unauthorized(w http.ResponseWriter, data, meta interface{}) error {
	return SendResponse(w, http.StatusUnauthorized, WrapPayload(data, meta), nil)
}

// PaymentRequired is a helper function used to send response data
// with StatusPaymentRequired status code (402)
func PaymentRequired(w http.ResponseWriter, data, meta interface{}) error {
	return SendResponse(w, http.StatusPaymentRequired, WrapPayload(data, meta), nil)
}

// Forbidden is a helper function used to send response data
// with StatusForbidden status code (403)
func Forbidden(w http.ResponseWriter, data, meta interface{}) error {
	return SendResponse(w, http.StatusForbidden, WrapPayload(data, meta), nil)
}

// BadRequest is a helper function used to send response data
// with StatusBadRequest status code (400)
func BadRequest(w http.ResponseWriter, data, meta interface{}) error {
	return SendResponse(w, http.StatusBadRequest, WrapPayload(data, meta), nil)
}

// NotFound is a helper function used to send response data
// with StatusNotFound status code (404)
func NotFound(w http.ResponseWriter, data, meta interface{}) error {
	return SendResponse(w, http.StatusNotFound, WrapPayload(data, meta), nil)
}

// InternalServerError is a helper function used to send response data
// with StatusInternalServerError status code (500)
func InternalServerError(w http.ResponseWriter, data, meta interface{}) error {
	return SendResponse(w, http.StatusInternalServerError, WrapPayload(data, meta), nil)
}

// response holds the handlerfunc response
type response struct {
	Data interface{} `json:"data,omitempty"`
	Meta Meta        `json:"meta"`
}

// Meta holds the status of the request informations
type Meta struct {
	Status  int         `json:"status_code"`
	Message string      `json:"error_message,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// TODO: can be improved

// WrapPayload is used to create a generic payload for the data
// and the metadata passed
func WrapPayload(data, meta interface{}) JSON {
	x := make(JSON)
	if data != nil {
		x["data"] = data
	}

	if meta != nil {
		x["meta"] = meta
	}

	return x
}

// JSON maps go type to a JSON struct
type JSON map[string]interface{}

// ToReader returns a io.Reader after encoding the map into JSON byte array
func (j JSON) ToReader() io.Reader {
	b := bytes.NewBuffer(nil)
	json.NewEncoder(b).Encode(j)
	return b
}

// Array corresponds to an array of elements with any type
type Array interface{}
