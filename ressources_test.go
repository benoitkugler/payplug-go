package payplug

import (
	"encoding/json"
	"testing"
)

func TestJSON(t *testing.T) {
	v := Metadata{
		"custom_id":    78.,
		"custom_value": "test",
	}
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	var out Metadata
	err = json.Unmarshal(b, &out)
	if err != nil {
		t.Fatal(err)
	}
	if i := out["custom_id"]; i != 78. {
		t.Fatalf("expected 78., got %v type %T", i, i)
	}
	if i := out["custom_value"]; i != "test" {
		t.Fatalf("expected 'test', got %v type %T", i, i)
	}
}

func TestPayment(t *testing.T) {
	jsonIn := `
	{
		"id": "pay_5iHMDxy4ABR4YBVW4UscIn",
		"object": "payment",
		"is_live": true,
		"amount": 3300,
		"amount_refunded": 0,
		"authorization": null,
		"currency": "EUR",
		"created_at": 1434010787,
		"refundable_after": 1449157171,
		"refundable_until": 1459157171,
		"installment_plan_id": null,
		"is_paid": true,
		"paid_at": 1555073519,
		"is_refunded": false,
		"is_3ds": false,
		"save_card": false,
		"card": {
		  "last4": "1800",
		  "country": "FR",
		  "exp_month": 9,
		  "exp_year": 2017,
		  "brand": "Mastercard",
		  "id": null
		},
		"billing": {
		  "title": "mr",
		  "first_name": "John",
		  "last_name": "Watson",
		  "email": "john.watson@example.net",
		  "mobile_phone_number": null,
		  "landline_phone_number": null,
		  "address1": "221B Baker Street",
		  "address2": null,
		  "postcode": "NW16XE",
		  "city": "London",
		  "state": null,
		  "country": "GB",
		  "language": "en"
		},
		"shipping": {
		  "title": "mr",
		  "first_name": "John",
		  "last_name": "Watson",
		  "email": "john.watson@example.net",
		  "mobile_phone_number": null,
		  "landline_phone_number": null,
		  "address1": "221B Baker Street",
		  "address2": null,
		  "postcode": "NW16XE",
		  "city": "London",
		  "state": null,
		  "country": "GB",
		  "language": "en",
		  "delivery_type": "BILLING"
		},
		"hosted_payment": {
		  "payment_url": "https://secure.payplug.com/pay/5iHMDxy4ABR4YBVW4UscIn",
		  "return_url": "https://example.net/success?id=42",
		  "cancel_url": "https://example.net/cancel?id=42",
		  "paid_at": 1434010827,
		  "sent_by": null
		},
		"notification": {
		  "url": "https://example.net/notifications?id=42",
		  "response_code": 200
		},
		"failure": null,
		"description": null,
		"metadata": {
		  "customer_id": 42
		}
	  }`
	var p Payment
	if err := json.Unmarshal([]byte(jsonIn), &p); err != nil {
		t.Fatal(err)
	}
}

func TestRefound(t *testing.T) {
	jsonIn := `
	{
		"id": "re_3NxGqPfSGMHQgLSZH0Mv3B",
		"payment_id": "pay_5iHMDxy4ABR4YBVW4UscIn",
		"object": "refund",
		"is_live": true,
		"amount": 358,
		"currency": "EUR",
		"created_at": 1434012358,
		"metadata": {
		  "customer_id": 42,
		  "reason": "The delivery was delayed"
		}
	  }`
	var r Refund
	if err := json.Unmarshal([]byte(jsonIn), &r); err != nil {
		t.Fatal(err)
	}
}
