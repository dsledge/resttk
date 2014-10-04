resttk
======

REST Toolkit for Go

The rest toolkit provides a simple way to bring up a restful web service in Go. It 
provides a base controller allowing you to override the method protocols you wish
to implement. We also provide a router which supports filters allowing you route
request to contollers, chaining filters which get applied before reaching the
controller logic.

###Running Test:
```bash
go test
```

###Basic Controller Example:
```go
package main

import (
  "github.com/dsledge/resttk"
  "fmt"
)

//implementing controller
type HelloController struct {
  resttk.BaseController
}

//quick method to return an instance of the implementing controller
func NewHelloController() resttk.ControllerInterface {
  return &HelloController{}
}

//implementing the http GET method handler
func (c *HelloController) Get() {
  vars := c.ParsePath()
  
  name, ok := vars["name"]
  if ok {
    c.SendResponse(fmt.Sprintf("Hello %s!", name))
    return
  }
  c.SendResponse("Hello world!")
}

func main() {
  //make a resttk router and adding a route to the implementing controller instance
  //first route has no parameters, second route has a paramater that will be place 
  //in a map with the key "name"
  router := resttk.NewRouter()
  router.AddRoute("/api/v1/hello", NewHelloController)
  router.AddRoute("/api/v1/hello/:name", NewHelloController)

  //create and run the resttk server
  resttk := &resttk.Server{SSL: false, Addr: 127.0.0.1, Port: 3000, Routes: router}
  fmt.Printf("Starting the resttk server at http://%s:%s", resttk.Addr, resttk.Port)
  if err := resttk.Run(); err != nil {
    panic(err)
  }
}
```
Calling **GET** method, no parameters
```bash
-bash-4.2# curl -v http://127.0.0.1:3000/api/v1/hello
* About to connect() to 127.0.0.1 port 3000 (#0)
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 3000 (#0)
> GET /api/v1/hello HTTP/1.1
> User-Agent: curl/7.29.0
> Host: 127.0.0.1:3000
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Sat, 04 Oct 2014 05:22:34 GMT
< Content-Length: 15
< 
"Hello world!"
* Connection #0 to host 127.0.0.1 left intact
```
Calling **GET** method, with parameters
```bash
-bash-4.2# curl -v http://127.0.0.1:3000/api/v1/hello/Jerry
* About to connect() to 127.0.0.1 port 3000 (#0)
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 3000 (#0)
> GET /api/v1/hello/Jerry HTTP/1.1
> User-Agent: curl/7.29.0
> Host: 127.0.0.1:3000
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Sat, 04 Oct 2014 05:12:25 GMT
< Content-Length: 15
< 
"Hello Jerry!"
* Connection #0 to host 127.0.0.1 left intact
```
Calling **POST** method, not implemented
```bash
-bash-4.2# curl -vd "Test Data" http://127.0.0.1:3000/api/v1/hello
* About to connect() to 127.0.0.1 port 3000 (#0)
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 3000 (#0)
> POST /api/v1/hello HTTP/1.1
> User-Agent: curl/7.29.0
> Host: 127.0.0.1:3000
> Accept: */*
> Content-Length: 9
> Content-Type: application/x-www-form-urlencoded
> 
* upload completely sent off: 9 out of 9 bytes
< HTTP/1.1 501 Not Implemented
< Date: Sat, 04 Oct 2014 05:18:24 GMT
< Content-Length: 0
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host 127.0.0.1 left intact
```

###Filter Example:
```go
package main

import (
  "github.com/dsledge/resttk"
  "fmt"
)

//implementing controller
type HelloController struct {
  resttk.BaseController
}

//quick method to return an instance of the implementing controller
func NewHelloController() resttk.ControllerInterface {
  return &HelloController{}
}

//implementing the http GET method handler
func (c *HelloController) Get() {
  vars := c.ParsePath()
  
  name, ok := vars["name"]
  if ok {
    c.SendResponse(fmt.Sprintf("Hello %s!", name))
    return
  }
  c.SendResponse("Hello world!")
}

func main() {
  //Function for that looks for a TOKEN header
  tokenFunc := func(w http.ResponseWriter, r *http.Request) bool {
    if r.Header.Get("TOKEN") != "1234" {
      w.WriteHeader(http.StatusUnauthorized)
      return false
    }
    return true
  }

  //Filter that requires the TOKEN to be valid
  tokenFilter := &resttk.Filter{Name: "TokenFilter", Apply: tokenFunc}

  //make a resttk router and adding a route to the implementing controller instance
  //first route has no parameters, second route has a paramater that will be place 
  //in a map with the key "name"
  router := resttk.NewRouter()
  router.AddRoute("/api/v1/hello", NewHelloController).ApplyFilter(tokenFilter)

  //create and run the resttk server
  resttk := &resttk.Server{SSL: false, Addr: 127.0.0.1, Port: 3000, Routes: router}
  fmt.Printf("Starting the resttk server at http://%s:%s", resttk.Addr, resttk.Port)
  if err := resttk.Run(); err != nil {
    panic(err)
  }
}
```
Calling **GET** method, missing required header
```bash
-bash-4.2# curl -v http://127.0.0.1:3000/api/v1/hello
* About to connect() to 127.0.0.1 port 3000 (#0)
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 3000 (#0)
> GET /api/v1/hello HTTP/1.1
> User-Agent: curl/7.29.0
> Host: 127.0.0.1:3000
> Accept: */*
> 
< HTTP/1.1 401 Unauthorized
< Date: Sat, 04 Oct 2014 06:53:51 GMT
< Content-Length: 0
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host 127.0.0.1 left intact
```
Calling **GET** method, with required header
```bash
-bash-4.2# curl -vH "TOKEN:1234" http://127.0.0.1:3000/api/v1/hello                                         
* About to connect() to 127.0.0.1 port 3000 (#0)
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 3000 (#0)
> GET /api/v1/hello HTTP/1.1
> User-Agent: curl/7.29.0
> Host: 127.0.0.1:3000
> Accept: */*
> TOKEN:1234
> 
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Sat, 04 Oct 2014 06:54:27 GMT
< Content-Length: 15
< 
"Hello world!"
* Connection #0 to host 127.0.0.1 left intact
```
###Session Example:
###Resource Example:



