package main

import (
	"fmt"
	// "os"
	"flag"
)

func main() {
	// argCount := len(os.Args)
	// fmt.Printf("Количество элементов %d\n", argCount)
	// fmt.Println("Аргументы: ", os.Args)

	// for _, value := range os.Args {
	// 	fmt.Printf(" %s\n", value)
	// }

	var (
		name   string
		age    int
		height float64
	)

	flag.StringVar(&name, "name", "John Dow", "Your name")
	flag.IntVar(&age, "age", 21, "Your age")
	flag.Float64Var(&height, "height", 180, "Ваш рост (см)")

	flag.Parse()

	fmt.Printf("Name: %s", name)
	fmt.Printf("Age: %d", age)
	fmt.Printf("Height: %.1f", height)

}
