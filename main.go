package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/cashier")
	{
		v1.GET("/", cashierInfoHandler())
		v1.POST("/purchase", purchaseHandler())
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func cashierInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, CashierInformation{
			"Cashier01",
			StorageInfo{
				[]Banknote{
					{
						"Coin10", 10, 5,
					},
					{
						"Coin5", 5, 2,
					},
				},
			},
		})
	}
}
func purchaseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pAmount float64
		var nAmount float64
		paymentAmount := c.PostForm("paymentAmount")
		pAmount, err := strconv.ParseFloat(paymentAmount, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Invalid PaymentAmount",
			})
			return
		}

		netAmount := c.PostForm("netAmount")
		nAmount, err = strconv.ParseFloat(netAmount, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Invalid NetAmount",
			})
			return
		}
		fmt.Println("Input ", pAmount, nAmount)

		c.JSON(200, gin.H{
			"Change": (pAmount - nAmount),
		})
	}
}

type CashierInformation struct {
	Name        string
	StorageInfo StorageInfo
}

type StorageInfo struct {
	Item []Banknote
}

type Banknote struct {
	Name     string
	Value    float64
	Quantity int
}
