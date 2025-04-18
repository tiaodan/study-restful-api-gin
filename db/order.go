// db order 相关操作
package db

import (
	"log"
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
		log.Println("创建失败:", result.Error)
		return result.Error
	} else {
		log.Println("创建成功:", order)
	}
	return nil
}

// 批量增
func OrderBatchAdd(orders []*models.Order) {
	for i, order := range orders {
		err := OrderAdd(order)
		if err == nil {
			log.Printf("批量创建第%d条成功, order: %v", i+1, &order)
		} else {
			log.Printf("批量创建第%d条失败, err: %v", i+1, err)
		}
	}
}

// 删
func OrderDelete(id uint) {
	var order models.Order
	result := DB.Delete(&order, id)
	if result.Error != nil {
		log.Println("删除失败:", result.Error)
	} else {
		log.Println("删除成功:", id)
	}
}

// 批量删
func OrdersBatchDelete(ids []uint) {
	var orders []models.Order
	result := DB.Delete(&orders, ids)
	if result.Error != nil {
		log.Println("批量删除失败:", result.Error)
	} else {
		log.Println("批量删除成功:", ids)
	}
}

// 改 - 参数用map
// func UpdateOrder(orderId uint, updates map[string]interface{}) {
// 	var order models.Order
// 	result := DB.Model(&order).Where("pdd_order_id = ?", orderId).Updates(updates)
// 	if result.Error != nil {
// 		log.Println("修改失败:", result.Error)
// 	} else {
// 		log.Println("修改成功:", orderId)
// 	}
// }

// 改 - 参数用结构体
func OrderUpdate(orderId string, order *models.Order) error {
	result := DB.Model(&order).Where("pdd_order_id = ?", orderId).Updates(order)
	if result.Error != nil {
		log.Println("修改失败:", result.Error)
		return result.Error
	} else {
		log.Println("修改成功:", orderId)
	}

	return nil
}

// 批量改
func OrdersBatchUpdate(updates map[uint]map[string]interface{}) {
	for orderId, update := range updates {
		var order models.Order
		result := DB.Model(&order).Where("order_id = ?", orderId).Updates(update)
		if result.Error != nil {
			log.Printf("更新订单 %d 失败: %v\n", orderId, result.Error)
		} else {
			log.Printf("更新订单 %d 成功\n", orderId)
		}
	}
}

// 查
func OrderQueryById(id uint) *models.Order {
	var order models.Order
	result := DB.First(&order, id)
	if result.Error != nil {
		log.Println("查询失败:", result.Error)
		return nil
	}
	log.Println("查询成功:", order)
	return &order
}

// 批量查
func OrdersBatchQuery(ids []uint) ([]*models.Order, error) {
	var orders []*models.Order
	result := DB.Find(&orders, ids)
	if result.Error != nil {
		log.Printf("批量查询失败: %v\n", result.Error)
		return orders, result.Error
	}
	log.Printf("批量查询成功, 查询到 %d 条记录\n", len(orders))
	return orders, nil
}

// 查所有
func OrdersQueryAll() ([]*models.Order, error) {
	var orders []*models.Order
	result := DB.Find(&orders)
	if result.Error != nil {
		log.Printf("批量查询失败: %v\n", result.Error)
		return orders, result.Error
	}
	log.Printf("批量查询成功, 查询到 %d 条记录\n", len(orders))
	return orders, nil
}
