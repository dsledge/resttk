package resttk

import (
	"bytes"
	"regexp"
	"strings"
)

type Route struct {
	path    string
	handler func() ControllerInterface
	regex   *regexp.Regexp
	filters []*Filter
}

func NewRoute(path string, controller func() ControllerInterface) *Route {
	route := &Route{path: path, handler: controller}
	route.buildRegex(path)
	return route
}

func (r *Route) ApplyFilter(filter *Filter) *Route {
	r.filters = append(r.filters, filter)
	return r
}

func (r *Route) buildRegex(path string) {
	var buffer bytes.Buffer
	buffer.WriteString("^")

	parts := strings.Split(path, "/")
	for _, part := range parts {
		if strings.Contains(part, ":") {
			buffer.WriteString("[A-Za-z0-9_@\\.\\-]+")
		} else {
			buffer.WriteString(part)
		}
		buffer.WriteString("/")
	}

	str := strings.TrimSuffix(buffer.String(), "/")
	r.regex = regexp.MustCompile(str + "$")
}
