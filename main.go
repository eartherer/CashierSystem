package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var banknoteStorage StorageSystem

func main() {

	banknoteStorage = NewStorageSystemWithName("Storage01")

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
			banknoteStorage,
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
		change := pAmount - nAmount
		if change < 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Not Enough payment amount",
			})
			return
		}
		result, ok := calculateBanknoteFromChange(change)
		fmt.Println(result)
		if !ok {
			c.JSON(200, gin.H{
				"Message": "Cannot purchase,storage not enough change money",
			})
			return
		}
		if len(result.BanknoteChange) == 0 {
			c.JSON(200, gin.H{
				"Message": "No change money",
			})
			return
		}
		c.JSON(200, result)
	}
}

type CashierInformation struct {
	Name        string
	StorageInfo StorageSystem
}

type BankNoteChangeInfo struct {
	Change         float64
	BanknoteChange []BankNoteInfo
}

type BankNoteInfo struct {
	Name     string
	Value    float64
	Quantity int
}
