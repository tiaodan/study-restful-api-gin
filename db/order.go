// db order 相关操作
package db

import (
	"study-restful-api-gin/logger"
	"study-restful-api-gin/models"

	// 导入 clause 包
	"gorm.io/gorm/clause"
)

// 增
func OrderAdd(order *models.Order) error {
	result := DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "OrderId"}, {Name: "DropShippingOrderId"}}, // 判断唯一索引: pddId + 代发id
		DoUpdates: clause.Assignments(map[string]interface{}{
			"pdd_order_time":               order.PddOrderTime,
			"pdd_order_price":              order.PddOrderPrice,
			"pdd_product_type":             order.PddProductType,
			"pdd_product_color":            order.PddProductColor,
			"pdd_order_status":             order.PddOrderStatus,
			"pdd_buyer_info":               order.PddBuyerInfo,
			"pdd_express_company":          order.PddExpressCompany,
			"pdd_express_id":               order.PddExpressId,
			"pdd_is_black_list":            order.PddIsBlackList,
			"pdd_remark":                   order.PddRemark,
			"drop_shipping_platform":       order.DropShippingPlatform,
			"drop_shipping_order_time":     order.DropShippingOrderTime,
			"drop_shipping_factory_name":   order.DropShippingFactoryName,
			"drop_shipping_real_price":     order.DropShippingRealPrice,
			"drop_shipping_price":          order.DropShippingPrice,
			"drop_shipping_discount_price": order.DropShippingDiscountPrice,
			"drop_shipping_remark":         order.DropShippingRemark,
		}),
	}).Create(order)
	if result.Error != nil {
		logger.Error("创建失败: %v", result.Error)
		return result.Error
	} else {
		logger.Debug("创建成功: %v", order)
	}
	return nil
}

// 批量增
func OrderBatchAdd(orders []*models.Order) {
	for i, order := range orders {
		err := OrderAdd(order)
		if err == nil {
			logger.Error("批量创建第%d条成功, order: %v", i+1, &order)
		} else {
			logger.Debug("批量创建第%d条失败, err: %v", i+1, err)
		}
	}
}

// 删
func OrderDelete(id uint) error {
	logger.Debug("删除订单, 参数id= %d", id)
	var order models.Order
	result := DB.Delete(&order, id)
	if result.Error != nil {
		logger.Error("删除失败: %v", result.Error)
		return result.Error
	} else {
		logger.Debug("删除成功: %d", id)
	}
	return nil
}

// 批量删
func OrdersBatchDelete(ids []uint) {
	var orders []models.Order
	result := DB.Delete(&orders, ids)
	if result.Error != nil {
		logger.Error("批量删除失败: %v", result.Error)
	} else {
		logger.Debug("批量删除成功: %v", ids)
	}
}

// 改 - 参数用map
// func UpdateOrder(orderId uint, updates map[string]interface{}) {
// 	var order models.Order
// 	result := DB.Model(&order).Where("pdd_order_id = ?", orderId).Updates(updates)
// 	if result.Error != nil {
// 		logger.Error("修改失败:", result.Error)
// 	} else {
// 		logger.Debug("修改成功:", orderId)
// 	}
// }

// 改 - 参数用结构体
func OrderUpdate(orderId string, order *models.Order) error {
	result := DB.Model(&order).Where("pdd_order_id = ?", orderId).Updates(order)
	if result.Error != nil {
		logger.Error("修改失败: %v", result.Error)
		return result.Error
	} else {
		logger.Debug("修改成功: %v", orderId)
	}

	return nil
}

// 批量改
func OrdersBatchUpdate(updates map[uint]map[string]interface{}) {
	for orderId, update := range updates {
		var order models.Order
		result := DB.Model(&order).Where("order_id = ?", orderId).Updates(update)
		if result.Error != nil {
			logger.Error("更新订单 %d 失败: %v", orderId, result.Error)
		} else {
			logger.Debug("更新订单 %d 成功", orderId)
		}
	}
}

// 查
func OrderQueryById(id uint) *models.Order {
	var order models.Order
	result := DB.First(&order, id)
	if result.Error != nil {
		logger.Error("查询失败: %v", result.Error)
		return nil
	}
	logger.Debug("查询成功: %v", order)
	return &order
}

// 批量查
func OrdersBatchQuery(ids []uint) ([]*models.Order, error) {
	var orders []*models.Order
	result := DB.Find(&orders, ids)
	if result.Error != nil {
		logger.Error("批量查询失败: %v", result.Error)
		return orders, result.Error
	}
	logger.Debug("批量查询成功, 查询到 %d 条记录", len(orders))
	return orders, nil
}

// 查所有
func OrdersQueryAll() ([]*models.Order, error) {
	var orders []*models.Order
	result := DB.Find(&orders)
	if result.Error != nil {
		logger.Error("批量查询失败: %v", result.Error)
		return orders, result.Error
	}
	logger.Debug("批量查询成功, 查询到 %d 条记录", len(orders))
	return orders, nil
}

// 查数据总数
func OrdersTotal() (int64, error) {
	var count int64
	result := DB.Model(&models.Order{}).Count(&count)
	if result.Error != nil {
		logger.Error("查询数据总数失败: %v", result.Error)
		return 0, result.Error
	}
	logger.Info("查询数据总数成功, 总数为 %d", count)
	return count, nil
}

// 分页查询
func OrdersPageQuery(pageNum, pageSize int) ([]*models.Order, error) {
	var orders []*models.Order
	result := DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&orders)
	if result.Error != nil {
		logger.Error("分页查询失败: %v", result.Error)
		return orders, result.Error
	}
	logger.Debug("分页查询成功, 查询到 %d 条记录", len(orders))
	return orders, result.Error
}
