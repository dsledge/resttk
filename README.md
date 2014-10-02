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
###Filter Example:
###Session Example:
###Resource Example:
###Authentication Example:



