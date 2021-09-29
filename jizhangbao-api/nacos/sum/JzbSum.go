package sum

import (
	"fmt"
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/logs"
	"github.com/maczh/mgcall"
	"github.com/maczh/utils"
	"ququ.im/jizhangbao-api/nacos/data"
	"ququ.im/jizhangbao-sum/model"
	"strconv"
)

const (
	SERVICE_JZB_SUM             = "jzb-sum"
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

func UpdateBillSummary(bookId, billTypeId, customerId, billTime string, inOut int, amount float64) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billTime"] = billTime
	params["billTypeId"] = billTypeId
	params["customerId"] = customerId
	params["inOut"] = strconv.Itoa(inOut)
	params["amount"] = fmt.Sprintf("%.2f", amount)
	res, err := mgcall.Call(SERVICE_JZB_SUM, URI_BILL_INCR, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_SUM, URI_BILL_INCR, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func UpdateAccrualSummary(bookId, billTypeId, customerId string, inOut int, amount float64) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billTypeId"] = billTypeId
	params["customerId"] = customerId
	params["inOut"] = strconv.Itoa(inOut)
	params["amount"] = fmt.Sprintf("%.2f", amount)
	res, err := mgcall.Call(SERVICE_JZB_SUM, URI_ACCRUAL_INCR, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_SUM, URI_ACCRUAL_INCR, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func GetBillDaySummary(bookId, day, billTypeId, customerId string) model.BillSummary {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["day"] = day
	params["billTypeId"] = billTypeId
	params["customerId"] = customerId
	res, err := mgcall.Call(SERVICE_JZB_SUM, URI_BILL_SUM_DAY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_SUM, URI_BILL_SUM_DAY, err.Error())
		return model.BillSummary{}
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return model.BillSummary{}
	}
	var billSummary model.BillSummary
	utils.FromJSON(utils.ToJSON(result.Data), &billSummary)
	return billSummary
}

func GetBillMonthSummary(bookId, month, billTypeId, customerId string) model.BillSummary {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["month"] = month
	params["billTypeId"] = billTypeId
	params["customerId"] = customerId
	res, err := mgcall.Call(SERVICE_JZB_SUM, URI_BILL_SUM_MONTH, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_SUM, URI_BILL_SUM_MONTH, err.Error())
		return model.BillSummary{}
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return model.BillSummary{}
	}
	var billSummary model.BillSummary
	utils.FromJSON(utils.ToJSON(result.Data), &billSummary)
	return billSummary
}

func GetBillYearSummary(bookId, year, billTypeId, customerId string) model.BillSummary {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["year"] = year
	params["billTypeId"] = billTypeId
	params["customerId"] = customerId
	res, err := mgcall.Call(SERVICE_JZB_SUM, URI_BILL_SUM_YEAR, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_SUM, URI_BILL_SUM_YEAR, err.Error())
		return model.BillSummary{}
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return model.BillSummary{}
	}
	var billSummary model.BillSummary
	utils.FromJSON(utils.ToJSON(result.Data), &billSummary)
	return billSummary
}

func GetBillWeekSummary(bookId, year, week, billTypeId, customerId string) model.BillSummary {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["year"] = year
	params["week"] = week
	params["billTypeId"] = billTypeId
	params["customerId"] = customerId
	res, err := mgcall.Call(SERVICE_JZB_SUM, URI_BILL_SUM_WEEK, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_SUM, URI_BILL_SUM_WEEK, err.Error())
		return model.BillSummary{}
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return model.BillSummary{}
	}
	var billSummary model.BillSummary
	utils.FromJSON(utils.ToJSON(result.Data), &billSummary)
	return billSummary
}

func ListTypeBillSummaryByDay(bookId, day string, inOut int) map[string]float64 {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["day"] = day
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_SUM, URI_BILL_SUM_TYPE_DAY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_SUM, URI_BILL_SUM_TYPE_DAY, err.Error())
		return nil
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return nil
	}
	billSummary := map[string]float64{}
	utils.FromJSON(utils.ToJSON(result.Data), &billSummary)
	return billSummary
}

func ListTypeBillSummaryByMonth(bookId, month string, inOut int) map[string]float64 {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["month"] = month
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_SUM, URI_BILL_SUM_TYPE_MONTH, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_SUM, URI_BILL_SUM_TYPE_MONTH, err.Error())
		return nil
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return nil
	}
	billSummary := map[string]float64{}
	utils.FromJSON(utils.ToJSON(result.Data), &billSummary)
	return billSummary
}

