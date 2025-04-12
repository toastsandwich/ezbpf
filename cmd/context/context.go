package context

import "fmt"

type Context map[string]any

func EmptyContext() Context {
	ctx := make(Context)
	return ctx
}

func (c Context) Put(k string, v any) {
	c[k] = v
}

func (c Context) Get(k string) any {
	v, ok := c[k]
	if !ok {
		panic(fmt.Sprintf("context %s was not able to find value for key %s", c, k))
	}
	return v
}

func ContextWithValue(k string, v any) Context {
	ctx := EmptyContext()
	ctx.Put(k, v)
	return ctx
}
