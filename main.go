package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func Sum(a, b int) int {
	return a + b
}

func handler(w http.ResponseWriter, r *http.Request) {
	total := Sum(rand.Intn(100), rand.Intn(100))

	fmt.Fprintln(w, total)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
