package model

type BillSummary struct {
	Income  float64 `json:"income" bson:"income"`   //收入统计
	Expense float64 `json:"expense" bson:"expense"` //支出统计
}

type AccrualSummary struct {
	Receiveable float64 `json:"receiveable" bson:"receiveable"` //应收合计
	Payable     float64 `json:"payable" bson:"payable"`         //应付合计
}
