package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	db, err := gorm.Open("mysql", "root:qujian123@tcp(47.98.179.41:13360)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//fmt.Println("db: ", db)
	defer db.Close()

	db.AutoMigrate(User{})

	// 创建
	/*
	u1 := User{Name: "chiling", Age: 48}
	u2 := User{Name: "qujian", Age: 34}
	db.Create(&u1)
	db.Create(&u2)
	 */

	// 查询
	// 把结果放到user里面
	//var user User // 声明模型结构体变量user
	//fmt.Println(user)   // {{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC <nil>}  0}

	// var user User
	var user = User{} // 声明模型结构体变量user
	fmt.Println(user)   // {{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC <nil>}  0}

	// 根据主键去查找的
	db.Debug().First(&user)  // 传值和传地址的区别   SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL ORDER BY `users`.`id` ASC LIMIT 1
	fmt.Println(user)

	user1 := new(User)  // 返回的是一个user结构体的指针
	fmt.Println(user1) // &{{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC <nil>}  0}
	db.First(user1)
	fmt.Println(user1) // &{{1 2020-09-04 10:32:21 +0800 CST 2020-09-04 10:32:21 +0800 CST <nil>} chiling 48}

	// // 随机获取一条记录
	var user2 User
	// 随机获取一条记录
	db.Take(&user2)
	fmt.Println(user2)

	// // 根据主键查询最后一条记录
	var user3 User
	db.Debug().Last(&user3)
	fmt.Println(user3)

	//查询所有的记录
	var user4 []*User
	db.Find(&user4)
	fmt.Println(user4)  // [0xc0002079e0 0xc000207b00 0xc000207c20 0xc000207d40]
	for _, v := range user4 {
		fmt.Println(*v)
	}

	var user5 []User
	db.Find(&user5)
	for _, v := range user5 {
		fmt.Println(v)
	}

	// // 查询指定的某条记录(仅当主键为整型时可用)
	var user6 User
	db.First(&user6, 3)
	fmt.Println(user6)  // {{3 2020-09-04 10:41:55 +0800 CST 2020-09-04 10:41:55 +0800 CST <nil>}  0}

	fmt.Println("-------------where------------")
	// Get first matched record
	var user7 User
	db.Where("name=?", "chiling").First(&user7)
	fmt.Println(user7)

	//// Get all matched records
	var user8 []User
	db.Where("name=?", "chiling").Find(&user8)
	fmt.Println(user8)

	// <>
	fmt.Println("<>")
	var user9 []User
	db.Where("name <> ?", "chiling").Find(&user9)
	fmt.Println(user9)

	// in
	fmt.Println("-----------in-------")
	var user10 []User
	db.Debug().Where("name in (?)", []string{"chiling", "qujian"}).Find(&user10)
	fmt.Println(user10)

	// like
	fmt.Println("-----------like-------")
	var user11 []User
	db.Where("name like ?", "%jia%").Find(&user11)
	fmt.Println(user11)


	// and
	fmt.Println("-----------and-------")
	var user12 []User
	db.Where("name = ? and age >= ?", "chiling", "20").Find(&user12)
	fmt.Println(user12)

	// time
	fmt.Println("-----------time-------")
	var user13 []User
	db.Where("updated_at > ?", "2020-09-04 10:30").Find(&user13)
	fmt.Println(user13)

	/*
	// between
	fmt.Println("-----------between-------")
	var user14 []User
	db.Where("created_at > between ? and ?", "lastWeek", "now()").Find(&user14)
	fmt.Println(user14)
	 */

	// struct & map查询
	fmt.Println("-----------struct & map -------")
	var users1 User
	db.Where(&User{Name: "chiling", Age: 48}).First(&users1)
	fmt.Println(users1)


	// attrs 如果记录未找到，将使用参数创建 struct 和记录.
	var users2 User
	db.Where(User{Name: "小红"}).Attrs(User{Age: 20}).FirstOrInit(&users2)
	fmt.Println(users2)

	// assign  不管记录是否找到，都将参数赋值给 struct 并保存至数据库.
	var users3 User
	db.Where(User{Name: "小红"}).Assign(User{Age: 20}).FirstOrCreate(&users3)
	fmt.Println(users3)

	// 更新
	var xxoo User
	xxoo.Name = "小东子"
	xxoo.Age = 90
	db.Save(&xxoo) // 默认会修改所有的字段

	// 更新当个属性
	//db.Model(&xxoo).Update("name", "小东子")

	var xxoo1 User
	db.Model(&xxoo1).Where("name = ?", "小东子").Update("name", "小桂子")

	//m1 := map[string]interface{}{
	//	"Name": "helloworld",
	//	"Age": 28,
	//}
	//db.Model(&xxoo1).Updates(m1) // 会更新所有的记录的name和age

	//db.Model(&xxoo1).Select("Age").Updates(m1) // 会更新所有记录的Age字段

	fmt.Println("---------------updatecolumn------------------")
	// 更新单个属性, 类似于update
	//m1 = map[string]interface{}{
	//	"Name": "ko",
	//	"Age": 28,
	//}
	//db.Debug().Model(&user).UpdateColumn("name", "小林子")

	//db.Table("users").Where("id IN (?)", []int{10, 11}).Updates(map[string]interface{}{"name": "hello_server", "age": 18})


	var usern User
	db.First(&usern,28)
	// 所有的 age + 2
	db.Model(&usern).Update("age", gorm.Expr("age + ?", 2))

	db.Debug().Model(&usern).Where("age > 10").UpdateColumn("age", gorm.Expr("age - ?", 10))

}

type User struct {
	gorm.Model
	Name string
	Age int64
}
