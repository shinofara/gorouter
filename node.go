package gorouter

import (
	"reflect"
)

type tree struct {
	methods map[string]*node
}

func NewTree() *tree {
	return &tree{
		methods: make(map[string]*node),
	}
}

type node struct {
	path map[string]*Handler
}

func NewNode() *node {
	return &node{
		path: make(map[string]*Handler),
	}
}

type Handler struct {
	Value reflect.Value
	Type reflect.Type
}

func NewHandler(h Handle) *Handler {
	v := reflect.ValueOf(h)
	t := reflect.TypeOf(h)
	return &Handler{
		Value: v,
		Type: t,
	}
}

func (t *tree) Add(method, path string, handle Handle) {

	if t.methods[method] == nil {
		t.methods[method] = NewNode()
	}

	t.methods[method].path[path] = NewHandler(handle)
}

func (t *tree) GetHandler(method, path string) *Handler {
	m := t.methods[method]
	return m.path[path]
}
