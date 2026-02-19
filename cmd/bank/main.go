package main

import "fmt"

func main() {
	categories := []string{"auto", "food", "happiness", "asd"}
	top3 := categories[:3]
	fmt.Printf("%T\n", top3)
	fmt.Println(len(top3))
	fmt.Println(cap(top3))

}
