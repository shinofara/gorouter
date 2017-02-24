package gorouter

import (
	"fmt"
	"net/http"
	"reflect"
	"log"
	"github.com/gorilla/schema"	
)

type Mux struct {
	// The radix trie router
	tree *tree
}

const (
	ArgWriter = 0
	ArgRequest = 1
	ArgParams = 2
	NumArgs   = 3
)


func NewMux() *Mux {
	mux := &Mux{tree: NewTree()}
	return mux
}

// GET is a shortcut for router.Handle("GET", path, handle)
func (mx *Mux) GET(path string, handle Handle) {
	mx.handle("GET", path, handle)
}

func (mx *Mux) handle(method, pattern string, h Handle) {
	// Validate arguments
	ft := reflect.TypeOf(h)
	if ft.NumIn() != NumArgs {
		name := reflect.ValueOf(h)
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

//	if r.nodes == nil {
//		r.nodes = make(map[string]*Node)
//	}
	
	//node := new(Node)
	//r.nodes[method] = node
	//node.addRoute(path, handle)

	
	mx.tree.Add(method, pattern, h)
}


// http servewrの為に必要
// todo: メモコメント
func (mx *Mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	
	if h := mx.tree.GetHandler(req.Method, path); h != nil {
		handler(w, req, h)
		return
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
