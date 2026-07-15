// package main

// import "fmt"

// func main() {
//     // We pre-allocate a slice with a length of 3 and a capacity of 10
//     s := make([]int, 3, 10) 
//     s[0], s[1], s[2] = 10, 20, 30

//     modifySlice(s)

//     fmt.Println("Main slice after function:", s)
//     fmt.Println("Main slice length:", len(s))
// }

// func modifySlice(x []int) {
//     x[0] = 999
//     x = append(x, 40) // within the capacity
//     fmt.Println("Inside function slice:", x)
// }


// //999, 20, 30, 40
// //999, 20, 30
// //3
package main