package main

import (
	"fmt"
	"net/http"

	"math/rand"
)

func helloClient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Client, Welcome to my new server")
}

func coinToss(w http.ResponseWriter, r *http.Request) {
	value := rand.Intn(2)
	if value == 1 {
		fmt.Fprintf(w, "Head is the call")
	} else if value == 0 {
		fmt.Fprintf(w, "Tail is the call")
	}
}
func main() {
	http.HandleFunc("/", helloClient)
	http.HandleFunc("/cointoss", coinToss)
	if err := http.ListenAndServe(":8090", nil); err != nil {
		fmt.Println(err)
	}

}
