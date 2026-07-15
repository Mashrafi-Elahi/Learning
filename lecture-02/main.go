package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic recovered!")
		}
	}()
	defer fmt.Print("one")
	defer fmt.Print("Two")
	defer fmt.Print("Three")
	fmt.Println("hello")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic recovered!")
			panic("ohh noooow i panicked again")
		}
	}()
	run()
	panic("noooooo")

}

func run() {

	for i := 10; i >= 0; i-- {
		fmt.Println(2 / i)
	}
}
