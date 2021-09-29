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
	"ququ.im/jizhangbao-data/controller"
	_ "ququ.im/jizhangbao-data/docs"
)

/**
统一路由映射入口
*/
func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	engine := gin.Default()

	//设置跟踪链路
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
	//资金账户管理
	engine.Any("/account/add", func(c *gin.Context) {
		res := controller.AddBillAccount(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/account/update", func(c *gin.Context) {
		res := controller.UpdateBillAccount(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/account/del", func(c *gin.Context) {
		res := controller.DeleteBillAccount(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/account/list", func(c *gin.Context) {
		res := controller.ListBillAccount(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})

	//账本管理
	engine.Any("/book/new", func(c *gin.Context) {
		res := controller.NewBillBook(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/book/update", func(c *gin.Context) {
		res := controller.UpdateBillBook(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/book/list", func(c *gin.Context) {
		res := controller.ListBillBooks(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})

	//收支类型管理
	engine.Any("/type/add/public", func(c *gin.Context) {
		res := controller.AddPublicBillType(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/type/add", func(c *gin.Context) {
		res := controller.AddBillType(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/type/get", func(c *gin.Context) {
		res := controller.GetBillType(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/type/update", func(c *gin.Context) {
		res := controller.UpdateBillType(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/type/del", func(c *gin.Context) {
		res := controller.DeleteBillType(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/type/list", func(c *gin.Context) {
		res := controller.ListBillTypes(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})

	//往来客户管理
	engine.Any("/customer/get", func(c *gin.Context) {
		res := controller.GetBusinessCustomer(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/customer/add", func(c *gin.Context) {
		res := controller.AddBusinessCustomer(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/customer/update", func(c *gin.Context) {
		res := controller.UpdateBusinessCustomer(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/customer/del", func(c *gin.Context) {
		res := controller.DeleteBusinessCustomer(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/customer/list", func(c *gin.Context) {
		res := controller.ListBusinessCustomer(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})

	//收支账单管理
	engine.Any("/flow/add", func(c *gin.Context) {
		res := controller.AddBillFlow(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/flow/get", func(c *gin.Context) {
		res := controller.GetBillFlow(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/flow/update", func(c *gin.Context) {
		res := controller.UpdateBillFlow(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/flow/del", func(c *gin.Context) {
		res := controller.DeleteBillFlow(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/flow/list", func(c *gin.Context) {
		res := controller.ListBillFlow(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})

	//应收应付管理
	engine.Any("/accrual/add", func(c *gin.Context) {
		res := controller.AddBillAccrual(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/accrual/get", func(c *gin.Context) {
		res := controller.GetBillAccrual(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/accrual/update", func(c *gin.Context) {
		res := controller.UpdateBillAccrual(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/accrual/del", func(c *gin.Context) {
		res := controller.DeleteBillAccrual(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/accrual/list", func(c *gin.Context) {
		res := controller.ListBillAccrual(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/accrual/record/add", func(c *gin.Context) {
		res := controller.AddBillRecord(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/accrual/record/update", func(c *gin.Context) {
		res := controller.UpdateBillRecord(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})
	engine.Any("/accrual/record/del", func(c *gin.Context) {
		res := controller.DeleteBillRecord(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})

	//操作记录管理
	engine.Any("/logs/list", func(c *gin.Context) {
		res := controller.ListOperateLog(utils.GinParamMap(c))
		c.JSON(http.StatusOK, res)
	})

	return engine
}

func recoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(http.StatusOK, *mgresult.Error(-1, "系统异常，请联系客服"))
}
