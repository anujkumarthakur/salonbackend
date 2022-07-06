package account

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"userservice/userservice/api/v1/logic"
	"userservice/userservice/errors"
	"userservice/userservice/model"
	"userservice/userservice/respond"
	"userservice/userservice/utils"

	validation "github.com/go-ozzo/ozzo-validation"
)

type twillioRequest struct {
	URL     string
	Payload io.Reader
}

func sendOtp(writer http.ResponseWriter, request *http.Request) *errors.AppError {
	var userOtpRequest model.UserOtpRequest
	decodeErr := utils.Decode(request, &userOtpRequest)
	if decodeErr != nil {
		return errors.BadRequest("error decoding request body")
	}
	err := validation.Validate(userOtpRequest)
	if err != nil {
		respond.BadRequest(writer, nil, err)
		return nil
	}
	verifyQuery := `select email, phone_number from users where email=$1`
	if RowExists(verifyQuery, userOtpRequest.Email, userOtpRequest.Phone) {
		//call twillo service api also response data store in database
		errEmail, isEmailVerified := logic.CheckEmailVerifiedOrNot(userOtpRequest.Email)
		if errEmail != nil {
			log.Println(errEmail)
			return errors.BadRequest("")
		}
		var payloadReq model.RequestBody
		if !isEmailVerified {
			payloadReq = model.RequestBody{
				To:      userOtpRequest.Email,
				Channel: "email",
				Code:    "6589",
			}
			twilloResp, errTwillo := requestSendOTP(payloadReq)
			if errTwillo != nil {
				log.Println("requestSendOTP Error:", errTwillo)
				return errors.BadRequest("Error From requestSendOTP Func..")
			}
			respond.OK(writer, twilloResp, nil)
			//return nil
		}
		errPhone, isPhoneVerified := logic.CheckContactVerifiedOrNot(userOtpRequest.Phone)
		if errPhone != nil {
			log.Println(errPhone)
			return errors.BadRequest("")
		}
		if !isPhoneVerified {
			payloadReq = model.RequestBody{
				To:      userOtpRequest.Phone,
				Channel: "sms/call",
				Code:    "6589",
			}
			twilloResp, errTwillo := requestSendOTP(payloadReq)
			if errTwillo != nil {
				log.Println("requestSendOTP Error:", errTwillo)
				return errors.BadRequest("Error From requestSendOTP Func..")
			}
			respond.OK(writer, twilloResp, nil)
			//return nil
		}

	} else {
		respond.OK(writer, "Inavlid Email and Phone Number", nil)
	}
	return nil
}

func requestSendOTP(ApiRequestBody model.RequestBody) (interface{}, error) {
	var (
		twillioResponse      *model.SendOTPResponse
		twillioErrorResponse *model.TwillioErrorResponse
	)
	url := os.Getenv("TWILLIO_BASE_URL") + os.Getenv("TWILLIO_SID") + "/Verifications"

	var payload *strings.Reader
	channelType := ApiRequestBody.Channel
	if channelType == "email" {
		payload = strings.NewReader("To=" + ApiRequestBody.To + "&Channel=" + ApiRequestBody.Channel)
	} else if channelType == "sms/call" {
		payload = strings.NewReader("To=%2B" + ApiRequestBody.To + "&Channel=" + ApiRequestBody.Channel)
	} else if channelType == "E_InvalidEmail" {
		return nil, errors.BadRequest("Bad Input: To-> InvalidEmailFound")
	}
	requestTwillio := twillioRequest{URL: url, Payload: payload}
	body, err := MakeRequest(requestTwillio)
	if err != nil {
		log.Println("twillio:request.go:RequestSendOTPService:TwillioReqError :", err.Error())
		return nil, err
	}
	err = json.Unmarshal(body, &twillioResponse)
	if err != nil {
		err := json.Unmarshal(body, &twillioErrorResponse)
		if err != nil {
			log.Println("twillio:request.go:RequestSendOTPService:TwillioResponseUnmarshall Error: ", err)
			return nil, err
		}
		log.Println("twillio:request.go:RequestSendOTPService:TwillioResponseUnmarshall: ", err)
		return twillioErrorResponse, err
	}
	return twillioResponse, nil
}

