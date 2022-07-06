package utils

import (
	"encoding/json"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/schema"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ok interface {
	Ok() error
}

// Decode - decodes the request body and extends the validator interface
func Decode(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}

	if payload, ok := v.(ok); ok {
		return payload.Ok()
	}
	return nil
}

// JustDecode just decodes the request body
func JustDecode(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}

	return nil
}

// QUERY DECODER
var Decoder *schema.Decoder

func init() {
	Decoder = schema.NewDecoder()
	Decoder.ZeroEmpty(true)
	Decoder.IgnoreUnknownKeys(true)
	Decoder.RegisterConverter(time.Time{}, parseFilterTime)
	Decoder.RegisterConverter(primitive.NilObjectID, parseFilterObjectID)
}

func parseFilterTime(date string) reflect.Value {
	if s, err := time.Parse(time.RFC3339, date); err == nil {
		return reflect.ValueOf(s)
	}

	return reflect.Value{}
}

func parseFilterObjectID(id string) reflect.Value {
	if s, err := primitive.ObjectIDFromHex(id); err == nil {
		return reflect.ValueOf(s)
	}

	return reflect.Value{}
}
