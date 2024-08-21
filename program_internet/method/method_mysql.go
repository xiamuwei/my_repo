package method 

import (
	"fmt"

	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
    "log"
)

var db *gorm.DB
func init(){
	dsn := "root:123456@tcp(127.0.0.1:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database: ", err)
    } 
}

func InsertData(user *User) error{
	fmt.Println("insert database user = ", *user)
    
	// 自动迁移模式，这将创建表、缺失的外键、约束等
    db.AutoMigrate(&User{})

    // 插入数据
    result := db.Create(user)

    if result.Error != nil {
        log.Fatal("failed to insert user: ", result.Error)
		return result.Error
    }

    log.Printf("User created successfully", user)
	return nil
}

func UpdateData(user *User) error{

    return nil
}

func GetData() error{
    return nil
}

func DeleteData() error{
    return nil
}