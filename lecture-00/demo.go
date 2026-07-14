package main

import (
	"fmt"
	"reflect"
	"strings"
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
		fmt.Printf("damn that's expensive %.2f \n", total)
	} else {
		fmt.Println("yess!")
	}

	sum := 0

	for i := 1; i < 101; i++ {
		sum = sum + i
	}
	fmt.Printf("the value of the sum is %d \n", sum)

	day := "monday"

	switch day {
	case strings.ToLower("Monday"):
		fmt.Println("this is Monday")
	case strings.ToLower("Sunday"):
		fmt.Println("this is Sunday")
	case strings.ToLower("Tuesday"):
		fmt.Println("this is Tuesday")
	case strings.ToLower("Friday"):
		fmt.Println("this is Friday")
	default:
		fmt.Println("not a valid week name")
	}

	fmt.Println()

	like, commentCount, shareCount := 23, 54, 100
	fmt.Println(like, commentCount, shareCount)

	fmt.Printf("%T \n", commentCount)
	fmt.Println(reflect.TypeOf(shareCount))

	fmt.Println(rating)
	fmt.Println(city)

	fmt.Println(year)

	// default declaration

	var a [5]int
	fmt.Println(a)

	// array literals
	//primes := [3]string{"shams", "pahlowan", "soad"}
	primes := [...]string{"shams", "pahlowan", "soad"} // prev line and this one gives the same result

	fmt.Println(primes)

	//sparse initialization
	b := [10]int{5: 37}
	fmt.Println(b)

	for _, v := range b {
		fmt.Println(v)
	}
}
