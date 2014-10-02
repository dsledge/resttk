package resttk

import (
	"testing"
	// "net/http"
	// "net/url"
	// "log"
	// "fmt"
)

func TestRoute(t *testing.T) {
	// t.Skip()
	// s := "http://bing.com/search/test/me?q=dotnet"
	// url, err := url.Parse(s)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// route := NewRoute(url.Path, &Controller{})
	// fmt.Printf("REGEX: %s\n", route.regex)
	// fmt.Printf("ROUTE: %v\n", route)
	// fmt.Printf("HANDLER: %v\n", route.handler)
	// route := NewRoute("/api/v1/", &BaseController{})
	// fmt.Printf("REGEX: %s\n", route.regex)

	// route2 := NewRoute("/api/v1/test1", &BaseController{})
	// fmt.Printf("REGEX: %s\n", route2.regex)

	// route3 := NewRoute("/api/v1/test2", &BaseController{})
	// fmt.Printf("REGEX: %s\n", route3.regex)

	// route4 := NewRoute("/api/v1/test3/:id", &BaseController{})
	// fmt.Printf("REGEX: %s\n", route4.regex)

	// handler := func() interface{} {
	// 	return new(BaseController)
	// }

	// NewRoute("/api/v1/", handler)

	// 	authFunc := func(w http.ResponseWriter, r *http.Request) bool {
	// 	_, ok := resources.SessionStore.Get(r.Header.Get("AUTH-TOKEN"))
	// 	if !ok {
	// 		scribble.Warn("AUTH-TOKEN not found, redirecting back to login")
	// 		http.Redirect(w, r, "/", http.StatusUnauthorized)
	// 		return false
	// 	}
	// 	return true
	// }

	// fmt.Printf("URL: %s\n", s)
	// fmt.Printf("ROUTE: %v\n", route)
	// fmt.Printf("REGEX: %s\n", route.regex)
	// fmt.Printf("HANDLER: %v\n", route.handler)
}

// func testHandler(http.ResponseWriter, *http.Request) { fmt.Println("TEST HANDLER FUNCTION") }
