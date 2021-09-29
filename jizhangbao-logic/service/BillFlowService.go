package service

import (
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/utils"
	"ququ.im/jizhangbao-logic/model"
	"ququ.im/jizhangbao-logic/nacos"
	"time"
)

func AddBillFlow(bookId, billTypeId, billAccountId, customerId, remark, photos, billTime string, inOut int, amount float64) mgresult.Result {
	billFlowResult := nacos.AddBillFlow(bookId, billTypeId, billAccountId, customerId, remark, photos, billTime, inOut, amount)
	if billFlowResult.Status == 1 {
		return nacos.UpdateBillSummary(bookId, billTypeId, customerId, billTime, inOut, amount)
	} else {
		return billFlowResult
	}
}

func DeleteBillFlow(bookId, billId string) mgresult.Result {
	billFlow := nacos.GetBillFlow(billId)
	if billFlow.BillId == "" {
		return *mgresult.Error(-1, "无此账单")
	}
	result := nacos.DeleteBillFlow(bookId, billId)
	if result.Status != 1 {
		return result
	}
	return nacos.UpdateBillSummary(bookId, billFlow.BillTypeId, billFlow.BusinessCustomerId, billFlow.BillTime, billFlow.InOut, -billFlow.Amount)
}

func UpdateBillFlow(bookId, billId, billTypeId, billAccountId, customerId, remark, photos, billTime string, amount float64) mgresult.Result {
	billFlow := nacos.GetBillFlow(billId)
	if billFlow.BillId == "" {
		return *mgresult.Error(-1, "无此账单")
	}
	result := nacos.UpdateBillFlow(bookId, billId, billTypeId, billAccountId, customerId, remark, photos, billTime, amount)
	if result.Status != 1 {
		return result
	}
	if amount == 0 {
		amount = billFlow.Amount
	}
	if billTypeId == "" {
		billTypeId = billFlow.BillTypeId
	}
	if customerId == "" {
		customerId = billFlow.BusinessCustomerId
	}
	if billTime == "" {
		billTime = billFlow.BillTime
	}
	if amount != billFlow.Amount || billTypeId != billFlow.BillTypeId || customerId != billFlow.BusinessCustomerId || billTime != billFlow.BillTime {
		nacos.UpdateBillSummary(bookId, billFlow.BillTypeId, billFlow.BusinessCustomerId, billFlow.BillTime, billFlow.InOut, -billFlow.Amount)
		nacos.UpdateBillSummary(bookId, billTypeId, customerId, billTime, billFlow.InOut, amount)
		return *mgresult.Success(nacos.GetBillFlow(billId))
	} else {
		return result
	}
}

func GetMonthBillFlowPage(bookId, month, billTypeId, customerId string, inOut int) mgresult.Result {
	var billFlowPage model.BillFlowPage
	if month == "" {
		month = utils.DateFormat(time.Now(), "yyyy-MM")
	}
	monthSummary := nacos.GetBillMonthSummary(bookId, month, billTypeId, customerId)
	billFlowPage.BillDuration = month
	billFlowPage.Expense = monthSummary.Expense
	billFlowPage.Income = monthSummary.Income
	billFlowGroupByDayList := make([]model.BillFlowGroupByDay, 0)
	billFlows := nacos.ListBillFlow(bookId, "", month, "", billTypeId, customerId, inOut, 0, 0)
	if billFlows == nil || len(billFlows) == 0 {
		billFlowPage.BillFlow = billFlowGroupByDayList
		return *mgresult.Success(billFlowPage)
	}
	parseDay := billFlows[0].BillTime[:10]
	billFlowGroupByDay := new(model.BillFlowGroupByDay)
	billList := make([]model.Bill, 0)
	for i, billFlow := range billFlows {
		theday := billFlow.BillTime[:10]
		if i == 0 || theday != parseDay {
			if i > 0 {
				billFlowGroupByDay.BillList = billList
				billFlowGroupByDayList = append(billFlowGroupByDayList, *billFlowGroupByDay)
				billFlowGroupByDay = new(model.BillFlowGroupByDay)
			}
			parseDay = theday
			billFlowGroupByDay.BillDate = theday
			daySummary := nacos.GetBillDaySummary(bookId, theday, billTypeId, customerId)
			billFlowGroupByDay.Expense = daySummary.Expense
			billFlowGroupByDay.Income = daySummary.Income
			billList = make([]model.Bill, 0)
		}
		var bill model.Bill
		utils.FromJSON(utils.ToJSON(billFlow), &bill)
		bill.BillDate = theday
		bill.BillTime = bill.BillTime[11:]
		billList = append(billList, bill)
	}
	billFlowGroupByDay.BillList = billList
	billFlowGroupByDayList = append(billFlowGroupByDayList, *billFlowGroupByDay)
	billFlowPage.BillFlow = billFlowGroupByDayList
	return *mgresult.Success(billFlowPage)
}

func ListTypeBillSummaryByMonth(bookId, month string, inOut int) mgresult.Result {
	if month == "" {
		month = utils.DateFormat(time.Now(), "yyyy-MM")
	}
	typeSummary := nacos.ListTypeBillSummaryByMonth(bookId, month, inOut)
	if typeSummary == nil || len(typeSummary) == 0 {
		return *mgresult.Success([]string{})
	}
	billTypeSummary := make(map[string]interface{})
	for k, v := range typeSummary {
		if v > 0.0 {
			billType := nacos.GetBillType(k)
			billTypeSummary[billType.BillType] = v
		}
	}
	result := utils.SortMapByValueDesc(billTypeSummary)
	return *mgresult.Success(result)
}

func ListCustomerBillSummaryByMonth(bookId, month string, inOut int) mgresult.Result {
	if month == "" {
		month = utils.DateFormat(time.Now(), "yyyy-MM")
	}
	summary := nacos.ListCustomerBillSummaryByMonth(bookId, month, inOut)
	if summary == nil || len(summary) == 0 {
		return *mgresult.Success([]string{})
	}
	customerSummary := make(map[string]interface{})
	for k, v := range summary {
		if v > 0.0 {
			businessCustomer := nacos.GetBusinessCustomer(bookId, k)
			customerSummary[businessCustomer.BusinessCustomer] = v
		}
	}
	result := utils.SortMapByValueDesc(customerSummary)
	return *mgresult.Success(result)
}
