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
	"ququ.im/jizhangbao-api/controller"
	_ "ququ.im/jizhangbao-api/docs"
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

	var result mgresult.AppResult
	//添加所需的路由映射
	//资金账户管理
	engine.Any("/account/add", func(c *gin.Context) {
		result = controller.AddBillAccount(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/account/update", func(c *gin.Context) {
		result = controller.UpdateBillAccount(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/account/del", func(c *gin.Context) {
		result = controller.DeleteBillAccount(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/account/get", func(c *gin.Context) {
		result = controller.GetBillAccount(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/account/list", func(c *gin.Context) {
		result = controller.ListBillAccount(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})

	//账本管理
	engine.Any("/book/new", func(c *gin.Context) {
		result = controller.NewBillBook(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/book/update", func(c *gin.Context) {
		result = controller.UpdateBillBook(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/book/list", func(c *gin.Context) {
		result = controller.ListBillBooks(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/book/get", func(c *gin.Context) {
		result = controller.GetBillBook(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})

	//收支类型管理
	engine.Any("/type/add/public", func(c *gin.Context) {
		result = controller.AddPublicBillType(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/type/add", func(c *gin.Context) {
		result = controller.AddBillType(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/type/get", func(c *gin.Context) {
		result = controller.GetBillType(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/type/update", func(c *gin.Context) {
		result = controller.UpdateBillType(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/type/del", func(c *gin.Context) {
		result = controller.DeleteBillType(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/type/list", func(c *gin.Context) {
		result = controller.ListBillTypes(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})

	//往来客户管理
	engine.Any("/customer/get", func(c *gin.Context) {
		result = controller.GetBusinessCustomer(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/customer/add", func(c *gin.Context) {
		result = controller.AddBusinessCustomer(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/customer/update", func(c *gin.Context) {
		result = controller.UpdateBusinessCustomer(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/customer/del", func(c *gin.Context) {
		result = controller.DeleteBusinessCustomer(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/customer/list", func(c *gin.Context) {
		result = controller.ListBusinessCustomer(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})

	//收支账单管理
	engine.Any("/flow/add", func(c *gin.Context) {
		result = controller.AddBillFlow(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/flow/get", func(c *gin.Context) {
		result = controller.GetBillFlow(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/flow/update", func(c *gin.Context) {
		result = controller.UpdateBillFlow(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/flow/del", func(c *gin.Context) {
		result = controller.DeleteBillFlow(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/flow/list/month", func(c *gin.Context) {
		result = controller.GetMonthBillFlowPage(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})

	//分组统计-收支
	engine.Any("/bill/list/type/month", func(c *gin.Context) {
		result = controller.ListTypeBillSummaryByMonth(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/bill/list/customer/month", func(c *gin.Context) {
		result = controller.ListCustomerBillSummaryByMonth(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})

	//应收应付管理
	engine.Any("/accrual/add", func(c *gin.Context) {
		result = controller.AddBillAccrual(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/accrual/get", func(c *gin.Context) {
		result = controller.GetBillAccrual(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/accrual/update", func(c *gin.Context) {
		result = controller.UpdateBillAccrual(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/accrual/del", func(c *gin.Context) {
		result = controller.DeleteBillAccrual(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/accrual/record/add", func(c *gin.Context) {
		result = controller.AddBillRecord(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/accrual/record/update", func(c *gin.Context) {
		result = controller.UpdateBillRecord(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/accrual/record/del", func(c *gin.Context) {
		result = controller.DeleteBillRecord(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/accrual/list/month", func(c *gin.Context) {
		result = controller.GetMonthBillAccrualPage(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	//分组统计--应收应付
	engine.Any("/accrual/list/type", func(c *gin.Context) {
		result = controller.ListTypeAccrualSummary(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/accrual/list/customer", func(c *gin.Context) {
		result = controller.ListCustomerAccrualSummary(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})

	//操作记录管理
	engine.Any("/logs/list", func(c *gin.Context) {
		result = controller.ListOperateLog(utils.GinParamMap(c))
		c.JSON(http.StatusOK, result)
	})

	return engine
}

func recoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(http.StatusOK, *mgresult.Error(-1, "系统异常，请联系客服"))
}
