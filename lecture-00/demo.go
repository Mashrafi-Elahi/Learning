package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("hello")

	// variable declaration syntaxt

	// var name type

	var year int16 = 2026
	var rating float64 = 4.5

	// the walrus operator

	const city = "dhaka🤣"

	mark := 78

	if mark > 90 {
		fmt.Println("great job!")
	} else if mark > 80 {
		fmt.Println("good job!")
	} else if mark > 60 {
		fmt.Println("hmm great")
	} else {
		fmt.Println("try hard man!")
	}

	item := 50
	pricePerItem := 32.35

	if total := float64(item) * pricePerItem; total > 100 {
		fmt.Printf("damn that's expensive %.2f", total)
	} else {
		fmt.Println("yess!")
	}

	fmt.Println()

	like, commentCount, shareCount := 23, 54, 100
	fmt.Println(like, commentCount, shareCount)

	fmt.Printf("%T \n", commentCount)
	fmt.Println(reflect.TypeOf(shareCount))

	fmt.Println(rating)
	fmt.Println(city)

	fmt.Println(year)
}
