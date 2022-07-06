package logic

import (
	"fmt"
	"userservice/userservice/db"
	"userservice/userservice/model"
)

func CheckUserAuth(authLoginReq model.UserEmailLoginRequest) (error, bool) {
	var email, password string
	sqlQuery := "select email, password from users where email=$1 and password=$2"

	err := db.Psql.QueryRow(sqlQuery, authLoginReq.Email, authLoginReq.Password).Scan(&email, &password)
	if err != nil {
		fmt.Println("Error:", err)
		return err, false
	}
	if email != authLoginReq.Email && password != authLoginReq.Password {
		return nil, false
	}
	return nil, true
}
