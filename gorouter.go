package gorouter


type Handle interface{}
//func(http.ResponseWriter, *http.Request, interface{}) error
//type Handle func(http.ResponseWriter, *http.Request)

func New() *Mux {
	return NewMux()
}
