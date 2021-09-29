package nacos

import (
	"fmt"
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/logs"
	"github.com/maczh/mgcall"
	"github.com/maczh/utils"
	"ququ.im/jizhangbao-data/model"
	"strconv"
)

const (
	SERVICE_JZB_DATA   = "jzb-data"
	URI_FLOW_LIST      = "/flow/list"
	URI_FLOW_ADD       = "/flow/add"
	URI_FLOW_GET       = "/flow/get"
	URI_FLOW_UPDATE    = "/flow/update"
	URI_FLOW_DELETE    = "/flow/del"
	URI_ACCRUAL_LIST   = "/accrual/list"
	URI_ACCRUAL_ADD    = "/accrual/add"
	URI_ACCRUAL_GET    = "/accrual/get"
	URI_ACCRUAL_UPDATE = "/accrual/update"
	URI_ACCRUAL_DELETE = "/accrual/del"
	URI_RECORD_ADD     = "/accrual/record/add"
	URI_RECORD_UPDATE  = "/accrual/record/update"
	URI_RECORD_DELETE  = "/accrual/record/del"
	URI_TYPE_GET       = "/type/get"
	URI_TYPE_LIST      = "/type/list"
	URI_CUSTOMER_GET   = "/customer/get"
	URI_CUSTOMER_LIST  = "/customer/list"
)

func ListBillFlow(bookId, year, month, week, billTypeId, customerId string, inOut, page, size int) []model.BillFlow {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["year"] = year
	params["month"] = month
	params["week"] = week
	params["billTypeId"] = billTypeId
	params["customerId"] = customerId
	params["inOut"] = strconv.Itoa(inOut)
	params["page"] = strconv.Itoa(page)
	params["size"] = strconv.Itoa(size)
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_FLOW_LIST, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_FLOW_LIST, err.Error())
		return nil
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return nil
	}
	billFlows := make([]model.BillFlow, 0)
	utils.FromJSON(utils.ToJSON(result.Data), &billFlows)
	return billFlows
}

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
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_FLOW_ADD, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_FLOW_ADD, err.Error())
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
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_FLOW_UPDATE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_FLOW_UPDATE, err.Error())
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
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_FLOW_DELETE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_FLOW_DELETE, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func GetBillFlow(billId string) model.BillFlow {
	params := make(map[string]string)
	params["billId"] = billId
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_FLOW_GET, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_FLOW_GET, err.Error())
		return model.BillFlow{}
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return model.BillFlow{}
	}
	var billFlow model.BillFlow
	utils.FromJSON(utils.ToJSON(result.Data), &billFlow)
	return billFlow
}

func ListBillAccrual(bookId, year, month, week, billTypeId, customerId string, excuteDate, inOut, hideFinished, page, size int) []model.BillAccrual {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["year"] = year
	params["month"] = month
	params["week"] = week
	params["billTypeId"] = billTypeId
	params["customerId"] = customerId
	params["excuteDate"] = strconv.Itoa(excuteDate)
	params["inOut"] = strconv.Itoa(inOut)
	params["hideFinished"] = strconv.Itoa(hideFinished)
	params["page"] = strconv.Itoa(page)
	params["size"] = strconv.Itoa(size)
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_ACCRUAL_LIST, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_ACCRUAL_LIST, err.Error())
		return nil
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return nil
	}
	billAccruals := make([]model.BillAccrual, 0)
	utils.FromJSON(utils.ToJSON(result.Data), &billAccruals)
	return billAccruals
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
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_ACCRUAL_ADD, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_ACCRUAL_ADD, err.Error())
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
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_ACCRUAL_UPDATE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_ACCRUAL_UPDATE, err.Error())
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
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_ACCRUAL_DELETE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_ACCRUAL_DELETE, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func GetBillAccrual(billId string) model.BillAccrual {
	params := make(map[string]string)
	params["billId"] = billId
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_ACCRUAL_GET, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_ACCRUAL_GET, err.Error())
		return model.BillAccrual{}
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return model.BillAccrual{}
	}
	var billAccrual model.BillAccrual
	utils.FromJSON(utils.ToJSON(result.Data), &billAccrual)
	return billAccrual
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
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_RECORD_ADD, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_RECORD_ADD, err.Error())
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
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_RECORD_DELETE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_RECORD_DELETE, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func UpdateBillRecord(bookId, billId, billAccountId, remark, photos string, amount float64, billRecordIndex int) model.BillAccrual {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billId"] = billId
	params["billAccountId"] = billAccountId
	params["remark"] = remark
	params["photos"] = photos
	params["amount"] = fmt.Sprintf("%.2f", amount)
	params["billRecordIndex"] = strconv.Itoa(billRecordIndex)
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_RECORD_UPDATE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_RECORD_UPDATE, err.Error())
		return model.BillAccrual{}
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return model.BillAccrual{}
	}
	var billAccrual model.BillAccrual
	utils.FromJSON(utils.ToJSON(result.Data), &billAccrual)
	return billAccrual
}

func ListBillTypes(bookId, operateType string) []model.BillType {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["operateType"] = operateType
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_TYPE_LIST, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_TYPE_LIST, err.Error())
		return nil
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return nil
	}
	billTypes := make([]model.BillType, 0)
	utils.FromJSON(utils.ToJSON(result.Data), &billTypes)
	return billTypes
}

func GetBillType(billTypeId string) model.BillType {
	params := make(map[string]string)
	params["billTypeId"] = billTypeId
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_TYPE_GET, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_TYPE_GET, err.Error())
		return model.BillType{}
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return model.BillType{}
	}
	var billType model.BillType
	utils.FromJSON(utils.ToJSON(result.Data), &billType)
	return billType
}

func GetBusinessCustomer(bookId, customerId string) model.BusinessCustomer {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["customerId"] = customerId
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_CUSTOMER_GET, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_CUSTOMER_GET, err.Error())
		return model.BusinessCustomer{}
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	if result.Status != 1 {
		return model.BusinessCustomer{}
	}
	var businessCustomer model.BusinessCustomer
	utils.FromJSON(utils.ToJSON(result.Data), &businessCustomer)
	return businessCustomer
}
