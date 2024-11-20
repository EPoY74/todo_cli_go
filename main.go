package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	dbw "todo_go/dbw"
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

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: \n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}

	flag.StringVar(&name, "name", "John Dow", "Your name")
	flag.IntVar(&age, "age", 21, "Your age")
	flag.Float64Var(&height, "height", 180, "Ваш рост (см)")

	flag.Parse()
	if len(os.Args) < 4 {
		flag.Usage()
		fmt.Scanln()
		os.Exit(1)
	}

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Height: %.2f\n", height)

	id_rec, err := dbw.ReadRec(1)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Прочитано: %v", id_rec)
	}

}
