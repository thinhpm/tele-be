package models

type Account struct {
	Name   string `json:"name" bson:"name"`
	Token  string `json:"token" bson:"token"`
	Server string `json:"server" bson:"server"`
	Proxy  string `json:"proxy" bson:"proxy"`
	Active bool   `json:"active" bson:"active"`
}
