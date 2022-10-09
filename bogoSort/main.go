package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func main() {
	list := []int{3, 2, 1}
	expected := []int{1, 2, 3}

	iteration := 0
	// if not sorted, keep sorting
	for !reflect.DeepEqual(list, expected) {
		fmt.Println("Sorting iteration", iteration)
		list = bogoSort(list)
		iteration++
	}
	fmt.Println("Sorted list:", list)

}

func bogoSort(list []int) []int {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})

	return list
}
