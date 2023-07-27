package main

import (
	"sort"
)

/*
LC #853
TC--> O(nlogn) because we are sorting.
SC--> O(n) as we are using space to store sorted arrangement of car in an array.

Note: we can also use stack to implement the same.
- Intuition
Use a stack based approach to detrmine the number of car fleets that will arrive at the destination.
Start off by sorting the cars based on their positions in descending order. Then, iterate through
the sorted cars and calculate the time it takes for each car to reach the target destination. If the
calculated time is greater than the time of the car fleet at the top of the stack (or if the stack is empty),
we add this time to the stack to represent a new car fleet. Finally, return the count of car fleets,
which is the length of the stack.

- Approach
1. Create a struct type car to store the position and speed of a car.

2. Create a slice of type car to store the cars and initialize it with the given positions and speeds.

3. Sort the cars based on their positions in descending order using the sort.Slice function.

4. Create an empty stack of type float64 to store the times of the car fleets.

5. Iterate through the sorted cars:

	Calculate the time it takes for the current car to reach the target destination using the formula
	(target - car.position) / car.speed.

	Check if the stack is empty or if the calculated time is greater than the time of the car fleet
	at the top of the stack.

	If true, append the calculated time to the stack.

6. Return the count of car fleets, which is the length of the stack.
*/
type Car struct {
	Position int
	Speed    int
}

func carFleet(target int, postion, speed []int) int {
	arr := []Car{}

	for i := 0; i < len(postion); i++ {
		arr = append(arr, Car{Position: postion[i], Speed: speed[i]})
	}

	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i].Position > arr[j].Position
	})

	prevTimeToReachTarget := 0.0
	fleet := 0
	for _, car := range arr {
		currTimeToReachTarget := float64(target-car.Position) / float64(car.Speed)

		if currTimeToReachTarget > prevTimeToReachTarget {
			fleet++
			prevTimeToReachTarget = currTimeToReachTarget
		}
	}
	return fleet
}

// func main() {
// 	carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3})
// }
