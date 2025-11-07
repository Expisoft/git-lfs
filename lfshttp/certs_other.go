// certs_other.go
//go:build !windows

package lfshttp

import (
	"crypto/tls"

	"github.com/git-lfs/git-lfs/v3/errors"
)

// Just a stub on other os:es
func getClientCertForHostFromSchannel(c *Client, host string) (*tls.Certificate, error) {
	return nil, errors.New("Function not supported")
}
