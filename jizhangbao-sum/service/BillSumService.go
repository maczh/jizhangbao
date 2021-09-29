package service

import (
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/mgconfig"
	"github.com/maczh/utils"
	"ququ.im/jizhangbao-sum/model"
	"strconv"
	"strings"
	"time"
)

var OperateInfo = map[int]string{1: "收入", 2: "支出", 3: "应收", 4: "应付"}

//更新实时统计数据
func UpdateBillSummary(bookId, billTypeId, customerId, billTime string, inOut int, amount float64) mgresult.Result {
	now := time.Now()
	if billTime != "" {
		if strings.Contains(billTime, ":") {
			now, _ = time.Parse("2006-01-02 15:04:05", billTime)
		} else {
			now, _ = time.Parse("2006-01-02", billTime)
		}
	}
	today := utils.DateFormat(now, "yyyy-MM-dd")
	month := utils.DateFormat(now, "yyyy-MM")
	year := utils.DateFormat(now, "yyyy")
	week := strconv.Itoa(utils.WeekByDate(now))
	//总数统计 totals
	//日统计
	mgconfig.Redis.HIncrByFloat("jzb:"+bookId+":"+OperateInfo[inOut], "total:day:"+today, amount)
	//月统计
	mgconfig.Redis.HIncrByFloat("jzb:"+bookId+":"+OperateInfo[inOut], "total:month:"+month, amount)
	//年统计
	mgconfig.Redis.HIncrByFloat("jzb:"+bookId+":"+OperateInfo[inOut], "total:year:"+year, amount)
	//周统计
	mgconfig.Redis.HIncrByFloat("jzb:"+bookId+":"+OperateInfo[inOut], "total:week:"+year+":"+week, amount)

	//按账单分类统计
	//日统计
	mgconfig.Redis.HIncrByFloat("jzb:"+bookId+":"+OperateInfo[inOut], "type:"+billTypeId+":day:"+today, amount)
	//月统计
	mgconfig.Redis.HIncrByFloat("jzb:"+bookId+":"+OperateInfo[inOut], "type:"+billTypeId+":month:"+month, amount)
	//年统计
	mgconfig.Redis.HIncrByFloat("jzb:"+bookId+":"+OperateInfo[inOut], "type:"+billTypeId+":year:"+year, amount)
	//周统计
	mgconfig.Redis.HIncrByFloat("jzb:"+bookId+":"+OperateInfo[inOut], "type:"+billTypeId+":week:"+year+":"+week, amount)

	//按往来客户统计
	if customerId != "" {
		//日统计
		mgconfig.Redis.HIncrByFloat("jzb:"+bookId+":"+OperateInfo[inOut], "customer:"+customerId+":day:"+today, amount)
		//月统计
		mgconfig.Redis.HIncrByFloat("jzb:"+bookId+":"+OperateInfo[inOut], "customer:"+customerId+":month:"+month, amount)
		//年统计
		mgconfig.Redis.HIncrByFloat("jzb:"+bookId+":"+OperateInfo[inOut], "customer:"+customerId+":year:"+year, amount)
		//周统计
		mgconfig.Redis.HIncrByFloat("jzb:"+bookId+":"+OperateInfo[inOut], "customer:"+customerId+":week:"+year+":"+week, amount)
	}
	summary := mgconfig.Redis.HGetAll("jzb:" + bookId + ":" + OperateInfo[inOut]).Val()
	return *mgresult.Success(summary)
}

func UpdateAccrualSummary(bookId, billTypeId, customerId string, inOut int, amount float64) mgresult.Result {
	//总额
	mgconfig.Redis.HIncrByFloat("jzb:"+bookId+":"+OperateInfo[inOut], "total", amount)
	//分类
	mgconfig.Redis.HIncrByFloat("jzb:"+bookId+":"+OperateInfo[inOut], "type:"+billTypeId, amount)
	//往来客户
	if customerId != "" {
		mgconfig.Redis.HIncrByFloat("jzb:"+bookId+":"+OperateInfo[inOut], "customer:"+customerId, amount)
	}
	summary := mgconfig.Redis.HGetAll("jzb:" + bookId + ":" + OperateInfo[inOut]).Val()
	return *mgresult.Success(summary)
}

