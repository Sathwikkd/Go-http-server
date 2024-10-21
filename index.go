package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/vedas", veda)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func index(writer http.ResponseWriter, request *http.Request) {
	msg := "hello from goo.."
	fmt.Fprintln(writer, msg)
}
func veda(writer http.ResponseWriter, request *http.Request) {
	var vedas = []string{"Rig", "Yaju", "Sama", "Atharva"}
	for _, data := range vedas {
		fmt.Fprintln(writer, data)
	}
}
