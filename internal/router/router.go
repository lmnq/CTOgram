package router

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

// Handler ..
type Handler func(*Context)

// Route ..
type Route struct {
	Method  string
	Pattern *regexp.Regexp
	Handler Handler
	Keys    []string
}

// Router ..
type Router struct {
	Routes       []Route
}

// NewRouter ..
func NewRouter() *Router {
	return &Router{
		// DefaultRoute: handlers.ErrorHandler(),
	}
}

// addRoute ..
func (r *Router) addRoute(method, pattern string, handler Handler) {
	re, keys := readPatternAndKeys(pattern)
	route := Route{Method: method, Pattern: re, Handler: handler, Keys: keys}
	r.Routes = append(r.Routes, route)
}

func readPatternAndKeys(pattern string) (*regexp.Regexp, []string) {
	var keys []string
	segments := strings.Split(pattern, "/")

	for i, v := range segments {
		switch {
		case strings.HasPrefix(v, ":"):
			keys = append(keys, v[1:])
			segments[i] = `([0-9]+)`
		}
	}

	path := fmt.Sprintf("^%s$", strings.Join(segments, "/"))

	return regexp.MustCompile(path), keys
}

// GET ..
func (r *Router) GET(pattern string, handler Handler) {
	r.addRoute("GET", pattern, handler)
}

// POST ..
func (r *Router) POST(pattern string, handler Handler) {
	r.addRoute("POST", pattern, handler)
}

// PUT ..
func (r *Router) PUT(pattern string, handler Handler) {
	r.addRoute("PUT", pattern, handler)
}

// DELETE ..
func (r *Router) DELETE(pattern string, handler Handler) {
	r.addRoute("DELETE", pattern, handler)
}

// Serve ..
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := &Context{
		ResponseWriter: w,
		Request:        req,
		Params:         make(map[string]string),
	}
	var allow []string
	for _, route := range r.Routes {
		matches := route.Pattern.FindStringSubmatch(req.URL.Path)
		if len(matches) > 0 {
			if req.Method != route.Method {
				allow = append(allow, route.Method)
				continue
			}

			if len(matches) > 1 && len(route.Keys) == len(matches[1:]) {
				ctx.setURLValues(route.Keys, matches[1:])
			}

			route.Handler(ctx)
			return
		}
	}

	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.NotFound(w, req)
}
