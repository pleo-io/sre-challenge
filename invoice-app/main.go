package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"pleo.io/invoice-app/db"

	"github.com/gin-gonic/gin"
)

func main() {
	dbClient = db.InitializeDatabase()

	router := setupRouter()

	err := router.Run(":8081")
	if err != nil {
		fmt.Printf("could not start server: %v", err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.New()
	r.POST("invoices/pay", pay)
	r.GET("invoices", getInvoices)
	return r
}

func getInvoices(c *gin.Context) {
	invoices := dbClient.GetInvoices()

	c.JSON(http.StatusOK, invoices)
}

func pay(c *gin.Context) {
	invoices := dbClient.GetUnpaidInvoices()
	for _, invoice := range invoices {
		client := http.Client{}
		req := payRequest{
			Id:       invoice.InvoiceId,
			Value:    invoice.Value,
			Currency: invoice.Currency,
		}
		b, err := json.Marshal(req)
		data := bytes.NewBuffer(b)
		_, err = client.Post(os.Getenv("PAYMENT_URL"), "application/json", data)

		if err != nil {
			fmt.Printf("Error %s", err)
			return
		}

		dbClient.PayInvoice(invoice.InvoiceId)
	}

	fmt.Printf("Invoices paid!\n")

	c.JSON(http.StatusOK, gin.H{})
}

var dbClient *db.Client

type payRequest struct {
	Id       string  `json:"id"`
	Value    float32 `json:"value"`
	Currency string  `json:"currency"`
}
