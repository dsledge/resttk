package resttk

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"
)

type ControllerInterface interface {
	Init(w http.ResponseWriter, r *http.Request, p ControllerInterface, path string)
	ParsePath() map[string]string
	Head()
	Options()
	Get()
	Put()
	Post()
	Patch()
	Delete()
	Websocket()
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	SendJSON(obj interface{})
	SendXML(obj interface{})
	SendResponse(obj interface{})
	SetHeader(key, value string)
	GetHeader(key string) string
	GetBodyJSON(obj interface{}) interface{}
	GetBodyXML(obj interface{}) interface{}
	GetBody(obj interface{}) interface{}
	SendStatus(code int)
}

type BaseController struct {
	parent   ControllerInterface
	path     string
	Request  *http.Request
	Response http.ResponseWriter
}

func (c *BaseController) Init(w http.ResponseWriter, r *http.Request, p ControllerInterface, path string) {
	c.parent = p
	c.path = path
	c.Request = r
	c.Response = w
}

func (c *BaseController) ParsePath() map[string]string {
	vars := make(map[string]string)
	index := 0
	value := strings.Split(c.Request.URL.Path, "/")
	parts := strings.Split(c.path, "/")

	for _, part := range parts {
		if strings.Contains(part, ":") {
			vars[strings.TrimPrefix(part, ":")] = value[index]
		}
		index += 1
	}

	return vars
}

func (c *BaseController) Head() {
}

func (c *BaseController) Options() {
}

func (c *BaseController) Get() {
	c.SendStatus(http.StatusNotImplemented)
}

func (c *BaseController) Put() {
	c.SendStatus(http.StatusNotImplemented)
}

func (c *BaseController) Post() {
	c.SendStatus(http.StatusNotImplemented)
}

func (c *BaseController) Patch() {
	c.SendStatus(http.StatusNotImplemented)
}

func (c *BaseController) Delete() {
	c.SendStatus(http.StatusNotImplemented)
}

func (c *BaseController) Websocket() {
	c.parent.Websocket()
}

func (c *BaseController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Upgrade") == "websocket" {
		c.Websocket()
	} else {
		switch r.Method {
		case "HEAD":
			c.parent.Head()
		case "OPTIONS":
			c.parent.Options()
		case "GET":
			c.parent.Options()
			c.parent.Get()
		case "PUT":
			c.parent.Options()
			c.parent.Put()
		case "POST":
			c.parent.Options()
			c.parent.Post()
		case "PATCH":
			c.parent.Options()
			c.parent.Patch()
		case "DELETE":
			c.parent.Options()
			c.parent.Delete()
		default:
			c.parent.SendStatus(http.StatusMethodNotAllowed)
		}
	}
}

func (c *BaseController) SendJSON(obj interface{}) {
	c.Response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Response).Encode(obj)
}

func (c *BaseController) SendXML(obj interface{}) {
	c.Response.Header().Set("Content-Type", "application/xml")
	xml.NewEncoder(c.Response).Encode(obj)
}

func (c *BaseController) SendResponse(obj interface{}) {
	switch c.Request.Header.Get("Content-Type") {
	case "application/json":
		c.SendJSON(obj)
	case "application/xml":
		c.SendXML(obj)
	default:
		c.SendJSON(obj)
	}
}

func (c *BaseController) SetHeader(key, value string) {
	c.Response.Header().Set(key, value)
}

func (c *BaseController) GetHeader(key string) string {
	return c.Request.Header.Get(key)
}

func (c *BaseController) GetBodyJSON(obj interface{}) interface{} {
	if c.Request.Body != nil {
		if err := json.NewDecoder(c.Request.Body).Decode(obj); err != nil {
			c.SendStatus(http.StatusBadRequest)
		}
		return obj
	}
	c.SendStatus(http.StatusBadRequest)
	return nil
}

func (c *BaseController) GetBodyXML(obj interface{}) interface{} {
	if c.Request.Body != nil {
		if err := xml.NewDecoder(c.Request.Body).Decode(obj); err != nil {
			c.SendStatus(http.StatusBadRequest)
		}
		return obj
	}
	c.SendStatus(http.StatusBadRequest)
	return nil
}

func (c *BaseController) GetBody(obj interface{}) interface{} {
	if c.Request.Body != nil {
		switch c.Request.Header.Get("Content-Type") {
		case "application/json":
			return c.GetBodyJSON(obj)
		case "application/xml":
			return c.GetBodyXML(obj)
		default:
			return c.GetBodyJSON(obj)
		}
	}
	c.SendStatus(http.StatusBadRequest)
	return nil
}

func (c *BaseController) SendStatus(code int) {
	c.Response.WriteHeader(code)
}
