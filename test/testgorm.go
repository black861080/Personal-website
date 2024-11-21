package main

import (
	"black/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Friend{})
	db.AutoMigrate(&models.Visitor{})
	db.AutoMigrate(&models.Photo{})
	db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.Project{})
	db.AutoMigrate(&models.Information{})
	db.AutoMigrate(&models.Skill{})
	db.AutoMigrate(&models.Experience{})
	db.AutoMigrate(&models.Blog{})
	friend := &models.Friend{
		Friend_name:  "李锦升",
		LoginTime:    time.Now(),
		LoginOutTime: time.Now(),
		Upload_time:  time.Now(),
	}
	db.Create(friend)

	// Read

	//fmt.Println(db.First(friend, 1)) // 根据整型主键查找
	//db.First(friend, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(friend).Update("times", 1)
	// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&product, 1)
}
