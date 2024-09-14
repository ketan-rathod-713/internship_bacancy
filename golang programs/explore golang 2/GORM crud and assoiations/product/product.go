package product

import (
	"fmt"

	"gorm.io/gorm"
)

type Product struct {
	Id           uint64 `gorm:"primaryKey"`
	Name         string
	Price        int64
	CategoryName string   // foreign key referencing Category.Name
	Category     Category // define the relationship // many to one so no need to use array here
}

type ProductCategoryStatus struct {
	Category   string
	TotalPrice int64
}

type Category struct {
	Name     string `gorm:"primaryKey"`
	Remark   string
	Products []Product // one to many relationship
}

func (p *Product) TableName() string {
	return "gormbasics1.product"
}

func (p *Category) TableName() string {
	return "gormbasics1.category"
}

func ProductPackage(db *gorm.DB) {
	db.AutoMigrate(&Category{})
	// var category1 Category
	// db.Create(&Category{Name: "electronic", Remark: "this is the best grocery ha ha"})

	db.AutoMigrate(&Product{}) // if it is first then it will violet foreign key constraint
	// insert some Dummy Data
	// var product Product = Product{
	// 	Id:           2, // keep changing id
	// 	Name:         "head phone 3",
	// 	Price:        40,
	// 	CategoryName: "grocery",
	// }

	// db.Create(&product)

	/* Now apply query operations */
	/* Get All Products */
	/*
		var products []Product
		db.Model(&Product{}).Find(&products)
		fmt.Println(products)
	*/

	// /* Get all categories with count of total price */
	// var productCategoryStatuses []ProductCategoryStatus
	// // TODO: by default gorm uses camelcase writting hence we should write total_price and not TotalPrice or any other thing
	// db.Model(&Product{}).Select("category, sum(price) as total_price").Group("category").Scan(&productCategoryStatuses)

	// fmt.Println(productCategoryStatuses)

	/* Joins of category and product */
	/* Insert Dummy Data */
	// db.Create(&Category{Name: "electronic", Remark: "It is used for home purpose"})

	// type result struct {
	// 	ProductName  string
	// 	CategoryName string
	// }
	// var rs []result
	// db.Model(&Product{}).Select("product.name as product_name ,category.name as category_name").Joins("inner join gormbasics1.category on category.name = product.category").Scan(&rs)
	// fmt.Println(rs)

	// db.Exec("select product.name, category.remark from gormbasics1.product inner join gormbasics1.category on product.category = category.name limit 1", rs)
	// fmt.Println(rs)

	/* Scan works similar as that of Find */

	/* */
	var product3 Product
	// var category2 Category
	db.Preload("Category").First(&product3)

	fmt.Println(product3)
}
