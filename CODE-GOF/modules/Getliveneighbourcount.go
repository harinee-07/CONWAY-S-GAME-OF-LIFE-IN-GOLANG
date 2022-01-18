package modules
//Function to check the neighbour count at a particular position of the grid

func Getliveneighbourcount(currgen [][]int, i, j, m, n int) int {
	count := 0
	for k := i - 1; k <= i+1; k++ {
		for l := j - 1; l <= j+1; l++ {
			x := k
			y := l
			if x < 0 {
				x = m - 1
			}

			if x > m-1 {
				x = 0
			}

			if y < 0 {
				y = n - 1
			}

			if y > n-1 {
				y = 0
			}

			if x == i && y == j {
				continue
			}

			if currgen[x][y] == 1 {
				count++
			}
		}
	}

	return count
}