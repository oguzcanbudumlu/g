package main

import "fmt"

const helloPrefix = "hello "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello(""))
}
