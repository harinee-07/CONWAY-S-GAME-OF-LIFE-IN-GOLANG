package modules

//Importing the Neccessary libararies
import (
	"fmt"
	"reflect"
	"testing"

)

//Function to test the count of the neighbour at a particular position of the grid
func TestCount(t *testing.T) {
	//Grid Size
	var testGridSize = 8
	//Variables to store the expected and current output
	var testWantCount,testNewCount int
	//Grid structure
	testGrid := [][]int{{0, 1, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0}}
	testWantCount=3
	//Get the values from the Function
	testNewCount=Getliveneighbourcount(testGrid,0,0, testGridSize, testGridSize)
	//Check for the outputs
	result := reflect.DeepEqual(testNewCount, testWantCount)
	fmt.Println(result)
	//Print the expected and the current output if false
	if result == false {
		t.Errorf("got %v, wanted %v", testNewCount, testWantCount)
	}
}