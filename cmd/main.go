package main

import (
	"fmt"

	utils "github.com/idylicaro/go-bank/internal/utils"
)

func Hello() string {
	return utils.ViperEnvVariable("HELLO_WORLD")
}

func main() {
	fmt.Print("Hello")
}
