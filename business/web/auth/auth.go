package auth

import (
	"errors"
	"sync"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

// ErrForbidden is returned when a auth issue is identified.
var ErrForbidden = errors.New("attempted action is not allowed")

// KeyLookup declares a method set of behavior for looking up
// private and public keys for JWT use.
type KeyLookup interface {
	PrivateKeyPEM(kid string) (pem string, err error)
	PublicKeyPEM(kid string) (pem string, err error)
}

// Config represents information required to initialize auth.
type Config struct {
	Log       *zap.SugaredLogger
	KeyLookup KeyLookup
}

// Auth is used to authenticate clients. It can generate a token for a
// set of user claims and recreate the claims by parsing the token.
type Auth struct {
	log       *zap.SugaredLogger
	keyLookup KeyLookup
	method    jwt.SigningMethod
	parser    *jwt.Parser
	mu        sync.RWMutex
	cache     map[string]string
}

// New creates an Auth to support authentication/authorization.
func New(cfg Config) (*Auth, error) {
	a := Auth{
		log:       cfg.Log,
		keyLookup: cfg.KeyLookup,
		method:    jwt.GetSigningMethod("RS256"),
		parser:    jwt.NewParser(jwt.WithValidMethods([]string{"RS256"})),
		cache:     make(map[string]string),
	}
	return &a, nil
}
