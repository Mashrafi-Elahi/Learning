package main

import "fmt"

func main() {
    scores := []int{1, 2, 3}
    
    realAppend(&scores)
    
    fmt.Println("Scores in main:", scores)
    fmt.Println("Scores length in main:", len(scores))
}

func realAppend(s *[]int) {
    // Dereferencing the pointer to modify the actual header
    *s = append(*s, 4)
    (*s)[0] = 100
}