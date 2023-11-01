package stack

/*
	LC #84
	TC---> O(n)
	SC---> O(n)
	Watch video for reference.
*/

type Dimensions struct {
	Height int
	Index  int
}

func LargestRectangleArea(heights []int) int {
	stack := make([]Dimensions, 0)
	maxArea := 0

	for i, h := range heights {
		start := i
		for len(stack) > 0 && stack[len(stack)-1].Height > h {
			stackTop := stack[len(stack)-1]
			maxArea = calculateMaxArea(stackTop.Height, i-stackTop.Index, maxArea)
			stack = stack[:len(stack)-1]
			start = stackTop.Index
		}

		stack = append(stack, Dimensions{Height: h, Index: start})
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		maxArea = calculateMaxArea(top.Height, len(heights)-top.Index, maxArea)
		stack = stack[:len(stack)-1]
	}

	return maxArea
}

func calculateMaxArea(height, width, maxArea int) int {
	area := height * width

	if area > maxArea {
		return area
	}
	return maxArea
}

/*
	TC---> O(n)
	SC---> O(n)
	for every histogram find previous smaller and next smaller histogram index (which will give the range upto
	which our currents histogram can expand i.e required width).
*/
func LargestRectangleArea2(heights []int) int {
	nextSmaller := make([]int, 0)
	prevSmaller := make([]int, 0)
	stack := []int{}

	for i := 0; i < len(heights); i++ {
		nextSmaller = append(nextSmaller, len(heights))
		prevSmaller = append(prevSmaller, -1)
	}

	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			stackTop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextSmaller[stackTop] = i
		}

		if len(stack) > 0 {
			prevSmaller[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	maxArea := 0
	for i, h := range heights {
		// for width we are doing nextSmaller[i]- prevSmaller[i] - 1, decrementing by 1 becoz we need to find
		// no of histograms b/w next and previous smaller index which also equals the required width of rectangle.
		maxArea = calculateMaxArea(h, nextSmaller[i]-prevSmaller[i]-1, maxArea)
	}
	return maxArea
}
