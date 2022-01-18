package modules

import "fmt"
//Function to get the initial values
func Get_input() (int,int,int,int){
	var m,n,gen,getlivecell int
	fmt.Printf("Enter the row value of the grid:")
	fmt.Scanln(&m)
	fmt.Printf("Enter the column value of the grid:")
	fmt.Scanln(&n)
	fmt.Printf("Enter the no of generations to run")
	fmt.Scanln(&gen)
	fmt.Printf("Enter the live cell count for the grid:")
	fmt.Scanln(&getlivecell)
	return m,n,gen,getlivecell
}
