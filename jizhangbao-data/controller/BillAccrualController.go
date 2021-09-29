package controller

import (
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/utils"
	"ququ.im/jizhangbao-data/service"
	"strconv"
)

// AddBillAccrual	godoc
// @Summary		添加应收应付账单
// @Description	添加应收应付账单
// @Tags	应收应付账单
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	billTypeId formData string true "账单类型编号"
// @Param	customerId formData string false "往来客户编号"
// @Param	remark formData string false "备注"
// @Param	photos formData string false "账单相关照片url数组，JSON数组格式"
// @Param	billTime formData string false "记账时间，格式 yyyy-MM-dd hh:mm:ss，不传为当前时间"
// @Param	excuteDate formData string true "预约收款/付款日期，格式为 yyyy-MM-dd"
// @Param	inOut formData int true "账单应收应付类型编号 3-应收 4-应付"
// @Param	amount formData float64 true "账单金额，单位为元，最多2位小数，正数"
// @Success 200 {string} string	"ok"
// @Router	/accrual/add [post]
func AddBillAccrual(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "billTypeId") {
		return *mgresult.Error(-1, "缺少账单类型编号参数")
	}
	if !utils.Exists(params, "billTypeId") {
		return *mgresult.Error(-1, "缺少账单类型编号参数")
	}
	if !utils.Exists(params, "excuteDate") {
		return *mgresult.Error(-1, "缺少账单预约收款/付款日期参数")
	}
	if !utils.Exists(params, "inOut") {
		return *mgresult.Error(-1, "缺少账单应收应付类型编号参数，3-应收 4-应付")
	}
	inOut, _ := strconv.Atoi(params["inOut"])
	if !utils.Exists(params, "amount") {
		return *mgresult.Error(-1, "缺少账单金额参数")
	}
	amount, _ := strconv.ParseFloat(params["amount"], 64)
	return service.AddBillAccrual(params["bookId"], params["billTypeId"], params["customerId"], params["remark"], params["photos"], params["billTime"], params["excuteDate"], inOut, amount)
}

// UpdateBillAccrual	godoc
// @Summary		修改应收应付账单
// @Description	修改应收应付账单
// @Tags	应收应付账单
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	billId formData string true "账单编号"
// @Param	billTypeId formData string false "账单类型编号"
// @Param	customerId formData string false "往来客户编号"
// @Param	remark formData string false "备注"
// @Param	billTime formData string false "记账时间，格式 yyyy-MM-dd hh:mm:ss"
// @Param	excuteDate formData string false "预约收款/付款日期，格式为 yyyy-MM-dd"
// @Param	photos formData string false "账单相关照片url数组，JSON数组格式"
// @Param	amount formData float64 false "账单金额，单位为元，最多2位小数，正数"
// @Success 200 {string} string	"ok"
// @Router	/accrual/update [post]
func UpdateBillAccrual(params map[string]string) mgresult.Result {
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

	return service.UpdateBillAccrual(params["bookId"], params["billId"], params["billTypeId"], params["customerId"], params["remark"], params["photos"], params["billTime"], params["excuteDate"], amount)
}

// DeleteBillAccrual	godoc
// @Summary		删除应收应付账单
// @Description	删除应收应付账单
// @Tags	应收应付账单
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	billId formData string true "账单编号"
// @Success 200 {string} string	"ok"
// @Router	/accrual/del [post]
func DeleteBillAccrual(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "billId") {
		return *mgresult.Error(-1, "缺少账单编号参数")
	}
	return service.DeleteBillAccrual(params["bookId"], params["billId"])
}

// GetBillAccrual	godoc
// @Summary		获取指定应收应付账单
// @Description	获取指定应收应付账单
// @Tags	应收应付账单
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	billId formData string true "账单编号"
// @Success 200 {string} string	"ok"
// @Router	/accrual/get [post]
func GetBillAccrual(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "billId") {
		return *mgresult.Error(-1, "缺少账单编号参数")
	}
	return service.GetBillAccrual(params["billId"])
}

