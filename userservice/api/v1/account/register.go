package account

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"userservice/userservice/api/v1/logic"
	"userservice/userservice/api/v1/middleware"
	"userservice/userservice/db"
	"userservice/userservice/errors"
	"userservice/userservice/model"
	"userservice/userservice/respond"
	"userservice/userservice/utils"

	validation "github.com/go-ozzo/ozzo-validation"
)

func registerUser(writer http.ResponseWriter, request *http.Request) *errors.AppError {
	var registerUserRequest model.RegisterUserRequest
	decodeErr := utils.Decode(request, &registerUserRequest)
	if decodeErr != nil {
		return errors.BadRequest("error decoding request body")
	}
	err := validation.Validate(registerUserRequest)
	if err != nil {
		respond.BadRequest(writer, nil, err)
		return nil
	}
	if !RowExists("SELECT email FROM users WHERE email=$1", registerUserRequest.Email) {
		registeredUser, registerErr := logic.RegisterNewUser(registerUserRequest)
		if registerErr != nil {
			return registerErr
		}
		respond.Created(writer, model.RegisterUserResponse{
			User: struct {
				Id       int           `json:"id"`
				FullName string        `json:"full_name"`
				Email    string        `json:"email"`
				Contact  model.Contact `json:"contact"`
			}{
				Id:       registeredUser.ID,
				FullName: registerUserRequest.FullName,
				Email:    registerUserRequest.Email,
				Contact:  registerUserRequest.Contact,
			},
		}, nil)
	} else {
		errEmailVerify, isEmailVerified := logic.CheckEmailVerifiedOrNot(registerUserRequest.Email)
		if errEmailVerify != nil {
			log.Println(errEmailVerify)
			return errors.BadRequest("Error From CheckEmailVerifiedOrNot")
		}
		if !isEmailVerified {
			respond.OK(writer, model.UserVerifiedResponse{
				Email:    registerUserRequest.Email,
				Verified: false,
				Message:  "Email Not Verified",
			}, nil)
			return nil

		}
		errContactVerify, isContactVerified := logic.CheckContactVerifiedOrNot(registerUserRequest.Contact.PhoneNumber)
		if errContactVerify != nil {
			log.Println(errContactVerify)
			return errors.BadRequest("Error From CheckEmailVerifiedOrNot")
		}
		if !isContactVerified {
			respond.OK(writer, model.UserPhoneVerifiedResponse{
				Contact:  registerUserRequest.Contact,
				Verified: false,
				Message:  "Phone Not Verified",
			}, nil)
			return nil

		}
	}

	return nil
}

func RowExists(query string, args ...interface{}) bool {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := db.Psql.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("error checking if row exists '%s' %v", args, err)
	}
	return exists
}

func authEmailLogin(writer http.ResponseWriter, request *http.Request) *errors.AppError {
	var authEmailRequest model.UserEmailLoginRequest
	decodeErr := utils.Decode(request, &authEmailRequest)
	if decodeErr != nil {
		return errors.BadRequest("error decoding request body")
	}
	err := validation.Validate(authEmailRequest)
	if err != nil {
		respond.BadRequest(writer, nil, err)
		return nil
	}
	_, okAuth := logic.CheckUserAuth(authEmailRequest)
	if okAuth {
		_, okEmailVerified := logic.CheckEmailVerifiedOrNot(authEmailRequest.Email)
		if okEmailVerified {
			token, err := middleware.GenerateJWT(authEmailRequest.Email)
			if err != nil {
				return errors.Forbidden("Internal Server Error")
			}
			respond.OK(writer, model.UserEmailLogedResponse{
				Email: authEmailRequest.Email,
				Token: token,
			}, nil)
			return nil
		}
	}
	return nil
}
func resetPassword(writer http.ResponseWriter, request *http.Request) *errors.AppError {
	return nil
}
