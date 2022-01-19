package models

import (
	"log"
	"time"

	db "derekshop.com/product/database"
)

type Product struct {
	ProductNo     int       `json:"product_no" form:"product_no"`
	ProductName   string    `json:"product_name" form:"product_name"`
	ProductPrice  int       `json:"product_price" form:"product_price"`
	ProductDesc   string    `json:"product_desc" form:"product_desc"`
	ImgUrl        string    `json:"img_url" form:"img_url"`
	CreateTime    time.Time `json:"create_time" form:"create_time"`
	ProductStatus int       `json:"product_status" form:"product_status"`
}

func (p *Product) AddProduct() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO product(product_name, product_price, product_desc, img_url, create_time, product_status) VALUES (?, ?, ?, ?, now(), 0)", p.ProductName, p.ProductPrice, p.ProductDesc, p.ImgUrl)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (p *Product) GetProducts() (products []Product, err error) {
	products = make([]Product, 0)
	rows, err := db.SqlDB.Query("SELECT * FROM product")

	if err != nil {
		log.Printf("Query failed,err:%v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product Product

		err = rows.Scan(&product.ProductNo, &product.ProductName, &product.ProductPrice,
			&product.ProductDesc, &product.ImgUrl, &product.CreateTime, &product.ProductStatus)
		if err != nil {
			log.Printf("Scan failed,err:%v\n", err)
			return
		}
		products = append(products, product)
	}

	log.Printf("products:%+v:", products)
	return
}

func (p *Product) GetProduct(productNo int) (product Product, err error) {

	err = db.SqlDB.QueryRow("SELECT * FROM product WHERE product_no=?", productNo).Scan(
		&product.ProductNo, &product.ProductName, &product.ProductPrice,
		&product.ProductDesc, &product.ImgUrl, &product.CreateTime,
		&product.ProductStatus)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (p *Product) ModProduct(productNo int) (product Product, err error) {
	stm, _ := db.SqlDB.Prepare("UPdate user set age=? where product_no=?")
	stm.Exec(productNo)
	stm.Close()

	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (p *Product) DelProduct(productNo int) (product Product, err error) {

	//err = db.SqlDB.Exec("DELETE FROM product WHERE product_no=?", productNo)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	return
}
