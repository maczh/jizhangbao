package model

type BusinessCustomer struct {
	Id                 string `json:"id" bson:"id"`
	BusinessCustomerId string `json:"businessCustomerId" bson:"businessCustomerId"` //往来客户id
	BusinessCustomer   string `json:"businessCustomer" bson:"businessCustomer"`     //往来客户名称
	BookId             string `json:"bookId" bson:"bookId"`                         //账本id
	UpdateTime         string `json:"updateTime" bson:"updateTime"`                 //更新时间，yyyy-MM-dd HH:mm:ss
}
