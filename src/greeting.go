package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	x := 0.0001
	for i := 0; i <= 1000000; i++ {
		x = math.Sqrt(x)
	}
	fmt.Fprintf(w, "%s", greeting("Code.education Rocks!"))
}

func greeting(str string) string {
	return "<b>" + str + "</b>"
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
