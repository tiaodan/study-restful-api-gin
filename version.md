# v0.0.0.0 
- 整合配置文件, from github项目: study-config-viper
- 整合日志文件, from github项目: study-log-go-original
- 整合错误处理, lian, from github项目: study-error-go-original
- 整合数据库  , from github项目: study-db-gorm

# v0.0.0.1
- 实现查询接口, 访问: http://127.0.0.1:8888/orders

# v0.0.0.1 2025.04.16
- 修改model, 增加字段-DropShippingOrderTime 代发订单时间

# v0.0.0.1 2025.04.16-v2
- model, 新增gorm相关配置,如Id，主键自增

# v0.0.0.2
- 订单号增、改、查已完成,把db方法都改成OrderAdd 这种Order动作形式

# v0.0.0.3
- 配置文件加上数据库配置

# v0.0.0.4
- 加上删除api

# v0.0.0.4 2025.04.16
待处理：
- 加上删除api，删除不成功，需调试
- logger打的日志，文件位置不对，都是logger.go
- 整合gin后，日志打印少，自定义W的logger + fmt.Println都不打印了

# v0.0.0.5
- 实现分页查询

# v0.0.0.6
- 对分页查询请求, 做参数校验。现在r.GET("/orders", order.OrdersPageQuery)检测范围太宽了。在代码逻辑修改

# v0.0.0.7
- 分页查询, 倒序排列
