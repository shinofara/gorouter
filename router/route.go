package router

import (
	"reflect"
	"net/http"
//	"github.com/gorilla/schema"
	"log"
	"fmt"
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

	//handleの検証
	log.Print("handlerの追加を開始")
	log.Print("handlerの引数の数を検証開始")

	//handlerの引数検証
	ft := reflect.TypeOf(handle)
	if ft.NumIn() != NumArgs {
		name := reflect.ValueOf(handle)
		panic(fmt.Sprintf("handler %sには%dつ存在すべき", name, NumArgs))
	}
	log.Print(ft.NumIn())

	//第三引数の検証
	///ptrである事
	arg := ft.In(ArgParams)
	if arg.Kind().String() != "ptr" {
		panic(fmt.Sprintf("第三引数がptrではなく、%s", arg.Kind().String()))
	}
	
	//構造体内のフィールドのタグチェック
	log.Print("handlerのフィールドタグ検証を開始")
	log.Print(&arg)

	elm := arg.Elem()
	for i :=0; i <  elm.NumField(); i++ {
		field := elm.Field(i)
		log.Print(field.Tag.Get("schema"))
	}

	if r.nodes == nil {
		r.nodes = make(map[string]*Node)
	}
	
	p := new(Node)
	r.nodes[method] = p
	p.addRoute(path, handle)
}

// http servewrの為に必要
// todo: メモコメント
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	
	if node := r.nodes[req.Method]; node != nil {
		if handle := node.GetHandle(path); handle != nil {
			log.Print(handle)
			handler(w, req, handle)
			return
		}
	}
}

func handler(w http.ResponseWriter, req *http.Request, handler Handle) {

	fv := reflect.ValueOf(handler)
	ft := reflect.TypeOf(handler)

	//第3引数の型を取得
	arg3 := ft.In(2)
	log.Print(arg3.Kind().String())

	//これはここじゃなくてregistのときに検証
	fuga := struct{
		Hoge string
	}{"test"}
	
	//fuga構造体をhoge構造体にcast
	hoge := reflect.ValueOf(&fuga).Convert(arg3)
	log.Print(hoge)
	arg1 := reflect.ValueOf(w)
	arg2 := reflect.ValueOf(req)
	
	result := fv.Call([]reflect.Value{arg1, arg2, hoge}) //[]reflect.Value ←戻り値も複数返せる為、スライスとなっている
	log.Print(result[0])
	log.Print("handlerの実行完了")
}
