package modules
import "fmt"

//Function to print the Grid
func PrinttheGrid(grid[][]int,m, n int)  {

	for k := 0; k < m; k++ {
		for l := 0; l < n; l++ {
			fmt.Printf("%d ", grid[k][l])
		}
		fmt.Printf("\n")
	}

}