func GetPayloadType(input string) string {
	if len(input) < 3 || len(input) > 254 {
		return "E_InvalidEmail"
	}
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if emailRegex.MatchString(input) {
		return "email"
	} else {
		return "sms/call"
	}
}

func MakeRequest(payloadReq twillioRequest) ([]byte, error) {
	//http client declaration and initialization
	client := &http.Client{}
	//creating request for the url
	req, err := http.NewRequest(http.MethodPost, payloadReq.URL, payloadReq.Payload)
	if err != nil {
		log.Println("twillio:request.go:MakeRequest:TwillioRequestError: ", err)
		return nil, err
	}
	//get environment variables required for twillio authentication
	accountSID := os.Getenv("TWILLIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILLIO_AUTH_TOKEN")
	req.SetBasicAuth(accountSID, authToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//make request
	res, err := client.Do(req)
	if err != nil {
		log.Println("twillio:request.go:MakeRequest:TwillioResponseError: ", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("twillio:request.go:RequestSendOTPService:TwillioResponseBody Unmarshal Error : ", err)
		return nil, err
	}
	return body, nil
}

/*method is used to make request to Twillio to verify otp of the user*/
func requestVerifyOTP(ApiRequestBody model.RequestBody) (interface{}, error) {
	//declaring entities for the request Body
	var (
		twillioResponse      *model.VerifyOTPResponse
		twillioErrorResponse *model.TwillioErrorResponse
	)
	//twillio verificationcheck url
	url := os.Getenv("TWILLIO_BASE_URL") + os.Getenv("TWILLIO_SID") + "/VerificationCheck"
	var payload *strings.Reader
	channelType := ApiRequestBody.Channel
	if channelType == "email" {
		payload = strings.NewReader("To=" + ApiRequestBody.To + "&Code=" + ApiRequestBody.Code)
	} else if channelType == "sms/call" {
		payload = strings.NewReader("To=%2B" + ApiRequestBody.To + "&Code=" + ApiRequestBody.Code)
	} else if channelType == "E_InvalidEmail" {
		return nil, errors.BadRequest("Bad Input: To-> InvalidEmailFound")
	}
	requestTwillio := twillioRequest{URL: url, Payload: payload}
	body, err := MakeRequest(requestTwillio)
	if err != nil {
		log.Println("twillio:request.go:RequestVerifyOTPService:TwillioReqError:", err.Error())
		return nil, err
	}
	err = json.Unmarshal(body, &twillioResponse)
	if err != nil {
		err := json.Unmarshal(body, &twillioErrorResponse)
		if err != nil {
			log.Println("twillio:twillioController.go:verifyOTP:TwillioResponseUnmarshall:", err)
			return nil, err
		}
		log.Println("twillio:twillioController.go:verifyOTP:TwillioResponseUnmarshall: ", err)
		return twillioErrorResponse, err
	}
	return twillioResponse, nil
}

//verify opt by email & phone
func verifyOtp(writer http.ResponseWriter, request *http.Request) *errors.AppError {
	//after verified update user profile - email-true,phone-true, updated -time.now()
	var userOtpRequest model.UserOtpRequest
	var data []interface{}
	decodeErr := utils.Decode(request, &userOtpRequest)
	if decodeErr != nil {
		return errors.BadRequest("error decoding request body")
	}
	err := validation.Validate(userOtpRequest)
	if err != nil {
		respond.BadRequest(writer, nil, err)
		return nil
	}
	if userOtpRequest.Email != "" && userOtpRequest.Phone != "" {
		verifyEmailOtp := model.RequestBody{
			To:      userOtpRequest.Email,
			Code:    "6589",
			Channel: "email",
		}
		otpVerifyEmailResp, errEmailVerify := requestVerifyOTP(verifyEmailOtp)
		if errEmailVerify != nil {
			log.Println(errEmailVerify)
			return errors.BadRequest("")
		}
		verifyPhoneOtp := model.RequestBody{
			To:      userOtpRequest.Phone,
			Code:    "6589",
			Channel: "sms/call",
		}
		otpVerifyPhoneResp, errPhoneVerify := requestVerifyOTP(verifyPhoneOtp)
		if errPhoneVerify != nil {
			log.Println(errPhoneVerify)
			return errors.BadRequest("")
		}
		data = append(data, otpVerifyEmailResp, otpVerifyPhoneResp)
		respond.OK(writer, data, nil)
	}

	return nil
}
