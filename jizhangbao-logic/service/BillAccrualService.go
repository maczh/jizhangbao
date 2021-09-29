package service

import (
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/utils"
	model2 "ququ.im/jizhangbao-data/model"
	"ququ.im/jizhangbao-logic/model"
	"ququ.im/jizhangbao-logic/nacos"
	"time"
)

func GetMonthBillAccrualPage(bookId, month, billTypeId, customerId string, inOut, excuteDate int) mgresult.Result {
	var billAccrualPage model.BillAccrualPage
	if month == "" {
		month = utils.DateFormat(time.Now(), "yyyy-MM")
	}
	billAccruals := nacos.ListBillAccrual(bookId, "", month, "", billTypeId, customerId, excuteDate, inOut, 0, 0, 0)
	accrualSummary := nacos.GetAccrualSummary(bookId, billTypeId, customerId)
	billAccrualPage.Receiveable = accrualSummary.Receiveable
	billAccrualPage.Payable = accrualSummary.Payable
	if billAccruals == nil || len(billAccruals) == 0 {
		return *mgresult.Success(billAccrualPage)
	}
	parseDay := billAccruals[0].BillTime[:10]
	if excuteDate == 1 {
		parseDay = billAccruals[0].ExcuteDate[:10]
	}
	billAccrualGroupByDayList := make([]model.BillAccrualGroupByDay, 0)
	billAccrualGroupByDay := new(model.BillAccrualGroupByDay)
	accrualList := make([]model.Accrual, 0)
	for i, billAccrual := range billAccruals {
		theday := billAccrual.BillTime[:10]
		if excuteDate == 1 {
			theday = billAccrual.ExcuteDate[:10]
		}
		if i == 0 || theday != parseDay {
			if i > 0 {
				billAccrualGroupByDay.AccrualList = accrualList
				billAccrualGroupByDayList = append(billAccrualGroupByDayList, *billAccrualGroupByDay)
				billAccrualGroupByDay = new(model.BillAccrualGroupByDay)
			}
			parseDay = theday
			billAccrualGroupByDay.GroupDate = parseDay
			accrualList = make([]model.Accrual, 0)
		}
		if billAccrual.InOut == 3 {
			billAccrualGroupByDay.Receiveable += billAccrual.RemainAmount
		} else if billAccrual.InOut == 4 {
			billAccrualGroupByDay.Payable += billAccrual.RemainAmount
		}
		var accrual model.Accrual
		utils.FromJSON(utils.ToJSON(billAccrual), &accrual)
		accrualList = append(accrualList, accrual)
	}
	billAccrualGroupByDay.AccrualList = accrualList
	billAccrualGroupByDayList = append(billAccrualGroupByDayList, *billAccrualGroupByDay)
	billAccrualPage.BillAccrual = billAccrualGroupByDayList
	return *mgresult.Success(billAccrualPage)
}

func AddBillAccrual(bookId, billTypeId, customerId, remark, photos, billTime, excuteDate string, inOut int, amount float64) mgresult.Result {
	addResult := nacos.AddBillAccrual(bookId, billTypeId, customerId, remark, photos, billTime, excuteDate, inOut, amount)
	if addResult.Status == 1 {
		return nacos.UpdateAccrualSummary(bookId, billTypeId, customerId, inOut, amount)
	} else {
		return addResult
	}
}

func UpdateBillAccrual(bookId, billId, billTypeId, customerId, remark, photos, billTime, excuteDate string, amount float64) mgresult.Result {
	accrual := nacos.GetBillAccrual(billId)
	updateResult := nacos.UpdateBillAccrual(bookId, billId, billTypeId, customerId, remark, photos, billTime, excuteDate, amount)
	if updateResult.Status == 1 {
		var updatedAccrual model2.BillAccrual
		utils.FromJSON(utils.ToJSON(updateResult.Data), &updatedAccrual)
		nacos.UpdateAccrualSummary(bookId, accrual.BillTypeId, accrual.BusinessCustomerId, accrual.InOut, -accrual.RemainAmount)
		return nacos.UpdateAccrualSummary(bookId, updatedAccrual.BillTypeId, updatedAccrual.BusinessCustomerId, accrual.InOut, updatedAccrual.RemainAmount)
	} else {
		return updateResult
	}
}

func DeleteBillAccrual(bookId, billId string) mgresult.Result {
	accrual := nacos.GetBillAccrual(billId)
	delResult := nacos.DeleteBillAccrual(bookId, billId)
	if delResult.Status == 1 {
		return nacos.UpdateAccrualSummary(bookId, accrual.BillTypeId, accrual.BusinessCustomerId, accrual.InOut, -accrual.RemainAmount)
	} else {
		return delResult
	}
}

func AddBillRecord(bookId, billId, billAccountId, remark, photos string, amount float64, sync int) mgresult.Result {
	accrual := nacos.GetBillAccrual(billId)
	result := nacos.AddBillRecord(bookId, billId, billAccountId, remark, photos, amount, sync)
	if result.Status == 1 {
		return nacos.UpdateAccrualSummary(bookId, accrual.BillTypeId, accrual.BusinessCustomerId, accrual.InOut, -amount)
	} else {
		return result
	}
}

func DeleteBillRecord(bookId, billId string, billRecordIndex int) mgresult.Result {
	accrual := nacos.GetBillAccrual(billId)
	result := nacos.DeleteBillRecord(bookId, billId, billRecordIndex)
	if result.Status == 1 {
		return nacos.UpdateAccrualSummary(bookId, accrual.BillTypeId, accrual.BusinessCustomerId, accrual.InOut, accrual.BillRecords[billRecordIndex].Amount)
	} else {
		return result
	}
}

func UpdateBillRecord(bookId, billId, billAccountId, remark, photos string, amount float64, billRecordIndex int) mgresult.Result {
	accrual := nacos.GetBillAccrual(billId)
	result := nacos.UpdateBillRecord(bookId, billId, billAccountId, remark, photos, amount, billRecordIndex)
	if result.BillId != "" {
		diff := result.RemainAmount - accrual.RemainAmount
		return nacos.UpdateAccrualSummary(bookId, accrual.BillTypeId, accrual.BusinessCustomerId, accrual.InOut, diff)
	} else {
		return *mgresult.Error(-1, "修改失败")
	}
}

func ListTypeAccrualSummary(bookId string, inOut int) mgresult.Result {
	accrualSummary := nacos.ListTypeAccrualSummary(bookId, inOut)
	if accrualSummary == nil || len(accrualSummary) == 0 {
		return *mgresult.Success([]string{})
	}
	summary := make(map[string]interface{})
	for k, v := range accrualSummary {
		if v > 0.0 {
			billType := nacos.GetBillType(k)
			summary[billType.BillType] = v
		}
	}
	result := utils.SortMapByValueDesc(summary)
	return *mgresult.Success(result)
}

func ListCustomerAccrualSummary(bookId string, inOut int) mgresult.Result {
	accrualSummary := nacos.ListCustomerAccrualSummary(bookId, inOut)
	if accrualSummary == nil || len(accrualSummary) == 0 {
		return *mgresult.Success([]string{})
	}
	summary := make(map[string]interface{})
	for k, v := range accrualSummary {
		if v > 0.0 {
			businessCustomer := nacos.GetBusinessCustomer(bookId, k)
			summary[businessCustomer.BusinessCustomer] = v
		}
	}
	result := utils.SortMapByValueDesc(summary)
	return *mgresult.Success(result)
}
