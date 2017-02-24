package main

import (
	"fmt"
  "net/http"

	"github.com/shinofara/gorouter"
	"log"
)

type Param struct {
	Name string `schema:"name"`
}

func viewHandler(w http.ResponseWriter, r *http.Request, params Param) error {
	fmt.Fprintf(w, "Hello %+v", params)
	return nil
}

func main() {
	route := gorouter.New()
	route.GET("/", viewHandler)
	log.Print("serve")
  http.ListenAndServe(":8080", route)
}
