package resttk

import (
	"net/http"
)

type Filter struct {
	Name  string
	Apply func(http.ResponseWriter, *http.Request) bool
}
