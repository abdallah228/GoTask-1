package main

import (
	"fmt"
	"tasksGo/otherpackage"
)

var hi string
var welcome string

func main() {
	msg := "hello world"
	fmt.Println("task1", msg)
	hiFunc()
	welcomeFunc()
	welcomeSamePackage()
	otherpackage.OtherPackage()
}

func welcomeFunc() {
	welcome = "welcome message from task 3"
	fmt.Println(welcome)
}

func hiFunc() {
	hi = "hi every body from task 2"
	fmt.Println(hi)
}
