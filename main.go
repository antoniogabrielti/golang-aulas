package main

import "fmt"

/*

Package main demonstrates a simple example of a Go program that prints "Hello, World!".
It provides a basic structure for a Go application.

Additional Environment Information:
	- GOOS: Defines the operating system target (e.g., windows, linux, darwin).
	- GOARCH: Specifies the architecture target (e.g., amd64, 386, arm, arm64).

These environment variables are critical when building cross-platform Go applications.

go env -w GOOS=windows
go env -w GOARCH=amd64
go build -o hello.exe main.go

*/

func main() {
    fmt.Println("hello world")

    a, b := 10, 5
    fmt.Printf("%d + %d = %d\n", a, b, a+b)
    fmt.Printf("%d - %d = %d\n", a, b, a-b)
    fmt.Printf("%d * %d = %d\n", a, b, a*b)
    fmt.Printf("%d / %d = %d\n", a, b, a/b)

    switch a {
        case 1:
            fmt.Println("One")
        case 2:
            fmt.Println("Two")
        default:
            fmt.Println("Another number")   
    }
}