// ListBillAccrual	godoc
// @Summary		查询应收应付账单列表API
// @Description	查询应收应付账单列表，按时间段，带分页
// @Tags	应收应付账单
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	year formData string false "按年份查询，格式 yyyy"
// @Param	month formData string false "按指定月份查询，格式 yyyy-MM"
// @Param	week formData string false "按某年的第week周查询，1-52周，必须与year同时使用"
// @Param	billTypeId formData string false "账单类型编号"
// @Param	customerId formData string false "往来客户编号"
// @Param	excuteDate formData int false "是否按预约收款/付款日期查询与排序，0-否 1-是，默认为0"
// @Param	inOut formData int false "账单应收应付类型编号 3-应收 4-应付 0或不传为全部"
// @Param	hideFinished formData int false "是否隐藏已完成应收/应付账单，0-不隐藏 1-隐藏，默认为隐藏"
// @Param	page formData int false "分页页号参数，第几页，若要分页，页号>=1，不分页则不传"
// @Param	size formData int false "分页的页大小参数，不传为不分页"
// @Success 200 {string} string	"ok"
// @Router	/accrual/list [post]
func ListBillAccrual(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	inOut := 0
	if utils.Exists(params, "inOut") {
		inOut, _ = strconv.Atoi(params["inOut"])
	}
	hideFinished := 1
	if utils.Exists(params, "hideFinished") {
		hideFinished, _ = strconv.Atoi(params["hideFinished"])
	}
	excuteDate := 0
	if utils.Exists(params, "excuteDate") {
		excuteDate, _ = strconv.Atoi(params["excuteDate"])
	}
	page := 1
	if utils.Exists(params, "page") {
		page, _ = strconv.Atoi(params["page"])
	}
	size := 0
	if utils.Exists(params, "size") {
		size, _ = strconv.Atoi(params["size"])
	}

	return service.ListBillAccrual(params["bookId"], params["year"], params["month"], params["week"], params["billTypeId"], params["customerId"], excuteDate, inOut, hideFinished, page, size)
}

// AddBillRecord	godoc
// @Summary		添加应收应付账单中的收付款记录
// @Description	添加应收应付账单中的收付款记录
// @Tags	应收应付账单
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	billId formData string true "账单编号"
// @Param	billAccountId formData string true "资金账户编号"
// @Param	remark formData string false "备注"
// @Param	photos formData string false "账单相关照片url数组，JSON数组格式"
// @Param	amount formData float64 true "账单金额，单位为元，最多2位小数，正数"
// @Param	sync formData int false "是否同步生成收支账单 0-不生成 1-生成，默认不生成"
// @Success 200 {string} string	"ok"
// @Router	/accrual/record/add [post]
func AddBillRecord(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "billId") {
		return *mgresult.Error(-1, "缺少应收应付账单编号参数")
	}
	if !utils.Exists(params, "billAccountId") {
		return *mgresult.Error(-1, "缺少资金账户编号参数")
	}
	sync := 0
	if utils.Exists(params, "sync") {
		sync, _ = strconv.Atoi(params["sync"])
	}
	if !utils.Exists(params, "amount") {
		return *mgresult.Error(-1, "缺少收款金额/付款金额参数")
	}
	amount, _ := strconv.ParseFloat(params["amount"], 64)
	return service.AddBillRecord(params["bookId"], params["billId"], params["billAccountId"], params["remark"], params["photos"], amount, sync == 1)
}

// UpdateBillRecord	godoc
// @Summary		修改应收应付账单中的收付款记录
// @Description	修改应收应付账单中的收付款记录
// @Tags	应收应付账单
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	billId formData string true "账单编号"
// @Param	billAccountId formData string false "资金账户编号"
// @Param	remark formData string false "备注"
// @Param	photos formData string false "账单相关照片url数组，JSON数组格式"
// @Param	amount formData float64 false "账单金额，单位为元，最多2位小数，正数"
// @Param	billRecordIndex formData int true "收付款记录号，数组下标，从0开始"
// @Success 200 {string} string	"ok"
// @Router	/accrual/record/update [post]
func UpdateBillRecord(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "billId") {
		return *mgresult.Error(-1, "缺少应收应付账单编号参数")
	}
	if !utils.Exists(params, "billRecordIndex") {
		return *mgresult.Error(-1, "缺少账单收付款记录的记录号")
	}
	billRecordIndex, _ := strconv.Atoi(params["billRecordIndex"])
	if !utils.Exists(params, "amount") {
		return *mgresult.Error(-1, "缺少收款金额/付款金额参数")
	}
	amount, _ := strconv.ParseFloat(params["amount"], 64)
	return service.UpdateBillRecord(params["bookId"], params["billId"], params["billAccountId"], params["remark"], params["photos"], amount, billRecordIndex)
}

// DeleteBillRecord	godoc
// @Summary		添加应收应付账单中的收付款记录
// @Description	添加应收应付账单中的收付款记录
// @Tags	应收应付账单
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	billId formData string true "账单编号"
// @Param	billRecordIndex formData int true "收付款记录号，数组下标，从0开始"
// @Success 200 {string} string	"ok"
// @Router	/accrual/record/del [post]
func DeleteBillRecord(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "billId") {
		return *mgresult.Error(-1, "缺少应收应付账单编号参数")
	}
	if !utils.Exists(params, "billRecordIndex") {
		return *mgresult.Error(-1, "缺少账单收付款记录的记录号")
	}
	billRecordIndex, _ := strconv.Atoi(params["billRecordIndex"])
	return service.DeleteBillRecord(params["bookId"], params["billId"], billRecordIndex)
}
