package arrays

func CanCompleteCircuit(gas []int, cost []int) int {
	var tank, curr_pos int
	for start := 0; start < len(gas); start++ {
		tank = gas[start]
		if cost[start] <= tank {
			if start == len(gas)-1 {
				curr_pos = 0
			} else {
				curr_pos = start + 1
			}
			tank = tank - cost[start] + gas[curr_pos]
			for curr_pos != start {
				if cost[curr_pos] > tank {
					break
				} else {
					if curr_pos == len(gas)-1 {
						tank = tank - cost[curr_pos] + gas[0]
						curr_pos = 0
					} else {
						tank = tank - cost[curr_pos] + gas[curr_pos+1]
						curr_pos++
					}
				}
			}
			if curr_pos == start {
				return start
			}
		}
	}
	return -1
}

func CanCompleteCircuit2(gas, cost []int) int {
	curr_gain, total_gain := 0, 0
	var pos int

	for i := 0; i < len(gas); i++ {
		total_gain += gas[i] - cost[i]
		curr_gain += gas[i] - cost[i]

		if curr_gain < 0 {
			pos = i + 1
			curr_gain = 0
		}
	}
	if total_gain >= 0 {
		return pos
	}
	return -1
}
