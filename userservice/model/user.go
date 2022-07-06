package model

type RegisterUserRequest struct {
	FullName string  `json:"full_name"`
	Contact  Contact `json:"contact"`
	Email    string  `json:"email" validate:"required,email"`
	Address  Address `json:"address"`
	Password string  `json:"password" validate:"required,min=10"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type Contact struct {
	DialCode    string `bson:"dialCode" json:"dial_code"`
	PhoneNumber string `bson:"phoneNumber" json:"phone_number"`
}

type RegisterUserResponse struct {
	User struct {
		Id       int     `json:"id"`
		FullName string  `json:"full_name"`
		Email    string  `json:"email"`
		Contact  Contact `json:"contact"`
	} `json:"user"`
}

type UserVerifiedResponse struct {
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
	Message  string `json:"message"`
}

type UserPhoneVerifiedResponse struct {
	Contact  Contact `json:"contact"`
	Verified bool    `json:"verified"`
	Message  string  `json:"message"`
}

type UserOtpRequest struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserEmailLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserEmailLogedResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

// type UserPhoneLoginRequest struct {
// 	Contact Contact `json:"contact"`
// }
