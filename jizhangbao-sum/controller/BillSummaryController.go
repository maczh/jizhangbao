package controller

import (
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/utils"
	"ququ.im/jizhangbao-sum/service"
	"strconv"
)

// UpdateBillSummary	godoc
// @Summary		根据收支账单金额更新账单统计数据
// @Description	根据收支账单金额更新账单统计数据
// @Tags	更新统计
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	billTypeId formData string true "账单类型编号"
// @Param	customerId formData string false "往来客户编号"
// @Param	billTime formData string false "记账时间，格式 yyyy-MM-dd hh:mm:ss，不传为当前时间"
// @Param	inOut formData int true "账单收支类型编号 1-收入 2-支出"
// @Param	amount formData float64 true "账单金额，单位为元，最多2位小数，可以是正负数"
// @Success 200 {string} string	"ok"
// @Router	/bill/incr [post]
func UpdateBillSummary(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "billTypeId") {
		return *mgresult.Error(-1, "缺少账单类型编号参数")
	}
	if !utils.Exists(params, "inOut") {
		return *mgresult.Error(-1, "缺少账单收支类型编号参数，1-收入 2-支出")
	}
	inOut, _ := strconv.Atoi(params["inOut"])
	if !utils.Exists(params, "amount") {
		return *mgresult.Error(-1, "缺少账单金额参数")
	}
	amount, _ := strconv.ParseFloat(params["amount"], 64)
	return service.UpdateBillSummary(params["bookId"], params["billTypeId"], params["customerId"], params["billTime"], inOut, amount)
}

// UpdateAccrualSummary	godoc
// @Summary		根据应收应付金额更新统计数据
// @Description	根据应收应付金额更新统计数据
// @Tags	更新统计
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	billTypeId formData string true "账单类型编号"
// @Param	customerId formData string false "往来客户编号"
// @Param	inOut formData int true "账单类型编号"
// @Param	amount formData float64 true "账单金额，单位为元，最多2位小数，可以是正负数"
// @Success 200 {string} string	"ok"
// @Router	/accrual/incr [post]
func UpdateAccrualSummary(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "billTypeId") {
		return *mgresult.Error(-1, "缺少账单类型编号参数")
	}
	if !utils.Exists(params, "inOut") {
		return *mgresult.Error(-1, "缺少账单收支类型编号参数，1-收入 2-支出")
	}
	inOut, _ := strconv.Atoi(params["inOut"])
	if !utils.Exists(params, "amount") {
		return *mgresult.Error(-1, "缺少账单金额参数")
	}
	amount, _ := strconv.ParseFloat(params["amount"], 64)
	return service.UpdateAccrualSummary(params["bookId"], params["billTypeId"], params["customerId"], inOut, amount)
}

// GetBillDaySummary	godoc
// @Summary		根据日期查询账单统计数据
// @Description	根据日期查询账单统计数据
// @Tags	查询统计
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	day formData string true "查询日期，格式 yyyy-MM-dd"
// @Param	billTypeId formData string false "账单类型编号"
// @Param	customerId formData string false "往来客户编号"
// @Success 200 {string} string	"ok"
// @Router	/bill/get/day [post]
func GetBillDaySummary(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "day") {
		return *mgresult.Error(-1, "缺少日期参数，如2021-03-25")
	}
	return service.GetBillDaySummary(params["bookId"], params["day"], params["billTypeId"], params["customerId"])
}

// GetBillMonthSummary	godoc
// @Summary		根据月份查询账单统计数据
// @Description	根据月份查询账单统计数据
// @Tags	查询统计
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	month formData string true "查询月份，格式 yyyy-MM"
// @Param	billTypeId formData string false "账单类型编号"
// @Param	customerId formData string false "往来客户编号"
// @Success 200 {string} string	"ok"
// @Router	/bill/get/month [post]
func GetBillMonthSummary(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "month") {
		return *mgresult.Error(-1, "缺少月份参数，如2021-03")
	}
	return service.GetBillMonthSummary(params["bookId"], params["month"], params["billTypeId"], params["customerId"])
}

// GetBillYearSummary	godoc
// @Summary		根据年份查询账单统计数据
// @Description	根据年份查询账单统计数据
// @Tags	查询统计
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	year formData string true "查询年份，格式 yyyy"
// @Param	billTypeId formData string false "账单类型编号"
// @Param	customerId formData string false "往来客户编号"
// @Success 200 {string} string	"ok"
// @Router	/bill/get/year [post]
func GetBillYearSummary(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "year") {
		return *mgresult.Error(-1, "缺少年份参数，如2021")
	}
	return service.GetBillYearSummary(params["bookId"], params["year"], params["billTypeId"], params["customerId"])
}

