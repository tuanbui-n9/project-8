package cookies

import "github.com/gorilla/securecookie"

type Cookies interface {
	Encode(name string, value string) (string, error)
	Decode(name string, value string) (string, error)
}

type SecureCookie struct {
	securecookie *securecookie.SecureCookie
}

// NewSecureCookie creates a new SecureCookie instance
func NewSecureCookie(secret string) (Cookies, error) {
	securecookie := securecookie.New([]byte(secret), nil)

	return &SecureCookie{
		securecookie: securecookie,
	}, nil
}

func (sc *SecureCookie) Encode(name string, value string) (string, error) {
	return sc.securecookie.Encode(name, value)
}

func (sc *SecureCookie) Decode(name string, value string) (string, error) {
	var decodedValue string
	err := sc.securecookie.Decode(name, value, &decodedValue)
	return decodedValue, err
}
