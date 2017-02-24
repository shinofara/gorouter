package router

import (
	"reflect"
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/schema"
)

const (
	ArgWriter = 0
	ArgRequest = 1
	ArgParams = 2
	NumArgs   = 3
)

type Handle interface{}
//func(http.ResponseWriter, *http.Request, interface{}) error
//type Handle func(http.ResponseWriter, *http.Request)

type Router struct {
	nodes map[string]*Node
}

func New() *Router {
	return &Router{}
}

// GET is a shortcut for router.Handle("GET", path, handle)
func (r *Router) GET(path string, handle Handle) {
	r.Handle("GET", path, handle)
}


func (r *Router) Handle(method, path string, handle Handle) {

	// Validate arguments
	ft := reflect.TypeOf(handle)
	if ft.NumIn() != NumArgs {
		name := reflect.ValueOf(handle)
		panic(fmt.Sprintf("handler %sには%dつ存在すべき", name, NumArgs))
	}

	//第三引数の検証
	///ptrである事
	arg := ft.In(ArgParams)
	if arg.Kind().String() != "struct" {
		panic(fmt.Sprintf("Absolute third argument is struct. But it is %s.", arg.Kind().String()))
	}

	// Validate the fields of the structure.
	//for i :=0; i <  atg.NumField(); i++ {
	//	field := arg.Field(i)
	//}

	if r.nodes == nil {
		r.nodes = make(map[string]*Node)
	}
	
	node := new(Node)
	r.nodes[method] = node
	node.addRoute(path, handle)
}

// http servewrの為に必要
// todo: メモコメント
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	
	if node := r.nodes[req.Method]; node != nil {
		if h := node.GetHandler(path); h != nil {
			handler(w, req, h)
			return
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request, h *Handler) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	//登録時に検証して生成しておくと、次のセッションまで引き継がれる為、handler実行時に評価される様にしておく
	rvp := reflect.New(h.Type.In(2))
	schema.NewDecoder().Decode(rvp.Interface(), r.Form)

	err := h.Value.Call([]reflect.Value{
		reflect.ValueOf(w),
		reflect.ValueOf(r),
		rvp.Elem(),
	})
	
	if err != nil {
		log.Printf("%+V", err)
	}
}
