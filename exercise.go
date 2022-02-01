package main

import "fmt"

func main() {
	inp := [][]string{
		{"X", "X", "X", "X", "X", "X"},
		{"X", "X", "0", "X", "X", "0"},
		{"0", "X", "0", "X", "0", "X"},
		{"0", "0", "0", "0", "0", "0"}}
	// 0 living X dead
	//We will compare according to these constraints
	/*• If the cell's current status is "alive": ○ At less than two living neighboring cells, the cell dies (subpopulation). ○ By two or three living neighboring cells, the cell will live on. ○ If the cell has more than three living neighboring cells, it will die (overpopulation) • If the cell is "dead": ○ The cell's status becomes "live" (reproduction) if it has exactly three living neighbors.*/
	for i := 1; i < 3; i++ {
		for j := 1; j < 5; j++ {
			al := 0
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if inp[i+x][j+y] == "0" {
						al = al + 1
					}
				}
			}
			if inp[i][j] == "0" {
				al = al - 1
			}
			// now for condition checking
			if al < 2 && inp[i][j] == "0" {
				inp[i][j] = "X"
			}
			if al == 2 || al == 3 {
				inp[i][j] = "0"
			}
			if al > 3 && inp[i][j] == "0" {
				inp[i][j] = "X"
			}
		}
	}

	//lets print the solution now
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(inp[i][j], "\t")
		}
		fmt.Println()
	}
}
