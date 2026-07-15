package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println("hello", reflect.TypeFor[string]())

	res := func(a int, b int) int {
		return a + b
	}

	fmt.Println(res(4, 5))

	if err := run(); err != nil {
		log.Fatal(err)
	}

}

func run() error {
	input := "3"

	val, err := parseLevel(input)

	if err != nil {
		return err
	}

	fmt.Println(val)

	return nil
}

func parseLevel(s string) (int, error) {
	n, err := strconv.Atoi(s)

	if err != nil {
		return 0, fmt.Errorf("level must be a number")
	}

	if n < 1 || n > 5 {
		return 0, fmt.Errorf("level must be between 1 to 5")
	}

	return n, nil
}
