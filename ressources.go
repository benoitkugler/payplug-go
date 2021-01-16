package payplug

import (
	"encoding/json"
)

type Timestamp uint // unix timestamp, zero corresponds to null value

type Currency string // three-letter ISO 4217

const (
	Eur Currency = "EUR"
)

type Brand string

const (
	Mastercard Brand = "Mastercard"
	Maestro    Brand = "Maestro"
	Visa       Brand = "Visa"
	CB         Brand = "CB"
)

type Delivery string // any of: BILLING VERIFIED NEW SHIP_TO_STORE DIGITAL_GOODS TRAVEL_OR_EVENT OTHER

const (
	Billing_      Delivery = "BILLING"         // Ship to cardholder’s billing address
	Verified      Delivery = "VERIFIED"        // Ship to another verified address on file with merchant
	New           Delivery = "NEW"             // Ship to an address that is different than the cardholder’s billing address
	ShipToStore   Delivery = "SHIP_TO_STORE"   // Pick-up at a local store (store address shall be populated in shipping address fields)
	DigitalGoods  Delivery = "DIGITAL_GOODS"   // Online services, electronic gift cards, redemption codes, and other digital goods
	TravelOrEvent Delivery = "TRAVEL_OR_EVENT" // Travel and Event tickets, not shipped
	Other         Delivery = "OTHER"           // Other (for example, Gaming, digital services not shipped, e-media subscriptions, etc.)
)

type Authorization struct {
	AuthorizedAmount uint      `json:"authorized_amount,omitempty"` //  The positive amount that was authorized.
	AuthorizedAt     Timestamp `json:"authorized_at,omitempty"`     //  Date at which the payment was authorized by the customer, null if the payment has not yet been authorized.
	ExpiresAt        Timestamp `json:"expires_at,omitempty"`        //  Date at which the authorization expires, null if the payment has not yet been authorized
}

type OptionnalAuthorization struct {
	Valid         bool
	Authorization Authorization
}

func (ot OptionnalAuthorization) MarshalJSON() ([]byte, error) {
	if !ot.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(ot.Authorization)
}

func (ot *OptionnalAuthorization) UnmarshalJSON(b []byte) error {
	p := new(Authorization)
	err := json.Unmarshal(b, &p)
	if err != nil {
		return err
	}
	if p == nil {
		*ot = OptionnalAuthorization{}
	} else {
		ot.Valid = true
		ot.Authorization = *p
	}
	return nil
}

type CardPayment struct {
	Last4    string `json:"last_4,omitempty"`    // Last 4 digits of the card number.
	Country  string `json:"country,omitempty"`   // Country code (two-letter ISO 3166).
	ExpYear  int    `json:"exp_year,omitempty"`  // Credit card expiration year.
	ExpMonth int    `json:"exp_month,omitempty"` // Credit card expiration month.
	Brand    Brand  `json:"brand,omitempty"`     // Credit card brand, can be Mastercard, Maestro, Visa or CB.
	Id       string `json:"id,omitempty"`        // Credit card ID, available when a payment has been created with `save_card=true`, or has been created with this id.
}

type Billing struct {
	Title               string `json:"title,omitempty"`                 //  Customer title, can be mr, mrs, miss, null
	FirstName           string `json:"first_name,omitempty"`            //  Customer first name.
	LastName            string `json:"last_name,omitempty"`             //  Customer last name.
	Email               string `json:"email,omitempty"`                 //  Customer email address.
	MobilePhoneNumber   string `json:"mobile_phone_number,omitempty"`   //  Customer mobile phone number (international format in the E.164 standard).
	LandlinePhoneNumber string `json:"landline_phone_number,omitempty"` //  Customer landline phone number (international format in the E.164 standard).
	Address1            string `json:"address1,omitempty"`              //  Customer address line 1 (Street address/PO Box/Company name).
	Address2            string `json:"address2,omitempty"`              //  Customer address line 2 (Apartment/Suite/Unit/Building).
	CompanyName         string `json:"company_name,omitempty"`          //  Customer company.
	Postcode            string `json:"postcode,omitempty"`              //  Customer Zip/Postal code.
	City                string `json:"city,omitempty"`                  //  Customer city.
	State               string `json:"state,omitempty"`                 //  Customer state.
	Country             string `json:"country,omitempty"`               //  Customer country code (two-letter ISO 3166).
	Language            string `json:"language,omitempty"`              //  Customer language code (two-letter ISO 639-1).
}

