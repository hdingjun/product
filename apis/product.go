package apis

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"derekshop.com/product/models"
	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddProductApi(c *gin.Context) {
	var p models.Product
	err := c.Bind(&p)
	ra, err := p.AddProduct()
	if err != nil {
		log.Println(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetProductApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var p models.Product
	product, err := p.GetProduct(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"product": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func GetProductsApi(c *gin.Context) {
	var p models.Product
	products, err := p.GetProducts()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"products": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func ModProductApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var p models.Product
	product, err := p.ModProduct(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"product": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func DelProductApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var p models.Product
	product, err := p.DelProduct(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"product": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}
