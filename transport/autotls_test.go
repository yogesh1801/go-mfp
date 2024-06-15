// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// TLS auto-detect

package transport

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"sync"
	"testing"
	"time"
)

// TLS certificate, for testing
var testAutoTLSCert = testAutoTLSCertGenerate()

// testAutoTLSAddr tests Addr() method of child listeners
func testAutoTLSAddr(t *testing.T, tr *Transport, l net.Listener) {
	p, e := NewAutoTLSListener(l)

	addr := l.Addr().String()
	if p.Addr().String() != addr {
		t.Errorf("plain.Addr(): expected %s, present %s",
			p.Addr(), addr)
	}

	if e.Addr().String() != addr {
		t.Errorf("encrypted.Addr(): expected %s, present %s",
			e.Addr(), addr)
	}
}

// testAutoTLSHTTP tests HTTP/HTTPS request over AutoTLS listener
func testAutoTLSHTTP(t *testing.T, tr *Transport, l net.Listener) {
	const response = "Hello, world!"

	// Build http/https URLs
	addr := l.Addr()
	urlHTTP := MustParseURL(fmt.Sprintf("http://%s/", addr))
	urlHTTPS := MustParseURL(fmt.Sprintf("https://%s/", addr))

	// Create a client
	clnt := &http.Client{
		Transport: tr,
		Timeout:   5 * time.Second,
	}

	// Create http.Server
	handler := func(w http.ResponseWriter, rq *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}

	srvr := &http.Server{
		Handler:      http.HandlerFunc(handler),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		TLSConfig: &tls.Config{
			//Certificates: []tls.Certificate{testAutoTLSCert},
			GetCertificate: func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
				return testAutoTLSCert, nil
			},
		},
	}

	// Run server in background
	var done sync.WaitGroup
	done.Add(2)

	p, e := NewAutoTLSListener(l)
	go func() {
		srvr.Serve(p)
		done.Done()
	}()

	go func() {
		srvr.ServeTLS(e, "", "")
		done.Done()
	}()

	// Perform HTTP requests
	for _, u := range []*url.URL{urlHTTP, urlHTTPS} {
		rq, err := NewRequest("GET", u, nil)
		if err != nil {
			t.Errorf("GET %s: %s", u, err)
			continue
		}

		rsp, err := clnt.Do(rq)
		if err != nil {
			t.Errorf("GET %s: %s", u, err)
			continue
		}

		if rsp.StatusCode != http.StatusOK {
			t.Errorf("GET %s: status expected %d, present %d",
				u, rsp.StatusCode, http.StatusOK)
		}

		rsp.Body.Close()
	}

	// Shutdown the server
	srvr.Close()
	done.Wait()
}

// testAutoTLSServerClose tests that incoming but not yet accepted connections
// are properly closed
func testAutoTLSServerClose(t *testing.T, tr *Transport, l net.Listener) {
	// Build http/https URLs
	addr := l.Addr()
	urlHTTP := MustParseURL(fmt.Sprintf("http://%s/", addr))
	urlHTTPS := MustParseURL(fmt.Sprintf("https://%s/", addr))

	// Create a client
	clnt := &http.Client{
		Transport: tr,
		Timeout:   5 * time.Second,
	}

	// Setup AutoTLS listener
	atl, p, e := newAutoTLSListener(l)

	// Initiate incoming requests
	var done sync.WaitGroup
	numConn := 0

	for i := 0; i < 64; i++ {
		for _, u := range []*url.URL{urlHTTP, urlHTTPS} {
			rq, err := NewRequest("GET", u, nil)
			if err != nil {
				t.Errorf("GET %s: %s", u, err)
				continue
			}

			numConn++
			go func() {
				done.Add(1)

				rsp, err := clnt.Do(rq)
				if err == nil {
					rsp.Body.Close()
				}

				done.Done()
			}()
		}
	}

	// Wait until all connections are internally accepted
	for {
		plain, encrypted, pending := atl.testCounters()
		total := plain + encrypted + pending

		if total == numConn {
			break
		}

		atl.acceptWait()
	}

	// Close listeners
	p.Close()
	e.Close()

	// Running requests MUST terminate
	done.Wait()
}

// TestAutoTLS performs a series of tests of the AutoTLS listener
func TestAutoTLS(t *testing.T) {
	var dialer net.Dialer

	// prep function for TCP connections
	prepTCP := func() (*Transport, net.Listener, error) {
		l, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			return nil, nil, err
		}

		template := (http.DefaultTransport.(*http.Transport)).Clone()
		template.DialContext = dialer.DialContext
		template.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

		tr := NewTransport(template)

		return tr, l, nil
	}

	// testData represents a single test
	type testData struct {
		prep func() (*Transport, net.Listener, error)
		test func(*testing.T, *Transport, net.Listener)
	}

	// tests contains a series of tests to be performed
	tests := []testData{
		{
			prep: prepTCP,
			test: testAutoTLSAddr,
		},

		{
			prep: prepTCP,
			test: testAutoTLSHTTP,
		},

		{
			prep: prepTCP,
			test: testAutoTLSServerClose,
		},
	}

	// Run tests in loop
	for _, test := range tests {
		// Setup things
		tr, l, err := test.prep()
		if err != nil {
			t.Errorf("%s", err)
			continue
		}

		// Run a test
		test.test(t, tr, l)

		// Cleanup
		l.Close()
	}
}

// testAutoTLSCertGenerate generates TLS certificate, for testing
func testAutoTLSCertGenerate() *tls.Certificate {
	// Generate private key
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	// Fill certificate template
	serialNumber := big.NewInt(12345)
	notBefore := time.Now()
	notAfter := notBefore.Add(time.Hour * 24 * 365)
	keyUsage := x509.KeyUsageDigitalSignature

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Test Certificate"},
		},

		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              keyUsage,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              []string{"localhost"},
	}

	// Generate certificate
	der, err := x509.CreateCertificate(rand.Reader,
		&template, &template, pub, priv)

	if err != nil {
		panic(err)
	}

	cert := &tls.Certificate{
		Certificate: [][]byte{der},
		PrivateKey:  priv,
	}

	return cert
}
