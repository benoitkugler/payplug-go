package payplug

import (
	"fmt"
	"path"
)

// API base url
const (
	API_BASE_URL = "https://api.payplug.com"
	API_VERSION  = "1"
	baseUrl      = API_BASE_URL + "/v" + API_VERSION
)

// Resources URL
const (
	PAYMENT_RESOURCE           = baseUrl + "/payments"
	REFUND_RESOURCE            = PAYMENT_RESOURCE + "/%s/refunds" // payment id
	CUSTOMER_RESOURCE          = baseUrl + "/customers"
	CARD_RESOURCE              = CUSTOMER_RESOURCE + "/%s/cards" // customer id
	ACCOUNTING_REPORT_RESOURCE = baseUrl + "/accounting_reports"
)

func (p *Payment) urlForConsistent() string {
	return path.Join(PAYMENT_RESOURCE, p.Id)
}

func (r *Refund) urlForConsistent() string {
	u := fmt.Sprintf(REFUND_RESOURCE, r.PaymentId)
	return path.Join(u, r.Id)
}

func (a *AccountingReport) urlForConsistent() string {
	return path.Join(ACCOUNTING_REPORT_RESOURCE, a.Id)
}
