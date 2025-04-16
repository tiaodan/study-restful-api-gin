// 拼多多订单数据模型, 存数据用的
package models

// 订单数据
type Order struct {
	// 拼多多相关信息
	PddOrderId        string  `json:"pddOrderId"`        // PDD订单号
	PddOrderTime      string  `json:"pddOrderTime"`      // 购买时间
	PddOrderPrice     float64 `json:"pddOrderPrice"`     // 购买价格
	PddProductType    string  `json:"pddProductType"`    // 产品类型
	PddProductColor   string  `json:"pddProductColor"`   // 颜色
	PddOrderStatus    string  `json:"pddOrderStatus"`    // 订单状态
	PddBuyerInfo      string  `json:"pddBuyerInfo"`      // 买家信息
	PddExpressCompany string  `json:"pddExpressCompany"` // 快递公司
	PddExpressId      string  `json:"pddExpressId"`      // 物流编号
	PddIsBlackList    bool    `json:"pddIsBlackList"`    // 买家拉黑
	PddRemark         string  `json:"pddRemark"`         // pdd备注

	// 代发平台相关信息
	DropShippingPlatform      string  `json:"dropShippingPlatform"`      // 代发平台
	DropShippingOrderId       string  `json:"dropShippingOrderId"`       // 代发订单号
	DropShippingFactoryName   string  `json:"dropShippingFactoryName"`   // 代发厂家名
	DropShippingRealPrice     float64 `json:"dropShippingRealPrice"`     // 代发实际价
	DropShippingPrice         float64 `json:"dropShippingPrice"`         // 购买价格
	DropShippingDiscountPrice float64 `json:"dropShippingDiscountPrice"` // 优惠
	DropShippingRemark        string  `json:"dropShippingRemark"`        // 代发备注
}
