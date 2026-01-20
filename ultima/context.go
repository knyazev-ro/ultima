package ultima

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	w http.ResponseWriter
	r *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		w: w,
		r: r,
	}
}

func (c *Context) Json(i any, status int) *Context {
	c.w.Header().Set("Content-Type", "application/json")
	c.w.WriteHeader(status)
	json.NewEncoder(c.w).Encode(i)
	return c
}
