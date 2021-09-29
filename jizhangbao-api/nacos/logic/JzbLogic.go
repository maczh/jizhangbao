package logic

import (
	"fmt"
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/logs"
	"github.com/maczh/mgcall"
	"github.com/maczh/utils"
	"strconv"
)

const (
	SERVICE_JZB_LOGIC      = "jzb-logic"
	URI_FLOW_LIST          = "/flow/list"
	URI_FLOW_ADD           = "/flow/add"
	URI_FLOW_LIST_MONTH    = "/flow/list/month"
	URI_FLOW_UPDATE        = "/flow/update"
	URI_FLOW_DELETE        = "/flow/del"
	URI_ACCRUAL_LIST       = "/accrual/list"
	URI_ACCRUAL_ADD        = "/accrual/add"
	URI_ACCRUAL_LIST_MONTH = "/accrual/list/month"
	URI_ACCRUAL_UPDATE     = "/accrual/update"
	URI_ACCRUAL_DELETE     = "/accrual/del"
	URI_RECORD_ADD         = "/accrual/record/add"
	URI_RECORD_UPDATE      = "/accrual/record/update"
	URI_RECORD_DELETE      = "/accrual/record/del"
	URI_TYPE_GET           = "/type/get"
	URI_TYPE_LIST          = "/type/list"
	URI_CUSTOMER_GET       = "/customer/get"
	URI_CUSTOMER_LIST      = "/customer/list"

	URI_BILL_SUM_DAY            = "/bill/get/day"
	URI_BILL_SUM_MONTH          = "/bill/get/month"
	URI_BILL_SUM_YEAR           = "/bill/get/year"
	URI_BILL_SUM_WEEK           = "/bill/get/week"
	URI_BILL_SUM_TYPE_DAY       = "/bill/list/type/day"
	URI_BILL_SUM_TYPE_MONTH     = "/bill/list/type/month"
	URI_BILL_SUM_TYPE_YEAR      = "/bill/list/type/year"
	URI_BILL_SUM_TYPE_WEEK      = "/bill/list/type/week"
	URI_BILL_SUM_CUSTOMER_DAY   = "/bill/list/customer/day"
	URI_BILL_SUM_CUSTOMER_MONTH = "/bill/list/customer/month"
	URI_BILL_SUM_CUSTOMER_YEAR  = "/bill/list/customer/year"
	URI_BILL_SUM_CUSTOMER_WEEK  = "/bill/list/customer/week"
	URI_BILL_INCR               = "/bill/incr"
	URI_ACCRUAL_INCR            = "/accrual/incr"
	URI_ACCRUAL_SUM             = "/accrual/get"
	URI_ACCRUAL_SUM_TYPE        = "/accrual/list/type"
	URI_ACCRUAL_SUM_CUSTOMER    = "/accrual/list/customer"
)

