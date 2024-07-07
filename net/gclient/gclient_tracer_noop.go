// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gclient

import (
	"crypto/tls"
	"net/http/httptrace"
	"net/textproto"
)

type clientTracerNoop struct{}

// newClientTracerNoop creates and returns object of httptrace.ClientTrace.
func newClientTracerNoop() *httptrace.ClientTrace {
	c := &clientTracerNoop{}
	return &httptrace.ClientTrace{
		GetConn:              c.GetConn,
		GotConn:              c.GotConn,
		PutIdleConn:          c.PutIdleConn,
		GotFirstResponseByte: c.GotFirstResponseByte,
		Got100Continue:       c.Got100Continue,
		Got1xxResponse:       c.Got1xxResponse,
		DNSStart:             c.DNSStart,
		DNSDone:              c.DNSDone,
		ConnectStart:         c.ConnectStart,
		ConnectDone:          c.ConnectDone,
		TLSHandshakeStart:    c.TLSHandshakeStart,
		TLSHandshakeDone:     c.TLSHandshakeDone,
		WroteHeaderField:     c.WroteHeaderField,
		WroteHeaders:         c.WroteHeaders,
		Wait100Continue:      c.Wait100Continue,
		WroteRequest:         c.WroteRequest,
	}
}

// GetConn is called before a connection is created or
// retrieved from an idle pool. The hostPort is the
// "host:port" of the target or proxy. GetConn is called even
// if there's already an idle cached connection available.
// ff:
// hostPort:
func (*clientTracerNoop) GetConn(hostPort string) {}

// GotConn is called after a successful connection is
// obtained. There is no hook for failure to obtain a
// connection; instead, use the error from
// Transport.RoundTrip.
// ff:
func (*clientTracerNoop) GotConn(httptrace.GotConnInfo) {}

// PutIdleConn is called when the connection is returned to
// the idle pool. If err is nil, the connection was
// successfully returned to the idle pool. If err is non-nil,
// it describes why not. PutIdleConn is not called if
// connection reuse is disabled via Transport.DisableKeepAlives.
// PutIdleConn is called before the caller's Response.Body.Close
// call returns.
// For HTTP/2, this hook is not currently used.
// ff:
// err:
func (*clientTracerNoop) PutIdleConn(err error) {}

// GotFirstResponseByte is called when the first byte of the response
// headers is available.
// ff:
func (*clientTracerNoop) GotFirstResponseByte() {}

// Got100Continue is called if the server replies with a "100
// Continue" response.
// ff:
func (*clientTracerNoop) Got100Continue() {}

// Got1xxResponse is called for each 1xx informational response header
// returned before the final non-1xx response. Got1xxResponse is called
// for "100 Continue" responses, even if Got100Continue is also defined.
// If it returns an error, the client request is aborted with that error value.
// ff:
// code:
// header:
func (*clientTracerNoop) Got1xxResponse(code int, header textproto.MIMEHeader) error {
	return nil
}

// DNSStart is called when a DNS lookup begins.
// ff:
func (*clientTracerNoop) DNSStart(httptrace.DNSStartInfo) {}

// DNSDone is called when a DNS lookup ends.
// ff:
func (*clientTracerNoop) DNSDone(httptrace.DNSDoneInfo) {}

// ConnectStart is called when a new connection's Dial begins.
// If net.Dialer.DualStack (IPv6 "Happy Eyeballs") support is
// enabled, this may be called multiple times.
// ff:
// network:
// addr:
func (*clientTracerNoop) ConnectStart(network, addr string) {}

// ConnectDone is called when a new connection's Dial
// completes. The provided err indicates whether the
// connection completed successfully.
// If net.Dialer.DualStack ("Happy Eyeballs") support is
// enabled, this may be called multiple times.
// ff:
// network:
// addr:
// err:
func (*clientTracerNoop) ConnectDone(network, addr string, err error) {}

// TLSHandshakeStart is called when the TLS handshake is started. When
// connecting to an HTTPS site via an HTTP proxy, the handshake happens
// after the CONNECT request is processed by the proxy.
// ff:
func (*clientTracerNoop) TLSHandshakeStart() {}

// TLSHandshakeDone is called after the TLS handshake with either the
// successful handshake's connection state, or a non-nil error on handshake
// failure.
// ff:
func (*clientTracerNoop) TLSHandshakeDone(tls.ConnectionState, error) {}

// WroteHeaderField is called after the Transport has written
// each request header. At the time of this call the values
// might be buffered and not yet written to the network.
// ff:
// key:
// value:
func (*clientTracerNoop) WroteHeaderField(key string, value []string) {}

// WroteHeaders is called after the Transport has written
// all request headers.
// ff:
func (*clientTracerNoop) WroteHeaders() {}

// Wait100Continue is called if the Request specified
// "Expect: 100-continue" and the Transport has written the
// request headers but is waiting for "100 Continue" from the
// server before writing the request body.
// ff:
func (*clientTracerNoop) Wait100Continue() {}

// WroteRequest is called with the result of writing the
// request and any body. It may be called multiple times
// in the case of retried requests.
// ff:
func (*clientTracerNoop) WroteRequest(httptrace.WroteRequestInfo) {}
