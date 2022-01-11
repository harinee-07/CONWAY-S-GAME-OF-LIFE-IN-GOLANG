package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGeneration(t *testing.T) {
	var testBoardSize = 8
	testBoard := [][]int{{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0}}
	testNewBoard := make([][]int, testBoardSize)
	for i := range testNewBoard {
		testNewBoard[i] = make([]int, testBoardSize)
	}
	testWantBoard := [][]int{{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0}}
	testNewBoard = ComputeNextGen(testBoard, testBoardSize, testBoardSize)
	result := reflect.DeepEqual(testNewBoard, testWantBoard)
	fmt.Println(result)
	if result == false {
		t.Errorf("got %v, wanted %v", testNewBoard, testWantBoard)
	}
}
