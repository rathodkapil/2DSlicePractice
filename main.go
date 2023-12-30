package main

import "fmt"

func create2DSlice(rows int, columns int) [][]int {

	slice := make([][]int, rows)

	for i := range slice {
		// populating each rows
		slice[i] = make([]int, columns)
	}

	return slice

}

func populate(slice [][]int) {

	for i := range slice {

		for j := range slice[i] {
			if i == 1 && j == 2 {
				slice[i][j] = 100
			}
		}
	}

	//fmt.Println(slice)
}

func iterateAndFetch(slice [][]int, target int) {

	for i := range slice {

		for j := range slice[i] {

			if slice[i][j] == target {
				fmt.Println(i, j)
			}
		}
	}

}

func createSlice2(rows int, col int) {

	s := make([][]int, rows)

	for r := range s {
		s[r] = make([]int, col)
	}

	//counter := 4

	for i := 0; i < rows; i++ {
		for j := 0; j < col; j++ {
			s[i][j+i] = 1
			break
		}

	}

	/*
		[1000]
		[0100]
		[0010]
		[0001]
	*/

	fmt.Println(s)
}

func main() {

	// mySlice := create2DSlice(3, 4)
	// populate(mySlice)
	// fmt.Println(mySlice)
	// iterateAndFetch(mySlice, 100)
	createSlice2(4, 4)

}
