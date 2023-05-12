package models

import (
	"fmt"
	"log"
	"time"
)

type Goods struct {
	Id         int
	Title      string
	Price      float64
	Stock      int
	Type       int
	CreateTime time.Time
}

func (Goods) TableName() string {
	return "goods"
}

func SvaeGoods(goods Goods) {
	GetDB().Create(&goods)
}

func UpdateGoods() {
	var goods Goods
	GetDB().Where("id = ?", "1").Take(&goods)
	goods.Price = 50
	//这里的save实际上是进行了更新,save是有则更新，无则插入 类似于upserte
	//GetDB().Save(&goods)
	//	Model(&Goods{}) 与 Model(&goods)是不同的。在上述查询中赋值于goods指针变量时附带有id，那么在进行update的时候没有使用Where条件时会自动识别id字段作为where条件
	//err := GetDB().Model(&goods).Update("title", "香蕉").Error
	//	1.更新单列
	//err := GetDB().Model(&Goods{}).Where("id=?", "1").Update("title", "香蕉").Error
	////	2.更新多列
	//err := GetDB().Model(&Goods{}).Where("id=?", "1").Updates(Goods{
	//	Price: 30,
	//	Title: "橘子",
	//}).Error
	//	3.也可以使用map更新
	//	4.选择更新字段
	//err := GetDB().Model(&Goods{}).Where("id=?", "1").Select("title").Updates(Goods{
	//	Price: 30,
	//	Title: "橘子",
	//}).Error
	//	5.排除更新字段
	//err := GetDB().Model(&Goods{}).Where("id=?", "1").Omit("title").Updates(Goods{
	//	Price: 30,
	//	Title: "橘子",
	//}).Error
	//	6.子查询更新字段-->在update中的value使用条件操作
	err := GetDB().Model(&goods).Where("id=?", "1").Update("title", GetDB().Model(&UserBasic{}).Select("name").Where("id = ?", 2)).Error

	if err != nil {
		log.Panicln("更新数据有误：", err)
	}
}

func DeleteGoods() {
	db := GetDB()
	db.Delete(&Goods{}, 2)
}

func FindGoods() {
	db := GetDB()
	//var goods Goods
	//var titles []string
	//	1.查询单条数据 如果没有查询到会报一个error的错误
	//db.Where("id = ?", "3").Take(&goods)
	//	2.查询第一条数据
	//db.Model(&Goods{}).Where("id > ?", "3").First(&goods)
	//	3.查询最后一条数据
	//db.Model(&Goods{}).Where("id > ?", "3").Last(&goods)
	//	4.查询多条数据或单条数据
	//db.Where("id > ?", "3").Find(&goods)
	//err := db.Where("id > ?", "3").Limit(1).Find(&goods).Error
	//	5.查询一列的值
	//err := db.Model(&Goods{}).Pluck("title", &titles).Error
	//	6.使用聚合函数进行查询
	//err := db.Model(&Goods{}).Select("count(*) as nums").Pluck("nums", &titles).Error
	//	7.分页查询
	//err := db.Model(&Goods{}).Limit(2).Offset(2).Find(&goods).Error
	//	8.分组查询
	//统计每个商品分类下面有多少个商品
	//定一个Result结构体类型，用来保存查询结果
	type Result struct {
		Type  int
		Total int
	}
	var results []Result

	////等价于: SELECT type, count(*) as  total FROM `goods` GROUP BY type HAVING (total > 0)
	//db.Model(Goods{}).Select("type, count(*) as  total").Group("type").Having("total > 0").Scan(&results)

	//scan类似Find都是用于执行查询语句，然后把查询结果赋值给结构体变量，区别在于scan不会从传递进来的结构体变量提取表名.
	//这里因为我们重新定义了一个结构体用于保存结果，但是这个结构体并没有绑定goods表，所以这里只能使用scan查询函数。
	//	9.直接运行sql语句 (不要做可变参数的拼接，防止sql注入，尽量使用如下的参数化查询)
	sql := "SELECT type, count(*) as  total FROM `goods` where id > ? GROUP BY type HAVING (total > 0) "
	db.Raw(sql, "1").Scan(&results)

	fmt.Println(results)
	//
	//if err != nil {
	//	log.Println("err:", err)
	//}
}