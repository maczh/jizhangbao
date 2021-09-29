package controller

import (
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/utils"
	"ququ.im/jizhangbao-data/service"
)

// NewBillBook	godoc
// @Summary		添加账本
// @Description	添加账本
// @Tags	账本
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户编号"
// @Param	userId formData string true "用户编号"
// @Param	bookName formData string false "账本名称"
// @Success 200 {string} string	"ok"
// @Router	/book/new [post]
func NewBillBook(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "shopId") {
		return *mgresult.Error(-1, "缺少商户编号参数")
	}
	if !utils.Exists(params, "userId") {
		return *mgresult.Error(-1, "缺少用户编号参数")
	}
	if !utils.Exists(params, "bookName") {
		params["bookName"] = "默认账本"
	}
	return service.NewBillBook(params["shopId"], params["userId"], params["bookName"])
}

// UpdateBillBook	godoc
// @Summary		账本改名
// @Description	账本改名
// @Tags	账本
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	bookName formData string true "账本名称"
// @Success 200 {string} string	"ok"
// @Router	/book/update [post]
func UpdateBillBook(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "bookName") {
		return *mgresult.Error(-1, "缺少账本名称参数")
	}
	return service.UpdateBillBook(params["bookId"], params["bookName"])
}

// ListBillBooks	godoc
// @Summary		列出用户所有账本
// @Description	列出用户所有账本
// @Tags	账本
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户编号"
// @Param	userId formData string true "用户编号"
// @Success 200 {string} string	"ok"
// @Router	/book/list [post]
func ListBillBooks(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "shopId") {
		return *mgresult.Error(-1, "缺少商户编号参数")
	}
	if !utils.Exists(params, "userId") {
		return *mgresult.Error(-1, "缺少用户编号参数")
	}
	return service.ListBillBooks(params["shopId"], params["userId"])
}

// GetBillBook	godoc
// @Summary		查询获取账本
// @Description	查询获取账本
// @Tags	账本
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Success 200 {string} string	"ok"
// @Router	/book/get [post]
func GetBillBook(params map[string]string) mgresult.Result {
	if !utils.Exists(params, "bookId") {
		return *mgresult.Error(-1, "缺少账本编号参数")
	}
	if !utils.Exists(params, "bookName") {
		return *mgresult.Error(-1, "缺少账本名称参数")
	}
	return service.GetBillBook(params["bookId"])
}
