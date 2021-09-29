package data

import (
	"fmt"
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/logs"
	"github.com/maczh/mgcall"
	"github.com/maczh/utils"
	"strconv"
)

const (
	SERVICE_JZB_DATA    = "jzb-data"
	URI_FLOW_LIST       = "/flow/list"
	URI_FLOW_ADD        = "/flow/add"
	URI_FLOW_GET        = "/flow/get"
	URI_FLOW_UPDATE     = "/flow/update"
	URI_FLOW_DELETE     = "/flow/del"
	URI_ACCRUAL_LIST    = "/accrual/list"
	URI_ACCRUAL_ADD     = "/accrual/add"
	URI_ACCRUAL_GET     = "/accrual/get"
	URI_ACCRUAL_UPDATE  = "/accrual/update"
	URI_ACCRUAL_DELETE  = "/accrual/del"
	URI_RECORD_ADD      = "/accrual/record/add"
	URI_RECORD_UPDATE   = "/accrual/record/update"
	URI_RECORD_DELETE   = "/accrual/record/del"
	URI_TYPE_PUBLIC_ADD = "/type/add/public"
	URI_TYPE_ADD        = "/type/add"
	URI_TYPE_UPDATE     = "/type/update"
	URI_TYPE_DELETE     = "/type/del"
	URI_TYPE_GET        = "/type/get"
	URI_TYPE_LIST       = "/type/list"
	URI_BOOK_NEW        = "/book/new"
	URI_BOOK_UPDATE     = "/book/update"
	URI_BOOK_LIST       = "/book/list"
	URI_BOOK_GET        = "/book/get"
	URI_ACCOUNT_ADD     = "/account/add"
	URI_ACCOUNT_UPDATE  = "/account/update"
	URI_ACCOUNT_DELETE  = "/account/del"
	URI_ACCOUNT_LIST    = "/account/list"
	URI_ACCOUNT_GET     = "/account/get"
	URI_CUSTOMER_ADD    = "/customer/add"
	URI_CUSTOMER_UPDATE = "/customer/update"
	URI_CUSTOMER_DELETE = "/customer/del"
	URI_CUSTOMER_GET    = "/customer/get"
	URI_CUSTOMER_LIST   = "/customer/list"
	URI_LOGS_LIST       = "/logs/list"
)

func ListBillFlow(bookId, year, month, week, billTypeId, customerId string, inOut, page, size int) mgresult.Result {
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
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
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

func GetBillFlow(billId string) mgresult.Result {
	params := make(map[string]string)
	params["billId"] = billId
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_FLOW_GET, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_FLOW_GET, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func ListBillAccrual(bookId, year, month, week, billTypeId, customerId string, excuteDate, inOut, hideFinished, page, size int) mgresult.Result {
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

func GetBillAccrual(billId string) mgresult.Result {
	params := make(map[string]string)
	params["billId"] = billId
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_ACCRUAL_GET, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_ACCRUAL_GET, err.Error())
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

func UpdateBillRecord(bookId, billId, billAccountId, remark, photos string, amount float64, billRecordIndex int) mgresult.Result {
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
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func ListBillTypes(bookId, operateType string) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["operateType"] = operateType
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_TYPE_LIST, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_TYPE_LIST, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func GetBillType(billTypeId string) mgresult.Result {
	params := make(map[string]string)
	params["billTypeId"] = billTypeId
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_TYPE_GET, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_TYPE_GET, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func AddPublicBillType(billType, billTypeIconId, operateType string) mgresult.Result {
	params := make(map[string]string)
	params["billType"] = billType
	params["billTypeIconId"] = billTypeIconId
	params["operateType"] = operateType
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_TYPE_PUBLIC_ADD, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_TYPE_PUBLIC_ADD, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func AddBillType(bookId, billType, billTypeIconId, operateType string) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billType"] = billType
	params["billTypeIconId"] = billTypeIconId
	params["operateType"] = operateType
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_TYPE_ADD, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_TYPE_ADD, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func UpdateBillType(bookId, billType, billTypeIconId, operateType string) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billType"] = billType
	params["billTypeIconId"] = billTypeIconId
	params["operateType"] = operateType
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_TYPE_UPDATE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_TYPE_UPDATE, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func DeleteBillType(bookId, billTypeId string) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billTypeId"] = billTypeId
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_TYPE_DELETE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_TYPE_DELETE, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func AddBusinessCustomer(bookId, customerName string) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["customerName"] = customerName
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_CUSTOMER_ADD, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_CUSTOMER_ADD, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func UpdateBusinessCustomer(bookId, customerId, customerName string) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["customerId"] = customerId
	params["customerName"] = customerName
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_CUSTOMER_UPDATE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_CUSTOMER_UPDATE, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func DeleteBusinessCustomer(bookId, customerId string) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["customerId"] = customerId
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_CUSTOMER_DELETE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_CUSTOMER_DELETE, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func ListBusinessCustomer(bookId, customerName string) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["customerName"] = customerName
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_CUSTOMER_LIST, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_CUSTOMER_LIST, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func GetBusinessCustomer(bookId, customerId string) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["customerId"] = customerId
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_CUSTOMER_GET, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_CUSTOMER_GET, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func NewBillBook(shopId, userId, bookName string) mgresult.Result {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["userId"] = userId
	params["bookName"] = bookName
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_BOOK_NEW, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_BOOK_NEW, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func UpdateBillBook(bookId, bookName string) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["bookName"] = bookName
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_BOOK_UPDATE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_BOOK_UPDATE, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func ListBillBooks(shopId, userId string) mgresult.Result {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["userId"] = userId
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_BOOK_LIST, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_BOOK_LIST, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func GetBillBook(bookId string) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_BOOK_GET, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_BOOK_GET, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func AddBillAccount(bookId, billAccountName string) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billAccountName"] = billAccountName
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_ACCOUNT_ADD, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_ACCOUNT_ADD, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func UpdateBillAccount(bookId, billAccountId, billAccountName string) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billAccountId"] = billAccountId
	params["billAccountName"] = billAccountName
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_ACCOUNT_UPDATE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_ACCOUNT_UPDATE, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func DeleteBillAccount(bookId, billAccountId string) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billAccountId"] = billAccountId
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_ACCOUNT_DELETE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_ACCOUNT_DELETE, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func GetBillAccount(bookId, billAccountId string) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["billAccountId"] = billAccountId
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_ACCOUNT_GET, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_ACCOUNT_GET, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func ListBillAccount(bookId string) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_ACCOUNT_LIST, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_ACCOUNT_LIST, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}

func ListOperateLog(bookId, startDate, endDate string) mgresult.Result {
	params := make(map[string]string)
	params["bookId"] = bookId
	params["startDate"] = startDate
	params["endDate"] = endDate
	res, err := mgcall.Call(SERVICE_JZB_DATA, URI_LOGS_LIST, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_JZB_DATA, URI_LOGS_LIST, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var result mgresult.Result
	utils.FromJSON(res, &result)
	return result
}
