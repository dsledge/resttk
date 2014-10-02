package resttk

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testMockController struct {
	BaseController
}

func (c *testMockController) Get() {
	c.SendStatus(201)
}

func (c *testMockController) Post() {
	c.SendStatus(202)
}

func (c *testMockController) Put() {
	c.SendStatus(203)
}

func (c *testMockController) Head() {
	c.SendStatus(204)
}

func (c *testMockController) Options() {
	// c.SendStatus(205)
}

func (c *testMockController) Patch() {
	c.SendStatus(206)
}

func (c *testMockController) Delete() {
	c.SendStatus(207)
}

func TestGet(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	c := &testMockController{}
	c.Init(rec, req, c, "/foo")
	c.ServeHTTP(rec, req)
	if rec.Code != 201 {
		t.Errorf("Got Response Code %i, expected 201", rec.Code)
	}
}

func TestPost(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://example.com/foo", nil)
	c := &testMockController{}
	c.Init(rec, req, c, "/foo")
	c.ServeHTTP(rec, req)
	if rec.Code != 202 {
		t.Errorf("Got Response Code %i, expected 202", rec.Code)
	}
}

func TestPut(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "http://example.com/foo", nil)
	c := &testMockController{}
	c.Init(rec, req, c, "/foo")
	c.ServeHTTP(rec, req)
	if rec.Code != 203 {
		t.Errorf("Got Response Code %i, expected 203", rec.Code)
	}
}

func TestHead(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("HEAD", "http://example.com/foo", nil)
	c := &testMockController{}
	c.Init(rec, req, c, "/foo")
	c.ServeHTTP(rec, req)
	if rec.Code != 204 {
		t.Errorf("Got Response Code %i, expected 204", rec.Code)
	}
}

func TestOptions(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", "http://example.com/foo", nil)
	c := &testMockController{}
	c.Init(rec, req, c, "/foo")
	c.ServeHTTP(rec, req)
	if rec.Code != 200 {
		t.Errorf("Got Response Code %i, expected 205", rec.Code)
	}
}

func TestPatch(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "http://example.com/foo", nil)
	c := &testMockController{}
	c.Init(rec, req, c, "/foo")
	c.ServeHTTP(rec, req)
	if rec.Code != 206 {
		t.Errorf("Got Response Code %i, expected 206", rec.Code)
	}
}

func TestDelete(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "http://example.com/foo", nil)
	c := &testMockController{}
	c.Init(rec, req, c, "/foo")
	c.ServeHTTP(rec, req)
	if rec.Code != 207 {
		t.Errorf("Got Response Code %i, expected 207", rec.Code)
	}
}

type testEncodingController struct {
	BaseController
}

type testEncodingStruct struct {
	Field1 int
	Field2 string
}

func (c *testEncodingController) Post() {
	s := &testEncodingStruct{2, "fieldvalue"}
	fmt.Printf("S is %q\n", s)
	c.SendJSON(s)
}

func (c *testEncodingController) Put() {
	s := &testEncodingStruct{2, "fieldvalue"}
	c.SendXML(s)
}

func TestUseJSON(t *testing.T) {
	t.Skip()
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://example.com/foo", nil)
	c := &testEncodingController{}
	c.Init(rec, req, c, "/foo")
	c.ServeHTTP(rec, req)
	dec := json.NewDecoder(rec.Body)
	s := &testEncodingStruct{}
	if err := dec.Decode(s); err != nil {
		t.Error(err.Error())
	}
	if s.Field1 != 2 {
		t.Errorf("Got field value %i, expected 2", s.Field1)
	}
	if s.Field2 != "fieldvalue" {
		t.Errorf("Got field value %s, expected \"fieldvalue\"", s.Field1)
	}
}

func TestUseXML(t *testing.T) {
	t.Skip()
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "http://example.com/foo", nil)
	c := &testEncodingController{}
	c.Init(rec, req, c, "/foo")
	c.ServeHTTP(rec, req)
	dec := xml.NewDecoder(rec.Body)
	s := &testEncodingStruct{}
	dec.Decode(s)
	if s.Field1 != 2 {
		t.Errorf("Got field value %i, expected 2", s.Field1)
	}
	if s.Field2 != "fieldvalue" {
		t.Errorf("Got field value %s, expected \"fieldvalue\"", s.Field1)
	}
}

type testMockController1 struct {
	BaseController
}

func (c *testMockController1) Options() {
	// c.SendStatus(205)
	return
}

func TestGet1(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	c := &testMockController1{}
	c.Init(rec, req, c, "/foo")
	c.ServeHTTP(rec, req)
	if rec.Code != 501 {
		t.Errorf("Got Response Code %i, expected 501", rec.Code)
	}
}
