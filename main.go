package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo 用户信息
type User struct {
	ID       uint `gorm:"type:int(10);column:id"`
	Username string
	Password string
	Age      int
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//通过db.SingularTable(true)，gorm会在创建表的时候去掉"s"的后缀
	db.SingularTable(true)

	// 自动迁移
	db.AutoMigrate(&User{})

	u1 := User{Username: fmt.Sprintf("u%d", time.Now().Unix()), Password: "123456", Age: 12}

	// 创建记录
	db.Create(&u1)

	//读取数据
	var u2 User
	db.First(&u2)
	fmt.Println(u2)
}