// GetBillWeekSummary	godoc
// @Summary		根据周查询账单统计数据
// @Description	根据周查询账单统计数据
// @Tags	查询统计
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	year formData string true "查询年份，格式 yyyy"
// @Param	week formData string true "查询周数，1-52"
// @Param	billTypeId formData string false "账单类型编号"
// @Param	customerId formData string false "往来客户编号"
// @Success 200 {string} string	"ok"
// @Router	/bill/get/week [post]
func GetBillWeekSummary(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "year") {
		return *mgresult.Error(-1, "缺少年份参数，如2021")
	}
	if !utils.Exists(params, "week") {
		return *mgresult.Error(-1, "缺少周数参数，如:6")
	}
	return service.GetBillWeekSummary(params["bookId"], params["year"], params["week"], params["billTypeId"], params["customerId"])
}

// ListTypeBillSummaryByDay	godoc
// @Summary		根据日期列出账单分类分组统计数据
// @Description	根据日期列出账单分类分组统计数据
// @Tags	查询统计
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	day formData string true "查询日期，格式 yyyy-MM-dd"
// @Param	inOut formData int true "账单收支类型编号 1-收入 2-支出"
// @Success 200 {string} string	"ok"
// @Router	/bill/list/type/day [post]
func ListTypeBillSummaryByDay(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "day") {
		return *mgresult.Error(-1, "缺少日期参数，如2021-03-25")
	}
	if !utils.Exists(params, "inOut") {
		return *mgresult.Error(-1, "缺少账单收支类型编号参数，1-收入 2-支出")
	}
	inOut, _ := strconv.Atoi(params["inOut"])
	return service.ListTypeBillSummaryByDay(params["bookId"], params["day"], inOut)
}

// ListTypeBillSummaryByMonth	godoc
// @Summary		根据月份列出账单分类分组统计数据
// @Description	根据月份列出账单分类分组统计数据
// @Tags	查询统计
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	month formData string true "查询月份，格式 yyyy-MM"
// @Param	inOut formData int true "账单收支类型编号 1-收入 2-支出"
// @Success 200 {string} string	"ok"
// @Router	/bill/list/type/month [post]
func ListTypeBillSummaryByMonth(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "month") {
		return *mgresult.Error(-1, "缺少月份参数，如2021-03")
	}
	if !utils.Exists(params, "inOut") {
		return *mgresult.Error(-1, "缺少账单收支类型编号参数，1-收入 2-支出")
	}
	inOut, _ := strconv.Atoi(params["inOut"])
	return service.ListTypeBillSummaryByMonth(params["bookId"], params["month"], inOut)
}

// ListTypeBillSummaryByYear	godoc
// @Summary		根据年份列出账单分类分组统计数据
// @Description	根据年份列出账单分类分组统计数据
// @Tags	查询统计
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	year formData string true "查询年份，格式 yyyy"
// @Param	inOut formData int true "账单收支类型编号 1-收入 2-支出"
// @Success 200 {string} string	"ok"
// @Router	/bill/list/type/year [post]
func ListTypeBillSummaryByYear(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "day") {
		return *mgresult.Error(-1, "缺少年份参数，如2021")
	}
	if !utils.Exists(params, "inOut") {
		return *mgresult.Error(-1, "缺少账单收支类型编号参数，1-收入 2-支出")
	}
	inOut, _ := strconv.Atoi(params["inOut"])
	return service.ListTypeBillSummaryByYear(params["bookId"], params["year"], inOut)
}

// ListTypeBillSummaryByWeek	godoc
// @Summary		根据周列出账单分类分组统计数据
// @Description	根据周列出账单分类分组统计数据
// @Tags	查询统计
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	year formData string true "查询年份，格式 yyyy"
// @Param	week formData string true "查询周数，1-52"
// @Param	inOut formData int true "账单收支类型编号 1-收入 2-支出"
// @Success 200 {string} string	"ok"
// @Router	/bill/list/type/week [post]
func ListTypeBillSummaryByWeek(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "year") {
		return *mgresult.Error(-1, "缺少年份参数，如2021")
	}
	if !utils.Exists(params, "week") {
		return *mgresult.Error(-1, "缺少周数参数，如:6")
	}
	if !utils.Exists(params, "inOut") {
		return *mgresult.Error(-1, "缺少账单收支类型编号参数，1-收入 2-支出")
	}
	inOut, _ := strconv.Atoi(params["inOut"])
	return service.ListTypeBillSummaryByWeek(params["bookId"], params["year"], params["week"], inOut)
}

// ListCustomerBillSummaryByDay	godoc
// @Summary		根据日期列出往来客户分组统计数据
// @Description	根据日期列出往来客户分组统计数据
// @Tags	查询统计
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	day formData string true "查询日期，格式 yyyy-MM-dd"
// @Param	inOut formData int true "账单收支类型编号 1-收入 2-支出"
// @Success 200 {string} string	"ok"
// @Router	/bill/list/customer/day [post]
func ListCustomerBillSummaryByDay(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "day") {
		return *mgresult.Error(-1, "缺少日期参数，如2021-03-25")
	}
	if !utils.Exists(params, "inOut") {
		return *mgresult.Error(-1, "缺少账单收支类型编号参数，1-收入 2-支出")
	}
	inOut, _ := strconv.Atoi(params["inOut"])
	return service.ListCustomerBillSummaryByDay(params["bookId"], params["day"], inOut)
}

