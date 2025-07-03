package sqliteorm

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Age      int    `gorm:"not null"`
}

func CRUD() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database test.db")
	}

	// 迁移 schema
	db.AutoMigrate(&User{})

	// 清空表数据, 方便多次执行测试
	db.Exec("delete from users where 1")

	// 创建
	user := User{Username: "johndoe", Password: "password123", Email: "abc@qq.com", Age: 30}
	result := db.Create(&user) // 通过数据的指针来创建

	// 打印 user 结构, 包括字段名称
	fmt.Printf("Created User: %+v\n", result)
	fmt.Printf("Created User: %+v\n", user)

	// 查询
	db.First(&user, user.ID) // 查询刚刚插入的用户
	fmt.Println("Select by ID:", user)

	db.First(&user, "username = ?", "johndoe") // 查询用户名为 "johndoe" 的用户
	fmt.Println("Select by username:", user)

	// 更新
	db.Model(&user).Update("Password", "newpassword123")                                                          // 更新单个字段
	db.Model(&user).Updates(User{Username: "john_doe", Password: "newpassword456", Email: "123@qq.com", Age: 31}) // 更新多个字段
	fmt.Println("Updated Password:", user.Password)
	fmt.Println("After password updated:", user)

	// 删除
	db.Delete(&user, user.ID) // 删除 ID 为 1 的用户
	fmt.Println("Deleted User:", user)
}
