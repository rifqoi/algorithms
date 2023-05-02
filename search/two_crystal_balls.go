package main

import (
	"fmt"
	"math"
)

func twoCrystalBalls(floors []bool) (bool, int) {
	floorLength := float64(len(floors))
	jumpRange := int(math.Floor(math.Sqrt(floorLength)))

	// Checking if there are no true values in the floors
	if floors[len(floors)-1] == false {
		return false, -1
	}

	i := jumpRange
	for ; i <= len(floors); i += jumpRange {
		if floors[i] == true {
			fmt.Println(jumpRange)
			fmt.Println(i, floors[i])
			break
		}
	}

	i -= jumpRange

	for j := 0; j <= jumpRange && i < len(floors); i, j = i+1, j+1 {
		if floors[i] {
			return true, i
		}
	}

	return false, -1
}

func main() {
	arr := []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true}

	fmt.Println(twoCrystalBalls(arr))
}
