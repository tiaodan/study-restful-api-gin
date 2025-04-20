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
/*
参数：
	dbType string 数据库类型 如 mysql、sqlite3、postgres 等
	dbName string 数据库名
	dbUser string 数据库用户名
	dbPass string 数据库密码
*/
func InitDB(dbType, dbName, dbUser, dbPass string) {
	once.Do(func() { // 使用 sync.Once 确保只执行一次
		// dsn := "root:password@tcp(127.0.0.1:3306)/pdd_order?charset=utf8mb4&parseTime=True&loc=Local"
		dsn := dbUser + ":" + dbPass + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
		var err error

		var dbOpen gorm.Dialector // 用什么数据库打开
		if dbType == "mysql" {
			dbOpen = mysql.Open(dsn)
		}
		// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		DB, err = gorm.Open(dbOpen, &gorm.Config{})
		if err != nil {
			log.Println("单例: 数据库连接失败:", err)
			panic(err)
		}
		log.Println("单例: 数据库连接成功")
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
