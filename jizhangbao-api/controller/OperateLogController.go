package controller

import (
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/utils"
	"ququ.im/jizhangbao-api/nacos/data"
	"time"
)

// ListOperateLog	godoc
// @Summary		操作日志列表
// @Description	操作日志列表
// @Tags	操作日志
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	bookId formData string true "账本编号"
// @Param	startDate formData string false "开始日期，格式为 yyyy-MM-dd，不传为今天"
// @Param	endDate formData string false "结束日期，格式为 yyyy-MM-dd，不传为今天"
// @Success 200 {string} string	"ok"
// @Router	/logs/list [post]
func ListOperateLog(params map[string]string) mgresult.AppResult {
	if !utils.Exists(params, "bookId") {
		return mgresult.AppError(-1, "缺少账本编号参数")
	}
	startDate := utils.DateFormat(time.Now(), "yyyy-MM-dd")
	if utils.Exists(params, "startDate") {
		startDate = params["startDate"]
	}
	endDate := utils.DateFormat(time.Now(), "yyyy-MM-dd")
	if utils.Exists(params, "endDate") {
		endDate = params["endDate"]
	}
	return mgresult.NewAppResult(data.ListOperateLog(params["bookId"], startDate, endDate))
}
