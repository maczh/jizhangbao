package model

import "gopkg.in/mgo.v2/bson"

//应收应付账单
type BillAccrual struct {
	Id                 bson.ObjectId `json:"id" bson:"_id"`
	BookId             string        `json:"bookId" bson:"bookId"`                         //账本id
	BillId             string        `json:"billId" bson:"billId"`                         //账单编号
	BillTime           string        `json:"billTime" bson:"billTime"`                     //记账时间 yyyy-MM-dd HH:mm:ss
	UpdateTime         string        `json:"updateTime" bson:"updateTime"`                 //更新时间 yyyy-MM-dd HH:mm:ss
	InOut              int           `json:"inOut" bson:"inOut"`                           //收支 3-应收 4-应付
	Amount             float64       `json:"amount" bson:"amount"`                         //￿应付/应收总金额，单位为元，float64
	RemainAmount       float64       `json:"remainAmount" bson:"remainAmount"`             //剩余未付/未收金额，单位为元，float64
	BillTypeId         string        `json:"billTypeId" bson:"billTypeId"`                 //账单分类id
	BillType           string        `json:"billType" bson:"billType"`                     //账单分类名称
	BillTypeIconId     string        `json:"billTypeIconId" bson:"billTypeIconId"`         //账单分类图标id
	BusinessCustomerId string        `json:"businessCustomerId" bson:"businessCustomerId"` //往来客户id
	BusinessCustomer   string        `json:"businessCustomer" bson:"businessCustomer"`     //往来客户名称
	Remark             string        `json:"remark" bson:"remark"`                         //备注
	Photos             []string      `json:"photos" bson:"photos"`                         //照片url地址列表
	ExcuteDate         string        `json:"excuteDate" bson:"excuteDate"`                 //预约付款/收款日期
	BillRecords        []BillRecord  `json:"billRecords" bson:"billRecords"`               //收付款记录列表
}

type BillRecord struct {
	Time          string   `json:"time" bson:"time"`                   //收款/付款时间 yyyy-MM-dd HH:mm:ss
	Amount        float64  `json:"amount" bson:"amount"`               //￿收款/付款金额，单位为元，float64
	BillAccountId string   `json:"billAccountId" bson:"billAccountId"` //资金账户id
	BillAccount   string   `json:"billAccount" bson:"billAccount"`     //资金账户名称
	Remark        string   `json:"remark" bson:"remark"`               //备注
	Photos        []string `json:"photos" bson:"photos"`               //照片url地址列表
}