func GetBillDaySummary(bookId, day, billTypeId, customerId string) mgresult.Result {
	var billSummary model.BillSummary
	if billTypeId != "" {
		billSummary.Income, _ = mgconfig.Redis.HGet("jzb:"+bookId+":收入", "type:"+billTypeId+":day:"+day).Float64()
		billSummary.Expense, _ = mgconfig.Redis.HGet("jzb:"+bookId+":支出", "type:"+billTypeId+":day:"+day).Float64()
		return *mgresult.Success(billSummary)
	}
	if customerId != "" {
		billSummary.Income, _ = mgconfig.Redis.HGet("jzb:"+bookId+":收入", "customer:"+customerId+":day:"+day).Float64()
		billSummary.Expense, _ = mgconfig.Redis.HGet("jzb:"+bookId+":支出", "customer:"+customerId+":day:"+day).Float64()
		return *mgresult.Success(billSummary)
	}
	billSummary.Income, _ = mgconfig.Redis.HGet("jzb:"+bookId+":收入", "total:day:"+day).Float64()
	billSummary.Expense, _ = mgconfig.Redis.HGet("jzb:"+bookId+":支出", "total:day:"+day).Float64()
	return *mgresult.Success(billSummary)
}

func GetBillMonthSummary(bookId, month, billTypeId, customerId string) mgresult.Result {
	var billSummary model.BillSummary
	if billTypeId != "" {
		billSummary.Income, _ = mgconfig.Redis.HGet("jzb:"+bookId+":收入", "type:"+billTypeId+":month:"+month).Float64()
		billSummary.Expense, _ = mgconfig.Redis.HGet("jzb:"+bookId+":支出", "type:"+billTypeId+":month:"+month).Float64()
		return *mgresult.Success(billSummary)
	}
	if customerId != "" {
		billSummary.Income, _ = mgconfig.Redis.HGet("jzb:"+bookId+":收入", "customer:"+customerId+":month:"+month).Float64()
		billSummary.Expense, _ = mgconfig.Redis.HGet("jzb:"+bookId+":支出", "customer:"+customerId+":month:"+month).Float64()
		return *mgresult.Success(billSummary)
	}
	billSummary.Income, _ = mgconfig.Redis.HGet("jzb:"+bookId+":收入", "total:month:"+month).Float64()
	billSummary.Expense, _ = mgconfig.Redis.HGet("jzb:"+bookId+":支出", "total:month:"+month).Float64()
	return *mgresult.Success(billSummary)
}

func GetBillYearSummary(bookId, year, billTypeId, customerId string) mgresult.Result {
	var billSummary model.BillSummary
	if billTypeId != "" {
		billSummary.Income, _ = mgconfig.Redis.HGet("jzb:"+bookId+":收入", "type:"+billTypeId+":year:"+year).Float64()
		billSummary.Expense, _ = mgconfig.Redis.HGet("jzb:"+bookId+":支出", "type:"+billTypeId+":year:"+year).Float64()
		return *mgresult.Success(billSummary)
	}
	if customerId != "" {
		billSummary.Income, _ = mgconfig.Redis.HGet("jzb:"+bookId+":收入", "customer:"+customerId+":year:"+year).Float64()
		billSummary.Expense, _ = mgconfig.Redis.HGet("jzb:"+bookId+":支出", "customer:"+customerId+":year:"+year).Float64()
		return *mgresult.Success(billSummary)
	}
	billSummary.Income, _ = mgconfig.Redis.HGet("jzb:"+bookId+":收入", "total:year:"+year).Float64()
	billSummary.Expense, _ = mgconfig.Redis.HGet("jzb:"+bookId+":支出", "total:year:"+year).Float64()
	return *mgresult.Success(billSummary)
}

