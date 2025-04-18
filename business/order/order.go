// 功能: 封装restfult api - order模块
package order

import (
	"study-restful-api-gin/db"
	"study-restful-api-gin/logger"
	"study-restful-api-gin/models"

	"github.com/gin-gonic/gin"
)

// 增
func OrderAdd(c *gin.Context) {
	logger.Debug("增加订单")
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		logger.Error("解析请求体失败, err: %v", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return // 必须保留 return，确保绑定失败时提前退出
	}
	err := db.OrderAdd(&order)
	if err != nil {
		logger.Error("增加订单失败, err: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "添加成功")
}

// 删
// 改
func OrderUpdate(c *gin.Context) {
	logger.Debug("修改订单")
	// 绑定前端数据
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		logger.Error("解析请求体失败, err: %v", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return // 必须保留 return，确保绑定失败时提前退出
	}
	err := db.OrderUpdate(order.PddOrderId, &order)

	if err != nil {
		logger.Error("修改订单失败, err: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "修改成功")
}

// 查
func OrdersQuery(c *gin.Context) {
	logger.Debug("查询所有订单")
	// 假的数据
	// var order = []models.Order{
	// 	{PddOrderId: "1111", PddOrderTime: "2222"},
	// }
	orders, _ := db.OrdersQueryAll()
	c.JSON(200, orders)
}
