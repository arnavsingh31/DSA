package main

/*
	LC #402
	Intuition:
	Maintaining digits in ascending order from most significant digit produces least possible number.
	traversing from left to right -> remove the previously seen peak element to flatten the curve.
	eg := 372181 and k=1
                   8
	  7           / \
	 / \         /   \
	/   \       /     \
	3    \     /       \
		  2   /         \
		   \ /           \
		    1             1

	here removing peak value which is 8 will result in 37211 but this is not the smallest number. If we remove
	7 we get the smallest possible number which is 32181. This is happening becoz 7 has more weightage than 8.
	First dip in the above curve is from 7->2 and second dip in the curve is from 8->1. We don't care what is the
	height of the peak, we want to process it dip-wise. Whenever the curve dips this mean we have already seen
	a maxima, if maxima is already seen then this mean it will have higher weightage than the maxima which will
	occur after this dip. So after removing 7 in this case 2nd digit of resulting number will be 2 while
	removing 8 will result in 2nd digit to be 7. So even after removing any digit either 7 or 8, the resulting
	number will have 1 digit less, but then leftmost digit is having the most weightage, so we want to decrement
	value of the leftmost digit as much as possible (depending upon k). So that is why will process elemnets from
	left to right and when we see first dip we have to remove previous peek element in order to faltten the curve.

				   8
	              / \
	             /   \
	    3       /     \
	     \     /       \
	  	  2   /         \
		   \ /           \
		    1             1

	After removing 7, suppose we have k=2, in above curve we see we have first peak at 3 and second peak at 8.
	But since 3 has more weightage since it more siginificant digit than 8, so we remove 3. We get resulting
	number = 2181
*/
func removeKDigits(s string, k int) string {
	if len(s) == k {
		return "0"
	}
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		for len(stack) > 0 && k > 0 && stack[len(stack)-1] > s[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, s[i])
	}

	// removing digits if after traversing k is still > 0
	stack = stack[:len(stack)-k]

	for len(stack) > 1 && stack[0] == '0' {
		stack = stack[1:]
	}

	return string(stack)
}