func GetBillWeekSummary(bookId, year, week, billTypeId, customerId string) mgresult.Result {
	var billSummary model.BillSummary
	if billTypeId != "" {
		billSummary.Income, _ = mgconfig.Redis.HGet("jzb:"+bookId+":收入", "type:"+billTypeId+":week:"+year+":"+week).Float64()
		billSummary.Expense, _ = mgconfig.Redis.HGet("jzb:"+bookId+":支出", "type:"+billTypeId+":week:"+year+":"+week).Float64()
		return *mgresult.Success(billSummary)
	}
	if customerId != "" {
		billSummary.Income, _ = mgconfig.Redis.HGet("jzb:"+bookId+":收入", "customer:"+customerId+":week:"+year+":"+week).Float64()
		billSummary.Expense, _ = mgconfig.Redis.HGet("jzb:"+bookId+":支出", "customer:"+customerId+":week:"+year+":"+week).Float64()
		return *mgresult.Success(billSummary)
	}
	billSummary.Income, _ = mgconfig.Redis.HGet("jzb:"+bookId+":收入", "total:week:"+year+":"+week).Float64()
	billSummary.Expense, _ = mgconfig.Redis.HGet("jzb:"+bookId+":支出", "total:week:"+year+":"+week).Float64()
	return *mgresult.Success(billSummary)
}

func ListTypeBillSummaryByDay(bookId, day string, inOut int) mgresult.Result {
	keys, _ := mgconfig.Redis.HScan("jzb:"+bookId+":"+OperateInfo[inOut], 0, "type:*:day:"+day, 100).Val()
	mgresultMap := make(map[string]float64)
	if len(keys) == 0 {
		return *mgresult.Success(mgresultMap)
	}
	for i := 0; i < len(keys)/2; i++ {
		s := strings.Split(keys[i*2], ":")
		mgresultMap[s[1]], _ = strconv.ParseFloat(keys[i*2+1], 64)
	}
	return *mgresult.Success(mgresultMap)
}

func ListTypeBillSummaryByMonth(bookId, month string, inOut int) mgresult.Result {
	keys, _ := mgconfig.Redis.HScan("jzb:"+bookId+":"+OperateInfo[inOut], 0, "type:*:month:"+month, 100).Val()
	mgresultMap := make(map[string]float64)
	if len(keys) == 0 {
		return *mgresult.Success(mgresultMap)
	}
	for i := 0; i < len(keys)/2; i++ {
		s := strings.Split(keys[i*2], ":")
		mgresultMap[s[1]], _ = strconv.ParseFloat(keys[i*2+1], 64)
	}
	return *mgresult.Success(mgresultMap)
}

func ListTypeBillSummaryByYear(bookId, year string, inOut int) mgresult.Result {
	keys, _ := mgconfig.Redis.HScan("jzb:"+bookId+":"+OperateInfo[inOut], 0, "type:*:year:"+year, 100).Val()
	mgresultMap := make(map[string]float64)
	if len(keys) == 0 {
		return *mgresult.Success(mgresultMap)
	}
	for i := 0; i < len(keys)/2; i++ {
		s := strings.Split(keys[i*2], ":")
		mgresultMap[s[1]], _ = strconv.ParseFloat(keys[i*2+1], 64)
	}
	return *mgresult.Success(mgresultMap)
}

func ListTypeBillSummaryByWeek(bookId, year, week string, inOut int) mgresult.Result {
	keys, _ := mgconfig.Redis.HScan("jzb:"+bookId+":"+OperateInfo[inOut], 0, "type:*:week:"+year+":"+week, 100).Val()
	mgresultMap := make(map[string]float64)
	if len(keys) == 0 {
		return *mgresult.Success(mgresultMap)
	}
	for i := 0; i < len(keys)/2; i++ {
		s := strings.Split(keys[i*2], ":")
		mgresultMap[s[1]], _ = strconv.ParseFloat(keys[i*2+1], 64)
	}
	return *mgresult.Success(mgresultMap)
}

func ListCustomerBillSummaryByDay(bookId, day string, inOut int) mgresult.Result {
	keys, _ := mgconfig.Redis.HScan("jzb:"+bookId+":"+OperateInfo[inOut], 0, "customer:*:day:"+day, 100).Val()
	mgresultMap := make(map[string]float64)
	if len(keys) == 0 {
		return *mgresult.Success(mgresultMap)
	}
	for i := 0; i < len(keys)/2; i++ {
		s := strings.Split(keys[i*2], ":")
		mgresultMap[s[1]], _ = strconv.ParseFloat(keys[i*2+1], 64)
	}
	return *mgresult.Success(mgresultMap)
}

