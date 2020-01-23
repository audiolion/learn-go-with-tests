package main

import "net/http"
import "github.com/audiolion/learn-go-with-tests/dependency-injection"

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(dependency_injection.MyGreeterHandler))
}
