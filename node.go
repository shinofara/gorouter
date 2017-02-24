package gorouter

import (
	"reflect"
)

type Node struct {
	path map[string]*Handler
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

func (n *Node) addRoute(path string, handle Handle) {

	if n.path == nil {
		n.path = make(map[string]*Handler)
	}	
	n.path[path] = NewHandler(handle)
}

func (n *Node) GetHandler(path string) *Handler {
	return n.path[path]
}