func AddBillFlow(bookId, billTypeId, billAccountId, customerId, remark, photos, billTime string, inOut int, amount float64) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billTypeId"] = billTypeId
	params["customerId"] = customerId
	params["billAccountId"] = billAccountId
	params["remark"] = remark
	params["photos"] = photos
	params["billTime"] = billTime
	params["inOut"] = strconv.Itoa(inOut)
	params["amount"] = fmt.Sprintf("%.2f", amount)
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_FLOW_ADD, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_FLOW_ADD, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func UpdateBillFlow(bookId, billId, billTypeId, billAccountId, customerId, remark, photos, billTime string, amount float64) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billId"] = billId
	params["billTypeId"] = billTypeId
	params["customerId"] = customerId
	params["billAccountId"] = billAccountId
	params["remark"] = remark
	params["photos"] = photos
	params["billTime"] = billTime
	params["amount"] = fmt.Sprintf("%.2f", amount)
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_FLOW_UPDATE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_FLOW_UPDATE, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func DeleteBillFlow(bookId, billId string) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billId"] = billId
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_FLOW_DELETE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_FLOW_DELETE, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func GetMonthBillFlowPage(bookId, month, billTypeId, customerId string, inOut int) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["month"] = month
	params["billTypeId"] = billTypeId
	params["customerId"] = customerId
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_FLOW_LIST_MONTH, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_FLOW_LIST_MONTH, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func AddBillAccrual(bookId, billTypeId, customerId, remark, photos, billTime, excuteDate string, inOut int, amount float64) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billTypeId"] = billTypeId
	params["customerId"] = customerId
	params["excuteDate"] = excuteDate
	params["remark"] = remark
	params["photos"] = photos
	params["billTime"] = billTime
	params["inOut"] = strconv.Itoa(inOut)
	params["amount"] = fmt.Sprintf("%.2f", amount)
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_ACCRUAL_ADD, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_ACCRUAL_ADD, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func UpdateBillAccrual(bookId, billId, billTypeId, customerId, remark, photos, billTime, excuteDate string, amount float64) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billId"] = billId
	params["billTypeId"] = billTypeId
	params["customerId"] = customerId
	params["excuteDate"] = excuteDate
	params["remark"] = remark
	params["photos"] = photos
	params["billTime"] = billTime
	params["amount"] = fmt.Sprintf("%.2f", amount)
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_ACCRUAL_UPDATE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_ACCRUAL_UPDATE, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func DeleteBillAccrual(bookId, billId string) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billId"] = billId
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_ACCRUAL_DELETE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_ACCRUAL_DELETE, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func GetMonthBillAccrualPage(bookId, month, billTypeId, customerId string, inOut, excuteDate int) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["month"] = month
	params["billTypeId"] = billTypeId
	params["customerId"] = customerId
	params["inOut"] = strconv.Itoa(inOut)
	params["excuteDate"] = strconv.Itoa(excuteDate)
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_ACCRUAL_LIST_MONTH, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_ACCRUAL_LIST_MONTH, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func AddBillRecord(bookId, billId, billAccountId, remark, photos string, amount float64, sync int) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billId"] = billId
	params["billAccountId"] = billAccountId
	params["remark"] = remark
	params["photos"] = photos
	params["amount"] = fmt.Sprintf("%.2f", amount)
	params["sync"] = strconv.Itoa(sync)
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_RECORD_ADD, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_RECORD_ADD, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func DeleteBillRecord(bookId, billId string, billRecordIndex int) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billId"] = billId
	params["billRecordIndex"] = strconv.Itoa(billRecordIndex)
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_RECORD_DELETE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_RECORD_DELETE, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func UpdateBillRecord(bookId, billId, billAccountId, remark, photos string, amount float64, billRecordIndex int) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billId"] = billId
	params["billAccountId"] = billAccountId
	params["remark"] = remark
	params["photos"] = photos
	params["amount"] = fmt.Sprintf("%.2f", amount)
	params["billRecordIndex"] = strconv.Itoa(billRecordIndex)
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_RECORD_UPDATE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_RECORD_UPDATE, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func ListTypeBillSummaryByDay(bookId, day string, inOut int) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["day"] = day
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_BILL_SUM_TYPE_DAY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_BILL_SUM_TYPE_DAY, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func ListTypeBillSummaryByMonth(bookId, month string, inOut int) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["month"] = month
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_BILL_SUM_TYPE_MONTH, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_BILL_SUM_TYPE_MONTH, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func ListTypeBillSummaryByYear(bookId, year string, inOut int) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["year"] = year
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_BILL_SUM_TYPE_YEAR, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_BILL_SUM_TYPE_YEAR, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func ListTypeBillSummaryByWeek(bookId, year, week string, inOut int) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["year"] = year
	params["week"] = week
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_BILL_SUM_TYPE_WEEK, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_BILL_SUM_TYPE_WEEK, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func ListCustomerBillSummaryByDay(bookId, day string, inOut int) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["day"] = day
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_BILL_SUM_CUSTOMER_DAY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_BILL_SUM_CUSTOMER_DAY, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func ListCustomerBillSummaryByMonth(bookId, month string, inOut int) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["month"] = month
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_BILL_SUM_CUSTOMER_MONTH, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_BILL_SUM_CUSTOMER_MONTH, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func ListCustomerBillSummaryByYear(bookId, year string, inOut int) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["year"] = year
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_BILL_SUM_CUSTOMER_YEAR, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_BILL_SUM_CUSTOMER_YEAR, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func ListCustomerBillSummaryByWeek(bookId, year, week string, inOut int) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["year"] = year
	params["week"] = week
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_BILL_SUM_CUSTOMER_WEEK, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_BILL_SUM_CUSTOMER_WEEK, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func ListTypeAccrualSummary(bookId string, inOut int) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_ACCRUAL_SUM_TYPE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_ACCRUAL_SUM_TYPE, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func ListCustomerAccrualSummary(bookId string, inOut int) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_LOGIC, URI_ACCRUAL_SUM_CUSTOMER, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_LOGIC, URI_ACCRUAL_SUM_CUSTOMER, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}
