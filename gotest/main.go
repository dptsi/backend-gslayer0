package main

import (
	"fmt"
	"log"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "halo ini server bintang")
}

func main() {
	http.HandleFunc("/", Welcome)
	log.Fatal(http.ListenAndServe("10.160.4.76:7777", nil))
}
