package controller

import (
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/utils"
	"ququ.im/jizhangbao-api/nacos/data"
)

// AddBusinessCustomer	godoc
// @Summary		添加往来客户
// @Description	添加往来客户
// @Tags	往来客户
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	customerName formData string true "往来客户名称"
// @Success 200 {string} string	"ok"
// @Router	/customer/add [post]
func AddBusinessCustomer(params map[string]string) mgresult.AppResult {
	if !utils.Exists(params, "bookId") {
		return mgresult.AppError(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "customerName") {
		return mgresult.AppError(-1, "缺少往来客户名称参数")
	}
	return mgresult.NewAppResult(data.AddBusinessCustomer(params["bookId"], params["customerName"]))
}

// UpdateBusinessCustomer	godoc
// @Summary		修改往来客户
// @Description	修改往来客户
// @Tags	往来客户
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	customerId formData string true "往来客户编号"
// @Param	customerName formData string true "往来客户名称"
// @Success 200 {string} string	"ok"
// @Router	/customer/update [post]
func UpdateBusinessCustomer(params map[string]string) mgresult.AppResult {
	if !utils.Exists(params, "bookId") {
		return mgresult.AppError(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "customerId") {
		return mgresult.AppError(-1, "缺少往来客户编号参数")
	}
	if !utils.Exists(params, "customerName") {
		return mgresult.AppError(-1, "缺少往来客户名称参数")
	}
	return mgresult.NewAppResult(data.UpdateBusinessCustomer(params["bookId"], params["customerId"], params["customerName"]))
}

// DeleteBusinessCustomer	godoc
// @Summary		删除往来客户
// @Description	删除往来客户
// @Tags	往来客户
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	customerId formData string true "往来客户名称"
// @Success 200 {string} string	"ok"
// @Router	/customer/del [post]
func DeleteBusinessCustomer(params map[string]string) mgresult.AppResult {
	if !utils.Exists(params, "bookId") {
		return mgresult.AppError(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "customerId") {
		return mgresult.AppError(-1, "缺少往来客户编号参数")
	}
	return mgresult.NewAppResult(data.DeleteBusinessCustomer(params["bookId"], params["customerId"]))
}

// ListBusinessCustomer	godoc
// @Summary		列表/搜索往来客户
// @Description	列表/搜索往来客户
// @Tags	往来客户
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	customerName formData string false "往来客户名称关键字或简拼，不传为全部"
// @Success 200 {string} string	"ok"
// @Router	/customer/list [post]
func ListBusinessCustomer(params map[string]string) mgresult.AppResult {
	if !utils.Exists(params, "bookId") {
		return mgresult.AppError(-1, "缺少账本编号参数")
	}
	return mgresult.NewAppResult(data.ListBusinessCustomer(params["bookId"], params["customerName"]))
}

// GetBusinessCustomer	godoc
// @Summary		获取往来客户
// @Description	获取往来客户
// @Tags	往来客户
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	customerId formData string true "往来客户名称"
// @Success 200 {string} string	"ok"
// @Router	/customer/get [post]
func GetBusinessCustomer(params map[string]string) mgresult.AppResult {
	if !utils.Exists(params, "bookId") {
		return mgresult.AppError(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "customerId") {
		return mgresult.AppError(-1, "缺少往来客户编号参数")
	}
	return mgresult.NewAppResult(data.GetBusinessCustomer(params["bookId"], params["customerId"]))
}
