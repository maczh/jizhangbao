package controller

import (
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/utils"
	"ququ.im/jizhangbao-api/nacos/data"
)

// AddBillAccount	godoc
// @Summary		在账本中添加自定义资金账户
// @Description	在账本中添加自定义资金账户
// @Tags	资金账户
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	billAccountName formData string true "资金账户名，如 建行卡-4068"
// @Success 200 {string} string	"ok"
// @Router	/account/add [post]
func AddBillAccount(params map[string]string) mgresult.AppResult {
	if !utils.Exists(params, "bookId") {
		return mgresult.AppError(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "billAccountName") {
		return mgresult.AppError(-1, "缺少资金账户名称参数")
	}
	return mgresult.NewAppResult(data.AddBillAccount(params["bookId"], params["billAccountName"]))
}

// UpdateBillAccount	godoc
// @Summary		在账本中修改自定义资金账户
// @Description	在账本中修改自定义资金账户
// @Tags	资金账户
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	billAccountId formData string true "资金账户编号"
// @Param	billAccountName formData string true "资金账户名，如 建行卡-4068"
// @Success 200 {string} string	"ok"
// @Router	/account/update [post]
func UpdateBillAccount(params map[string]string) mgresult.AppResult {
	if !utils.Exists(params, "bookId") {
		return mgresult.AppError(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "billAccountId") {
		return mgresult.AppError(-1, "缺少资金账户编号参数")
	}
	if !utils.Exists(params, "billAccountName") {
		return mgresult.AppError(-1, "缺少资金账户名称参数")
	}
	return mgresult.NewAppResult(data.UpdateBillAccount(params["bookId"], params["billAccountId"], params["billAccountName"]))
}

// DeleteBillAccount	godoc
// @Summary		在账本中删除自定义资金账户
// @Description	在账本中删除自定义资金账户
// @Tags	资金账户
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	billAccountId formData string true "资金账户编号"
// @Success 200 {string} string	"ok"
// @Router	/account/del [post]
func DeleteBillAccount(params map[string]string) mgresult.AppResult {
	if !utils.Exists(params, "bookId") {
		return mgresult.AppError(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "billAccountId") {
		return mgresult.AppError(-1, "缺少资金账户编号参数")
	}
	return mgresult.NewAppResult(data.DeleteBillAccount(params["bookId"], params["billAccountId"]))
}

// ListBillAccount	godoc
// @Summary		列出在账本中所有资金账户
// @Description	列出在账本中所有资金账户
// @Tags	资金账户
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Success 200 {string} string	"ok"
// @Router	/account/list [post]
func ListBillAccount(params map[string]string) mgresult.AppResult {
	if !utils.Exists(params, "bookId") {
		return mgresult.AppError(-1, "缺少账本编号参数")
	}
	return mgresult.NewAppResult(data.ListBillAccount(params["bookId"]))
}

// GetBillAccount	godoc
// @Summary		在账本中获取资金账户
// @Description	在账本中获取资金账户
// @Tags	资金账户
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	billAccountId formData string true "资金账户编号"
// @Success 200 {string} string	"ok"
// @Router	/account/del [post]
func GetBillAccount(params map[string]string) mgresult.AppResult {
	if !utils.Exists(params, "bookId") {
		return mgresult.AppError(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "billAccountId") {
		return mgresult.AppError(-1, "缺少资金账户编号参数")
	}
	return mgresult.NewAppResult(data.GetBillAccount(params["bookId"], params["billAccountId"]))
}
