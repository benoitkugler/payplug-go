// This package is a library to ease the use of the Payplug payment services.
package payplug

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"runtime"
)

const version = "1.0.0"

// Session enables to create requests
// to Payplug server.
type Session struct {
	secretKey  string
	apiVersion string

	client *http.Client
}

func NewSession(token string) Session {
	return Session{secretKey: token, client: http.DefaultClient}
}

// NewSessionCert use `cert` content as a CA bundle
func NewSessionCert(token string, cert io.Reader) (Session, error) {
	caCert, err := ioutil.ReadAll(cert)
	if err != nil {
		return Session{}, fmt.Errorf("missing CA: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{RootCAs: caCertPool}}}

	return Session{secretKey: token, client: client}, nil
}

// SetApiVersion set the desired `version`, as an ISO-8601 date
func (s *Session) SetApiVersion(version string) {
	s.apiVersion = version
}

// Perform an HTTP request, by marshalling `body` as JSON, and unmarshal the response in `out`, which must be
// a pointer type.
// The status code is also checked, meaning that if `err` is nil, then `status` is valid (in the 2XX range).
func (s Session) Request(method, url string, body interface{}, out interface{}) (status int, err error) {
	b, err := json.Marshal(body)
	if err != nil {
		return 0, ClientError{err: err}
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(b))
	if err != nil {
		return 0, ClientError{err: err}
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	userAgent := fmt.Sprintf("Payplug-Go/%s (Go/%s)", version, runtime.Version())
	req.Header.Set("User-Agent", userAgent)

	if s.secretKey == "" {
		return 0, SecretKeyNotSet
	}
	req.Header.Set("Authorization", "Bearer "+s.secretKey)

	if s.apiVersion != "" {
		req.Header.Set("PayPlug-Version", s.apiVersion)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, ClientError{err: err}
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, ClientError{err: err}
	}

	if !(200 <= resp.StatusCode && resp.StatusCode < 300) {
		return resp.StatusCode, HttpError{code: resp.StatusCode, err: string(content)}
	}

	if err := json.Unmarshal(content, out); err != nil {
		return resp.StatusCode, unexpectedAPIResponseErr(err)
	}
	return resp.StatusCode, nil
}

// CreatePayment is a shortcut to add `payment`.
func (s Session) CreatePayment(payment Payment) (Payment, error) {
	var out Payment
	_, err := s.Request(http.MethodPost, PAYMENT_RESOURCE, payment, &out)
	return out, err
}
