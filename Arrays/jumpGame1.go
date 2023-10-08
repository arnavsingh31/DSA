package arrays

/*
	LC #55
	TC--->O(n)
	SC--->O(1)
*/
func CanJump(arr []int) bool {
	can_reach := arr[0]

	for i := 0; i <= can_reach; i++ {
		if i == len(arr)-1 {
			return true
		}
		if can_reach < i+arr[i] {
			can_reach = i + arr[i]
		}

		if can_reach >= len(arr)-1 {
			return true
		}
	}
	return false
}
