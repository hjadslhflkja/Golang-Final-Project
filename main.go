package main

import (
	"fmt"
	"time"
)

func timer(x int) {
	for {
		// check if end condition is met
		if x <= 0 {
			break
		}
		if x > 0 {
			fmt.Println(x)
			time.Sleep(1 * time.Second) // wait 1 sec
			x--                         // reduce time
		}
	}
}

func main() {

	var array [50]int
	var newArray [50]int
	var total, count int
	fmt.Print("How many numbers you want to enter, cowboy: ")
	fmt.Scanln(&count)
	timer(4)
	for i := 0; i < count; i++ {
		fmt.Print("Enter value : ")
		fmt.Scanln(&array[i])
		total += array[i]
		timer(4)
	}

	top := 0 //array[0]
	bottom := array[0]
	for i := 0; i < count; i++ {
		if array[i] >= top {
			top = array[i]
		}
		if array[i] <= bottom {
			bottom = array[i]
		}
	}

	for v := bottom; v <= top; v++ {
		for a := 0; a <= count; a++ {
			if v == array[a] {
				newArray[a] = v
			}
		}
	}
	medSpotArray := count / 2
	median := 0.0
	if count%2 == 1 {
		median = newArray[medSpotArray]
	}
	if count%2 == 0 {
		median = (newArray[medSpotArray] + newArray[medSpotArray-1]) / 2
	}

	//function above doesn't word with odd array values; needs casting

	rge := (top - bottom)
	average := (total / count) * 1.0
	fmt.Printf("Average is  %.2f ", average)
	fmt.Printf(" range is  %d", rge)
	fmt.Printf(" median is  %.2f ", median)
}
http.ListenAndServe(":8080", nil)

