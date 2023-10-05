package hashmaps

/*
	read below explanations to understand T.C of algorithm.
 	idea is to detect a loop cycle while generating next number for eg:
	n =8 --> 8->64->52->29->85->89->145->42->20->4->16->37->58->89...now we stop as we reached 89 second time
	now we are in a loop so we return false.

	https://leetcode.com/problems/happy-number/editorial/?envType=study-plan-v2&envId=top-interview-150
	https://stackoverflow.com/questions/50261364/explain-why-time-complexity-for-summing-digits-in-a-number-of-length-n-is-ologn
*/

func IsHappy(n int) bool {
	hashMap := make(map[int]bool)
	for {
		if n == 1 {
			return true
		}
		if _, exist := hashMap[n]; exist {
			return false
		}
		hashMap[n] = true
		n = getNext(n)
	}
}

func getNext(n int) int {
	sum := 0
	for n > 0 {
		d := n % 10
		n = n / 10
		sum += d * d
	}

	return sum
}
