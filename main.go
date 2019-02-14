package main

import (
	"fmt"
	"net/http"

	hanlders "gitlab.com/didikprabowo/crud-go/handlers"
)

func main() {

	http.HandleFunc("/index", hanlders.Pindex)
	http.HandleFunc("/posts/add", hanlders.Padd)
	http.HandleFunc("/posts/edit/", hanlders.Pedit)
	http.HandleFunc("/posts/update", hanlders.Pupdate)
	http.HandleFunc("/posts/delete", hanlders.Pdelete)
	fmt.Println("starting web server at http://localhost:8090/")
	http.ListenAndServe(":8090", nil)
}