type Shipping struct {
	Title               string   `json:"title,omitempty"`                 // Recipient title, can be mr, mrs, miss, null
	FirstName           string   `json:"first_name,omitempty"`            // Recipient first name.
	LastName            string   `json:"last_name,omitempty"`             // Recipient last name.
	Email               string   `json:"email,omitempty"`                 // Recipient email address.
	MobilePhoneNumber   string   `json:"mobile_phone_number,omitempty"`   // Mobile phone number (international format in the E.164 standard).
	LandlinePhoneNumber string   `json:"landline_phone_number,omitempty"` // Landline phone number (international format in the E.164 standard).
	Address1            string   `json:"address1,omitempty"`              // Shipping address line 1 (Street address/PO Box/Company name).
	Address2            string   `json:"address2,omitempty"`              // Shipping address line 2 (Apartment/Suite/Unit/Building).
	CompanyName         string   `json:"company_name,omitempty"`          // Company company.
	Postcode            string   `json:"postcode,omitempty"`              // Shipping Zip/Postal code.
	City                string   `json:"city,omitempty"`                  // Shipping city.
	State               string   `json:"state,omitempty"`                 // Shipping state.
	Country             string   `json:"country,omitempty"`               // Shipping country code (two-letter ISO 3166).
	Language            string   `json:"language,omitempty"`              // Shipping language code (two-letter ISO 639-1).
	DeliveryType        Delivery `json:"delivery_type,omitempty"`         // Type of delivery.
}

type HostedPayment struct {
	PaymentUrl string `json:"payment_url,omitempty"` // The payment URL you should redirect your customer to.
	ReturnUrl  string `json:"return_url,omitempty"`  // The URL the customer will be redirected to after the payment page whether it succeeds or not.
	CancelUrl  string `json:"cancel_url,omitempty"`  // The URL the customer will redirected to after a click on ‘Cancel Payment’.
	SentBy     string `json:"sent_by,omitempty"`     // By what means the payment URL was sent to the customer, if any.
}

type NotificationState struct {
	Url          string `json:"url,omitempty"`           // The URL PayPlug will send notifications to.
	ResponseCode int    `json:"response_code,omitempty"` // Integer http response code received when calling the URL of your notification page.
}

type PaymentFailureCode string

const (
	ProcessingError                              PaymentFailureCode = "processing_error"                                    //	Error while processing the card.
	CardDeclined                                 PaymentFailureCode = "card_declined"                                       //	The card has been rejected.
	InsufficientFunds                            PaymentFailureCode = "insufficient_funds"                                  //	Insufficient funds to cover the payment.
	Declined3ds                                  PaymentFailureCode = "3ds_declined"                                        //	The 3D Secure authentication request has failed.
	IncorrectNumber                              PaymentFailureCode = "incorrect_number"                                    //	The card number is incorrect.
	FraudSuspected                               PaymentFailureCode = "fraud_suspected"                                     //	Payment rejected because a fraud has been detected.
	MethodUnsupported                            PaymentFailureCode = "method_unsupported"                                  //	The payment method is not supported (e.g. e-carte bleue).
	CardSchemeMismatch                           PaymentFailureCode = "card_scheme_mismatch"                                //	The card number does not match with the selected brand.
	CardExpirationDatePriorToLastInstallmentDate PaymentFailureCode = "card_expiration_date_prior_to_last_installment_date" //	The card expiration date is prior to the last installment date.
	Aborted                                      PaymentFailureCode = "aborted"                                             //	Payment has been aborted with the Abort a payment feature.
	Timeout                                      PaymentFailureCode = "timeout"                                             //	The customer has not tried to pay and left the payment page.
)

type Failure struct {
	Code    PaymentFailureCode `json:"code,omitempty"`    // Payment failure code if the payment has failed.
	Message string             `json:"message,omitempty"` // Descriptive message explaining the reason of the payment failure.
}

type OptionnalFailure struct {
	Valid   bool
	Failure Failure
}

func (ot OptionnalFailure) MarshalJSON() ([]byte, error) {
	if !ot.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(ot.Failure)
}

func (ot *OptionnalFailure) UnmarshalJSON(b []byte) error {
	p := new(Failure)
	err := json.Unmarshal(b, &p)
	if err != nil {
		return err
	}
	if p == nil {
		*ot = OptionnalFailure{}
	} else {
		ot.Valid = true
		ot.Failure = *p
	}
	return nil
}

// Metadata are custom key/value pairs added by the user when creating
// objects.
type Metadata map[string]interface{}

type OneyPaiement struct {
	Type      string `json:"type,omitempty"`       // the type of payment made, either oney_x3_with_fees or oney_x4_with_fees.
	IsPending bool   `json:"is_pending,omitempty"` // whether the payment is in a pending state for Oney. If this is true, it means that the payer has successfully filled out Oney’s payment form, but Oney is still analyzing the payer’s file. In this case, the payment is neither authorized nor paid yet, but in a pending state. If a notification_url is set, then the payment resource will be posted to it once Oney has made its decision. See Interpreting an Oney payment status.
}

