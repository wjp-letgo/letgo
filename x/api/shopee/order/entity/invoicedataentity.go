package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//InvoiceDataEntity
type InvoiceDataEntity struct {
	Number             string  `json:"number"`
	SeriesNumber       string  `json:"series_number"`
	AccessKey          string  `json:"access_key"`
	IssueDate          int     `json:"issue_date"`
	TotalValue         float32 `json:"total_value"`
	ProductsTotalValue float32 `json:"products_total_value"`
	TaxCode            string  `json:"tax_code"`
}

//String
func (i InvoiceDataEntity) String() string {
	return lib.ObjectToString(i)
}
