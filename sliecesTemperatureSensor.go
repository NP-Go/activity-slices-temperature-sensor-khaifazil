package main

import "fmt"

func main() {
	//slices template
	room1 := []float64{20, 21, 23, 25, 22}
	room2 := []float64{27, 23, 25, 20, 30, 24}
	room3 := []float64{22, 23, 24, 22}

	//Insert code here
	sliceInSlice := [][]float64{}

	sliceInSlice = append(sliceInSlice, room1, room2, room3)
	fmt.Println(sliceInSlice)

	for i, room := range sliceInSlice {
		var total float64
		for _, temp := range room {
			total += temp
		}
		avetemp := total / float64(len(room))
		fmt.Printf("Average temperature in room %v is %.2f\n", i+1, avetemp)
	}
}
