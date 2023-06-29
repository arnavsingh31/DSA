package main

import "log"

// T.C---> Theta(d) d= no. of bits from LSB to MSB.
func count(n int) int {
	count := 0
	for n > 0 {
		if n%2 != 0 {
			count++
		}
		n = n / 2
	}
	return count
}

/*
	Brian Kerningam's algorithm.
	T.C--> Theta(no. of set bits)
*/
func count2(n uint32) int {
	count := 0
	log.Print(n)
	for n > 0 {
		n = n & (n - 1)
		count++
	}
	log.Print(count)
	return count
}

/*
	Lookup Table solution T.C---> O(1)
	Assumption: we have 32 bit numbers
	first we will pre-compute for 8 bit numbers (0->255) and store result in a table
	since number is 32 bit it can be broken into 4-> 8-bit parts and we can get number of set bits in 8 bit no
	using precomputed lookup-table.
	in order to get 8 bit number from 32 bit, we take largest 8-bit no which is 255 which contains 8 1's.
	and do "&" operation with n which gives us last 8 bits of n(How??) and we use these  8 bits to lookup for no.
	of 1's and inorder to get next 8 bits of n we simply shift n by 8, 16, 24.

	how??
	for eg n is any 32 bit number
	n   =  11010001 00001101 10101011 11001000
	255 =  00000000 00000000 00000000 11111111

	when we perform "&" operation between n & 255 we get only last 8 bit of n remaining 24 bit becomes 0.

*/

func count3(n uint32) int {

	setBitCountMap := make(map[uint32]int)
	res := 0
	setBitCountMap[0] = 0
	var i uint32
	for i = 1; i <= 255; i++ {
		setBitCountMap[i] = setBitCountMap[i&(i-1)] + 1
	}
	res = setBitCountMap[n&255] + setBitCountMap[(n>>8)&255] + setBitCountMap[(n>>16)&255] + setBitCountMap[(n>>24)&255]
	log.Print(res)
	return res
}
