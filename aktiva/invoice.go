package aktiva

import "github.com/shopspring/decimal"

type InvoiceRow struct {
	Item struct {
		Code        string
		Description string
		Type        int
		UOMName     string `json:",omitempty"`
	}
	Quantity       decimal.Decimal
	Price          decimal.Decimal
	DiscountPct    decimal.Decimal `json:",omitempty"`
	DiscountAmount decimal.Decimal `json:",omitempty"`
	TaxId          string
	LocationCode   string `json:",omitempty"`
}

type TaxAmount struct {
	TaxId  string
	Amount decimal.Decimal
}

type Payment struct {
	PaymentMethod string
	PaidAmount    decimal.Decimal
	PaymDate      string
}

type Invoice struct {
	Customer struct {
		Id              string `json:",omitempty"`
		Name            string
		RegNo           string          `json:",omitempty"`
		NotTDCustomer   bool            `json:",omitempty"`
		VatRegNo        string          `json:",omitempty"`
		CurrencyCode    string          `json:",omitempty"`
		PaymentDeadLine int             `json:",omitempty"`
		OverDueCharge   decimal.Decimal `json:",omitempty"`
		Address         string          `json:",omitempty"`
		City            string          `json:",omitempty"`
		Country         string          `json:",omitempty"`
		PostalCode      string          `json:",omitempty"`
		CountryCode     string
		PhoneNo         string `json:",omitempty"`
		PhoneNo2        string `json:",omitempty"`
		HomePage        string `json:",omitempty"`
		Email           string `json:",omitempty"`
	}
	DocDate        string
	DueDate        string
	InvoiceNo      string `json:",omitempty"`
	RefNo          string `json:",omitempty"`
	CurrencyCode   string `json:",omitempty"`
	DepartmentCode string `json:",omitempty"`
	ProjectCode    string `json:",omitempty"`
	InvoiceRow     []InvoiceRow
	TaxAmount      []TaxAmount
	RoundingAmount decimal.Decimal `json:",omitempty"`
	TotalAmount    decimal.Decimal
	Payment        *Payment `json:",omitempty"`
	Hcomment       string   `json:",omitempty"`
	Fcomment       string   `json:",omitempty"`
}

type SendInvoiceResponse struct {
	CustomerId string
	InvoiceId  string
	InvoiceNo  string
	RefNo      string
}

func (a *Aktiva) SendInvoice(i Invoice) (*SendInvoiceResponse, error) {
	var resp SendInvoiceResponse

	if err := a.httpPost(epSendInvoice, i, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
