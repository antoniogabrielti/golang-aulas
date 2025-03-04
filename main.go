package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

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
	fmt.Println("Jogo da Adivinhação")

	fmt.Println(
		"Um número aleatório será sorteado. Tente acertar. O número é um inteiro entre 0 a 100.",
	)

	x := rand.Int64N(101)

	fmt.Println("O número sorteado é", x)

	scanner := bufio.NewScanner(os.Stdin)

	chutes := [10]int64{}

	for i := range chutes {
		fmt.Println("Qual é o seu chute?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)

		if err != nil {
			fmt.Println("O seu chute tem que ser um número inteiro.")
			return
		}

		switch {
		case chuteInt < x:
			fmt.Println("Você errou. O número sorteado é maior que", chuteInt)
		case chuteInt > x:
			fmt.Println("Você errou. O número sorteado é menor que", chuteInt)
		case chuteInt == x:
			return
		}

		chutes[i] = chuteInt
	}
}
