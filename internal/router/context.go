package router

import (
	"net/http"
)

// Context ..
type Context struct {
	http.ResponseWriter
	*http.Request
	Params map[string]string
}

// setURLValues ..
func (ctx *Context) setURLValues(keys, values []string) {
	for i, key := range keys {
		ctx.SetParam(key, values[i])
	}
}

// SetParam ..
func (ctx *Context) SetParam(key, value string) {
	ctx.Params[key] = value
}
