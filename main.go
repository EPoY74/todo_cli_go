package main

import (
	"fmt"
	"os"
)

func main() {
	argCount := len(os.Args)
	fmt.Printf("Количество элементов %d/n", argCount)
	fmt.Println("Аргументы: ", os.Args)
}