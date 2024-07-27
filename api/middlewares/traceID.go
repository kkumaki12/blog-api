package middlewares

import (
	"context"
	"sync"
)

var (
	logNo int = 1
	mu    sync.Mutex
)

type traceIDKey struct{}

func newTraceID() int {
	var no int

	mu.Lock()
	no = logNo
	logNo++
	mu.Unlock()

	return no
}

func SetTraceID(cxt context.Context, traceID int) context.Context {
	return context.WithValue(cxt, traceIDKey{}, traceID)
}

func GetTraceID(cxt context.Context) int {
	id := cxt.Value(traceIDKey{})

	if idInt, ok := id.(int); ok {
		return idInt
	}
	return 0
}
