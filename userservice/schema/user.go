package schema

import (
	"fmt"
	"time"
)

// NotificationType is a type used as a enum type to store all the notification types
type NotificationType string

// All the constants used to denote the notification types
const (
	EMAIL NotificationType = "Email"
	PUSH  NotificationType = "Push"
	SMS   NotificationType = "SMS"
	PHONE NotificationType = "Phone"
)

// Contact struct contains the phone number details of a particular user
type Contact struct {
	DialCode    string `bson:"dialCode" json:"dial_code"`
	PhoneNumber string `bson:"phoneNumber" json:"phone_number"`
}

// User struct maps to a single user of the mongoDB schema
type User struct {
	ID        int     `bson:"_id,omitempty" json:"id"`
	FirstName string  `bson:"firstName" json:"first_name"`
	LastName  string  `bson:"lastName" json:"last_name"`
	Email     string  `bson:"email" json:"email"`
	Contact   Contact `bson:"contact" json:"contact"`

	Password string `bson:"password" json:"-"`
	// ResetAt tells the time when the user reset their password
	// This data is used to invalidate all the tokens that were
	// issued before resetting
	ResetAt       time.Time `bson:"resetAt,omitempty" json:"-"`
	EmailVerified bool      `bson:"emailVerified" json:"email_verified"`
	PhoneVerified bool      `bson:"phoneVerified" json:"phone_verified"`
	// InGracePeriod is used to temporarily enable notifications for users till they verify phonenumber
	Timezone string `bson:"timeZone" json:"time_zone"`
}

func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}
