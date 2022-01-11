package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGeneration(t *testing.T) {
	var testGridSize = 8
	testGrid := [][]int{{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0}}
	testNewGrid := make([][]int, testGridSize)
	for i := range testNewGrid {
		testNewGrid[i] = make([]int, testGridSize)
	}
	//expected output
	testWantGrid := [][]int{{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0}}
	testNewGrid = getnextgen(testGrid, testGridSize, testGridSize)
	result := reflect.DeepEqual(testNewGrid, testWantGrid)
	fmt.Println(result)
	if result == false {
		t.Errorf("got %v, wanted %v", testNewGrid, testWantGrid)
	}
}
