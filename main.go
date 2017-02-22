package main

import (
	"fmt"
  "net/http"

	"github.com/shinofara/router/router"
	"log"
)

type Param struct {
	Hoge string `schema:"hoge"`
}

func viewHandler(w http.ResponseWriter, r *http.Request, params *Param) error {
	fmt.Fprintf(w, "Hello World")
	return nil
}

func main() {
	route := router.New()
	route.GET("/", viewHandler)
	log.Print("serve")
  http.ListenAndServe(":8080", route)
}
