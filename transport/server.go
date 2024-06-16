// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// HTTP server

package transport

import (
	"net"
	"net/http"
	"sync"
)

// Server wraps [http.Server]
type Server struct {
	http.Server
}

// NewServer creates a new [Server].
//
// The [http.Server] used only as configuration template. If it
// is nil, the reasonable defaults will be used instead.
func NewServer(template *http.Server) *Server {
	if template == nil {
		template = &http.Server{}
	}

	srvr := &Server{
		Server: *template,
	}

	return srvr
}

// ServeAutoTLS is similar to the [http.Server.Serve] and
// [http.Server.ServeTLS].
//
// It accepts incoming connections on the [http.Listener] l,
// automatically detect encrypted (TLS) and non-encrypted
// connections and serves them respectively to their type.
//
// This function runs "forewer" and returns only in a case
// of error. Use Server.Shutdown or Server.Close to force
// this function to exit.
func (srvr *Server) ServeAutoTLS(l net.Listener) error {
	plain, encrypted := NewAutoTLSListener(l)

	errchan := make(chan error, 2)
	var done sync.WaitGroup

	done.Add(2)

	go func() {
		err := srvr.Serve(plain)
		errchan <- err
		done.Done()
	}()

	go func() {
		err := srvr.ServeTLS(encrypted, "", "")
		errchan <- err
		done.Done()
	}()

	err := <-errchan

	plain.Close()
	encrypted.Close()

	done.Wait()

	return err
}