// ListCustomerBillSummaryByMonth	godoc
// @Summary		根据月份列出往来客户分组统计数据
// @Description	根据月份列出往来客户分组统计数据
// @Tags	查询统计
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	month formData string true "查询月份，格式 yyyy-MM"
// @Param	inOut formData int true "账单收支类型编号 1-收入 2-支出"
// @Success 200 {string} string	"ok"
// @Router	/bill/list/customer/month [post]
func ListCustomerBillSummaryByMonth(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "month") {
		return *mgresult.Error(-1, "缺少月份参数，如2021-03")
	}
	if !utils.Exists(params, "inOut") {
		return *mgresult.Error(-1, "缺少账单收支类型编号参数，1-收入 2-支出")
	}
	inOut, _ := strconv.Atoi(params["inOut"])
	return service.ListCustomerBillSummaryByMonth(params["bookId"], params["month"], inOut)
}

// ListCustomerBillSummaryByYear	godoc
// @Summary		根据年份列出往来客户分组统计数据
// @Description	根据年份列出往来客户分组统计数据
// @Tags	查询统计
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	year formData string true "查询年份，格式 yyyy"
// @Param	inOut formData int true "账单收支类型编号 1-收入 2-支出"
// @Success 200 {string} string	"ok"
// @Router	/bill/list/customer/year [post]
func ListCustomerBillSummaryByYear(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "day") {
		return *mgresult.Error(-1, "缺少年份参数，如2021")
	}
	if !utils.Exists(params, "inOut") {
		return *mgresult.Error(-1, "缺少账单收支类型编号参数，1-收入 2-支出")
	}
	inOut, _ := strconv.Atoi(params["inOut"])
	return service.ListCustomerBillSummaryByYear(params["bookId"], params["year"], inOut)
}

// ListCustomerBillSummaryByWeek	godoc
// @Summary		根据周列出往来客户分组统计数据
// @Description	根据周列出往来客户分组统计数据
// @Tags	查询统计
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	year formData string true "查询年份，格式 yyyy"
// @Param	week formData string true "查询周数，1-52"
// @Param	inOut formData int true "账单收支类型编号 1-收入 2-支出"
// @Success 200 {string} string	"ok"
// @Router	/bill/list/customer/week [post]
func ListCustomerBillSummaryByWeek(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "year") {
		return *mgresult.Error(-1, "缺少年份参数，如2021")
	}
	if !utils.Exists(params, "week") {
		return *mgresult.Error(-1, "缺少周数参数，如:6")
	}
	if !utils.Exists(params, "inOut") {
		return *mgresult.Error(-1, "缺少账单收支类型编号参数，1-收入 2-支出")
	}
	inOut, _ := strconv.Atoi(params["inOut"])
	return service.ListCustomerBillSummaryByWeek(params["bookId"], params["year"], params["week"], inOut)
}

// GetAccrualSummary	godoc
// @Summary		根据账单类型或往来客户获取应收应付统计数据
// @Description	根据账单类型或往来客户获取应收应付统计数据
// @Tags	查询统计
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	billTypeId formData string false "账单类型编号"
// @Param	customerId formData string false "往来客户编号"
// @Success 200 {string} string	"ok"
// @Router	/accrual/get [post]
func GetAccrualSummary(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	return service.GetAccrualSummary(params["bookId"], params["billTypeId"], params["customerId"])
}

// ListTypeAccrualSummary	godoc
// @Summary		根据账单类型分组统计应收/应付数据
// @Description	根据账单类型分组统计应收/应付数据
// @Tags	查询统计
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	inOut formData int true "账单收支类型编号 3-应收 4-应付"
// @Success 200 {string} string	"ok"
// @Router	/accrual/list/type [post]
func ListTypeAccrualSummary(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "inOut") {
		return *mgresult.Error(-1, "缺少账单收支类型编号参数，3-应收 4-应付")
	}
	inOut, _ := strconv.Atoi(params["inOut"])
	return service.ListTypeAccrualSummary(params["bookId"], inOut)
}

// ListCustomerAccrualSummary	godoc
// @Summary		根据往来客户分组统计应收/应付数据
// @Description	根据往来客户分组统计应收/应付数据
// @Tags	查询统计
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	inOut formData int true "账单收支类型编号 3-应收 4-应付"
// @Success 200 {string} string	"ok"
// @Router	/accrual/list/customer [post]
func ListCustomerAccrualSummary(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "inOut") {
		return *mgresult.Error(-1, "缺少账单收支类型编号参数，3-应收 4-应付")
	}
	inOut, _ := strconv.Atoi(params["inOut"])
	return service.ListCustomerAccrualSummary(params["bookId"], inOut)
}
