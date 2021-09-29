package controller

import (
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/utils"
	"ququ.im/jizhangbao-api/nacos/data"
)

// AddPublicBillType	godoc
// @Summary		添加公共账单类型
// @Description	添加公共账单类型
// @Tags	账单类型
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	billType formData string true "账单类型名称"
// @Param	billTypeIconId formData string true "账单类型图标编号"
// @Param	operateType formData string true "操作类型：收入/支出/应收/应付"
// @Success 200 {string} string	"ok"
// @Router	/type/add/public [post]
func AddPublicBillType(params map[string]string) mgresult.AppResult {
	if !utils.Exists(params, "billType") {
		return mgresult.AppError(-1, "缺少账单类型名称参数")
	}
	if !utils.Exists(params, "billTypeIconId") {
		return mgresult.AppError(-1, "缺少用户编号参数")
	}
	if !utils.Exists(params, "operateType") {
		return mgresult.AppError(-1, "缺少操作类型参数：收入/支出/应收/应付")
	}
	return mgresult.NewAppResult(data.AddPublicBillType(params["billType"], params["billTypeIconId"], params["operateType"]))
}

// AddBillType	godoc
// @Summary		添加自定义账单类型
// @Description	添加自定义账单类型
// @Tags	账单类型
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	billType formData string true "账单类型名称"
// @Param	billTypeIconId formData string true "账单类型图标编号"
// @Param	operateType formData string true "操作类型：收入/支出/应收/应付"
// @Success 200 {string} string	"ok"
// @Router	/type/add [post]
func AddBillType(params map[string]string) mgresult.AppResult {
	if !utils.Exists(params, "bookId") {
		return mgresult.AppError(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "billType") {
		return mgresult.AppError(-1, "缺少账单类型名称参数")
	}
	if !utils.Exists(params, "billTypeIconId") {
		return mgresult.AppError(-1, "缺少用户编号参数")
	}
	if !utils.Exists(params, "operateType") {
		return mgresult.AppError(-1, "缺少操作类型参数：收入/支出/应收/应付")
	}
	return mgresult.NewAppResult(data.AddBillType(params["bookId"], params["billType"], params["billTypeIconId"], params["operateType"]))
}

// UpdateBillType	godoc
// @Summary		修改自定义账单类型
// @Description	修改自定义账单类型
// @Tags	账单类型
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	billTypeId formData string true "账单类型编号"
// @Param	billType formData string true "账单类型名称"
// @Param	billTypeIconId formData string true "账单类型图标编号"
// @Success 200 {string} string	"ok"
// @Router	/type/update [post]
func UpdateBillType(params map[string]string) mgresult.AppResult {
	if !utils.Exists(params, "billType") {
		return mgresult.AppError(-1, "缺少账单类型名称参数")
	}
	if !utils.Exists(params, "billTypeIconId") {
		return mgresult.AppError(-1, "缺少用户编号参数")
	}
	if !utils.Exists(params, "billTypeId") {
		return mgresult.AppError(-1, "缺少账单类型编号参数")
	}
	if !utils.Exists(params, "bookId") {
		return mgresult.AppError(-1, "缺少账本编号参数")
	}
	return mgresult.NewAppResult(data.UpdateBillType(params["bookId"], params["billTypeId"], params["billType"], params["billTypeIconId"]))
}

// DeleteBillType	godoc
// @Summary		删除自定义账单类型
// @Description	删除自定义账单类型
// @Tags	账单类型
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	billTypeId formData string true "账单类型编号"
// @Success 200 {string} string	"ok"
// @Router	/type/del [post]
func DeleteBillType(params map[string]string) mgresult.AppResult {
	if !utils.Exists(params, "billTypeId") {
		return mgresult.AppError(-1, "缺少账单类型编号参数")
	}
	if !utils.Exists(params, "bookId") {
		return mgresult.AppError(-1, "缺少账本编号参数")
	}
	return mgresult.NewAppResult(data.DeleteBillType(params["bookId"], params["billTypeId"]))
}

// ListBillTypes	godoc
// @Summary		列出账本的账单类型
// @Description	列出账本的账单类型
// @Tags	账单类型
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	operateType formData string true "操作类型：收入/支出/应收/应付"
// @Success 200 {string} string	"ok"
// @Router	/type/list [post]
func ListBillTypes(params map[string]string) mgresult.AppResult {
	if !utils.Exists(params, "bookId") {
		return mgresult.AppError(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "operateType") {
		return mgresult.AppError(-1, "缺少操作类型参数：收入/支出/应收/应付")
	}
	return mgresult.NewAppResult(data.ListBillTypes(params["bookId"], params["operateType"]))
}

// GetBillType	godoc
// @Summary		获取账单类型
// @Description	获取账单类型
// @Tags	账单类型
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	billTypeId formData string true "账单类型编号"
// @Success 200 {string} string	"ok"
// @Router	/type/get [post]
func GetBillType(params map[string]string) mgresult.AppResult {
	if !utils.Exists(params, "billTypeId") {
		return mgresult.AppError(-1, "缺少账单类型编号参数")
	}
	return mgresult.NewAppResult(data.GetBillType(params["billTypeId"]))
}
