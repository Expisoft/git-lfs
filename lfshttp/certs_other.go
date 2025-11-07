// certs_other.go
//go:build !windows

package lfshttp

// Just a stub on other os:es
func getClientCertForHostFromSchannel(c *Client, host string) (*tls.Certificate, error) {
}