func ListTypeBillSummaryByYear(bookId, year string, inOut int) map[string]float64 {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["year"] = year
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_SUM, URI_BILL_SUM_TYPE_YEAR, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_SUM, URI_BILL_SUM_TYPE_YEAR, err.Error())
		return nil
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return nil
	}
	billSummary := map[string]float64{}
	utils.FromJSON(utils.ToJSON(result.Data), &billSummary)
	return billSummary
}

func ListTypeBillSummaryByWeek(bookId, year, week string, inOut int) map[string]float64 {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["year"] = year
	params["week"] = week
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_SUM, URI_BILL_SUM_TYPE_WEEK, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_SUM, URI_BILL_SUM_TYPE_WEEK, err.Error())
		return nil
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return nil
	}
	billSummary := map[string]float64{}
	utils.FromJSON(utils.ToJSON(result.Data), &billSummary)
	return billSummary
}

func ListCustomerBillSummaryByDay(bookId, day string, inOut int) map[string]float64 {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["day"] = day
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_SUM, URI_BILL_SUM_CUSTOMER_DAY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_SUM, URI_BILL_SUM_CUSTOMER_DAY, err.Error())
		return nil
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return nil
	}
	billSummary := map[string]float64{}
	utils.FromJSON(utils.ToJSON(result.Data), &billSummary)
	return billSummary
}

func ListCustomerBillSummaryByMonth(bookId, month string, inOut int) map[string]float64 {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["month"] = month
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_SUM, URI_BILL_SUM_CUSTOMER_MONTH, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_SUM, URI_BILL_SUM_CUSTOMER_MONTH, err.Error())
		return nil
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return nil
	}
	billSummary := map[string]float64{}
	utils.FromJSON(utils.ToJSON(result.Data), &billSummary)
	return billSummary
}

func ListCustomerBillSummaryByYear(bookId, year string, inOut int) map[string]float64 {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["year"] = year
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_SUM, URI_BILL_SUM_CUSTOMER_YEAR, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_SUM, URI_BILL_SUM_CUSTOMER_YEAR, err.Error())
		return nil
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return nil
	}
	billSummary := map[string]float64{}
	utils.FromJSON(utils.ToJSON(result.Data), &billSummary)
	return billSummary
}

func ListCustomerBillSummaryByWeek(bookId, year, week string, inOut int) map[string]float64 {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["year"] = year
	params["week"] = week
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_SUM, URI_BILL_SUM_CUSTOMER_WEEK, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_SUM, URI_BILL_SUM_CUSTOMER_WEEK, err.Error())
		return nil
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return nil
	}
	billSummary := map[string]float64{}
	utils.FromJSON(utils.ToJSON(result.Data), &billSummary)
	return billSummary
}

func GetAccrualSummary(bookId, billTypeId, customerId string) model.AccrualSummary {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billTypeId"] = billTypeId
	params["customerId"] = customerId
	res, err := mgcall.Call(SERVICE_JZB_SUM, data.URI_ACCRUAL_GET, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_SUM, data.URI_ACCRUAL_GET, err.Error())
		return model.AccrualSummary{}
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return model.AccrualSummary{}
	}
	var accrualSummary model.AccrualSummary
	utils.FromJSON(utils.ToJSON(result.Data), &accrualSummary)
	return accrualSummary
}

func ListTypeAccrualSummary(bookId string, inOut int) map[string]float64 {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_SUM, URI_ACCRUAL_SUM_TYPE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_SUM, URI_ACCRUAL_SUM_TYPE, err.Error())
		return nil
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return nil
	}
	accrualSummary := map[string]float64{}
	utils.FromJSON(utils.ToJSON(result.Data), &accrualSummary)
	return accrualSummary
}

func ListCustomerAccrualSummary(bookId string, inOut int) map[string]float64 {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["inOut"] = strconv.Itoa(inOut)
	res, err := mgcall.Call(SERVICE_JZB_SUM, URI_ACCRUAL_SUM_CUSTOMER, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_SUM, URI_ACCRUAL_SUM_CUSTOMER, err.Error())
		return nil
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return nil
	}
	summary := map[string]float64{}
	utils.FromJSON(utils.ToJSON(result.Data), &summary)
	return summary
}
