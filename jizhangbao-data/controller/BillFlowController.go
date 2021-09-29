package controller

import (
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/utils"
	"ququ.im/jizhangbao-data/service"
	"strconv"
)

// AddBillFlow	godoc
// @Summary		添加收支账单
// @Description	添加收支账单
// @Tags	收支账单
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	billTypeId formData string true "账单类型编号"
// @Param	billAccountId formData string true "资金账户编号"
// @Param	customerId formData string false "往来客户编号"
// @Param	billTime formData string false "记账时间，格式 yyyy-MM-dd hh:mm:ss，不传为当前时间"
// @Param	remark formData string false "备注"
// @Param	photos formData string false "账单相关照片url数组，JSON数组格式"
// @Param	inOut formData int true "账单收支类型编号 1-收入 2-支出"
// @Param	amount formData float64 true "账单金额，单位为元，最多2位小数，正数"
// @Success 200 {string} string	"ok"
// @Router	/flow/add [post]
func AddBillFlow(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "billTypeId") {
		return *mgresult.Error(-1, "缺少账单类型编号参数")
	}
	if !utils.Exists(params, "billAccountId") {
		return *mgresult.Error(-1, "缺少资金账户编号参数")
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
	return service.AddBillFlow(params["bookId"], params["billTypeId"], params["billAccountId"], params["customerId"], params["remark"], params["photos"], "", params["billTime"], inOut, amount)
}

// UpdateBillFlow	godoc
// @Summary		修改收支账单
// @Description	修改收支账单
// @Tags	收支账单
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	billId formData string true "账单编号"
// @Param	billTypeId formData string false "账单类型编号"
// @Param	billAccountId formData string false "资金账户编号"
// @Param	customerId formData string false "往来客户编号"
// @Param	billTime formData string false "记账时间，格式 yyyy-MM-dd hh:mm:ss"
// @Param	remark formData string false "备注"
// @Param	photos formData string false "账单相关照片url数组，JSON数组格式"
// @Param	amount formData float64 false "账单金额，单位为元，最多2位小数，正数"
// @Success 200 {string} string	"ok"
// @Router	/flow/update [post]
func UpdateBillFlow(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "billId") {
		return *mgresult.Error(-1, "缺少账单编号参数")
	}
	amount := 0.0
	if utils.Exists(params, "amount") {
		amount, _ = strconv.ParseFloat(params["amount"], 64)
	}
	return service.UpdateBillFlow(params["bookId"], params["billId"], params["billTypeId"], params["billAccountId"], params["customerId"], params["remark"], params["photos"], params["billTime"], amount)
}

// DeleteBillFlow	godoc
// @Summary		删除收支账单
// @Description	删除收支账单
// @Tags	收支账单
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	billId formData string true "账单编号"
// @Success 200 {string} string	"ok"
// @Router	/flow/del [post]
func DeleteBillFlow(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "billId") {
		return *mgresult.Error(-1, "缺少账单编号参数")
	}
	return service.DeleteBillFlow(params["bookId"], params["billId"])
}

// GetBillFlow	godoc
// @Summary		获取指定收支账单
// @Description	获取指定收支账单
// @Tags	收支账单
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	billId formData string true "账单编号"
// @Success 200 {string} string	"ok"
// @Router	/flow/get [post]
func GetBillFlow(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "billId") {
		return *mgresult.Error(-1, "缺少账单编号参数")
	}
	return service.GetBillFlow(params["billId"])
}

// ListBillFlow	godoc
// @Summary		查询收支账单列表API
// @Description	查询收支账单列表，按时间段，带分页
// @Tags	收支账单
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	year formData string false "按年份查询，格式 yyyy"
// @Param	month formData string false "按指定月份查询，格式 yyyy-MM"
// @Param	week formData string false "按某年的第week周查询，1-52周，必须与year同时使用"
// @Param	billTypeId formData string false "账单类型编号"
// @Param	customerId formData string false "往来客户编号"
// @Param	inOut formData int false "账单收支类型编号 1-收入 2-支出 0或不传为全部"
// @Param	page formData int false "分页页号参数，第几页，若要分页，页号>=1，不分页则不传"
// @Param	size formData int false "分页的页大小参数，不传为不分页"
// @Success 200 {string} string	"ok"
// @Router	/flow/list [post]
func ListBillFlow(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	inOut := 0
	if utils.Exists(params, "inOut") {
		inOut, _ = strconv.Atoi(params["inOut"])
	}
	page := 1
	if utils.Exists(params, "page") {
		page, _ = strconv.Atoi(params["page"])
	}
	size := 0
	if utils.Exists(params, "size") {
		size, _ = strconv.Atoi(params["size"])
	}
	return service.ListBillFlow(params["bookId"], params["year"], params["month"], params["week"], params["billTypeId"], params["customerId"], inOut, page, size)
}
