package main

func candy(ratings []int) int {
	if len(ratings) <= 1 {
		return 1
	}

	candy_array := []int{}

	for i := 0; i < len(ratings); i++ {
		candy_array = append(candy_array, 1)
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candy_array[i] = candy_array[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candy_array[i] <= candy_array[i+1] {
			candy_array[i] = candy_array[i+1] + 1
		}
	}

	min_candy := 0
	for i := 0; i < len(candy_array); i++ {
		min_candy += candy_array[i]
	}

	return min_candy
}
