// 功能: 封装restfult api - order模块
package order

import (
	"strconv"
	"study-restful-api-gin/db"
	"study-restful-api-gin/errorutil"
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
func OrderDelete(c *gin.Context) {
	// 提取前端传递的 id 参数
	idStr := c.Param("id")
	logger.Debug("删除订单, 参数= %v", idStr)
	id, err := strconv.ParseUint(idStr, 10, 64) // 转换为 ​十进制 64 位无符号整数
	if err != nil {
		logger.Error("删除订单, 参数错误")
		c.JSON(400, gin.H{"error": "删除订单, 参数错误"})
		return
	}

	// 调用数据库删除方法
	err = db.OrderDelete(uint(id))
	// err := db.OrderDelete(1)
	if err != nil {
		logger.Error("删除订单失败, err: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "删除成功")
}

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
/*
返回: json对象
{
	"total": 0,
	"data": []
}
*/
func OrdersQuery(c *gin.Context) {
	logger.Debug("查询所有订单")
	total, err := db.OrdersTotal() // 补充总数获取
	errorutil.ErrorPrint(err, "查询订单总数失败")
	orders, _ := db.OrdersQueryAll()

	c.JSON(200, gin.H{
		"total": total,
		"data":  orders,
	})
}

// 查-分页
/*
返回: json对象
{
	"total": 0,
	"data": []
}
*/
func OrdersPageQuery(c *gin.Context) {
	logger.Debug("分页查询订单")
	logger.Debug("c.DefaultQuery= %v", c.DefaultQuery)
	pageStr := c.DefaultQuery("page", "1")  // 默认为 1
	sizeStr := c.DefaultQuery("size", "10") // 默认为 10 ,所以不存在类型不是string类型
	logger.Debug("前端传参, page=%v, size=%v", pageStr, sizeStr)

	total, err := db.OrdersTotal() // 总数
	errorutil.ErrorPrint(err, "查询订单总数失败")

	page, _ := strconv.Atoi(pageStr) // 因为默认都是数字str了，所以不存在报错情况
	size, _ := strconv.Atoi(sizeStr) // 因为默认都是数字str了，所以不存在报错情况
	orders, _ := db.OrdersPageQuery(page, size)

	// 构造指定的返回结构
	c.JSON(200, gin.H{
		"total": total,
		"data":  orders,
	})
}
