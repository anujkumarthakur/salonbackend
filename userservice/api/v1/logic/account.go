package logic

import (
	"fmt"
	"log"
	"time"
	"userservice/userservice/db"
	"userservice/userservice/errors"
	"userservice/userservice/model"
	"userservice/userservice/schema"
)

func RegisterNewUser(user model.RegisterUserRequest) (*schema.User, *errors.AppError) {
	var respUser schema.User
	sqlStatement := `INSERT INTO users (name, dial_code, phone_number, email, street, city, state, created, updated) VALUES ($1, $2, $3, $4, $5, $6, $7,$8,$9) RETURNING id`
	id := 0
	err := db.Psql.QueryRow(sqlStatement, user.FullName, user.Contact.DialCode, user.Contact.PhoneNumber, user.Email, user.Address.State, user.Address.City, user.Address.State, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		log.Println("Error:", err)
		return nil, errors.BadRequest("HandleRegistration QueryError")
	}
	respUser.ID = id
	return &respUser, nil
}

func CheckEmailVerifiedOrNot(email string) (error, bool) {
	var verifyemail bool
	sqlQuery := "select verified_email from users where email=$1"

	err := db.Psql.QueryRow(sqlQuery, email).Scan(&verifyemail)
	if err != nil {
		fmt.Println("Error:", err)
		return err, false
	}
	if !verifyemail {
		return nil, false
	}
	return nil, true
}

func CheckContactVerifiedOrNot(phone string) (error, bool) {
	var verifyephone bool
	sqlQuery := `select verified_phone from users where phone_number=$1`

	err := db.Psql.QueryRow(sqlQuery, phone).Scan(&verifyephone)
	if err != nil {
		return err, false
	}
	if !verifyephone {
		return nil, false
	}
	return nil, true
}
