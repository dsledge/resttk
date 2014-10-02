package resttk

import (
	"net/http"
)

func Cookie(name, value, path, domain string, maxage int, secure, httponly bool) *http.Cookie {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		Domain:   domain,
		MaxAge:   maxage,
		Secure:   secure,
		HttpOnly: httponly,
	}

	return cookie
}
