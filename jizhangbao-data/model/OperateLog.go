package model

import "gopkg.in/mgo.v2/bson"

//操作日志
type OperateLog struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	BookId  string        `json:"bookId" bson:"bookId"`     //账本id
	BillId  string        `json:"billId" bson:"billId"`     //账单编号
	Date    string        `json:"date" bson:"date"`         //操作日期 yyyy-MM-dd
	Time    string        `json:"time" bson:"time"`         //操作时间 hh:mm:ss
	Operate string        `json:"operator" bson:"operator"` //操作
	InOut   int           `json:"inOut" bson:"inOut"`       //收支 1-收入 2-支出 3-应收 4-应付
	Amount  float64       `json:"amount" bson:"amount"`     //操作金额
}