// Payment is the Payplug payment object.
type Payment struct {
	Id                string                 `json:"id,omitempty"`                  // Payment ID.
	Object            string                 `json:"object,omitempty"`              // Value is: payment.
	IsLive            bool                   `json:"is_live,omitempty"`             // true for a payment in LIVE mode, false in TEST mode.
	Amount            uint                   `json:"amount,omitempty"`              // Positive amount of the payment in cents.
	AmountRefunded    uint                   `json:"amount_refunded,omitempty"`     // Positive or zero	amount that has been refunded in cents.
	Authorization     OptionnalAuthorization `json:"authorization,omitempty"`       // Information about the authorization in case of a deferred payment, null otherwise.
	Currency          Currency               `json:"currency,omitempty"`            // Currency code (three-letter ISO 4217) in which the payment was made.
	InstallmentPlanId string                 `json:"installment_plan_id,omitempty"` // ID of the installment plan related to this payment, if any.
	IsPaid            bool                   `json:"is_paid,omitempty"`             // true if the payment has been paid, false if not.
	PaidAt            Timestamp              `json:"paid_at,omitempty"`             // Date at which the payment has been paid, null if the payment has not yet been paid.
	IsRefunded        bool                   `json:"is_refunded,omitempty"`         // true if this payment has been fully refunded, false if not.
	Is3ds             bool                   `json:"is_3ds,omitempty"`              // true if the payment was processed using 3-D Secure.
	SaveCard          bool                   `json:"save_card,omitempty"`           // true if the payment was used to save a card. On the payment page, saving a card was mendatory.
	AllowSaveCard     bool                   `json:"allow_save_card,omitempty"`     // Alternative to save_card. true if the payment gave the possibility to a customer to save a card. On the payment page, saving a card was optional.
	CreatedAt         Timestamp              `json:"created_at,omitempty"`          // Creation date.
	RefundableAfter   Timestamp              `json:"refundable_after,omitempty"`    // Date after which the payment can be refunded, null if the payment has not yet been paid.
	RefundableUntil   Timestamp              `json:"refundable_until,omitempty"`    // Date until which the payment can be refunded, null if the payment has not yet been paid.
	Card              CardPayment            `json:"card,omitempty"`                // Information about the card used for the payment.
	Billing           Billing                `json:"billing,omitempty"`             // Information about billing.
	Shipping          Shipping               `json:"shipping,omitempty"`            // Information about shipping.
	HostedPayment     HostedPayment          `json:"hosted_payment,omitempty"`      // Information about the payment.
	Failure           OptionnalFailure       `json:"failure,omitempty"`             // Information for unsuccessful payments.
	Description       string                 `json:"description,omitempty"`         // OPTIONAL Description shown to the customer.
	Metadata          Metadata               `json:"metadata,omitempty"`            // Custom metadata object added when creating the payment.
	PaymentMethod     OneyPaiement           `json:"payment_method,omitempty"`      // Data about the payment method, only available for Oney payments.
	NotificationUrl   string                 `json:"notification_url,omitempty"`    // The URL PayPlug will send notifications to.

	Notification NotificationState `json:"notification,omitempty"` // Data related to notifications
}

type Refund struct {
	Id        string    `json:"id,omitempty"`         // Refund ID
	PaymentId string    `json:"payment_id,omitempty"` // ID of the payment refunded.
	Object    string    `json:"object,omitempty"`     // Value is: refund.
	IsLive    bool      `json:"is_live,omitempty"`    // true for a payment in LIVE mode, false in TEST mode.
	Amount    uint      `json:"amount,omitempty"`     // Positive amount of the refund in cents and must be highier than 0.10€ (10 cents).
	Currency  Currency  `json:"currency,omitempty"`   // Currency code three-letter ISO 4217 in which the refund was made.
	CreatedAt Timestamp `json:"created_at,omitempty"` // Date of creation of the refund.
	Metadata  Metadata  `json:"metadata,omitempty"`   // Custom metadata object added to the request the object.
}

type AccountingReport struct {
	Id                 string    `json:"id,omitempty"`                   // The accounting report’s unique identifier.
	Object             string    `json:"object,omitempty"`               // Value is: accounting_report.
	IsLive             bool      `json:"is_live,omitempty"`              // true for an accounting report in LIVE mode, false in TEST mode.
	TemporaryUrl       string    `json:"temporary_url,omitempty"`        // The URL to download your accounting report file from once it’s available.
	FileAvailableUntil Timestamp `json:"file_available_until,omitempty"` // End date of your accounting report file’s availability. Note that files are available for 24 hours after their creation.
	StartDate          string    `json:"start_date,omitempty"`           // date (ISO 8601)	Your report’s start date. The report will cover all operations from the beginning of that day (UTC).
	EndDate            string    `json:"end_date,omitempty"`             // date (ISO 8601)	Your report’s end date. The report will cover all operations until the end of that day (UTC).
	NotificationUrl    string    `json:"notification_url,omitempty"`     // OPTIONAL	The URL PayPlug will send a notification to.
}
