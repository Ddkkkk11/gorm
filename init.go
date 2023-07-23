package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func init() {
	username := "root"   // 账号
	password := "123456" // 密码
	host := "127.0.0.1"  // 数据库地址, 也可以是IP或者域名
	port := 3306
	Dbname := "gorm" // 数据库名称
	timeout := "10s"
	//
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	// 连接MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "f_", // 表名前缀
			SingularTable: true, // 是否单数表名
			NoLowerCase:   true, // 该大写就大写 小写就小写
		},
	})
	if err != nil {
		panic("connect failed ......, error=" + err.Error())
	}
	// 连接成功
	//fmt.Println("db", db)
	DB = db
}

type Roll struct {
	ID   uint
	Name string
	Age  int
}

func main() {
	fmt.Println("Db", DB)
	DB.AutoMigrate(&Roll{})

}
