package modules

//Function to get the next generation
func Getnextgen(currgen [][]int, m, n int) [][]int {
	
	nextGen := make([][]int, m)
	for i := range nextGen {
		nextGen[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnt := Getliveneighbourcount(currgen, i, j, m, n)

			if currgen[i][j] == 1 {
				if cnt < 2 || cnt > 3 {
					nextGen[i][j] = 0
				}

				if cnt == 2 || cnt == 3 {
					nextGen[i][j] = 1
				}
			} else {
				if cnt == 3 {
					nextGen[i][j] = 1
				}
			}

		}
	}

	return nextGen
}

