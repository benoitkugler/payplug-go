package payplug

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const notificationMaxSize = 1000000 // 1 MB should be largely sufficient

// must be implemented by pointer types, since
// JSON content will be unmarshalled into the receiver
type notificationTarget interface {
	// returns the urlForConsistent to GET the consistent data
	urlForConsistent() string
}

// handles a request sent by PayPlug to your server to notify your system that some object (a payment, an installment plan, etc) was updated.
func (s Session) handleNotification(body io.Reader, n notificationTarget) error {
	content, err := ioutil.ReadAll(io.LimitReader(body, notificationMaxSize))
	if err != nil {
		return fmt.Errorf("can't read notification body : %s", err)
	}

	if err = json.Unmarshal(content, n); err != nil {
		return unexpectedAPIResponseErr(err)
	}

	_, err = s.Request(http.MethodGet, n.urlForConsistent(), nil, n) // fetch the true data
	return err                                                       // if err is nil, `n` is now completed and trusted
}

// HandleNotificationPayment reads the `body` of a notification,
// and fetch the completed and trusted data from PayPlug.
func (s Session) HandleNotificationPayment(body io.Reader) (Payment, error) {
	var r Payment
	err := s.handleNotification(body, &r)
	return r, err
}

// HandleNotificationRefund reads the `body` of a notification,
// and fetch the completed and trusted data from PayPlug.
func (s Session) HandleNotificationRefund(body io.Reader) (Refund, error) {
	var r Refund
	err := s.handleNotification(body, &r)
	return r, err
}

// HandleNotificationAccountingReport reads the `body` of a notification,
// and fetch the completed and trusted data from PayPlug.
func (s Session) HandleNotificationAccountingReport(body io.Reader) (AccountingReport, error) {
	var r AccountingReport
	err := s.handleNotification(body, &r)
	return r, err
}

// payment: Payment
// refund: Refund
// accounting_report: AccoutingReport

// not verifiable
// customer: Customer
// card: Card
