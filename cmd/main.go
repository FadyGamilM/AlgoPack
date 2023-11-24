package main

import (
	"log"

	"github.com/FadyGamilM/algopack/searching"
)

func main() {
	arr := make([]int, 0)
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}

	result := searching.BinarySearch(arr, 10)
	if result == nil {
		log.Println("not found")
	} else {
		log.Println("➜ ", *result)
	}

}
