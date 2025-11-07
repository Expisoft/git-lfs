// certs_other.go
//go:build !windows

package lfshttp

import (
	"crypto/tls"
)

// Just a stub on other os:es
func getClientCertForHostFromSchannel(c *Client, host string) (*tls.Certificate, error) {
	return nil, err
}
