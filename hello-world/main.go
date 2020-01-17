package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
 	englishHelloSuffix = "World"
)

func Hello(str string) string {
	if len(str) == 0 {
		return englishHelloPrefix + englishHelloSuffix
	}
	return englishHelloPrefix + str
}

func main() {
	fmt.Println(Hello("world"))
}
