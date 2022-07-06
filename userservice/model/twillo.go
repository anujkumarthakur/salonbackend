package model

import "time"

/*
struct to get the json request body for the twillio apis
*/
type RequestBody struct {
	To      string `json:"to"`
	Code    string `json:"code"`
	Channel string `json:"channel"`
}

/*
struct to send json response after error sending/verifying otp of user
*/
type TwillioErrorResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}

/*
Carrier struct
*/
type Carrier struct {
	MobileCountryCode string      `json:"mobile_country_code"`
	Type              string      `json:"type"`
	ErrorCode         interface{} `json:"error_code"`
	MobileNetworkCode string      `json:"mobile_network_code"`
	Name              string      `json:"name"`
}

/*
Lookup
*/
type Lookup struct {
	Carrier Carrier `json:"carrier"`
}

/*
send code attempts
*/
type SendCodeAttempts struct {
	AttemptSid string `json:"attempt_sid"`
	Channel    string `json:"channel"`
	//Time       time.Time `json:"time"`
}

/*
struct to send json response after successfully sending otp to user
*/
type SendOTPResponse struct {
	Status           string             `json:"status"`
	SendCodeAttempts []SendCodeAttempts `json:"send_code_attempts"`
	To               string             `json:"to"`
	Valid            bool               `json:"valid"`
	Lookup           Lookup             `json:"lookup"`
	Channel          string             `json:"channel"`
}

/*
struct to send json response after successfully verifying otp of user
*/
type VerifyOTPResponse struct {
	Status      string    `json:"status"`
	DateUpdated time.Time `json:"date_updated"`
	To          string    `json:"to"`
	Valid       bool      `json:"valid"`
	DateCreated time.Time `json:"date_created"`
	Channel     string    `json:"channel"`
}
