package router

import (
	"encoding/json"
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

func (ctx *Context) WriteError(status int, err string) {
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter.WriteHeader(status)
	data, _ := json.Marshal(struct {
		Error string
	}{
		Error: err,
	})
	ctx.ResponseWriter.Write(data)
}
