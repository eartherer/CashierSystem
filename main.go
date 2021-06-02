package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var banknoteStorage *StorageSystem

func main() {

	banknoteStorage = NewStorageSystemWithName("Storage01")
	banknoteStorage.initBankNoteStorageWithDefault()

	r := gin.Default()
	r = setupRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func setupRouter(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/cashier")
	{
		v1.GET("", cashierInfoHandler())
		v1.POST("/purchase", purchaseHandler())
		v1.POST("/refill", storageRefillHandler())
	}
	return r
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
		paymentAmount := c.Query("paymentAmount")
		pAmount, err := strconv.ParseFloat(paymentAmount, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Invalid PaymentAmount",
			})
			return
		}

		netAmount := c.Query("netAmount")
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
		result, ok := calculateBanknoteFromChange(banknoteStorage, change)
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

func storageRefillHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		nameInput := c.Query("name")
		if nameInput == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Invalid Banknote Name",
			})
			return
		}
		quantityInput := c.Query("quantity")
		quantity, err := strconv.Atoi(quantityInput)
		if (quantity < 0) || (err != nil) {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Invalid Quantity",
			})
			return
		}

		err = banknoteStorage.refillBankNoteStorage(nameInput, quantity)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": err.Error(),
			})
			return
		} else {
			c.JSON(200, gin.H{
				"Message": "Refill Banknote success",
			})
		}
	}
}

type CashierInformation struct {
	Name        string         `json:"CashierName"`
	StorageInfo *StorageSystem `json:"Storage"`
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
