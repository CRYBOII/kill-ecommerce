package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	ID         primitive.ObjectID `bosn:"_id,omitempty" `
	Username   string             `bosn:"username" json:"username,omitempty"`
	Email      string             `bosn:"email" json:"email" valid:"email"`
	Password   string             `bosn:"password" json:"password"`
	First_name string             `bosn:"firstname" json:"firstname,omitempty"`
	Last_name  string             `bosn:"lastname" json:"lastname,omitempty"`
	Address    []struct {
		Type    string `bosn:"type" json:"type,omitempty"`
		Street  string `bosn:"street" json:"street,omitempty"`
		City    string `bosn:"city" json:"city,omitempty"`
		State   string `bosn:"state" json:"state"`
		Pincode int64  `bosn:"pincode" json:"pincode,omitempty"`
	}
	Payments string `bson:"payment" json:"payment,omitempty"`
}
