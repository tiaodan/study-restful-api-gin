// 功能: 封装restfult api - order模块
package order

import (
	"study-restful-api-gin/db"
	"study-restful-api-gin/logger"

	"github.com/gin-gonic/gin"
)

// 增
// 删
// 改
// 查
func QueryOrders(c *gin.Context) {
	logger.Debug("查询所有订单")
	// 假的数据
	// var order = []models.Order{
	// 	{PddOrderId: "1111", PddOrderTime: "2222"},
	// }
	orders, _ := db.QueryAllOrders()
	c.JSON(200, orders)
}
