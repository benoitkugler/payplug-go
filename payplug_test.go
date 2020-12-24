package payplug

import (
	"crypto/x509"
	"io/ioutil"
	"log"
	"testing"
)

func TestCA(t *testing.T) {
	caCert, err := ioutil.ReadFile("certs/cacert.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	ok := caCertPool.AppendCertsFromPEM(caCert)
	if !ok {
		t.Fatal("invalid certificats")
	}
}

func TestWBadAuth(t *testing.T) {
	s := NewSession("invalid token")
	_, err := s.CreatePayment(Payment{})
	if asHttpError, ok := err.(HttpError); ok {
		if asHttpError.code != 401 {
			t.Fatalf("wrong error code, expected 401, got %d", asHttpError.code)
		}
	} else {
		t.Fatalf("wrong error, expected HttpError, got %T (%v)", err, err)
	}
}

type payList struct {
	Object  string    `json:"object,omitempty"`
	Page    int       `json:"page,omitempty"`
	PerPage int       `json:"per_page,omitempty"`
	HasMore bool      `json:"has_more,omitempty"`
	Data    []Payment `json:"data,omitempty"`
}
