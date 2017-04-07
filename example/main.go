package main

import (
	"fmt"
	"time"

	"github.com/akosmarton/merit-aktiva-api-go/aktiva"
	"github.com/shopspring/decimal"
)

func main() {
	a := aktiva.NewAktiva(aktiva.URLEstonia, "ApiId", "ApiKey")

	i := aktiva.Invoice{}
	i.Customer.Name = "New Customer"
	i.Customer.CountryCode = "US"
	i.DocDate = aktiva.TimeToString(time.Now())
	i.DueDate = aktiva.TimeToString(time.Now())
	ir := aktiva.InvoiceRow{}
	ir.Item.Code = "New Code"
	ir.Item.Description = "New Description"
	ir.Item.Type = 3
	ir.TaxId = "b9b25735-6a15-4d4e-8720-25b254ae3d21"
	ir.Quantity = decimal.NewFromFloat(1.0)
	ir.Price = decimal.NewFromFloat(1000.)
	i.InvoiceRow = append(i.InvoiceRow, ir)
	tr := aktiva.TaxAmount{}
	tr.TaxId = "b9b25735-6a15-4d4e-8720-25b254ae3d21"
	tr.Amount = decimal.NewFromFloat(200.0)
	i.TaxAmount = append(i.TaxAmount, tr)
	i.InvoiceNo = "INV0000001"
	i.TotalAmount = decimal.NewFromFloat(1000.)

	i.Payment = &aktiva.Payment{}
	i.Payment.PaymDate = aktiva.TimeToString(time.Now())
	i.Payment.PaymentMethod = "PayPal"
	i.Payment.PaidAmount = decimal.NewFromFloat(1200.0)

	r, err := a.SendInvoice(i)
	if err != nil {
		panic(err)
	}

	fmt.Println("InvoiceId:", r.InvoiceId)
}
