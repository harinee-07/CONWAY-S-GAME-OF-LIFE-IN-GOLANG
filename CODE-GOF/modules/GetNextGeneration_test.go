package modules

//Import neccessary libraries
import (
	"fmt"
	"reflect"
	"testing"

)
//Function to test the next generation grid
func TestGeneration(t *testing.T) {
	//size of the Grid
	var testGridSize = 8
	//Sample Input
	testGrid := [][]int{{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0}}
	//make a grid to store the current output
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
	//function call to get the output values
	testNewGrid =Getnextgen(testGrid, testGridSize, testGridSize)
	//Check for the result
	result := reflect.DeepEqual(testNewGrid, testWantGrid)
	fmt.Println(result)
	//If fails print the expected and the current output
	if result == false {
		t.Errorf("got %v, wanted %v", testNewGrid, testWantGrid)
	}
}