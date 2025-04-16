package db

import (
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once // 使用 sync.Once 确保单例

// 初始化数据库连接
func InitDB() {
	once.Do(func() { // 使用 sync.Once 确保只执行一次
		dsn := "root:password@tcp(127.0.0.1:3306)/pdd_order?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println("数据库连接失败:", err)
			panic(err)
		}
		log.Println("数据库连接成功")
	})
}

// 插入默认数据
/*
思路:
1. 插入默认数据-1
2. 插入默认数据-2
*/
func InsertDefaultData() {
	// 插入默认数据-website
	// websiteDefaultNoClass := &models.Website{Name: "待分类", NameId: 0, URL: "未知"}
	// websiteDefaultJ88d := &models.Website{Name: "j88d", NameId: 1, URL: "www.j88d.com"}         // 请求url 时带上http://
	// websiteDefaultXimalaya := &models.Website{Name: "喜马拉雅", NameId: 2, URL: "www.ximalaya.com"} // 请求url 时带上http://
	// defaultWebsites := []*models.Website{websiteDefaultNoClass, websiteDefaultJ88d, websiteDefaultXimalaya}
	// BatchAddWebsite(defaultWebsites)

}
