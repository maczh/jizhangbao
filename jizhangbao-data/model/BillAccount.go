package model

import "gopkg.in/mgo.v2/bson"

type BillAccount struct {
	Id            bson.ObjectId `json:"id" bson:"_id"`
	BillAccountId string        `json:"billAccountId" bson:"billAccountId"` //资金账户id
	BillAccount   string        `json:"billAccount" bson:"billAccount"`     //资金账户名称
	BookId        string        `json:"bookId" bson:"bookId"`               //账本id
	UpdateTime    string        `json:"updateTime" bson:"updateTime"`       //更新时间，yyyy-MM-dd HH:mm:ss
}
