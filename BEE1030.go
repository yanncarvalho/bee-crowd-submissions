package main

import (
	"fmt"
)

func RemoveElement(slice []int, index int) []int {
	result := make([]int, 0)
	result = append(result, slice[:index]...)
	result = append(result, slice[index+1:]...)
	return result
}

func CreateArray(lenght int) []int {
	result := make([]int, 0)
	for i := 0; i < lenght; i++ {
		result = append(result, i)
	}
	return result
}

func main() {

	var cases, numPeople, leap = 0, 0, 0
	fmt.Scanf("%d\n", &cases)
	result := make([]int, 0)
	for i := 0; i < cases; i++ {
		fmt.Scanf("%d %d\n", &numPeople, &leap)
		people := CreateArray(numPeople)
		leap--
		var index = 0
		for 1 < len(people) {
			index += leap
			if len(people) <= index {
				index = index % len(people)
			}
			people = RemoveElement(people, index)
		}
		result = append(result, people[0]+1)
	}

	for i := 0; i < len(result); i++ {
		fmt.Printf("Case %d: %d\n", (i + 1), result[i])
	}

}
