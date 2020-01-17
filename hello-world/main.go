package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(str string) string {
	return englishHelloPrefix + str
}

func main() {
	fmt.Println(Hello("world"))
}
