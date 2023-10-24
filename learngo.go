// Explaination:
// 1. package main: package declaration, every go file must start with a package declaration
// 2. import "fmt": import declaration, import packages that we need to use fmt extends to format
// 3.  main function declaration, the entry point of the program is the main function

// how to compile and run the program?
// 1. go build learngo.go
// 2. ./learngo

// Go is a statically typed language and a dynamically typed language
// example:

// package declaration, every go file must start with a package declaration
package main

import "fmt"

func main() {
	var name string = "Hello World"
	country := "Cabo Verde"
	fmt.Println(name + ", " + country + "! Welcome to Go World!")
}

// start git repository
// git init
// git add .
// git commit -m "first commit"
// git remote add origin
// git push -u origin master
