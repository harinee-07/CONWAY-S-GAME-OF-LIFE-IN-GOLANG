package modules

//Function to check the Extinction of the grid at each generation
func CheckExtinct(grid[][]int ,m,n int,extinct bool) bool {
	for k := 0; k < m; k++ {
		for l := 0; l < n; l++ {
			if grid[k][l] == 1 {
				extinct = false
				break
			}
		}
	}

return extinct
}