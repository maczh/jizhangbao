package model

type Bill struct {
	BillId             string   `json:"billId" bson:"billId"`                         //账单编号
	BillDate           string   `json:"billDate" bson:"billDate"`                     //记账时间 yyyy-MM-dd
	BillTime           string   `json:"billTime" bson:"billTime"`                     //记账时间 hh:mm:ss
	UpdateTime         string   `json:"updateTime" bson:"updateTime"`                 //更新时间 hh:mm:ss
	InOut              int      `json:"inOut" bson:"inOut"`                           //收支 1-收入 2-支出
	Amount             float64  `json:"amount" bson:"amount"`                         //￿金额，单位为元，float64
	BillTypeId         string   `json:"billTypeId" bson:"billTypeId"`                 //账单分类id
	BillType           string   `json:"billType" bson:"billType"`                     //账单分类名称
	BillTypeIconId     string   `json:"billTypeIconId" bson:"billTypeIconId"`         //账单分类图标id
	BillAccountId      string   `json:"billAccountId" bson:"billAccountId"`           //资金账户id
	BillAccount        string   `json:"billAccount" bson:"billAccount"`               //资金账户名称
	BusinessCustomerId string   `json:"businessCustomerId" bson:"businessCustomerId"` //往来客户id
	BusinessCustomer   string   `json:"businessCustomer" bson:"businessCustomer"`     //往来客户名称
	Remark             string   `json:"remark" bson:"remark"`                         //备注
	Photos             []string `json:"photos" bson:"photos"`                         //照片url地址列表
}

type BillFlowGroupByDay struct {
	BillDate string  `json:"billDate" bson:"billDate"` //记账时间 MM月dd日
	Expense  float64 `json:"expense" bson:"expense"`   //当天总支出
	Income   float64 `json:"income" bson:"income"`     //当天总收入
	BillList []Bill  `json:"billList" bson:"billList"` //当天收支账单列表
}

type BillFlowPage struct {
	BillDuration string               `json:"billDuration" bson:"billDuration"` //记账周期 具体年份/月份/日期/周
	Expense      float64              `json:"expense" bson:"expense"`           //本记账周期内总支出
	Income       float64              `json:"income" bson:"income"`             //本记账周期内总收入
	BillFlow     []BillFlowGroupByDay `json:"billFlow" bson:"billFlow"`         //按天的收支流水
}
