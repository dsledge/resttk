package resttk

import (
	"net/http"
	"testing"
)

func TestAppendRoute(t *testing.T) {
	router := &Router{}
	router.AddRoute("/api/v1/test", func() ControllerInterface { return new(BaseController) })

	// request := newRequest("GET", "http://localhost/api/v1/test?q=dotnet")
	// route := router.findRoute(request)
}

func newRequest(method, url string) *http.Request {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	return req
}
