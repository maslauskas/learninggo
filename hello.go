package main

import "fmt"

func main() {
    fmt.Println(hello("world"))
}

func hello(name string) string {
    if name == "" {
        name = "world"
    }
    return "Hello, " + name
}

func Add(i int, i2 int) int {
    return i + i2
}