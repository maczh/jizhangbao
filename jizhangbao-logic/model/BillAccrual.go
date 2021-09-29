package model

import (
	model2 "ququ.im/jizhangbao-data/model"
)

type Accrual struct {
	BookId             string              `json:"bookId" bson:"bookId"`                         //账本id
	BillId             string              `json:"billId" bson:"billId"`                         //账单编号
	BillTime           string              `json:"billTime" bson:"billTime"`                     //记账时间 yyyy-MM-dd HH:mm:ss
	UpdateTime         string              `json:"updateTime" bson:"updateTime"`                 //更新时间 yyyy-MM-dd HH:mm:ss
	InOut              int                 `json:"inOut" bson:"inOut"`                           //收支 3-应收 4-应付
	Amount             float64             `json:"amount" bson:"amount"`                         //￿应付/应收总金额，单位为元，float64
	RemainAmount       float64             `json:"remainAmount" bson:"remainAmount"`             //剩余未付/未收金额，单位为元，float64
	BillTypeId         string              `json:"billTypeId" bson:"billTypeId"`                 //账单分类id
	BillType           string              `json:"billType" bson:"billType"`                     //账单分类名称
	BillTypeIconId     string              `json:"billTypeIconId" bson:"billTypeIconId"`         //账单分类图标id
	BusinessCustomerId string              `json:"businessCustomerId" bson:"businessCustomerId"` //往来客户id
	BusinessCustomer   string              `json:"businessCustomer" bson:"businessCustomer"`     //往来客户名称
	Remark             string              `json:"remark" bson:"remark"`                         //备注
	Photos             []string            `json:"photos" bson:"photos"`                         //照片url地址列表
	ExcuteDate         string              `json:"excuteDate" bson:"excuteDate"`                 //预约付款/收款日期
	BillRecords        []model2.BillRecord `json:"billRecords" bson:"billRecords"`               //收付款记录列表
}

type BillAccrualGroupByDay struct {
	GroupDate   string    `json:"groupDate" bson:"groupDate"`     //排序分组的日期，可对应于billTime或excuteDate字段
	Payable     float64   `json:"payable" bson:"payable"`         //当天合计应付总金额，单位为元，float64
	Receiveable float64   `json:"receiveable" bson:"receiveable"` //当天合计应收总金额，单位为元，float64
	AccrualList []Accrual `json:"accrualList" bson:"accrualList"` //应收应付账单列表
}

type BillAccrualPage struct {
	Payable     float64                 `json:"payable" bson:"payable"`         //所有合计应付总金额，单位为元，float64
	Receiveable float64                 `json:"receiveable" bson:"receiveable"` //所有合计应收总金额，单位为元，float64
	BillAccrual []BillAccrualGroupByDay `json:"billAccrual" bson:"billAccrual"` //按天的应收付账单列表
}
