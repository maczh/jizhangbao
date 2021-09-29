package model

import "gopkg.in/mgo.v2/bson"

type BillFlow struct {
	Id                 bson.ObjectId `json:"id" bson:"_id"`
	BookId             string        `json:"bookId" bson:"bookId"`                         //账本id
	BillId             string        `json:"billId" bson:"billId"`                         //账单编号
	BillTime           string        `json:"billTime" bson:"billTime"`                     //记账时间 yyyy-MM-dd HH:mm:ss
	UpdateTime         string        `json:"updateTime" bson:"updateTime"`                 //更新时间 yyyy-MM-dd HH:mm:ss
	InOut              int           `json:"inOut" bson:"inOut"`                           //收支 1-收入 2-支出
	Amount             float64       `json:"amount" bson:"amount"`                         //￿金额，单位为元，float64
	BillTypeId         string        `json:"billTypeId" bson:"billTypeId"`                 //账单分类id
	BillType           string        `json:"billType" bson:"billType"`                     //账单分类名称
	BillTypeIconId     string        `json:"billTypeIconId" bson:"billTypeIconId"`         //账单分类图标id
	BillAccountId      string        `json:"billAccountId" bson:"billAccountId"`           //资金账户id
	BillAccount        string        `json:"billAccount" bson:"billAccount"`               //资金账户名称
	BusinessCustomerId string        `json:"businessCustomerId" bson:"businessCustomerId"` //往来客户id
	BusinessCustomer   string        `json:"businessCustomer" bson:"businessCustomer"`     //往来客户名称
	Remark             string        `json:"remark" bson:"remark"`                         //备注
	Photos             []string      `json:"photos" bson:"photos"`                         //照片url地址列表
	BillAccrualId      string        `json:"billAccrualId" bson:"billAccrualId"`           //关联的应收付账单id
}
