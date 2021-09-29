package model

import "gopkg.in/mgo.v2/bson"

type BillBook struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	ShopId     string        `json:"shopId" bson:"shopId"`         //
	UserId     string        `json:"userId" bson:"userId"`         //
	BookId     string        `json:"bookId" bson:"bookId"`         //账本id
	BookName   string        `json:"bookName" bson:"bookName"`     //账本名称
	UpdateTime string        `json:"updateTime" bson:"updateTime"` //更新时间，yyyy-MM-dd HH:mm:ss
}
