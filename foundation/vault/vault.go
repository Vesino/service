// Package vault provides support for accessing Hashicorp's vault service
// to access private keys.
package vault

import (
	"net"
	"net/http"
	"sync"
	"time"
)

// This provides a default client configuration, but it's recommended
// this is replaced by the user with application specific settings using
// the WithClient function at the time a GraphQL is constructed.
var defaultClient = http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          1,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
}

// Config represents the mandatory settings needed to work with Vault.
type Config struct {
	Address   string
	MountPath string
	Token     string
	Client    *http.Client
}

// Vault provides support to access Hashicorp's Vault product for keys.
type Vault struct {
	address   string
	token     string
	mountPath string
	client    *http.Client
	mu        sync.RWMutex
	store     map[string]string
}

// New constructs a vault for use.
func New(cfg Config) (*Vault, error) {
	if cfg.Client == nil {
		cfg.Client = &defaultClient
	}
	return &Vault{
		address:   cfg.Address,
		token:     cfg.Token,
		mountPath: cfg.MountPath,
		client:    cfg.Client,
		store:     make(map[string]string),
	}, nil
}

// SetToken allows the user to change out the token to use on calls.
func (v *Vault) SetToken(token string) {
	v.token = token
}