func ListCustomerBillSummaryByMonth(bookId, month string, inOut int) mgresult.Result {
	keys, _ := mgconfig.Redis.HScan("jzb:"+bookId+":"+OperateInfo[inOut], 0, "customer:*:month:"+month, 100).Val()
	mgresultMap := make(map[string]float64)
	if len(keys) == 0 {
		return *mgresult.Success(mgresultMap)
	}
	for i := 0; i < len(keys)/2; i++ {
		s := strings.Split(keys[i*2], ":")
		mgresultMap[s[1]], _ = strconv.ParseFloat(keys[i*2+1], 64)
	}
	return *mgresult.Success(mgresultMap)
}

func ListCustomerBillSummaryByYear(bookId, year string, inOut int) mgresult.Result {
	keys, _ := mgconfig.Redis.HScan("jzb:"+bookId+":"+OperateInfo[inOut], 0, "customer:*:year:"+year, 100).Val()
	mgresultMap := make(map[string]float64)
	if len(keys) == 0 {
		return *mgresult.Success(mgresultMap)
	}
	for i := 0; i < len(keys)/2; i++ {
		s := strings.Split(keys[i*2], ":")
		mgresultMap[s[1]], _ = strconv.ParseFloat(keys[i*2+1], 64)
	}
	return *mgresult.Success(mgresultMap)
}

func ListCustomerBillSummaryByWeek(bookId, year, week string, inOut int) mgresult.Result {
	keys, _ := mgconfig.Redis.HScan("jzb:"+bookId+":"+OperateInfo[inOut], 0, "customer:*:week:"+year+":"+week, 100).Val()
	mgresultMap := make(map[string]float64)
	if len(keys) == 0 {
		return *mgresult.Success(mgresultMap)
	}
	for i := 0; i < len(keys)/2; i++ {
		s := strings.Split(keys[i*2], ":")
		mgresultMap[s[1]], _ = strconv.ParseFloat(keys[i*2+1], 64)
	}
	return *mgresult.Success(mgresultMap)
}

func GetAccrualSummary(bookId, billTypeId, customerId string) mgresult.Result {
	var accrualSummary model.AccrualSummary
	if billTypeId != "" {
		accrualSummary.Receiveable, _ = mgconfig.Redis.HGet("jzb:"+bookId+":应收", "type:"+billTypeId).Float64()
		accrualSummary.Payable, _ = mgconfig.Redis.HGet("jzb:"+bookId+":应付", "type:"+billTypeId).Float64()
		return *mgresult.Success(accrualSummary)
	}
	if customerId != "" {
		accrualSummary.Receiveable, _ = mgconfig.Redis.HGet("jzb:"+bookId+":应收", "customer:"+customerId).Float64()
		accrualSummary.Payable, _ = mgconfig.Redis.HGet("jzb:"+bookId+":应付", "customer:"+customerId).Float64()
		return *mgresult.Success(accrualSummary)
	}
	accrualSummary.Receiveable, _ = mgconfig.Redis.HGet("jzb:"+bookId+":应收", "total").Float64()
	accrualSummary.Payable, _ = mgconfig.Redis.HGet("jzb:"+bookId+":应付", "total").Float64()
	return *mgresult.Success(accrualSummary)
}

func ListTypeAccrualSummary(bookId string, inOut int) mgresult.Result {
	keys, _ := mgconfig.Redis.HScan("jzb:"+bookId+":"+OperateInfo[inOut], 0, "type:*", 100).Val()
	mgresultMap := make(map[string]float64)
	if len(keys) == 0 {
		return *mgresult.Success(mgresultMap)
	}
	for i := 0; i < len(keys)/2; i++ {
		s := strings.Split(keys[i*2], ":")
		mgresultMap[s[1]], _ = strconv.ParseFloat(keys[i*2+1], 64)
	}
	return *mgresult.Success(mgresultMap)
}

func ListCustomerAccrualSummary(bookId string, inOut int) mgresult.Result {
	keys, _ := mgconfig.Redis.HScan("jzb:"+bookId+":"+OperateInfo[inOut], 0, "customer:*", 100).Val()
	mgresultMap := make(map[string]float64)
	if len(keys) == 0 {
		return *mgresult.Success(mgresultMap)
	}
	for i := 0; i < len(keys)/2; i++ {
		s := strings.Split(keys[i*2], ":")
		mgresultMap[s[1]], _ = strconv.ParseFloat(keys[i*2+1], 64)
	}
	return *mgresult.Success(mgresultMap)
}
