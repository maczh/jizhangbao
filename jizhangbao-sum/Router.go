package main

import (
	"github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	"github.com/maczh/gintool"
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/mgtrace"
	"github.com/maczh/utils"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"ququ.im/jizhangbao-sum/controller"
	_ "ququ.im/jizhangbao-sum/docs"
)

/**
统一路由映射入口
*/
func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	engine := gin.Default()

	//添加跟踪日志
	engine.Use(mgtrace.TraceId())

	//设置接口日志
	engine.Use(gintool.SetRequestLogger())
	//添加跨域处理
	engine.Use(gintool.Cors())

	//添加swagger支持
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//engine.GET("/docs/*any", swagger_skin.HandleReDoc)

	//处理全局异常
	engine.Use(nice.Recovery(recoveryHandler))

	//设置404返回的内容
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, *mgresult.Error(-1, "请求的方法不存在"))
	})

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, *mgresult.Success(nil))
	})

	//添加所需的路由映射
	//实时统计
	engine.Any("/bill/incr", func(c *gin.Context) {
		result := controller.UpdateBillSummary(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/accrual/incr", func(c *gin.Context) {
		result := controller.UpdateAccrualSummary(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	//统计查询
	engine.Any("/bill/get/day", func(c *gin.Context) {
		result := controller.GetBillDaySummary(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/bill/get/month", func(c *gin.Context) {
		result := controller.GetBillMonthSummary(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/bill/get/year", func(c *gin.Context) {
		result := controller.GetBillYearSummary(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/bill/get/week", func(c *gin.Context) {
		result := controller.GetBillWeekSummary(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/accrual/get", func(c *gin.Context) {
		result := controller.GetAccrualSummary(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})

	//分组统计--账单
	engine.Any("/bill/list/type/day", func(c *gin.Context) {
		result := controller.ListTypeBillSummaryByDay(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/bill/list/type/month", func(c *gin.Context) {
		result := controller.ListTypeBillSummaryByMonth(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/bill/list/type/year", func(c *gin.Context) {
		result := controller.ListTypeBillSummaryByYear(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/bill/list/type/week", func(c *gin.Context) {
		result := controller.ListTypeBillSummaryByWeek(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/bill/list/customer/day", func(c *gin.Context) {
		result := controller.ListCustomerBillSummaryByDay(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/bill/list/customer/month", func(c *gin.Context) {
		result := controller.ListCustomerBillSummaryByMonth(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/bill/list/customer/year", func(c *gin.Context) {
		result := controller.ListCustomerBillSummaryByYear(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/bill/list/customer/week", func(c *gin.Context) {
		result := controller.ListCustomerBillSummaryByWeek(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	//分组统计--应收应付
	engine.Any("/accrual/list/type", func(c *gin.Context) {
		result := controller.ListTypeAccrualSummary(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/accrual/list/customer", func(c *gin.Context) {
		result := controller.ListCustomerAccrualSummary(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})

	return engine
}

func recoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(http.StatusOK, *mgresult.Error(-1, "系统异常，请联系客服"))
}
