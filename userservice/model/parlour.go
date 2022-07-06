package model

import "time"

type ServiceBookRequest struct {
	ShopName         string    `json:"shop_name"`
	ShopAddress      string    `json:"shop_address"`
	BookTimeFrom     time.Time `json:"book_time_from"`
	BookTimeTo       time.Time `json:"book_time_to"`
	OwnerPhoneNumber string    `json:"owner_phone_number"`
	Service          string    `json:"service"`
}
